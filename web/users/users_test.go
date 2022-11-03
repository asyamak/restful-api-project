package users

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"testing"

	"restapi/pkg/user"

	_ "github.com/DATA-DOG/go-sqlmock"
)

func TestCreateUsers(t *testing.T) {
	u := user.UserDb{
		Id: 33,
		Data: user.Data{
			FirstName: "Ola",
			LastName:  "Sun",
			Interests: "golang",
		},
	}
	obj, err := json.Marshal(&u)
	if err != nil {
		log.Printf("error in create test:in marshalling: %v", err)
	}
	resp, err := http.Post("http://localhost:4040/user/33", "application/json", bytes.NewBuffer(obj))
	if err != nil {
		log.Printf("error in create test: http post: %v", err)
	}
	if resp.StatusCode != http.StatusCreated {
		t.Errorf("error in create test:not the same status codes:%v, %d", err, resp.StatusCode)
	}
}

func TestRead(t *testing.T) {
	u := struct {
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
		Interests string `json:"interests"`
	}{}
	resp, err := http.Get("http://localhost:4040/user/33")
	if err != nil {
		log.Printf("error in read test: http get: %v", err)
		// return nil
	}
	bb, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Errorf("error in read test:ioutil:%v", err)
	}
	err = json.Unmarshal(bb, &u)
	if err != nil {
		t.Error(err)
	}
	if u.FirstName != "Ola" {
		t.Error("error in read test: not the same data")
	}
	if u.LastName != "Sun" {
		t.Error("error in read test: not the same data")
	}
	if u.Interests != "golang" {
		t.Error("error in read test: not the same data")
	}
}

func TestUpdate(t *testing.T) {
	u := &user.UserDb{
		Id: 33,
		Data: user.Data{
			FirstName: "Ola",
			LastName:  "Sun",
			Interests: "golang",
		},
	}
	b, err := json.Marshal(u)
	if err != nil {
		t.Errorf("error in update test:marshalling: %v", err)
	}
	req, err := http.NewRequest("PUT", (fmt.Sprintf("http://localhost:4040/user/%d", u.Id)), bytes.NewBuffer(b))
	if err != nil {
		t.Errorf("error in update test: http new request: %v", err)
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Errorf("error in update test: DefaultClient: %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Errorf("error not the same status codes in delete unit test:%v, %d", err, resp.StatusCode)
	}
}

func TestDelete(t *testing.T) {
	u := &user.UserDb{
		Id: 9,
	}
	req, err := http.NewRequest("DELETE", (fmt.Sprintf("http://localhost:4040/user/%d", u.Id)), nil)
	if err != nil {
		log.Printf("error in delete test: http request: %v", err)
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Errorf("error in delete test: default client: %v", err)
	}
	if resp.StatusCode != http.StatusNoContent {
		t.Errorf("error not the same status codes in delete unit test:%v, %d", err, resp.StatusCode)
	}
}
