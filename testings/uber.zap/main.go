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

	type Cs struct {
		Id   []byte
		Vars map[string]string
		Arr  []Cs
	}

	anything := Anything{Data: "some data", num: 42}
	arrstr := []string{"s1", "s2"}
	cs := &Cs{
		Id: []byte("some"),
		Vars: map[string]string{
			"a": "letter a",
		},
		Arr: []Cs{
			{Id: []byte("some1")},
			{Id: []byte("some2")},
		},
	}
	logger.Debug(
		"WOrks",
		zap.Any("any", anything),
		zap.Any("any_ptr", &anything),
		zap.Any("arrstr", arrstr),
		zap.Any("cs", cs),
		zap.Object("obj", AL(anything)),
	)
	logger.Info("Somethig else")
}
