package configuration

import "os"

// GetEnvironment: Get environment from environment variable
func GetEnvironment() string {
	if env := os.Getenv("ENV"); env != "" {
		return env
	}
	return "dev"
}
