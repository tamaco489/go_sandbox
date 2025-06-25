package middleware

import (
	"context"
	"errors"
	"net/http"
)

type Auth struct{}

// Authorize checks if the request is authorized
func (a *Auth) Authorize(ctx context.Context, r *http.Request) error {

	// NOTE Implement authorization logic

	// ログの検証のため、ここでエラーが発生した場合はエラーを返す
	// return nil
	return errors.New("not authorized")
}
