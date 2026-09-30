package main

import (
	"context"
	"errors"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"microservices/auth/internal/config"
	"microservices/auth/internal/httpapi"
	"microservices/auth/internal/service"
	"microservices/auth/internal/token"
	"microservices/auth/internal/user"
)

func main() {
	log := slog.New(slog.NewJSONHandler(os.Stdout, nil)).With("service", "auth")
	cfg := config.Load()

	tokens := token.NewManager(cfg.JWTSecret, cfg.AccessTTL, cfg.RefreshTTL)
	auth := service.NewAuth(user.NewMemoryRepository(), tokens)

	srv := &http.Server{
		Addr:              cfg.Addr,
		Handler:           httpapi.New(auth, log),
		ReadHeaderTimeout: 5 * time.Second,
	}

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	go func() {
		log.Info("listening", "addr", cfg.Addr)
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Error("server failed", "err", err)
			stop()
		}
	}()

	<-ctx.Done()
	log.Info("shutting down")
	shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := srv.Shutdown(shutdownCtx); err != nil {
		log.Error("shutdown failed", "err", err)
	}
}
