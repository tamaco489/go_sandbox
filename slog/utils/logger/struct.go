package logger

import (
	"fmt"
	"log/slog"
	"net/http"
	"os"
)

// AppLogger: Application logger struct
type AppLogger struct {
	*slog.Logger
}

// SystemInfo: System information
type SystemInfo struct {
	Environment string `json:"environment"`
	Service     string `json:"service"`
	Hostname    string `json:"hostname"`
}

// NewSystemInfo: Create new SystemInfo instance
func NewSystemInfo(env string) SystemInfo {
	hostname, _ := os.Hostname()
	return SystemInfo{
		Environment: env,
		Service:     fmt.Sprintf("%s-slog-server", env),
		Hostname:    hostname,
	}
}

// AuthorizedInfo: Authorized information
type AuthorizedInfo struct {
	Role     string `json:"role"`
	TenantID string `json:"tenant_id"`
	MemberID string `json:"member_id"`
}

// NewInitialAuthorizedInfo: Create new AuthorizedInfo instance
func NewInitialAuthorizedInfo() AuthorizedInfo {
	return AuthorizedInfo{
		Role:     "anonymous",
		TenantID: "default",
		MemberID: "unknown",
	}
}

// HTTPRequestInfo: HTTP request information
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

// NewHTTPRequestInfo: Create new HTTPRequestInfo instance
func NewHTTPRequestInfo(r *http.Request, statusCode int, latency string, requestID string) HTTPRequestInfo {
	return HTTPRequestInfo{
		Method:     r.Method,
		Path:       r.URL.Path,
		Status:     statusCode,
		Latency:    latency,
		UserAgent:  r.UserAgent(),
		Referer:    r.Referer(),
		RemoteAddr: r.RemoteAddr,
		RequestID:  requestID,
	}
}
