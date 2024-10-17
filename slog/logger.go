package slog

import (
	"context"
	"log/slog"
	"os"
	"sync/atomic"
)

var defaultLogger atomic.Pointer[Logger]
var defaultLevel *Level = new(Level)

func init() {
	defaultLevel = LevelInfo
	defaultLogger.Store(New(os.Stdout, JSONHandler, *defaultLevel))
}

func SetDefault(l *Logger) {
	defaultLogger.Store(l)
}

func SetDefaultLevel(l Level) {
	*defaultLevel = l
}

func Default() *Logger { return defaultLogger.Load() }

func Trace(msg string, args ...any) {
	Default().log(context.Background(), LevelTrace, msg, args...)
}

func Tracef(format string, args ...any) {
	Default().logf(context.Background(), LevelTrace, format, args...)
}

func TraceContext(ctx context.Context, msg string, args ...any) {
	Default().log(ctx, LevelTrace, msg, args...)
}

func Debug(msg string, args ...any) {
	Default().log(context.Background(), LevelDebug, msg, args...)
}

func Debugf(format string, args ...any) {
	Default().logf(context.Background(), LevelDebug, format, args...)
}

func DebugContext(ctx context.Context, msg string, args ...any) {
	Default().log(ctx, LevelDebug, msg, args...)
}

func Info(msg string, args ...any) {
	Default().log(context.Background(), LevelInfo, msg, args...)
}

func Infof(format string, args ...any) {
	Default().logf(context.Background(), LevelInfo, format, args...)
}

func InfoContext(ctx context.Context, msg string, args ...any) {
	Default().log(ctx, LevelInfo, msg, args...)
}

func Warn(msg string, args ...any) {
	Default().log(context.Background(), LevelWarn, msg, args...)
}

func Warnf(format string, args ...any) {
	Default().logf(context.Background(), LevelWarn, format, args...)
}

func WarnContext(ctx context.Context, msg string, args ...any) {
	Default().log(ctx, LevelWarn, msg, args...)
}

func Error(msg string, args ...any) {
	Default().log(context.Background(), LevelError, msg, args...)
}

func Errorf(format string, args ...any) {
	Default().logf(context.Background(), LevelError, format, args...)
}

func ErrorContext(ctx context.Context, msg string, args ...any) {
	Default().log(ctx, LevelError, msg, args...)
}

func Fatal(msg string, args ...any) {
	Default().log(context.Background(), LevelFatal, msg, args...)
}

func Fatalf(format string, args ...any) {
	Default().logf(context.Background(), LevelFatal, format, args...)
}

func FatalContext(ctx context.Context, msg string, args ...any) {
	Default().log(ctx, LevelFatal, msg, args...)
}

func Log(ctx context.Context, level Level, msg string, args ...any) {
	Default().log(ctx, level, msg, args...)
}

func Logf(ctx context.Context, level Level, format string, args ...any) {
	Default().logf(ctx, level, format, args...)
}

func LogAttrs(ctx context.Context, level Level, msg string, attrs ...slog.Attr) {
	Default().logAttrs(ctx, level, msg, attrs...)
}
