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
	otelapi "go.opentelemetry.io/otel/metric"
	"go.uber.org/zap"

	"github.com/flashbots/chain-monitor/config"
	"github.com/flashbots/chain-monitor/httplogger"
	"github.com/flashbots/chain-monitor/logutils"
	"github.com/flashbots/chain-monitor/metrics"
)

type Server struct {
	cfg *config.Config

	failure chan error

	logger *zap.Logger
	server *http.Server

	l1 *L1
	l2 *L2
}

func New(cfg *config.Config) (*Server, error) {
	l1, err := newL1(cfg.L1)
	if err != nil {
		return nil, err
	}

	l2, err := newL2(cfg.L2)
	if err != nil {
		return nil, err
	}

	s := &Server{
		cfg:     cfg,
		failure: make(chan error, 1),
		l1:      l1,
		l2:      l2,
		logger:  zap.L(),
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

	if err := metrics.Setup(ctx, s.observe); err != nil {
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

	{ // run the monitors
		s.l1.run(ctx)
		s.l2.run(ctx)
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

	{ // stop the monitors
		s.l2.stop()
		s.l1.stop()
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

	{ // close the rpc client
		s.l2.rpc.Close()
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

func (s *Server) observe(ctx context.Context, o otelapi.Observer) error {
	errL1 := s.l1.observeWallets(ctx, o)
	errL2 := s.l2.observeWallets(ctx, o)

	switch {
	case errL1 != nil && errL2 != nil:
		return errors.Join(errL1, errL2)
	case errL1 != nil:
		return errL1
	case errL2 != nil:
		return errL2
	}

	return nil
}
