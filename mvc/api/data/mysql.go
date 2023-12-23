package data

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

// Database settings

var (
	user     = os.Getenv("MYSQL_USER")
	password = os.Getenv("MYSQL_PASS")
	database = os.Getenv("MYSQL_DATABASE")
)

func Connect() *sql.DB {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@/%s", user, password, database))
	if err != nil {
		log.Fatal(err)
	}
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}
	return db
}
