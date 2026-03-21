package main

import (
	"fmt"
	authsvc "github.com/swlee3306/go-api-crud/auth"
	"github.com/swlee3306/go-api-crud/config"
	"github.com/swlee3306/go-api-crud/health"
	"github.com/swlee3306/go-api-crud/middleware"
	"github.com/swlee3306/go-api-crud/models"
	"github.com/swlee3306/go-api-crud/routes"
	"github.com/gin-gonic/gin"
	"log"
	"os"
	"strconv"
	"time"

)

var (
	X_buildDatetime, X_buildRevision, X_buildRevisionShort, X_buildBranch, X_buildTag string
)

func main() {
	log.Printf("go-api-crud: build info")
	log.Printf("\t buildDatetime: %s", X_buildDatetime)
	log.Printf("\t buildRevision: %s (%s)", X_buildRevisionShort, X_buildRevision)
	log.Printf("\t buildBranch: %s", X_buildBranch)
	log.Printf("\t buildTag: %s", X_buildTag)

	dbManager, err := config.NewDatabaseManager()
	if err != nil {
		log.Fatalf("failed to create database manager: %v", err)
	}
	defer dbManager.Close()

	config.DB = dbManager.GetDB()

	if err := models.AutoMigrate(config.DB); err != nil {
		log.Fatalf("failed to run auto migration: %v", err)
	}

	authService := authsvc.NewAuthService(getEnv("JWT_SECRET", ""))
	rateLimiter := middleware.NewRateLimiter(
		getEnvInt("RATE_LIMIT_REQUESTS", 100),
		time.Duration(getEnvInt("RATE_LIMIT_WINDOW_MINUTES", 1))*time.Minute,
	)

	router := gin.Default()
	router.Use(gin.Recovery())
	router.Use(middleware.IPRateLimitMiddleware(rateLimiter))

	health.SetupHealthRoutes(router, appVersion())
	routes.SetupRoutes(router, authService)

	addr := fmt.Sprintf("%s:%s", getEnv("SERVER_HOST", ""), getEnv("SERVER_PORT", "8080"))
	if addr == ":"+getEnv("SERVER_PORT", "8080") || addr == "" {
		addr = ":" + getEnv("SERVER_PORT", "8080")
	}

	log.Printf("starting server on %s", addr)
	if err := router.Run(addr); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}

func getEnv(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}

func getEnvInt(key string, fallback int) int {
	value := os.Getenv(key)
	if value == "" {
		return fallback
	}
	parsed, err := strconv.Atoi(value)
	if err != nil {
		return fallback
	}
	return parsed
}

func appVersion() string {
	if X_buildTag != "" {
		return X_buildTag
	}
	if X_buildRevisionShort != "" {
		return X_buildRevisionShort
	}
	return "dev"
}
