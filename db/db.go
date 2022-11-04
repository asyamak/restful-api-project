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

const (
	port     = "5432"
	host     = "localhost"
	dbname   = "postgres"
	userdb   = "postgres"
	password = "postgres"
	ssl      = "disable"
)

var DB *sqlx.DB

func init() {
	conn := fmt.Sprintf("user=%s dbname=%s host=%s port=%s password=%s sslmode=%s", userdb, dbname, host, port, password, ssl)
	db, err := sqlx.Connect("postgres", conn)
	if err != nil {
		log.Fatalf("error initialise database: %v", err)
	}
	// if err = db.Ping(); err != nil {
	// 	db.Close()
	// 	log.Fatalf("error in ping db: %v", err)
	// }

	_, err = db.Exec(user)
	if err != nil {
		log.Fatalf("error exec table: %v", err)
	}
	// CreateTables(db)

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
	DB = db
}

// func CreateTables(db *sqlx.DB) {

// 	// if err != nil {
// 	// 	return err
// 	// }
// 	// log.Println("asd", res)
// }
