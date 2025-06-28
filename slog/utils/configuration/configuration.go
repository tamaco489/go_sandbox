package configuration

import "os"

// GetEnvironment: 環境変数から環境を取得
func GetEnvironment() string {
	if env := os.Getenv("ENV"); env != "" {
		return env
	}
	return "dev"
}
