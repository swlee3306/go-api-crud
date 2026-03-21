//go:build legacy

package main

import (
	"os"
	"strconv"

	"github.com/swlee3306/go-api-crud/internal/sysenv"
)

func main_LoadEnvDb() error {

	// database
	{
		if val, ok := os.LookupEnv("BATON_DB_DSN"); ok && (len(val) > 0) {
			sysenv.Database.Dsn = val
		}
		if val, ok := os.LookupEnv("BATON_DATABASE_MAX_IDLE_CONNS"); ok && (len(val) > 0) {
			sysenv.Database.MaxIdleConns, _ = strconv.Atoi(val)
		}
		if val, ok := os.LookupEnv("BATON_DATABASE_MAX_LIFETIME_HOUR"); ok && (len(val) > 0) {
			sysenv.Database.MaxLifetimeHour, _ = strconv.Atoi(val)
		}
	}

	return nil
}
