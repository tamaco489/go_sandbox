package auth

import (
	"context"
	"net/http"

	"github.com/tamaco489/go_sandbox/slog/utils/logger"
)

type Auth struct{}

func NewAuth() *Auth {
	return &Auth{}
}

// Authorize checks if the request is authorized
func (a *Auth) Authorize(ctx context.Context, r *http.Request) (*logger.AuthorizedInfo, error) {
	// NOTE: Return authorization success for testing
	// In a real application, this would be implemented with JWT token validation and database authorization checks

	return &logger.AuthorizedInfo{
		Role:     "user",
		TenantID: "tenant123",
		MemberID: "member456", // unique member_id
	}, nil
}
