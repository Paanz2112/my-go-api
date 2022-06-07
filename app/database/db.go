package database

import (
	"database/sql"
	"fmt"
	"log"
	"workspace/config"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func Connect() (*sql.DB, error) {
	var err error

	// Connection string
	connStr := fmt.Sprintf("postgresql://%s:%s@%s:%s/postgres?sslmode=disable", config.Config("DB_USER"), config.Config("DB_PASS"), config.Config("DB_HOST"), config.Config("DB_PORT"))
	// Connect to database
	DB, err := sql.Open("postgres", connStr)
	//Test DB instance
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	if err = DB.Ping(); err != nil {
		log.Fatal(err)
		return nil, err
	}
	log.Println("database initialize.")
	return DB, nil
}
