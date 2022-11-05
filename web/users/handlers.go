package users

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

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
		log.Printf("error occured:post handler:json decode: %v ", err)
	}
	// BEFORE
	// temp := []string{u.Data.FirstName, u.Data.LastName, u.Data.Interests}
	// data := strings.Join(temp, " ")

	// AFTER
	data := fmt.Sprintf("%s %s %s", u.Data.FirstName, u.Data.LastName, u.Data.Interests)

	err = u.Create(data)
	if err != nil {
		log.Printf("error occured:post handler:create method: %v", err)
	}
	w.WriteHeader(http.StatusCreated)
}

func GetUserHandler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		w.WriteHeader(404)
	}
	u := &user.UserDb{Id: id}

	new, err := u.Read()
	if err != nil {
		log.Printf("error occured:get handler:read method: %v", err)
	}
	// array := strings.Fields(new)
	// d := &user.Data{
	// 	FirstName: array[0],
	// 	LastName:  array[1],
	// 	Interests: array[2],
	// }
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(new)
	if err != nil {
		log.Printf("error occured:get handler: json encoder: %v", err)
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
		log.Printf("error occured:put handler:json decode: %s", err)
	}
	// fmt.Println(s)

	// BEFORE
	// s := user.FirstName + " " + user.LastName + " " + user.Interests

	// AFTER
	data := fmt.Sprintf("%s %s %s", user.Data.FirstName, user.Data.LastName, user.Data.Interests)

	err = user.Update(data)
	if err != nil {
		log.Printf("error occured:put handler:update method: %v", err)
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
		log.Printf("error occured:delete handler: delete method: %v", err)
	}
	w.WriteHeader(http.StatusOK)
}
