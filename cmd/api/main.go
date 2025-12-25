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
	)

	r := gin.Default()
	r.RedirectTrailingSlash = true

	userRepo := repositories.NewUserRepo(db)
	jwt := jwtpkg.New("super-secret")
	productRepo := repositories.NewProductRepo(db)
	cartRepo := repositories.NewCartRepo(db)

	userService := services.NewUserService(userRepo)
	authService := services.NewAuthService(userRepo, jwt)
	productService := services.NewProductService(productRepo)
	cartService := services.NewCartService(cartRepo, productRepo)

	authHandler := handlers.NewAuthHandler(userService, authService)
	productHandler := handlers.NewProductHandler(productService)
	cartHandler := handlers.NewCartHandler(cartService)

	routes.RegisterRoutes(r, authHandler, productHandler, cartHandler, jwt, userRepo)

	//for _, route := range r.Routes() {
	//	println(route.Method, route.Path)
	//}
	//r.Use(gin.CustomRecovery(func(c *gin.Context, recovered any) {
	//	fmt.Println("🔥 PANIC:", recovered)
	//	c.AbortWithStatusJSON(500, gin.H{"error": "panic"})
	//}))

	r.Run(":8080")
}
