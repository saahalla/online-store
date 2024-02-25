// 🚀 Fiber is an Express inspired web framework written in Go with 💖
// 📌 API Documentation: https://docs.gofiber.io
// 📝 Github Repository: https://github.com/gofiber/fiber
package main

import (
	"database/sql"
	"fmt"
	"log"
	"online-store/router"
	"os"
	"sync"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

// Database instance
var db *sql.DB
var dbMutex sync.Mutex

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

// Connect function
func Connect() error {
	host := os.Getenv("MYSQL_HOST")
	user := os.Getenv("MYSQL_USER")
	port := 3306
	password := os.Getenv("MYSQL_PASSWORD")
	dbname := os.Getenv("MYSQL_DBNAME")

	host = "localhost"
	user = "saahalla"
	port = 3306
	password = "sahal07seven"
	dbname = "online_store"

	var err error
	// Use DSN string to open
	db, err = sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%v)/%s?parseTime=true", user, password, host, port, dbname))
	if err != nil {
		return err
	}

	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)

	// Store the database connection in the global variable
	dbMutex.Lock()
	defer dbMutex.Unlock()
	db = db

	if err = db.Ping(); err != nil {
		return err
	}
	return nil
}

func Migration() error {
	err := CreateTableUsers()
	if err != nil {
		return err
	}

	return nil
}

func CreateTableUsers() error {
	schedules := `
		CREATE TABLE IF NOT EXISTS users (
			id int auto_increment not null primary key,
			username varchar(255) not null,
			password varchar(255) not null,
			email varchar(255) not null,
			phone varchar(255),
			user_role_id int not null,
			created_at timestamp not null default current_timestamp,
			created_by varchar(255),
			modified_at timestamp not null default current_timestamp,
			modified_by varchar(255),
			deleted_at timestamp not null default current_timestamp,
			deleted_by varchar(255)
		  );
	`

	_, err := db.Exec(schedules)
	if err != nil {
		return err
	}

	return nil
}

func main() {
	// Connect with database
	if err := Connect(); err != nil {
		log.Fatal(err)
	}
	// migration
	err := Migration()
	if err != nil {
		log.Fatal(err)
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
