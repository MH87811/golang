package main

import (
	"github.com/gin-gonic/gin"
	"shop/internal/handlers"
	"shop/internal/repositories"
	"shop/internal/routes"
	"shop/internal/services"
	"shop/pkg/jwtpkg"
)

func main() {
	r := gin.Default()

	repo := repositories.NewInMemoryUserRepo()
	jwt := jwtpkg.New("super-secret")

	userService := services.NewUserService(repo)
	authService := services.NewAuthService(repo, jwt)

	authHandler := handlers.NewAuthHandler(userService, authService)

	routes.RegisterRoutes(r, authHandler, jwt, repo)

	r.Run(":8080")
}
