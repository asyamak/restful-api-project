package db

import (
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
)

var Table = `CREATE TABLE IF NOT EXISTS users (
	name VARCHAR(100),
	surname VARCHAR(100),
	interests VARCHAR(250)
);`

var user = `CREATE TABLE users (
	id INT PRIMARY KEY, 
	data VARCHAR
);`

func InitDB() *sqlx.DB {
	conn := fmt.Sprintf("user=%s dbname=%s host=%s port=%s password=%s sslmode=%s", "postgres", "Asyaa", "localhost", "8888", "admin", "disable")
	db, err := sqlx.Connect("postgres", conn)
	if err != nil {
		log.Printf("error initiatedatabase: %v", err)
	}

	// _, err = migrate.New(
	// 	"file://./db/migration",
	// 	conn)
	// if err != nil {
	// 	fmt.Println("failed to make migrate", err)
	// 	log.Printf("error migrate new: %v", err)
	// }
	// if err := m.Up(); err != nil && err != migrate.ErrNoChange {
	// 	log.Printf("failed to m.Up() in migrate: %v", err)
	// }
	return db
}

func CreateTables(db *sqlx.DB) error {
	db.MustExec(Table)
	return nil
}
