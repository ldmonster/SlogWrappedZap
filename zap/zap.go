package zap

import (
	"context"
	"log/slog"

	uberzap "go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Extends default slog with new log levels
type WrappedLogger struct {
	*slog.Logger
}

func NewSlogLogger(h slog.Handler) *WrappedLogger {
	return &WrappedLogger{
		Logger: slog.New(h),
	}
}

// New LogLevels
const (
	LevelFatal = slog.Level(12)
)

// Fatal logs at LevelFatal.
func (l *WrappedLogger) Fatal(msg string, args ...any) {
	l.Log(context.Background(), LevelFatal, msg, args...)
}

func NewZapLogger(logLevel string) (*uberzap.Logger, error) {
	loggerConfig := uberzap.NewProductionConfig()

	level, err := uberzap.ParseAtomicLevel(logLevel)
	if err != nil {
		return nil, err
	}

	loggerConfig.Level = level

	zapLogger, err := loggerConfig.Build()
	if err != nil {
		return nil, err
	}

	return zapLogger, nil
}

var _ slog.Handler = (*ZapHandler)(nil)

type ZapHandler struct {
	logger *uberzap.Logger
}

func NewZapHandler(logger *uberzap.Logger) *ZapHandler {
	return &ZapHandler{
		logger: logger,
	}
}

func (h *ZapHandler) Enabled(_ context.Context, level slog.Level) bool {
	var zapLevel zapcore.Level
	switch level {
	case slog.LevelDebug:
		zapLevel = zapcore.DebugLevel
	case slog.LevelInfo:
		zapLevel = zapcore.InfoLevel
	case slog.LevelWarn:
		zapLevel = zapcore.WarnLevel
	case slog.LevelError:
		zapLevel = zapcore.ErrorLevel
	case LevelFatal:
		zapLevel = zapcore.FatalLevel
	default:
		zapLevel = zapcore.ErrorLevel // to catch default
	}

	return h.logger.Core().Enabled(zapLevel)
}

func (h *ZapHandler) Handle(_ context.Context, rec slog.Record) error {
	fields := make([]uberzap.Field, 0, rec.NumAttrs())

	rec.Attrs(func(a slog.Attr) bool {
		fields = append(fields, SlogAttToZapField(a))

		return len(fields) != cap(fields)
	})

	entry := h.logger.With(fields...)

	switch rec.Level {
	case slog.LevelDebug:
		entry.Debug(rec.Message)
	case slog.LevelInfo.Level():
		entry.Info(rec.Message)
	case slog.LevelWarn:
		entry.Warn(rec.Message)
	case slog.LevelError:
		entry.Error(rec.Message)
	case LevelFatal:
		entry.Fatal(rec.Message)
	default:
		entry.Error(rec.Message) // to catch default
	}

	return nil
}

func (h *ZapHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	fields := make([]uberzap.Field, 0, len(attrs))

	for _, attr := range attrs {
		fields = append(fields, SlogAttToZapField(attr))
	}

	return &ZapHandler{
		logger: h.logger.With(fields...),
	}
}

func (h *ZapHandler) WithGroup(name string) slog.Handler {
	return &ZapHandler{
		logger: h.logger.Named(name),
	}
}

func SlogAttToZapField(a slog.Attr) uberzap.Field {
	switch a.Value.Kind() {
	case slog.KindBool:
		return uberzap.Bool(a.Key, a.Value.Bool())
	case slog.KindDuration:
		return uberzap.Duration(a.Key, a.Value.Duration())
	case slog.KindFloat64:
		return uberzap.Float64(a.Key, a.Value.Float64())
	case slog.KindInt64:
		return uberzap.Int64(a.Key, a.Value.Int64())
	case slog.KindString:
		return uberzap.String(a.Key, a.Value.String())
	case slog.KindTime:
		return uberzap.Time(a.Key, a.Value.Time())
	case slog.KindUint64:
		return uberzap.Uint64(a.Key, a.Value.Uint64())
	default:
		return uberzap.Any(a.Key, a.Value.Any())
	}
}
