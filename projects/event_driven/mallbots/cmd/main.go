package main

import (
	"database/sql"
	"log"

	"github.com/go-chi/chi/v5"
	_ "github.com/jackc/pgx/v5"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"mallbots/internal/config"
	"mallbots/internal/logger"
	"mallbots/internal/rpc"
	"mallbots/internal/waiter"
	"mallbots/internal/web"
)

func main() {
	if err := run(); err != nil {
		log.Fatalln(err)
	}
}

func run() (err error) {
	var cfg config.AppConfig
	cfg, err = config.InitConfig()
	if err != nil {
		return err
	}

	app := App{cfg: cfg}

	// db
	app.db, err = sql.Open("pgx", cfg.DBConfig.Conn)
	if err != nil {
		return err
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			return
		}
	}(app.db)

	// logger
	app.logger = logger.New(cfg.Env, cfg.LoggerConfig)

	app.rpc = initRpc(cfg.RPCConfig)
	app.mux = initMux(cfg.WebConfig)
	app.waiter = waiter.New(waiter.CatchSignals())

	// init modules
	// m.modules = []modules.Module{
	// 	&baskets.Module{},
	// 	&customers.Module{},
	// 	&depot.Module{},
	// 	&notifications.Module{},
	// 	&ordering.Module{},
	// 	&payments.Module{},
	// 	&stores.Module{},
	// }
	return
}

func initRpc(_ rpc.RpcConfig) *grpc.Server {
	server := grpc.NewServer()
	reflection.Register(server)

	return server
}

func initMux(_ web.WebConfig) *chi.Mux {
	return chi.NewMux()
}
