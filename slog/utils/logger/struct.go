package logger

import (
	"fmt"
	"log/slog"
	"os"
)

// Logger: ログ出力機能を提供する構造体
type Logger struct {
	*slog.Logger
}

// SystemInfo: システム情報を保持する構造体
type SystemInfo struct {
	Environment string `json:"environment"`
	Service     string `json:"service"`
	Hostname    string `json:"hostname"`
}

// NewSystemInfo: SystemInfoの新しいインスタンスを作成
func NewSystemInfo(env string) SystemInfo {
	hostname, _ := os.Hostname()
	return SystemInfo{
		Environment: env,
		Service:     fmt.Sprintf("%s-slog-server", env),
		Hostname:    hostname,
	}
}

// HTTPRequestInfo: HTTPリクエスト情報を保持する構造体
type HTTPRequestInfo struct {
	Method     string `json:"method"`
	Path       string `json:"path"`
	Status     int    `json:"status"`
	Latency    string `json:"latency"`
	UserAgent  string `json:"user_agent"`
	Referer    string `json:"referer"`
	RemoteAddr string `json:"remote_addr"`
	RequestID  string `json:"request_id"`
}

// AuthorizedInfo: 認可後に得られる情報を保持する構造体
type AuthorizedInfo struct {
	Role     string `json:"role"`
	TenantID string `json:"tenant_id"`
	MemberID string `json:"member_id"`
}

// NewInitialAuthorizedInfo: AuthorizedInfoの新しいインスタンスを作成
func NewInitialAuthorizedInfo() AuthorizedInfo {
	return AuthorizedInfo{
		Role:     "anonymous",
		TenantID: "default",
		MemberID: "unknown",
	}
}
