package main

import (
	"fmt"
	"log/slog"
	stdslog "log/slog"
	"time"

	"slog-test/unilogger"
)

func main() {
	initWrappedSlogSlog()
}

type Alert struct {
	Msg   string
	Time  string
	Level string
}

func initWrappedSlogSlog() {
	logger := unilogger.NewLogger(unilogger.Options{
		Level: unilogger.LevelInfo.Level(),
		// Output: os.Stdout,
	})

	newLogger := logger.WithGroup("module")

	newLogger.Info("we can see dat", slog.Any("helm", &Alert{
		Msg:   "some message",
		Time:  "some time",
		Level: "some level",
	}))
	newLogger.Info("we can see dat", slog.Any("helm", map[string]string{
		"msg":   "some message",
		"time":  "some time",
		"level": "some level",
	}))
	newLogger.Info("we can see dat", slog.String("raw;json;some log", `{"time":"value"}`))

	unilogger.SetDefault(logger)

	fmt.Println("INFO LEVEL")
	unilogger.Debug("we cant see dat", "int attr", 42, stdslog.String("str arg", "lol"))
	unilogger.Info("we can see dat", "int attr", 42, stdslog.String("str arg", "lol"))
	unilogger.Warn("we can see dat", "int attr", 42, stdslog.String("str arg", "lol"))
	unilogger.Error("we can see dat", "int attr", 42, stdslog.String("str arg", "lol"))

	unilogger.SetDefaultLevel(unilogger.LevelDebug)
	fmt.Println("DEBUG LEVEL")
	logger.Debugf("we can see dat %s %d %v", "int attr", 42, stdslog.String("str arg", "lol"))
	unilogger.Info("we can see dat" /* "int attr", 42, stdslog.String("str arg", "lol")*/)
	unilogger.Warn("we can see dat", "int attr", 42, stdslog.String("str arg", "lol"))
	unilogger.Error("we can see dat", "int attr", 42, stdslog.String("str arg", "lol"))

	unilogger.SetDefaultLevel(unilogger.LevelError)
	fmt.Println("ERROR LEVEL")
	unilogger.Debug("we cant see dat", "int attr", 42, stdslog.String("str arg", "lol"))
	unilogger.Info("we cant see dat", "int attr", 42, stdslog.String("str arg", "lol"))
	unilogger.Warn("we cant see dat", "int attr", 42, stdslog.String("str arg", "lol"))
	unilogger.Error("we can see dat", "int attr", 42, stdslog.String("str arg", "lol"))
	// unilogger.Fatal("we can see dat", "int attr", 42, stdslog.String("str arg", "lol"))

	unilogger.SetDefaultLevel(unilogger.LevelInfo)
	fmt.Println("INFO LEVEL")
	unilogger.Debug("we cant see dat", "int attr", 42, stdslog.String("str arg", "lol"))
	unilogger.Info("we can see dat", "int attr", 42, stdslog.String("str arg", "lol"))
	unilogger.Warn("we can see dat", "int attr", 42, stdslog.String("str arg", "lol"))
	unilogger.Error("we can see dat", "int attr", 42, stdslog.String("str arg", "lol"))

	first := logger.Named("first")
	second := first.Named("second")
	third := second.Named("third")

	third.Error("well shit")
	second.Error("well shit")
	first.Error("well shit")
}

func parallelSourceTest() {
	logger := unilogger.NewLogger(unilogger.Options{
		Level: unilogger.LevelInfo.Level(),
		// Output: os.Stdout,
	})

	go func() {
		for range 1000000 {
			go func() {
				logger.SetLevel(unilogger.LevelDebug)
				logger.Info("stub msg", slog.Any("time", "kekerity"))
				logger.SetLevel(unilogger.LevelInfo)
				logger.Info("stub msg", slog.Any("time", "kekerity"))
				logger.SetLevel(unilogger.LevelDebug)
				logger.Info("stub msg", slog.Any("time", "kekerity"))
				logger.SetLevel(unilogger.LevelInfo)
				logger.Info("stub msg", slog.Any("time", "kekerity"))
			}()
		}
	}()

	time.Sleep(5 * time.Minute)
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
