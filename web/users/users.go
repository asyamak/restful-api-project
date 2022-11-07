package users

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func Router() error {
	r := mux.NewRouter()
	r.HandleFunc("/user/{id}", PostUserHandler).Methods("POST")
	r.HandleFunc("/user/{id}", GetUserHandler).Methods("GET")
	r.HandleFunc("/user/{id}", DeleteUserHandler).Methods("DELETE")
	r.HandleFunc("/user/{id}", PutUserHandler).Methods("PUT")

	Port := 4040
	log.Printf("server is on; port: %v\n", Port)
	return http.ListenAndServe(fmt.Sprintf(":%v", Port), r)
}
