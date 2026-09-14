package config

import (
	"strings"
	"testing"
)

func TestDatabasePasswordHasNoDefault(t *testing.T) {
	t.Setenv("DB_PASSWORD", "")
	if LoadDatabaseConfig().Password != "" {
		t.Fatal("database password must not have a built-in default")
	}
}

func TestRuntimeSecrets(t *testing.T) {
	for _, tc := range []struct {
		name, driver, jwt, password string
		valid                       bool
	}{
		{"missing jwt", "sqlite", "", "", false},
		{"short jwt", "sqlite", "short", "", false},
		{"placeholder jwt", "sqlite", "REPLACE_WITH_A_UNIQUE_LOCAL_SECRET", "", false},
		{"sqlite", "sqlite", strings.Repeat("z", 32), "", true},
		{"mysql missing password", "mysql", strings.Repeat("z", 32), "", false},
		{"postgres placeholder", "postgres", strings.Repeat("z", 32), "REPLACE_WITH_YOUR_DATABASE_PASSWORD", false},
		{"weak password", "mysql", strings.Repeat("z", 32), "password", false},
		{"mysql supplied", "mysql", strings.Repeat("z", 32), "SYNTHETIC-test-only", true},
	} {
		t.Run(tc.name, func(t *testing.T) {
			t.Setenv("DB_DRIVER", tc.driver)
			t.Setenv("JWT_SECRET", tc.jwt)
			t.Setenv("DB_PASSWORD", tc.password)
			err := ValidateRuntimeSecrets()
			if (err == nil) != tc.valid {
				t.Fatalf("unexpected validation outcome; want valid=%v", tc.valid)
			}
			if err != nil && ((tc.jwt != "" && strings.Contains(err.Error(), tc.jwt)) || (tc.password != "" && strings.Contains(err.Error(), tc.password))) {
				t.Fatal("error exposed a value")
			}
		})
	}
}
