package main

import (
	"log"

	"restapi/config"
	"restapi/internal/app"
)

func main() {
	config, err := config.NewConfig()
	if err != nil {
		log.Printf("error parsing config: %v", err)
	}
	app.NewApp(config).Run()
}
