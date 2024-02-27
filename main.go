// 🚀 Fiber is an Express inspired web framework written in Go with 💖
// 📌 API Documentation: https://docs.gofiber.io
// 📝 Github Repository: https://github.com/gofiber/fiber
package main

import (
	"log"
	"online-store/database"
	"online-store/router"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

// var (
// 	cachecc = cache.New(5*time.Minute, 10*time.Minute) // Waktu kadaluarsa 5 menit, pembersihan cache setiap 10 menit
// )

// Database settings

func init() {

	// err := godotenv.Load()
	// if err != nil {
	// 	log.Fatalf("Some error occured. Err: %s", err)
	// }

}

func main() {

	db, err := database.GetDBConnection()
	if err != nil {
		log.Fatal(err)
	}

	// migration
	err = database.Migration(db)
	if err != nil {
		log.Println(err)
	}

	// Create a Fiber app
	app := fiber.New()
	app.Use(cors.New())
	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.Send([]byte("Welcome to the online-store mongo book shop!"))
	})

	router.SetupRoutes(app)

	app.Get("/register", func(c *fiber.Ctx) error {

		// Return in JSON format
		return nil
	})

	log.Fatal(app.Listen(":3030"))
}
