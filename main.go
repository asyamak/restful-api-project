package main

import (
	"log"

	"restapi/web/users"
)

func main() {
	if err := users.Router(); err != nil {
		log.Printf("error initialise server: %v", err)
	}
}
