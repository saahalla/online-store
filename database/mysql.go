package database

import (
	"fmt"
	"os"
	"sync"

	"github.com/jmoiron/sqlx"
)

// Database instance
var DB *sqlx.DB
var dbMutex sync.Mutex

func GetDBConnection() (*sqlx.DB, error) {
	if DB == nil {
		err := Connect()
		if err != nil {
			return DB, err
		}
	}

	if err := DB.Ping(); err != nil {
		err := Connect()
		if err != nil {
			return DB, err
		}
	}

	return DB, nil
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
	DB, err = sqlx.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%v)/%s?parseTime=true", user, password, host, port, dbname))
	if err != nil {
		return err
	}

	DB.SetMaxIdleConns(10)
	DB.SetMaxOpenConns(100)

	// Store the database connection in the global variable
	dbMutex.Lock()
	defer dbMutex.Unlock()

	if err = DB.Ping(); err != nil {
		return err
	}

	return nil
}
