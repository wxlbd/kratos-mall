package log

import (
	"context"
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/wxlbd/kit/log/handler"
	"gorm.io/gorm/logger"
	"io"
	"log/slog"
	"os"
	"runtime"
	"time"
)

var (
	_ logger.Interface = (*Logger)(nil)
	_ log.Logger       = (*Logger)(nil)
)

type Options struct {
	Level      slog.Level
	Writer     io.Writer
	TimeFormat string
}

type Option interface {
	Apply(*Options)
}

type optionFunc func(*Options)

func (f optionFunc) Apply(opts *Options) {
	f(opts)
}

func WithLevel(level slog.Level) Option {
	return optionFunc(func(opts *Options) {
		opts.Level = level
	})
}

func WithWriter(writer io.Writer) Option {
	return optionFunc(func(opts *Options) {
		opts.Writer = writer
	})
}

func WithTimeFormat(format string) Option {
	return optionFunc(func(opts *Options) {
		opts.TimeFormat = format
	})
}

type Logger struct {
	*slog.Logger
	*handler.Handler
}

var defaultOptions = Options{
	Level:      slog.LevelInfo,
	Writer:     os.Stdout,
	TimeFormat: time.DateTime,
}

func NewLogger(opts ...Option) *Logger {
	for _, opt := range opts {
		opt.Apply(&defaultOptions)
	}
	h := handler.NewHandler(defaultOptions.Writer, &handler.Options{
		TimeFormat: defaultOptions.TimeFormat,
		Level:      defaultOptions.Level,
	})
	return &Logger{
		Logger:  slog.New(h),
		Handler: h,
	}
}

func (h *Logger) Log(level log.Level, keyAndValues ...any) error {
	var pcs [1]uintptr
	runtime.Callers(4, pcs[:])
	pc := pcs[0]
	var r slog.Record
	switch level {
	case log.LevelDebug:
		r = slog.NewRecord(time.Now(), slog.LevelDebug, "", pc)
		r.Add(keyAndValues...)
	case log.LevelInfo:
		r = slog.NewRecord(time.Now(), slog.LevelInfo, "", pc)
		r.Add(keyAndValues...)
	case log.LevelWarn:
		r = slog.NewRecord(time.Now(), slog.LevelWarn, "", pc)
		r.Add(keyAndValues...)
	case log.LevelError:
		r = slog.NewRecord(time.Now(), slog.LevelError, "", pc)
		r.Add(keyAndValues...)
	case log.LevelFatal:
		r = slog.NewRecord(time.Now(), slog.LevelError, "", pc)
		r.Add(keyAndValues...)
	}
	return h.Handle(context.TODO(), r)
}
func (h *Logger) LogMode(_ logger.LogLevel) logger.Interface {
	return h
}

func (h *Logger) Info(ctx context.Context, s string, i ...any) {
	if h.Handler.Enabled(ctx, slog.LevelInfo) {
		var pcs [1]uintptr
		runtime.Callers(4, pcs[:])
		pc := pcs[0]
		r := slog.NewRecord(time.Now(), slog.LevelInfo, "", pc)
		r.AddAttrs(slog.String("msg", s))
		r.Add(i...)
		_ = h.Handle(ctx, r)
	}
}

func (h *Logger) Warn(ctx context.Context, s string, i ...interface{}) {
	if h.Handler.Enabled(ctx, slog.LevelWarn) {
		var pcs [1]uintptr
		runtime.Callers(4, pcs[:])
		pc := pcs[0]
		r := slog.NewRecord(time.Now(), slog.LevelInfo, "", pc)
		r.AddAttrs(slog.String("msg", s))
		r.Add(i...)
		_ = h.Handle(ctx, r)
	}
}

func (h *Logger) Error(ctx context.Context, s string, i ...interface{}) {
	if h.Handler.Enabled(ctx, slog.LevelError) {
		var pcs [1]uintptr
		runtime.Callers(4, pcs[:])
		pc := pcs[0]
		r := slog.NewRecord(time.Now(), slog.LevelInfo, "", pc)
		r.AddAttrs(slog.String("msg", s))
		r.Add(i...)
		_ = h.Handle(ctx, r)
	}
}

func (h *Logger) Trace(ctx context.Context, begin time.Time, fc func() (sql string, rowsAffected int64), err error) {
	if h.Handler.Enabled(ctx, slog.LevelInfo) {
		var pcs [1]uintptr
		runtime.Callers(4, pcs[:])
		pc := pcs[0]
		r := slog.NewRecord(time.Now(), slog.LevelInfo, "", pc)
		sql, rows := fc()
		elapsed := time.Since(begin)
		if err != nil {
			r.AddAttrs(handler.Err(err))
		}
		if rows != -1 {
			r.AddAttrs(
				slog.String("time", fmt.Sprintf("%.3fms", float64(elapsed.Nanoseconds())/1e6)),
				slog.String("sql", sql),
			)
			_ = h.Handle(ctx, r)
		}
	}
}
