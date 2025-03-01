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

	logger *zap.Logger
	server *http.Server

	optBuilderAddr ethcommon.Address
	optReorgWindow int
	optWallets     map[string]ethcommon.Address

	opt       *ethclient.Client
	optTicker *time.Ticker

	optBlockHeight  uint64
	optBlocks       *types.RingBuffer[bool]
	optBlocksLanded int64
	optBlocksSeen   int64
}

func New(cfg *config.Config) (*Server, error) {
	var (
		optBuilderAddr ethcommon.Address
		optWallets     = make(map[string]ethcommon.Address, len(cfg.Opt.WalletAddresses))
	)

	{ // builder address
		addr, err := ethcommon.ParseHexOrString(cfg.Opt.BuilderAddress)
		if err != nil {
			return nil, err
		}
		if len(addr) != 20 {
			return nil, errors.New("invalid length for the builder address")
		}
		copy(optBuilderAddr[:], addr)
	}

	for name, wa := range cfg.Opt.WalletAddresses {
		var addr ethcommon.Address
		_addr, err := ethcommon.ParseHexOrString(wa)
		if err != nil {
			return nil, err
		}
		if len(_addr) != 20 {
			return nil, errors.New("invalid length for the wallet address")
		}
		copy(addr[:], _addr)
		optWallets[name] = addr
	}

	opt, err := ethclient.Dial(cfg.Opt.RPC)
	if err != nil {
		return nil, err
	}

	optBlockHeight, err := opt.BlockNumber(context.Background())
	if err != nil {
		return nil, err
	}

	s := &Server{
		cfg:     cfg,
		failure: make(chan error, 1),
		logger:  zap.L(),

		opt:            opt,
		optBlockHeight: optBlockHeight - 1,
		optBuilderAddr: optBuilderAddr,
		optReorgWindow: int(cfg.Opt.ReorgWindow/cfg.Opt.BlockTime) + 1,
		optWallets:     optWallets,
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

	{ // start the optimism block ticker
		s.optTicker = time.NewTicker(s.cfg.Opt.BlockTime)
		go func() {
			for {
				<-s.optTicker.C
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
		s.optTicker.Stop()
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
		s.opt.Close()
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
