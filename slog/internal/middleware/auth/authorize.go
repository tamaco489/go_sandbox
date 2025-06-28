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
	// NOTE: テスト用に認可成功を返す
	// 実際の実装では、JWTトークンの検証やデータベースでの認可チェックを行う

	return &logger.AuthorizedInfo{
		Role:     "user",
		TenantID: "tenant123",
		MemberID: "member456", // 固有のmember_id
	}, nil
}
