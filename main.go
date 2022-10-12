package main

import (
	"database/sql"
	"log"
)

func main() {
	connS := "user=postgres dbname=postgrestest host=localhost port=5432 password=qwer1234"
	db, err := sql.Open("postgres", connS)
	if err != nil {
		log.Printf("error open sql: %v", err)
	}
	if err = db.Ping(); err != nil {
		log.Printf("error in ping: %v", err)
	}
	// develop
}
