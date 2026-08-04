package main

import (
	"context"
	"crypto/tls"
	"errors"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	slog.SetDefault(slog.New(slog.NewJSONHandler(os.Stdout, nil)))

	cfg, err := loadConfig()
	if err != nil {
		slog.Error("invalid configuration", "error", err)
		os.Exit(1)
	}

	limiter := newRateLimiter(cfg.RateLimitRPS, cfg.RateLimitBurst)
	defer limiter.Stop()

	mux := http.NewServeMux()
	mux.HandleFunc("GET /healthz", healthHandler)
	mux.HandleFunc("GET /readyz", readyHandler)
	mux.Handle("GET /v1/example", requireAPIKey(cfg.APIKey)(http.HandlerFunc(exampleHandler)))

	handler := requestIDMiddleware(
		loggingMiddleware(
			securityHeaders(
				cors(cfg.AllowedOrigins)(
					maxBytes(cfg.MaxRequestBytes)(
						limiter.Middleware(mux),
					),
				),
			),
		),
	)
	handler = recoverMiddleware(handler)

	srv := &http.Server{
		Addr:              ":" + cfg.Port,
		Handler:           http.TimeoutHandler(handler, cfg.RequestTimeout, `{"error":"request timeout"}`),
		ReadHeaderTimeout: cfg.ReadHeaderTimeout,
		ReadTimeout:       cfg.ReadTimeout,
		WriteTimeout:      cfg.WriteTimeout,
		IdleTimeout:       cfg.IdleTimeout,
		TLSConfig:         &tls.Config{MinVersion: tls.VersionTLS12},
	}

	go func() {
		slog.Info("starting server", "addr", srv.Addr)
		// Terminate TLS here (srv.ListenAndServeTLS(certFile, keyFile)) or,
		// more commonly for a microservice, at an upstream load balancer /
		// service-mesh sidecar and keep this listener on the internal network.
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			slog.Error("server failed", "error", err)
			os.Exit(1)
		}
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
	<-stop

	slog.Info("shutting down")
	ctx, cancel := context.WithTimeout(context.Background(), cfg.ShutdownTimeout)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		slog.Error("graceful shutdown failed", "error", err)
	}
}
