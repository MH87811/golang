package routes

import (
	"shop/internal/handlers"
	"shop/internal/middlewares"
	"shop/internal/repositories"
	"shop/pkg/jwtpkg"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine, authHandler *handlers.AuthHandler, productHandler *handlers.ProductHandler, jwt *jwtpkg.JWT, repo repositories.UserRepository) {
	api := r.Group("/api")

	auth := api.Group("/auth")
	{
		auth.POST("/register", authHandler.Register)
		auth.POST("/login", authHandler.Login)
		auth.POST("/refresh", authHandler.Refresh)
	}

	protected := api.Group("/protected")
	protected.Use(middlewares.AuthMiddleware(jwt, repo))
	{
		protected.GET("/profile", func(c *gin.Context) {
			user, _ := c.Get("user")
			c.JSON(200, gin.H{"user": user})
		})
		protected.POST("/product", productHandler.Create)
		protected.GET("/product", productHandler.List)
		protected.PATCH("/product/:id", productHandler.Update)
		protected.DELETE("/product/:id", productHandler.Delete)
	}
}
