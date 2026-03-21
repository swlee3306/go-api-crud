package routes

import (
	authsvc "github.com/swlee3306/go-api-crud/auth"
	"github.com/swlee3306/go-api-crud/handlers"
	"github.com/swlee3306/go-api-crud/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine, authService *authsvc.AuthService) {
	handlers.SetAuthService(authService)

	authGroup := r.Group("/auth")
	{
		authGroup.POST("/login", handlers.Login)
		authGroup.POST("/register", handlers.Register)
	}

	users := r.Group("/users")
	users.Use(middleware.AuthMiddleware(authService))
	{
		users.GET("", handlers.ListUsers)
		users.POST("", handlers.CreateUser)
		users.GET("/:id", handlers.GetUser)
		users.PUT("/:id", handlers.UpdateUser)
		users.DELETE("/:id", handlers.DeleteUser)
	}

	api := r.Group("/api/v1")
	api.Use(middleware.AuthMiddleware(authService))
	{
		users := api.Group("/users")
		{
			users.GET("", handlers.ListUsers)
			users.POST("", handlers.CreateUser)
			users.GET("/:id", handlers.GetUser)
			users.PUT("/:id", handlers.UpdateUser)
			users.DELETE("/:id", handlers.DeleteUser)
		}
	}
}
