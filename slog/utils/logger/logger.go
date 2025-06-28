package logger

import (
	"context"
	"log/slog"
	"net/http"
	"os"
)

// globalLogger: global logger instance
var globalLogger *Logger

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

// GetLogger: get global logger instance
func GetLogger() *Logger {
	if globalLogger == nil {
		globalLogger = New()
	}
	return globalLogger
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

// ResponseWriterWrapper: wrapper to record status code and context
type ResponseWriterWrapper struct {
	http.ResponseWriter
	statusCode int
	ctx        *context.Context
}

// NewResponseWriterWrapper: create new ResponseWriterWrapper
func NewResponseWriterWrapper(w http.ResponseWriter) *ResponseWriterWrapper {
	defaultStatusCode := http.StatusOK
	return &ResponseWriterWrapper{
		ResponseWriter: w,
		statusCode:     defaultStatusCode,
	}
}

func (rw *ResponseWriterWrapper) WriteHeader(statusCode int) {
	rw.statusCode = statusCode
	rw.ResponseWriter.WriteHeader(statusCode)
}

func (rw *ResponseWriterWrapper) Write(data []byte) (int, error) {
	if rw.statusCode == 0 {
		rw.statusCode = 200
	}
	return rw.ResponseWriter.Write(data)
}

func (rw *ResponseWriterWrapper) UpdateContext(ctx context.Context) {
	if rw.ctx == nil {
		rw.ctx = &ctx
	} else {
		*rw.ctx = ctx
	}
}

func (rw *ResponseWriterWrapper) GetContext() *context.Context {
	return rw.ctx
}

func (rw *ResponseWriterWrapper) GetStatusCode() int {
	return rw.statusCode
}
