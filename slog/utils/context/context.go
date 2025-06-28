package context

import "context"

// contextKey: 汎用的なコンテキストキー
type contextKey string

const (
	// 汎用的なコンテキストキー
	requestIDKey contextKey = "request_id"
	userIDKey    contextKey = "user_id"
	tenantIDKey  contextKey = "tenant_id"
)

// WithRequestID: コンテキストにリクエストIDを設定
func WithRequestID(ctx context.Context, requestID string) context.Context {
	return context.WithValue(ctx, requestIDKey, requestID)
}

// GetRequestID: コンテキストからリクエストIDを取得
func GetRequestID(ctx context.Context) (string, bool) {
	requestID, ok := ctx.Value(requestIDKey).(string)
	return requestID, ok
}

// WithUserID: コンテキストにユーザーIDを設定
func WithUserID(ctx context.Context, userID string) context.Context {
	return context.WithValue(ctx, userIDKey, userID)
}

// GetUserID: コンテキストからユーザーIDを取得
func GetUserID(ctx context.Context) (string, bool) {
	userID, ok := ctx.Value(userIDKey).(string)
	return userID, ok
}

// WithTenantID: コンテキストにテナントIDを設定
func WithTenantID(ctx context.Context, tenantID string) context.Context {
	return context.WithValue(ctx, tenantIDKey, tenantID)
}

// GetTenantID: コンテキストからテナントIDを取得
func GetTenantID(ctx context.Context) (string, bool) {
	tenantID, ok := ctx.Value(tenantIDKey).(string)
	return tenantID, ok
}
