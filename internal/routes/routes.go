package routes

import (
	"net/http"
	"shop/internal/config"
	"shop/internal/handlers"
	"shop/internal/middlewares"
	"shop/internal/repositories"
	"shop/internal/services"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	api := r.Group("api/")

	api.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "ok",
		})
	})

	userRepo := repositories.NewInMemoryUserRepo()
	userService := services.NewUserService(userRepo)
	api.POST("user/register", func(c *gin.Context) {
		handlers.RegisterUser(c, userService)
	})

	authService := services.NewAuthService(userRepo, config.LoadConfig())

	api.POST("/auth/login", func(c *gin.Context) {
		handlers.Login(c, authService)
	})

	protected := api.Group("/protected")
	protected.Use(middlewares.AuthMiddleware(config.LoadConfig()))

	protected.GET("/profile", func(c *gin.Context) {
		userClaims, _ := c.Get("user")
		c.JSON(http.StatusOK, gin.H{
			"user": userClaims,
		})
	})

	api.POST("auth/refresh", func(c *gin.Context) {
		handlers.Refresh(c, authService)
	})
}
