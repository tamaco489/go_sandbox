package logger

import (
	"context"
	"log/slog"
	"net/http"
	"os"
	"sync"
)

var (
	globalLogger *AppLogger
	once         sync.Once
)

// New: create new logger instance
func New() *AppLogger {
	once.Do(func() {
		handler := slog.NewJSONHandler(
			os.Stdout, &slog.HandlerOptions{
				Level: slog.LevelDebug,
			},
		)

		globalLogger = &AppLogger{
			Logger: slog.New(handler),
		}
	})
	return globalLogger
}

// GetLogger: get global logger instance
func GetLogger() *AppLogger {
	if globalLogger == nil {
		return New()
	}
	return globalLogger
}

// DebugContext: output debug log
func (l *AppLogger) DebugContext(ctx context.Context, msg string, args ...any) {
	l.Logger.DebugContext(ctx, msg, args...)
}

// InfoContext: output info log
func (l *AppLogger) InfoContext(ctx context.Context, msg string, args ...any) {
	l.Logger.InfoContext(ctx, msg, args...)
}

// WarnContext: output warn log
func (l *AppLogger) WarnContext(ctx context.Context, msg string, args ...any) {
	l.Logger.WarnContext(ctx, msg, args...)
}

// ErrorContext: output error log
func (l *AppLogger) ErrorContext(ctx context.Context, msg string, args ...any) {
	l.Logger.ErrorContext(ctx, msg, args...)
}

// FatalContext: output fatal log
func (l *AppLogger) FatalContext(ctx context.Context, msg string, args ...any) {
	l.Logger.ErrorContext(ctx, msg, args...)
	os.Exit(1)
}

// LogRequestCompletion: Log request completion with appropriate level based on status code
func (l *AppLogger) LogRequestCompletion(ctx context.Context, statusCode int, httpInfo HTTPRequestInfo, systemInfo SystemInfo, authInfo AuthorizedInfo) {
	// Create structured log attributes using structures directly
	attrs := []any{
		"http_info", httpInfo,
		"system_info", systemInfo,
		"auth_info", authInfo,
	}

	// Determine log level and log with appropriate method
	switch {
	// status: 5xx, level: error
	case statusCode >= http.StatusInternalServerError:
		l.ErrorContext(ctx, "request failed", attrs...)

	// status: 4xx, level: warn
	case statusCode >= http.StatusBadRequest:
		l.WarnContext(ctx, "request failed", attrs...)

	// status: 2xx, level: info
	default:
		l.InfoContext(ctx, "request completed", attrs...)
	}
}
