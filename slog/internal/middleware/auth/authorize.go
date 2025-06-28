package auth

import (
	"context"
	"errors"
	"net/http"

	"github.com/tamaco489/go_sandbox/slog/utils/logger"
)

type Auth struct{}

func NewAuth() *Auth {
	return &Auth{}
}

// Authorize checks if the request is authorized
func (a *Auth) Authorize(ctx context.Context, r *http.Request) (*logger.AuthorizedInfo, error) {
	// NOTE Implement authorization logic

	// ログの検証のため、ここでエラーが発生した場合はエラーを返す
	// return &logger.AuthorizedInfo{
	//     Role:     "user",
	//     TenantID: "tenant123",
	//     MemberID: "member456",
	// }, nil
	return nil, errors.New("not authorized")
}
