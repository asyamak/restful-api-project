package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func main() {
	connS := "user=postgres dbname=Asya host=localhost port=8888 password=admin sslmode=disable"
	db, err := sql.Open("postgres", connS)
	if err != nil {
		log.Printf("error open sql: %v", err)
	}
	if err = db.Ping(); err != nil {
		log.Printf("error in ping: %v", err)
	}
	// develop
}
