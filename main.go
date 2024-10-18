package main

import (
	"fmt"
	stdslog "log/slog"

	"slog-test/slog"
)

func main() {
	initWrappedSlogSlog()
}

func initWrappedSlogSlog() {
	logger := slog.NewLogger(slog.Options{
		Level: slog.LevelInfo.Level(),
		// Output: os.Stdout,
	})
	// newLogger := logger.WithGroup("module")

	// slog.SetDefault(logger)

	fmt.Println("INFO LEVEL")
	slog.Debug("we cant see dat", "int attr", 42, stdslog.String("str arg", "lol"))
	slog.Info("we can see dat", "int attr", 42, stdslog.String("str arg", "lol"))
	slog.Warn("we can see dat", "int attr", 42, stdslog.String("str arg", "lol"))
	slog.Error("we can see dat", "int attr", 42, stdslog.String("str arg", "lol"))

	slog.SetDefaultLevel(slog.LevelDebug)
	fmt.Println("DEBUG LEVEL")
	slog.Debug("we can see dat", "int attr", 42, stdslog.String("str arg", "lol"))
	slog.Info("we can see dat" /* "int attr", 42, stdslog.String("str arg", "lol")*/)
	slog.Warn("we can see dat", "int attr", 42, stdslog.String("str arg", "lol"))
	slog.Error("we can see dat", "int attr", 42, stdslog.String("str arg", "lol"))

	slog.SetDefaultLevel(slog.LevelError)
	fmt.Println("ERROR LEVEL")
	slog.Debug("we cant see dat", "int attr", 42, stdslog.String("str arg", "lol"))
	slog.Info("we cant see dat", "int attr", 42, stdslog.String("str arg", "lol"))
	slog.Warn("we cant see dat", "int attr", 42, stdslog.String("str arg", "lol"))
	slog.Error("we can see dat", "int attr", 42, stdslog.String("str arg", "lol"))
	slog.Fatal("we can see dat", "int attr", 42, stdslog.String("str arg", "lol"))

	first := logger.Named("first")
	second := first.Named("second")
	third := second.Named("third")

	third.Error("well shit")
	second.Error("well shit")
	first.Error("well shit")

}

// func initZapInSlog() {
// 	zapLogger, err := zap.NewZapLogger("debug")
// 	if err != nil {
// 		panic("error when make zap logger")
// 	}

// 	log := zap.NewSlogLogger(zap.NewZapHandler(zapLogger))

// 	log.Info("info test", stdslog.String("string key", "string value"))

// 	log.WithGroup("monitoring").Info("metrics was collected")
// 	log.With(stdslog.String("first arg", "first"), stdslog.String("second arg", "second")).Info("with args")
// 	log.Info("log level debug is enabled?", stdslog.Bool("check", log.Enabled(context.Background(), stdslog.LevelDebug)))

// 	log.Warn("warn test", stdslog.Bool("string key", false))
// 	log.Debug("debug test", stdslog.Int("integer key", 100))
// 	log.Error("error test", stdslog.String("error key", errors.New("super_error").Error()))
// 	log.Fatal("fatal test", stdslog.Time("time key", time.Now()))
// }
