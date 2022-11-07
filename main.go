package main

import (
	"log"

	"restapi/web/users"
)

func main() {
	if err := users.Router(); err != nil {
		log.Printf("error in main: initialisation server: %v", err)
	}
}
