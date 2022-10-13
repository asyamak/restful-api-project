package database

import (
	"fmt"
	"log"

	"restapi/config"

	"github.com/golang-migrate/migrate"
	"github.com/jmoiron/sqlx"
	_ "github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var Table = `CREATE TABLE  users (
	id int PRIMARY KEY,
	name VARCHAR(100),
	surname VARCHAR(100),
	interests VARCHAR(250)
);`

func InitDB(c *config.Config) (*sqlx.DB, error) {
	conn := fmt.Sprintf("user=%s dbname=%s host=%s port=%s password=%s sslmode=%s", c.User, c.Dbname, c.Host, c.Port, c.Password, c.Ssl)
	// db, err := sql.Open("postgres", conn)
	// if err != nil {
	// 	log.Printf("error open sql: %v", err)
	// }
	// if err = db.Ping(); err != nil {
	// 	log.Printf("error in ping: %v", err)
	// }
	db, err := sqlx.Connect("postgres", conn)
	if err != nil {
		log.Printf("error initiatedatabase: %v", err)
	}

	m, err := migrate.New(
		"file://migrates",
		conn)
	if err != nil {
		fmt.Println("failed to make migrate", err)
		log.Printf("error")
	}
	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Printf("failed to m.Up() in migrate: %v", err)
	}
	return db, nil
}

func CreateTables(db *sqlx.DB) error {
	// tables := []string{Table}
	// for _, table := range tables {
	// 	_, err := db.Exec(table)
	// 	if err != nil {
	// 		return err
	// 	}
	// }
	db.MustExec(Table)
	return nil
}
