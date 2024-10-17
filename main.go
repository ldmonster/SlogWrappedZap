package main

import (
	"context"
	"errors"
	"os"
	"slog-test/slog"
	"slog-test/zap"
	"time"

	stdslog "log/slog"
)

func main() {
	initSlogSlog()
}

func initSlogSlog() {
	logger := slog.New(os.Stdout, slog.JSONHandler, slog.LevelInfo)
	newLogger := logger.WithGroup("module")
	logger.Fatal("kek")
	newLogger.Trace("kek")
}

func initZapInSlog() {
	zapLogger, err := zap.NewZapLogger("debug")
	if err != nil {
		panic("error when make zap logger")
	}

	log := zap.NewSlogLogger(zap.NewZapHandler(zapLogger))

	log.Info("info test", stdslog.String("string key", "string value"))

	log.WithGroup("monitoring").Info("metrics was collected")
	log.With(stdslog.String("first arg", "first"), stdslog.String("second arg", "second")).Info("with args")
	log.Info("log level debug is enabled?", stdslog.Bool("check", log.Enabled(context.Background(), stdslog.LevelDebug)))

	log.Warn("warn test", stdslog.Bool("string key", false))
	log.Debug("debug test", stdslog.Int("integer key", 100))
	log.Error("error test", stdslog.String("error key", errors.New("super_error").Error()))
	log.Fatal("fatal test", stdslog.Time("time key", time.Now()))
}
