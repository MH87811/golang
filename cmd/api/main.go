package main

import (
	"shop/internal/config"
	"shop/internal/handlers"
	"shop/internal/models"
	"shop/internal/repositories"
	"shop/internal/routes"
	"shop/internal/services"
	"shop/pkg/jwtpkg"

	"github.com/gin-gonic/gin"
)

func main() {

	dsn := "host=localhost user=postgres password=1234 dbname=shop port=5432 sslmode=disable"
	db, err := config.ConnectDB(dsn)
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(
		&models.User{},
		&models.Product{},
		&models.Cart{},
		&models.CartItems{},
		&models.Order{},
		&models.OrderItems{},
	)

	r := gin.Default()
	r.RedirectTrailingSlash = true

	userRepo := repositories.NewUserRepo(db)
	jwt := jwtpkg.New("super-secret")
	productRepo := repositories.NewProductRepo(db)
	cartRepo := repositories.NewCartRepo(db)
	orderRepo := repositories.NewOrderRepo(db)

	userService := services.NewUserService(userRepo)
	authService := services.NewAuthService(userRepo, jwt)
	productService := services.NewProductService(productRepo)
	cartService := services.NewCartService(cartRepo, productRepo)
	orderService := services.NewOrderService(db, orderRepo, cartRepo)

	authHandler := handlers.NewAuthHandler(userService, authService)
	productHandler := handlers.NewProductHandler(productService)
	cartHandler := handlers.NewCartHandler(cartService)
	orderHandler := handlers.NewOrderHandler(orderService)

	routes.RegisterRoutes(r, authHandler, productHandler, cartHandler, orderHandler, jwt, userRepo)

	r.Run(":8080")
}
