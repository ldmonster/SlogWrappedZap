package main

import (
	"context"
	"errors"
	"slog-test/zap"
	"time"

	"golang.org/x/exp/slog"
)

func main() {
	zapLogger, err := zap.NewZapLogger("debug")
	if err != nil {
		panic("error when make zap logger")
	}

	log := zap.NewSlogLogger(zap.NewZapHandler(zapLogger))

	log.Info("info test", slog.String("string key", "string value"))

	log.WithGroup("monitoring").Info("metrics was collected")
	log.With(slog.String("first arg", "first"), slog.String("second arg", "second")).Info("with args")
	log.Info("log level debug is enabled?", slog.Bool("check", log.Enabled(context.Background(), slog.LevelDebug)))

	log.Warn("warn test", slog.Bool("string key", false))
	log.Debug("debug test", slog.Int("integer key", 100))
	log.Error("error test", slog.String("error key", errors.New("super_error").Error()))
	log.Fatal("fatal test", slog.Time("time key", time.Now()))
}
