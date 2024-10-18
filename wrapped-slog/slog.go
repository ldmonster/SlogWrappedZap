package wrappedslog

import (
	"context"
	"fmt"
	"io"
	"log/slog"
	"runtime"
	"time"
)

type logger = slog.Logger

type Logger struct {
	*logger

	opts *slog.HandlerOptions
}

type HandlerType int

const (
	JSONHandler HandlerType = iota
	TextHandler
)

func New(w io.Writer, ht HandlerType, level Level, addSource bool) *Logger {
	logger := &Logger{
		opts: GetSlogOpts(),
	}
	logger.opts.AddSource = addSource
	logger.opts.Level = level

	switch ht {
	case JSONHandler:
		logger.logger = slog.New(slog.NewJSONHandler(w, logger.opts))
	case TextHandler:
		logger.logger = slog.New(slog.NewTextHandler(w, logger.opts))
	}

	return logger
}

func (l *Logger) log(ctx context.Context, level Level, msg string, args ...any) {
	if !l.Enabled(ctx, slog.Level(level)) {
		return
	}
	var pc uintptr

	var pcs [1]uintptr
	// skip [runtime.Callers, this function, this function's caller]
	runtime.Callers(3, pcs[:])
	pc = pcs[0]

	r := slog.NewRecord(time.Now(), slog.Level(level), msg, pc)
	r.Add(args...)
	if ctx == nil {
		ctx = context.Background()
	}

	_ = l.Handler().Handle(ctx, r)
}

func (l *Logger) logf(ctx context.Context, level Level, format string, args ...any) {
	msg := fmt.Sprintf(format, args...)

	if !l.Enabled(ctx, slog.Level(level)) {
		return
	}
	var pc uintptr

	var pcs [1]uintptr
	// skip [runtime.Callers, this function, this function's caller]
	runtime.Callers(3, pcs[:])
	pc = pcs[0]

	r := slog.NewRecord(time.Now(), slog.Level(level), msg, pc)
	r.Add(args...)
	if ctx == nil {
		ctx = context.Background()
	}

	_ = l.Handler().Handle(ctx, r)
}

func (l *Logger) logAttrs(ctx context.Context, level Level, msg string, attrs ...slog.Attr) {
	if !l.Enabled(ctx, slog.Level(level)) {
		return
	}
	var pc uintptr
	var pcs [1]uintptr
	// skip [runtime.Callers, this function, this function's caller]
	runtime.Callers(3, pcs[:])
	pc = pcs[0]

	r := slog.NewRecord(time.Now(), slog.Level(level), msg, pc)
	r.AddAttrs(attrs...)
	if ctx == nil {
		ctx = context.Background()
	}

	_ = l.Handler().Handle(ctx, r)
}

func (l *Logger) SetLevel(level Level) {
	l.opts.Level = level
}

func (l *Logger) With(args ...any) *Logger {
	return &Logger{
		logger: l.logger.With(args...),
		opts:   l.opts,
	}
}

func (l *Logger) WithGroup(name string) *Logger {
	return &Logger{
		logger: l.logger.WithGroup(name),
		opts:   l.opts,
	}
}

func (l *Logger) Log(ctx context.Context, level Level, msg string, args ...any) {
	l.log(ctx, level, msg, args...)
}

func (l *Logger) Logf(ctx context.Context, level Level, format string, args ...any) {
	l.logf(ctx, level, format, args...)
}

func (l *Logger) LogAttrs(ctx context.Context, level Level, msg string, attrs ...slog.Attr) {
	l.logAttrs(ctx, level, msg, attrs...)
}

func (l *Logger) Trace(msg string, args ...any) {
	l.log(context.Background(), LevelTrace, msg, args...)
}

func (l *Logger) Tracef(format string, args ...any) {
	l.logf(context.Background(), LevelTrace, format, args...)
}

func (l *Logger) TraceContext(ctx context.Context, msg string, args ...any) {
	l.log(ctx, LevelTrace, msg, args...)
}

func (l *Logger) Debug(msg string, args ...any) {
	l.log(context.Background(), LevelDebug, msg, args...)
}

func (l *Logger) Debugf(format string, args ...any) {
	l.logf(context.Background(), LevelDebug, format, args...)
}

func (l *Logger) DebugContext(ctx context.Context, msg string, args ...any) {
	l.log(ctx, LevelDebug, msg, args...)
}

func (l *Logger) Info(msg string, args ...any) {
	l.log(context.Background(), LevelInfo, msg, args...)
}

func (l *Logger) Infof(format string, args ...any) {
	l.logf(context.Background(), LevelInfo, format, args...)
}

func (l *Logger) InfoContext(ctx context.Context, msg string, args ...any) {
	l.log(ctx, LevelInfo, msg, args...)
}

func (l *Logger) Warn(msg string, args ...any) {
	l.log(context.Background(), LevelWarn, msg, args...)
}

func (l *Logger) Warnf(format string, args ...any) {
	l.logf(context.Background(), LevelWarn, format, args...)
}

func (l *Logger) WarnContext(ctx context.Context, msg string, args ...any) {
	l.log(ctx, LevelWarn, msg, args...)
}

func (l *Logger) Error(msg string, args ...any) {
	l.log(context.Background(), LevelError, msg, args...)
}

func (l *Logger) Errorf(format string, args ...any) {
	l.logf(context.Background(), LevelError, format, args...)
}

func (l *Logger) ErrorContext(ctx context.Context, msg string, args ...any) {
	l.log(ctx, LevelError, msg, args...)
}

func (l *Logger) Fatal(msg string, args ...any) {
	l.log(context.Background(), LevelFatal, msg, args...)
}

func (l *Logger) Fatalf(format string, args ...any) {
	l.logf(context.Background(), LevelFatal, format, args...)
}

func (l *Logger) FatalContext(ctx context.Context, msg string, args ...any) {
	l.log(ctx, LevelFatal, msg, args...)
}
