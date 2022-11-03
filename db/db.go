package db

import (
	"fmt"
	"log"

	// _ "github.com/golang-migrate/migrate"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var user = `CREATE TABLE IF NOT EXISTS users (
	id SERIAL PRIMARY KEY, 
	data VARCHAR
);`

var DB *sqlx.DB

func init() {
	DB = initDB()
}

func initDB() *sqlx.DB {
	conn := fmt.Sprintf("user=%s dbname=%s host=%s port=%s password=%s sslmode=%s", "postgres", "postgres", "localhost", "8000", "postgres", "disable")
	db, err := sqlx.Connect("postgres", conn)
	// log.Println("========", db)
	if err != nil {
		log.Printf("error initialise database: %v", err)
	}
	CreateTables(db)

	// m, err = migrate.New(
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

func CreateTables(db *sqlx.DB) {
	db.MustExec(user)
	// if err != nil {
	// 	return err
	// }
	// log.Println("asd", res)
}
