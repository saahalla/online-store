package database

import (
	"fmt"
	"online-store/config"
	"sync"

	_ "github.com/go-sql-driver/mysql"
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
	host := config.GetConfig("MYSQL_HOST")
	user := config.GetConfig("MYSQL_USER")
	port := config.GetConfig("MYSQL_PORT")
	password := config.GetConfig("MYSQL_PASSWORD")
	dbname := config.GetConfig("MYSQL_DBNAME")

	var err error
	// Use DSN string to open
	DB, err = sqlx.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%v)/%s?multiStatements=true&parseTime=true&charset=utf8;", user, password, host, port, dbname))
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
