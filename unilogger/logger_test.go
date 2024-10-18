package unilogger_test

import (
	"bytes"
	"context"
	"log/slog"
	"slog-test/unilogger"
	"testing"
	"time"

	"github.com/alecthomas/assert/v2"
)

func Test_Logger(t *testing.T) {
	t.Parallel()

	defaultLogFn := func(logger *unilogger.Logger) {
		logger.Trace("stub msg", slog.String("stub_arg", "arg"))
		logger.Debug("stub msg", slog.String("stub_arg", "arg"))
		logger.Info("stub msg", slog.String("stub_arg", "arg"))
		logger.Warn("stub msg", slog.String("stub_arg", "arg"))
		logger.Error("stub msg", slog.String("stub_arg", "arg"))
		//test fatal
		logger.Log(context.Background(), unilogger.LevelFatal.Level(), "stub msg", slog.String("stub_arg", "arg"))
	}

	type meta struct {
		name    string
		enabled bool
	}

	type fields struct {
		logfn func(logger *unilogger.Logger)
	}

	type args struct {
		addSource bool
		level     unilogger.Level
	}

	type wants struct {
		shouldContains    []string
		shouldNotContains []string
	}

	tests := []struct {
		meta   meta
		fields fields
		args   args
		wants  wants
	}{
		{
			meta: meta{
				name:    "logger default options is level info and add source false",
				enabled: true,
			},
			fields: fields{
				logfn: func(logger *unilogger.Logger) {},
			},
			args: args{},
			wants: wants{
				shouldContains: []string{
					`{"level":"info","msg":"stub msg","stub_arg":"arg","time":"2006-01-02T15:04:05Z"}`,
					`{"level":"warn","msg":"stub msg","stub_arg":"arg","time":"2006-01-02T15:04:05Z"}`,
					`{"level":"error","msg":"stub msg","stub_arg":"arg","time":"2006-01-02T15:04:05Z"}`,
					`{"level":"fatal","msg":"stub msg","stub_arg":"arg","time":"2006-01-02T15:04:05Z"}`,
				},
				shouldNotContains: []string{
					`"level":"debug"`,
					`"level":"trace"`,
				},
			},
		},
		{
			meta: meta{
				name:    "logger change to debug level should contains addsource and debug level",
				enabled: true,
			},
			fields: fields{
				logfn: func(logger *unilogger.Logger) {
					logger.SetLevel(unilogger.LevelDebug)
				},
			},
			args: args{
				level: unilogger.LevelInfo,
			},
			wants: wants{
				shouldContains: []string{
					`{"level":"debug","msg":"stub msg","source":"unilogger/logger_test.go:19","stub_arg":"arg","time":"2006-01-02T15:04:05Z"}`,
					`{"level":"info","msg":"stub msg","source":"unilogger/logger_test.go:20","stub_arg":"arg","time":"2006-01-02T15:04:05Z"}`,
					`{"level":"warn","msg":"stub msg","source":"unilogger/logger_test.go:21","stub_arg":"arg","time":"2006-01-02T15:04:05Z"}`,
					`{"level":"error","msg":"stub msg","source":"unilogger/logger_test.go:22","stub_arg":"arg","time":"2006-01-02T15:04:05Z"}`,
					`{"level":"fatal","msg":"stub msg","source":"unilogger/logger_test.go:24","stub_arg":"arg","time":"2006-01-02T15:04:05Z"}`,
				},
				shouldNotContains: []string{
					`"level":"trace"`,
				},
			},
		},
	}

	for _, tt := range tests {
		if !tt.meta.enabled {
			continue
		}

		t.Run(tt.meta.name, func(t *testing.T) {
			buf := bytes.NewBuffer([]byte{})

			logger := unilogger.NewLogger(unilogger.Options{
				AddSource: tt.args.addSource,
				Level:     tt.args.level.Level(),
				Output:    buf,
				TimeFunc: func(_ time.Time) time.Time {
					t, err := time.Parse(time.DateTime, "2006-01-02 15:04:05")
					if err != nil {
						panic(err)
					}

					return t
				},
			})

			tt.fields.logfn(logger)

			defaultLogFn(logger)

			for _, v := range tt.wants.shouldContains {
				assert.Contains(t, buf.String(), v)
			}

			for _, v := range tt.wants.shouldNotContains {
				assert.NotContains(t, buf.String(), v)
			}
		})
	}
}
