package logger

import (
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/pkgerrors"
)

type Level = string

const (
	Trace Level = "TRACE"
	Debug Level = "DEBUG"
	Info  Level = "INFO"
	Warn  Level = "WARN"
	Error Level = "ERROR"
	Panic Level = "PANIC"
)

func IsValidLevel(lvl Level) bool {
	return lvl == Trace ||
		lvl == Debug ||
		lvl == Info ||
		lvl == Warn ||
		lvl == Error ||
		lvl == Panic
}

type LogConfig struct {
	LogLevel Level
}

func New(env string, cfg LogConfig) zerolog.Logger {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnixMs
	zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack

	switch env {
	case "prod":
		return zerolog.New(os.Stdout).
			Level(logLevelToZero(cfg.LogLevel)).
			With().
			Timestamp().
			Logger()
	case "dev":
		return zerolog.New(zerolog.NewConsoleWriter(formatTime)).
			Level(logLevelToZero(cfg.LogLevel)).
			With().
			Timestamp().
			Logger()
	default:
		panic("invalid env")
	}
}

func logLevelToZero(level Level) zerolog.Level {
	switch level {
	case Panic:
		return zerolog.PanicLevel
	case Error:
		return zerolog.ErrorLevel
	case Warn:
		return zerolog.WarnLevel
	case Info:
		return zerolog.InfoLevel
	case Debug:
		return zerolog.DebugLevel
	case Trace:
		return zerolog.TraceLevel
	default:
		return zerolog.InfoLevel
	}
}

func formatTime(w *zerolog.ConsoleWriter) {
	w.TimeFormat = "03:04:05.000PM"
}
