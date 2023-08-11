package logging

import (
	"context"
	"io"
	"os"
	"time"

	"log/slog"
)

type logger struct {
	*slog.Logger
	StartTime time.Time
	Logs      []LogEntry
}

type LogEntry struct {
	Level     string
	Message   string
	Timestamp time.Time
}

type Logger interface {
	Debug(msg string, args ...interface{})
	DebugContext(ctx context.Context, msg string, args ...interface{})
	Error(msg string, args ...interface{})
	ErrorContext(ctx context.Context, msg string, args ...interface{})
	Info(msg string, args ...interface{})
	InfoContext(ctx context.Context, msg string, args ...interface{})
	Log(ctx context.Context, level slog.Level, msg string, args ...interface{})
	LogAttrs(ctx context.Context, level slog.Level, msg string, attrs ...slog.Attr)
	Warn(msg string, args ...interface{})
	WarnContext(ctx context.Context, msg string, args ...interface{})
	Flush() error
}

func GetLogger(ctx context.Context) Logger {
	return loggerFromContext(ctx)
}

func NewLogger() *logger {
	handlerOpts := slog.HandlerOptions{
		Level: slog.LevelInfo,
	}

	handler := slog.NewTextHandler(os.Stdout, &handlerOpts)

	slogger := slog.New(handler)

	return &logger{
		Logger: slogger,
	}
}

func (l *logger) Log(ctx context.Context, level slog.Level, msg string, args ...interface{}) {
	l.Logger.Log(ctx, level, msg, args...)
}

func (l *logger) LogAttrs(ctx context.Context, level slog.Level, msg string, attrs ...slog.Attr) {
	l.Logger.LogAttrs(ctx, level, msg, attrs...)
}

func (l *logger) Debug(msg string, args ...interface{}) {
	l.Logger.Debug(msg, args...)
}

func (l *logger) DebugContext(ctx context.Context, msg string, args ...interface{}) {
	l.Logger.DebugContext(ctx, msg, args...)
}

func (l *logger) Error(msg string, args ...interface{}) {
	l.Logger.Error(msg, args...)
}

func (l *logger) ErrorContext(ctx context.Context, msg string, args ...interface{}) {
	l.Logger.ErrorContext(ctx, msg, args...)
}

func (l *logger) Info(msg string, args ...interface{}) {
	l.Logger.Info(msg, args...)
}

func (l *logger) InfoContext(ctx context.Context, msg string, args ...interface{}) {
	l.Logger.InfoContext(ctx, msg, args...)
}

func (l *logger) Warn(msg string, args ...interface{}) {
	l.Logger.Warn(msg, args...)
}

func (l *logger) WarnContext(ctx context.Context, msg string, args ...interface{}) {
	l.Logger.WarnContext(ctx, msg, args...)
}

func (l *logger) Flush() error {
	if f, ok := l.Logger.Handler().(io.Writer); ok {
		return f.(*os.File).Sync()
	}
	return nil
}
