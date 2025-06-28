package logger

import (
	"context"
	"log/slog"
	"net/http"
	"os"
)

// New: create new logger instance
func New() *Logger {
	handler := slog.NewJSONHandler(
		os.Stdout, &slog.HandlerOptions{
			Level: slog.LevelDebug,
		},
	)

	return &Logger{
		Logger: slog.New(handler),
	}
}

// DebugContext: output debug log
func (l *Logger) DebugContext(ctx context.Context, msg string, args ...any) {
	l.Logger.DebugContext(ctx, msg, args...)
}

// InfoContext: output info log
func (l *Logger) InfoContext(ctx context.Context, msg string, args ...any) {
	l.Logger.InfoContext(ctx, msg, args...)
}

// WarnContext: output warn log
func (l *Logger) WarnContext(ctx context.Context, msg string, args ...any) {
	l.Logger.WarnContext(ctx, msg, args...)
}

// ErrorContext: output error log
func (l *Logger) ErrorContext(ctx context.Context, msg string, args ...any) {
	l.Logger.ErrorContext(ctx, msg, args...)
}

// FatalContext: output fatal log
func (l *Logger) FatalContext(ctx context.Context, msg string, args ...any) {
	l.Logger.ErrorContext(ctx, msg, args...)
	os.Exit(1)
}

// LogRequestCompletion: Log request completion with appropriate level based on status code
func (l *Logger) LogRequestCompletion(ctx context.Context, statusCode int, httpInfo HTTPRequestInfo, systemInfo SystemInfo, authInfo AuthorizedInfo) {

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
