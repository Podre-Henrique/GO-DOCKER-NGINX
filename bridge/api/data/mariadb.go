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
	password = os.Getenv("MYSQL_PASS")
	database = os.Getenv("MYSQL_DATABASE")
)

func Connect() *sql.DB {
	conn := fmt.Sprintf("root:%s@tcp(mariadb:3306)/%s", password, database)
	db, err := sql.Open("mysql", conn)
	if err != nil {
		log.Fatal(err)
	}
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}
	return db
}
