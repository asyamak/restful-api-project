package db

import (
	"fmt"
	"log"

	// _ "github.com/golang-migrate/migrate"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
)

var user = `CREATE TABLE IF NOT EXISTS users (
	id SERIAL PRIMARY KEY, 
	data VARCHAR
);`
var DB *sqlx.DB

func Configs() string {
	viper.SetConfigName("config")
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

func init() {
	conn := Configs()
	db, err := sqlx.Connect("postgres", conn)
	if err != nil {
		log.Fatalf("error initialise database: %v", err)
	}
	_, err = db.Exec(user)
	if err != nil {
		log.Fatalf("error exec table: %v", err)
	}
	DB = db
}
