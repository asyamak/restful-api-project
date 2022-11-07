package db

import (
	"embed"
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/pressly/goose/v3"
	"github.com/spf13/viper"
)

var user = `CREATE TABLE IF NOT EXISTS users (
	id SERIAL PRIMARY KEY, 
	data VARCHAR
);`

var DB *sqlx.DB

func Configs() string {
	viper.SetConfigName("development")
	viper.AddConfigPath("./")
	viper.SetConfigType("yaml")
	err := viper.ReadInConfig()
	if err != nil {
		log.Printf("error viper get port: %v", err)
	}
	slice := fmt.Sprintf("user=%v dbname=%s host=%s port=%s password=%s sslmode=%s",
		viper.Get("services.postgres.user"),
		viper.Get("services.postgres.db_name"),
		viper.Get("services.postgres.hostname"),
		viper.Get("services.postgres.port"),
		viper.Get("services.postgres.password"),
		viper.Get("services.postgres.ssl"))
	return slice
}

var embedMigrations embed.FS

func init() {
	conn := Configs()
	db, err := sqlx.Connect("postgres", conn)
	if err != nil {
		log.Fatalf("error initialise database: %v", err)
	}
	goose.SetBaseFS(embedMigrations)
	if err := goose.SetDialect("postgres"); err != nil {
		panic(err)
	}
	if err := goose.Up(db.DB, "./migration"); err != nil {
		log.Printf("error goose up: %v", err)
		// panic(err)
	}
	// _, err = db.Exec(user)
	// if err != nil {
	// 	log.Fatalf("error exec table: %v", err)
	// }
	DB = db
}
