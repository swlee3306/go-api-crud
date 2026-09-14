package config

import (
	"errors"
	"os"
	"strings"
)

// ValidateRuntimeSecrets checks startup configuration before connecting to a DB.
func ValidateRuntimeSecrets() error {
	jwt := os.Getenv("JWT_SECRET")
	if len(strings.TrimSpace(jwt)) < 32 || placeholderSecret(jwt) {
		return errors.New("JWT_SECRET must be supplied as a non-placeholder secret of at least 32 bytes")
	}
	driver := getEnv("DB_DRIVER", "mysql")
	if driver == "mysql" || driver == "postgres" {
		password := os.Getenv("DB_PASSWORD")
		if strings.TrimSpace(password) == "" || placeholderSecret(password) {
			return errors.New("DB_PASSWORD must be supplied for the selected database driver")
		}
	}
	return nil
}

func placeholderSecret(value string) bool {
	value = strings.ToLower(strings.TrimSpace(value))
	return strings.HasPrefix(value, "replace_with_") || strings.HasPrefix(value, "your-") || value == "change-me" || value == "changeme" || value == "password"
}
