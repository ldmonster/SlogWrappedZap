package wrappedslog

import (
	"log/slog"
	"time"
)

func GetSlogOpts() *slog.HandlerOptions {
	opts := &slog.HandlerOptions{
		AddSource: true,
		Level:     LevelInfo,
		ReplaceAttr: func(_ []string, a slog.Attr) slog.Attr {
			if a.Key == slog.LevelKey {
				level := a.Value.Any().(slog.Level)
				a.Value = slog.StringValue(Level(level).String())
			}

			if a.Key == slog.TimeKey {
				timeval := a.Value.Time()

				a.Value = slog.StringValue(timeval.Format(time.RFC3339))
			}

			return a
		},
	}

	return opts
}
