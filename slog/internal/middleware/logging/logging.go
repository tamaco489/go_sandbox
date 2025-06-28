package logging

import (
	"net/http"
	"time"

	"github.com/tamaco489/go_sandbox/slog/utils/configuration"
	"github.com/tamaco489/go_sandbox/slog/utils/logger"
)

// RequestMiddleware: Manage request start and end
func RequestMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Record request start time
		startTime := time.Now()

		// Initialize request context
		r = initializeRequestContext(r)

		// Create ResponseWriterWrapper (keep context pointer)
		wrappedWriter := logger.NewResponseWriterWrapper(w)
		wrappedWriter.UpdateContext(r.Context())

		// defer for request end logging
		defer func() {
			// Get latest context from wrappedWriter.ctx (updated by authorization middleware)
			requestID, _ := logger.GetRequestIDContext(*wrappedWriter.GetContext())
			systemInfo, _ := logger.GetSystemInfoContext(*wrappedWriter.GetContext())
			authorizedInfo, _ := logger.GetAuthorizedInfoContext(*wrappedWriter.GetContext())

			// Calculate processing time
			latency := time.Since(startTime)

			// Create HTTP information
			httpInfo := logger.NewHTTPRequestInfo(
				r,
				wrappedWriter.GetStatusCode(),
				latency.String(),
				requestID,
			)

			// Log request completion
			logger.GetLogger().LogRequestCompletion(
				*wrappedWriter.GetContext(),
				wrappedWriter.GetStatusCode(),
				httpInfo,
				systemInfo,
				authorizedInfo,
			)
		}()

		// Execute next handler
		next.ServeHTTP(wrappedWriter, r)
	}
}

// initializeRequestContext:
func initializeRequestContext(r *http.Request) *http.Request {
	// Generate request ID and set it in context
	ctx := logger.SetRequestIDContext(r.Context())

	// Initialize system information
	env := configuration.GetEnvironment()
	systemInfo := logger.NewSystemInfo(env)
	ctx = logger.SetSystemInfoContext(ctx, systemInfo)

	// Set initial authorized information
	authInfo := logger.NewInitialAuthorizedInfo()
	ctx = logger.SetAuthorizedInfoContext(ctx, authInfo)

	// Update request with updated context
	return r.WithContext(ctx)
}
