package postgres

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	_"github.com/lib/pq"
	_"github.com/joho/godotenv/autoload"

)

func ConnectPostgres() (*sql.DB, error) {
	dbInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", os.Getenv("POSTGRES_HOST"), os.Getenv("POSTGRES_PORT"), os.Getenv("POSTGRES_USER"), os.Getenv("POSTGRES_PASSWORD"), os.Getenv("POSTGRES_DB"))
	db, err := sql.Open("postgres", dbInfo)
	if err != nil {
		log.Println("Not connected to the database...", err)
		return nil, err 
	}

	err = db.Ping()
	if err != nil {
		log.Println("Error on ping...", err)
		return nil, err 
	}

	return db, nil 
}
