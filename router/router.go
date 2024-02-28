package router

import (
	"log"
	"online-store/common/middleware"
	"online-store/database"
	"online-store/modules/auth"
	carts "online-store/modules/cart"
	"online-store/modules/categories"
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
	authApi := api.Group("/auth")
	authServices := auth.NewService(db)

	authApi.Post("/register", auth.HandlerRegister(authServices))
	authApi.Post("/login", auth.HandlerLogin(authServices))

	// Product
	productApi := api.Group("/product")
	productServices := products.NewService(db)

	productApi.Get("/", products.HandlerList(productServices))
	productApi.Get("/:id", products.HandlerGet(productServices))
	productApi.Post("/", middleware.Protected(), products.HandlerAdd(productServices))
	productApi.Delete("/:id", middleware.Protected(), products.HandlerDelete(productServices))
	productApi.Put("/:id", middleware.Protected(), products.HandlerUpdate(productServices))

	// Category
	categoryApi := api.Group("/category")
	categoryServices := categories.NewService(db)

	categoryApi.Get("/", categories.HandlerList(categoryServices))
	categoryApi.Get("/:id", categories.HandlerGet(categoryServices))
	categoryApi.Post("/", middleware.Protected(), categories.HandlerAdd(categoryServices))
	categoryApi.Delete("/:id", middleware.Protected(), categories.HandlerDelete(categoryServices))
	categoryApi.Put("/:id", middleware.Protected(), categories.HandlerUpdate(categoryServices))

	cartApi := api.Group("/cart")
	cartServices := carts.NewService(db)

	cartApi.Get("/:id", middleware.Protected(), carts.HandlerGet(cartServices))
	cartApi.Put("/:id", middleware.Protected(), carts.HandlerUpdate(cartServices))
	cartApi.Post("/", middleware.Protected(), carts.HandlerAdd(cartServices))
	// product.Get("/", handler.GetAllProducts)
	// product.Get("/:id", handler.GetProduct)
	// product.Post("/", middleware.Protected(), handler.CreateProduct)
	// product.Delete("/:id", middleware.Protected(), handler.DeleteProduct)
}
