package auth

import (
	"context"
	"net/http"

	"github.com/tamaco489/go_sandbox/slog/utils/logger"
)

type Authorizer interface {
	Authorize(ctx context.Context, r *http.Request) (*logger.AuthorizedInfo, error)
}

// WithAuth: Authorization middleware
func WithAuth(authorizer Authorizer, next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		// Check if it's a ResponseWriterWrapper
		wrappedWriter, isWrapped := w.(*logger.ResponseWriterWrapper)

		// NOTE: Return authorization failure for testing
		// authInfo, err := authorizer.Authorize(r.Context(), r)
		// logger.GetLogger().DebugContext(r.Context(), "Authorization failed", "error", fmt.Errorf("Authorization failed"))
		// http.Error(w, fmt.Errorf("Authorization failed").Error(), http.StatusUnauthorized)
		// return

		// Execute authorization process
		authInfo, err := authorizer.Authorize(r.Context(), r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}

		// Set authorized information to context
		ctx := logger.SetAuthorizedInfoContext(r.Context(), *authInfo)

		// Update request context
		r = r.WithContext(ctx)

		// If it's a ResponseWriterWrapper, update the context
		if isWrapped {
			wrappedWriter.UpdateContext(ctx)
		}

		// Pass the updated request to the next handler
		next.ServeHTTP(w, r)
	}
}
