package server

import (
	"context"
	"errors"
	"net/http"
	"os"
	"os/signal"

	"net/http/pprof"
	"syscall"
	"time"

	"github.com/prometheus/client_golang/prometheus/promhttp"
	otelapi "go.opentelemetry.io/otel/metric"
	"go.uber.org/zap"

	"github.com/flashbots/chain-monitor/config"
	"github.com/flashbots/chain-monitor/httplogger"
	"github.com/flashbots/chain-monitor/logutils"
	"github.com/flashbots/chain-monitor/metrics"
	"github.com/flashbots/chain-monitor/utils"
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

	if cfg.Server.EnablePprof {
		mux.HandleFunc("/debug/pprof/", pprof.Index)
		mux.HandleFunc("/debug/pprof/cmdline", pprof.Cmdline)
		mux.HandleFunc("/debug/pprof/profile", pprof.Profile)
		mux.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
		mux.HandleFunc("/debug/pprof/trace", pprof.Trace)
		mux.Handle("/debug/pprof/heap", pprof.Handler("heap"))
		mux.Handle("/debug/pprof/goroutine", pprof.Handler("goroutine"))
		mux.Handle("/debug/pprof/block", pprof.Handler("block"))
		mux.Handle("/debug/pprof/threadcreate", pprof.Handler("threadcreate"))
	}

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

	if err := metrics.Setup(ctx, s.cfg.L2.ProbeTx, s.observe); err != nil {
		return err
	}

	go func() { // run the server
		l.Info("Chain monitor server is going up...",
			zap.String("server_listen_address", s.cfg.Server.ListenAddress),
			zap.Int("pid", os.Getpid()),
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

	{ // close the rpc clients
		s.l1.rpc.Close()
		s.l2.rpc.Close()
	}

	return utils.FlattenErrors(errs)
}

func (s *Server) observe(ctx context.Context, o otelapi.Observer) error {
	return errors.Join(
		s.l1.observeWallets(ctx, o),
		s.l2.observeBlockHeight(ctx, o),
		s.l2.observeWallets(ctx, o),
		s.l2.observerProbes(ctx, o),
	)
}
