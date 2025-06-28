package logging

import (
	"context"
	"net/http"
	"time"

	"github.com/tamaco489/go_sandbox/slog/utils/configuration"
	"github.com/tamaco489/go_sandbox/slog/utils/logger"
)

// logRequestCompletion: Log request completion with appropriate level based on status code
func logRequestCompletion(ctx context.Context, statusCode int, httpInfo logger.HTTPRequestInfo, systemInfo logger.SystemInfo, authInfo logger.AuthorizedInfo) {
	logger := logger.GetLogger()

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
		logger.ErrorContext(ctx, "Request completed", attrs...)

	// status: 4xx, level: warn
	case statusCode >= http.StatusBadRequest:
		logger.WarnContext(ctx, "Request completed", attrs...)

	// status: 2xx, level: info
	default:
		logger.InfoContext(ctx, "Request completed", attrs...)
	}
}

// RequestMiddleware: Manage request start and end
func RequestMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Record request start time
		startTime := time.Now()

		// Initialize request context
		r = initializeRequestContext(r)

		// Create ResponseWriterWrapper (keep context pointer)
		wrappedWriter := logger.NewResponseWriterWrapper(w)

		// Initialize ctx field and then call UpdateContext
		wrappedWriter.UpdateContext(r.Context())

		// defer for request end logging
		defer func() {
			// Get latest context from wrappedWriter.ctx (updated by authorization middleware)
			finalAuthInfo, _ := logger.GetAuthorizedInfoCtx(*wrappedWriter.GetContext())
			systemInfo, _ := logger.GetSystemInfoCtx(*wrappedWriter.GetContext())
			requestID, _ := logger.GetRequestIDCtx(*wrappedWriter.GetContext())

			// Calculate processing time
			latency := time.Since(startTime)

			// Create HTTP information
			httpInfo := logger.HTTPRequestInfo{
				Method:     r.Method,
				Path:       r.URL.Path,
				Status:     wrappedWriter.GetStatusCode(),
				Latency:    latency.String(),
				UserAgent:  r.UserAgent(),
				Referer:    r.Referer(),
				RemoteAddr: r.RemoteAddr,
				RequestID:  requestID,
			}

			// Log request completion with simplified structure
			ctx := *wrappedWriter.GetContext()
			logRequestCompletion(ctx, wrappedWriter.GetStatusCode(), httpInfo, systemInfo, finalAuthInfo)
		}()

		// Execute next handler
		next.ServeHTTP(wrappedWriter, r)
	}
}

// initializeRequestContext:
func initializeRequestContext(r *http.Request) *http.Request {
	// Generate request ID and set it in context
	ctx := logger.SetRequestIDCtx(r.Context())

	// Initialize system information
	env := configuration.GetEnvironment()
	systemInfo := logger.NewSystemInfo(env)
	ctx = logger.SetSystemInfoCtx(ctx, systemInfo)

	// Set initial authorized information
	authInfo := logger.NewInitialAuthorizedInfo()
	ctx = logger.SetAuthorizedInfoCtx(ctx, authInfo)

	// Update request with updated context
	return r.WithContext(ctx)
}
