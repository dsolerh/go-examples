package main

import (
	"database/sql"
	"mallbots/internal/config"
	"mallbots/internal/waiter"

	"github.com/go-chi/chi/v5"
	"github.com/rs/zerolog"
	"google.golang.org/grpc"
)

type App struct {
	cfg    config.AppConfig
	db     *sql.DB
	logger zerolog.Logger
	// modules []monolith.Module
	mux    *chi.Mux
	rpc    *grpc.Server
	waiter waiter.Waiter
}

func (a *App) Config() config.AppConfig {
	return a.cfg
}

func (a *App) DB() *sql.DB {
	return a.db
}

func (a *App) Logger() zerolog.Logger {
	return a.logger
}

func (a *App) Mux() *chi.Mux {
	return a.mux
}

func (a *App) RPC() *grpc.Server {
	return a.rpc
}

func (a *App) Waiter() waiter.Waiter {
	return a.waiter
}
