package router

import (
	"log"
	"online-store/database"
	"online-store/modules/products"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

// SetupRoutes setup router api
func SetupRoutes(app *fiber.App) {

	// database
	db, err := database.GetDBConnection()
	if err != nil {
		log.Fatal(err)
	}

	// Middleware
	api := app.Group("/", logger.New())
	// api.Get("/", handler.Hello)

	// Auth
	// auth := api.Group("/auth")
	// auth.Post("/login", handler.Login)

	// User
	// user := api.Group("/user")
	// user.Get("/:id", handler.GetUser)
	// user.Post("/", handler.CreateUser)
	// user.Patch("/:id", middleware.Protected(), handler.UpdateUser)
	// user.Delete("/:id", middleware.Protected(), handler.DeleteUser)

	// Product
	product := api.Group("/product")
	productServices := products.NewService(db)

	product.Get("/", products.List(productServices))
	product.Get("/:id", products.Get(productServices))
	product.Post("/", products.Add(productServices))

	// product.Get("/", handler.GetAllProducts)
	// product.Get("/:id", handler.GetProduct)
	// product.Post("/", middleware.Protected(), handler.CreateProduct)
	// product.Delete("/:id", middleware.Protected(), handler.DeleteProduct)
}
