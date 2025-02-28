package server

import (
	"context"
	"errors"
	"net/http"
	"os"
	"os/signal"

	"syscall"
	"time"

	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.uber.org/zap"

	"github.com/flashbots/chain-monitor/config"
	"github.com/flashbots/chain-monitor/httplogger"
	"github.com/flashbots/chain-monitor/logutils"
	"github.com/flashbots/chain-monitor/metrics"
	"github.com/flashbots/chain-monitor/types"

	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

type Server struct {
	cfg *config.Config

	failure chan error

	eth    *ethclient.Client
	logger *zap.Logger
	server *http.Server
	ticker *time.Ticker

	builderAddr ethcommon.Address

	blockHeight uint64

	blocksSeen   int64
	blocksLanded int64

	blocks  *types.RingBuffer[bool]
	wallets map[string]ethcommon.Address
}

func New(cfg *config.Config) (*Server, error) {
	var (
		builderAddr ethcommon.Address
		wallets     = make(map[string]ethcommon.Address, len(cfg.Eth.WalletAddresses))
	)

	{ // builder address
		addr, err := ethcommon.ParseHexOrString(cfg.Eth.BuilderAddress)
		if err != nil {
			return nil, err
		}
		if len(addr) != 20 {
			return nil, errors.New("invalid length for the builder address")
		}
		copy(builderAddr[:], addr)
	}

	for name, wa := range cfg.Eth.WalletAddresses {
		_addr, err := ethcommon.ParseHexOrString(wa)
		if err != nil {
			return nil, err
		}
		if len(_addr) != 20 {
			return nil, errors.New("invalid length for the wallet address")
		}
		var addr ethcommon.Address
		copy(addr[:], _addr)
		wallets[name] = addr
	}

	eth, err := ethclient.Dial(cfg.Eth.RPC)
	if err != nil {
		return nil, err
	}

	blockHeight, err := eth.BlockNumber(context.Background())
	if err != nil {
		return nil, err
	}

	s := &Server{
		blockHeight: blockHeight - 1,
		builderAddr: builderAddr,
		cfg:         cfg,
		eth:         eth,
		failure:     make(chan error, 1),
		logger:      zap.L(),
		wallets:     wallets,
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/", s.handleHealthcheck)
	mux.Handle("/metrics", promhttp.Handler())
	handler := httplogger.Middleware(s.logger, mux)

	s.server = &http.Server{
		Addr:              cfg.Server.ListenAddress,
		ErrorLog:          logutils.NewHttpServerErrorLogger(s.logger),
		Handler:           handler,
		MaxHeaderBytes:    1024,
		ReadHeaderTimeout: 30 * time.Second,
		ReadTimeout:       30 * time.Second,
		WriteTimeout:      30 * time.Second,
	}

	return s, nil
}

func (s *Server) Run() error {
	l := s.logger
	ctx := logutils.ContextWithLogger(context.Background(), l)

	if err := metrics.Setup(ctx, s.observeWallets); err != nil {
		return err
	}

	go func() { // run the server
		l.Info("Chain monitor server is going up...",
			zap.String("server_listen_address", s.cfg.Server.ListenAddress),
		)
		if err := s.server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			s.failure <- err
		}
		l.Info("Chain monitor server is down")
	}()

	{ // start the block ticker
		s.ticker = time.NewTicker(s.cfg.Eth.BlockTime)
		go func() {
			for {
				<-s.ticker.C
				s.processNewBlocks(ctx)
			}
		}()
	}

	errs := []error{}
	{ // wait until termination or internal failure
		terminator := make(chan os.Signal, 1)
		signal.Notify(terminator, os.Interrupt, syscall.SIGTERM)

		select {
		case stop := <-terminator:
			l.Info("Stop signal received; shutting down...",
				zap.String("signal", stop.String()),
			)
		case err := <-s.failure:
			l.Error("Internal failure; shutting down...",
				zap.Error(err),
			)
			errs = append(errs, err)
		exhaustErrors:
			for { // exhaust the errors
				select {
				case err := <-s.failure:
					l.Error("Extra internal failure",
						zap.Error(err),
					)
					errs = append(errs, err)
				default:
					break exhaustErrors
				}
			}
		}
	}

	{ // stop the block ticker
		s.ticker.Stop()
	}

	{ // stop the server
		ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
		defer cancel()
		if err := s.server.Shutdown(ctx); err != nil {
			l.Error("Chain monitor server shutdown failed",
				zap.Error(err),
			)
		}
	}

	{ // close the eth client
		s.eth.Close()
	}

	switch len(errs) {
	default:
		return errors.Join(errs...)
	case 1:
		return errs[0]
	case 0:
		return nil
	}
}
