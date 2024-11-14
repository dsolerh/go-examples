package config

import (
	"errors"
	"flag"
	"mallbots/internal/logger"
	"mallbots/internal/rpc"
	"mallbots/internal/web"
	"time"
)

const (
	DefaultEnvironment     = DevEnv
	DefaultShutdownTimeout = 30 * time.Second
	DefaultLogLevel        = logger.Debug
)

var (
	ErrInvalidDBConn      = errors.New("invalid database connection string")
	ErrInvalidLoggerLevel = errors.New("invalid logger level config")
	ErrInvalidEnvironment = errors.New("invalid environment config")
)

type Environment = string

const (
	DevEnv  Environment = "dev"
	ProdEnv Environment = "prod"
)

func IsValidEnv(env Environment) bool {
	return env == DevEnv ||
		env == ProdEnv
}

type DBConfig struct {
	Conn string
}

type AppConfig struct {
	Env             Environment
	LoggerConfig    logger.LogConfig
	DBConfig        DBConfig
	RPCConfig       rpc.RpcConfig
	WebConfig       web.WebConfig
	ShutdownTimeout time.Duration `envconfig:"SHUTDOWN_TIMEOUT" default:"30s"`
}

func InitConfig() (cfg AppConfig, err error) {
	flag.StringVar(&cfg.Env, "env", DefaultEnvironment, "the environment of the app")
	// logger
	flag.StringVar(&cfg.LoggerConfig.LogLevel, "logLevel", string(DefaultLogLevel), "the log level of the app")
	// db
	flag.StringVar(&cfg.DBConfig.Conn, "dbConn", "", "the connection string to the database")
	// rpc
	flag.StringVar(&cfg.RPCConfig.Host, "rpcHost", rpc.DefaultHost, "sets the rpc host address")
	flag.StringVar(&cfg.RPCConfig.Port, "rpcPort", rpc.DefaultPort, "sets the rpc host port")
	// web
	flag.StringVar(&cfg.WebConfig.Host, "webHost", web.DefaultHost, "sets the web host address")
	flag.StringVar(&cfg.WebConfig.Port, "webPort", web.DefaultPort, "sets the web host port")
	// shutdown
	flag.DurationVar(&cfg.ShutdownTimeout, "shutdownTimeout", DefaultShutdownTimeout, "controls the timeout for the shutdown of the app")

	flag.Parse()

	// validate the config
	if cfg.DBConfig.Conn == "" {
		err = ErrInvalidDBConn
		return
	}

	if !logger.IsValidLevel(cfg.LoggerConfig.LogLevel) {
		err = ErrInvalidLoggerLevel
		return
	}

	if !IsValidEnv(cfg.Env) {
		err = ErrInvalidEnvironment
		return
	}

	return
}
