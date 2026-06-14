package main

import (
	"fmt"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Anything struct {
	Data string
	num  int
}

type AL Anything

// MarshalLogObject implements [zapcore.ObjectMarshaler].
func (a AL) MarshalLogObject(enc zapcore.ObjectEncoder) error {
	fmt.Printf("a: %v\n", a)
	enc.AddString("Data", a.Data)
	enc.AddInt("num", a.num)
	return nil
}

var _ zapcore.ObjectMarshaler = AL{}

func main() {
	output := os.Stdout
	level := zapcore.DebugLevel
	enc := zapcore.NewJSONEncoder(zapcore.EncoderConfig{
		TimeKey:        "ts",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.StringDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	})

	core := zapcore.NewCore(enc, zapcore.Lock(output), level)
	options := []zap.Option{zap.AddCaller()}

	logger := zap.New(core, options...)

	anything := Anything{Data: "some data", num: 42}
	logger.Debug("WOrks", zap.Any("any", anything), zap.Object("obj", AL(anything)))
	logger.Info("Somethig else")
}
