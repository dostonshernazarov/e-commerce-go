package main

import (
	"comment-service/storage/postgres"
	"log"
)

func main() {
	db, err := postgres.ConnectPostgres()

	if err != nil {
		log.Println("Not connectedddd... ", err)
		return
	}
	log.Println("Connected to Postgres...", db)
}
