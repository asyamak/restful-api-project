package users

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"restapi/pkg/user"

	"github.com/gorilla/mux"
)

func PostUserHandler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		w.WriteHeader(404)
	}
	u := &user.UserDb{
		Id: id,
	}
	err = json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		log.Printf("error json decode: %v ", err)
	}
	s := u.FirstName + " " + u.LastName + " " + u.Interests
	err = u.Create(s)
	if err != nil {
		log.Printf("error in create handler: %v", err)
	}
	w.WriteHeader(http.StatusCreated)
}

func GetUserHandler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		w.WriteHeader(404)
	}
	u := &user.UserDb{Id: id}
	var new string
	new, err = u.Read()
	if err != nil {
		log.Printf("error read method: %v", err)
	}
	array := strings.Fields(new)
	d := &user.Data{
		FirstName: array[0],
		LastName:  array[1],
		Interests: array[2],
	}
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(d)
	if err != nil {
		log.Printf("error encoder json: %v", err)
	}
	w.WriteHeader(http.StatusOK)
}

func PutUserHandler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		w.WriteHeader(404)
	}
	user := &user.UserDb{
		Id: id,
		Data: user.Data{
			FirstName: r.FormValue("first_name"),
			LastName:  r.FormValue("last_name"),
			Interests: r.FormValue("interests"),
		},
	}
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		log.Printf("error in decoding json: %s", err)
	}
	s := user.FirstName + " " + user.LastName + " " + user.Interests
	fmt.Println(s)
	err = user.Update(s)
	if err != nil {
		log.Printf("error serialisation handler update method: %v", err)
	}
	w.WriteHeader(http.StatusOK)
}

func DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		w.WriteHeader(404)
	}
	new := &user.UserDb{
		Id: id,
	}
	err = new.Delete()
	if err != nil {
		log.Printf("error delete method in handler: %v", err)
	}
	w.WriteHeader(http.StatusNoContent)
}
