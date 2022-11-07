package main

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

var (
	port, host, dbname, userdb, password, ssl interface{}
	config                                    map[string]string
)

func Config() {
	// var err error
	viper.AddConfigPath("./docker-compose.yml")
	err := viper.ReadInConfig()
	if err != nil {
		log.Printf("error viper get port: %v", err)
	}

	// viper.SetConfigName("config")
	// viper.SetConfigType("yaml")
	// port = viper.Get("postgres.ports")
	// ssl = viper.GetString("postgres.environment")
	config = viper.GetStringMapString("postgres.environment")
	fmt.Println(port, host, dbname, userdb, password, ssl)
	fmt.Println(config)
	for k, v := range config {
		fmt.Printf("key:%v, value:%v\n", k, v)
	}
}
