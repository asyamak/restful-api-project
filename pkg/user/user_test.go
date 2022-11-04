package user

import (
	"encoding/json"
	"log"
	"strings"
	"testing"

	d "restapi/db"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
)

func TestCRUD(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatalf("an error occured '%s' expected when opening db connection", err)
	}
	sqlxDB := sqlx.NewDb(db, "sqlmock")
	d.DB = sqlxDB

	// creating user struct to compare
	u := &UserDb{
		Data: Data{
			FirstName: "Mia",
			LastName:  "Sunny",
			Interests: "sun",
		},
	}
	bytes, err := json.Marshal(u.Data)
	if err != nil {
		log.Printf("error marshal:%v", err)
	}

	mock.ExpectExec(`INSERT INTO users`).WithArgs(bytes).WillReturnResult(sqlmock.NewResult(1, 1))
	// checing create method
	err = u.Create(string(bytes))
	if err != nil {
		t.Errorf("error on create method: %v", err)
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were infulfilled expectations: %s", err)
	}
}

func TestReadMethod(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatalf("an error occured '%s' expected when opening db connection", err)
	}
	sqlxDB := sqlx.NewDb(db, "sqlmock")
	d.DB = sqlxDB
	// creating user struct to compare
	u := &UserDb{
		Id: 33,
		Data: Data{
			FirstName: "Mia",
			LastName:  "Sunny",
			Interests: "sun",
		},
	}
	// temp := []string{u.Data.FirstName, u.Data.LastName, u.Data.Interests}
	// data := strings.Join(temp, " ")
	bytes, err := json.Marshal(u.Data)
	if err != nil {
		log.Printf("error marshal:%v", err)
	}
	// checking read method
	mock.ExpectQuery(`SELECT`).WithArgs(u.Id).WillReturnRows(sqlmock.NewRows([]string{"bytes"}).AddRow(bytes))
	_, err = u.Read()
	if err != nil {
		t.Errorf("there were error in read: %v", err)
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were infulfilled expectations: %s", err)
	}
}

func TestUpdateMethod(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatalf("an error occured '%s' expected when opening db connection", err)
	}
	sqlxDB := sqlx.NewDb(db, "sqlmock")
	d.DB = sqlxDB
	// creating user struct to compare
	u := &UserDb{
		Id: 33,
		Data: Data{
			FirstName: "Mia",
			LastName:  "Sunny",
			Interests: "sun",
		},
	}
	bytes, err := json.Marshal(u.Data)
	if err != nil {
		log.Printf("error marshal:%v", err)
	}
	// checking update method
	// temp := []string{u.Data.FirstName, u.Data.LastName, u.Data.Interests}
	// data := strings.Join(temp, " ")
	mock.ExpectExec(`UPDATE users SET`).WithArgs(bytes, u.Id).WillReturnResult(sqlmock.NewResult(1, 1))
	err = u.Update(string(bytes))
	if err != nil {
		t.Errorf("error UPDATE name with givenname: %v", err)
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were infulfilled expectations: %s", err)
	}
	// read
	mock.ExpectQuery(`SELECT`).WithArgs(u.Id).WillReturnRows(sqlmock.NewRows([]string{"data"}).AddRow(bytes))
	_, err = u.Read()
	if err != nil {
		t.Errorf("there were error in read: %v", err)
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were infulfilled expectations: %s", err)
	}
}

func TestDeleteMethod(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatalf("an error occured '%s' expected when opening db connection", err)
	}
	sqlxDB := sqlx.NewDb(db, "sqlmock")
	d.DB = sqlxDB
	// creating user struct to compare
	u := &UserDb{
		Id: 33,
		Data: Data{
			FirstName: "Mia",
			LastName:  "Sunny",
			Interests: "sun",
		},
	}
	// checking delete method§
	mock.ExpectExec(`DELETE FROM users`).WithArgs(u.Id).WillReturnResult(sqlmock.NewResult(1, 1))
	err = u.Delete()
	if err != nil {
		log.Printf("error in DELETE method mock: %v", err)
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were infulfilled expectations: %s", err)
	}
	// read
	temp := []string{u.Data.FirstName, u.Data.LastName, u.Data.Interests}
	data := strings.Join(temp, " ")
	mock.ExpectQuery(`SELECT`).WithArgs(u.Id).WillReturnRows(sqlmock.NewRows([]string{"data"}).AddRow(data))
	_, err = u.Read()
	if err != nil {
		t.Errorf("there were error in read: %v", err)
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were infulfilled expectations: %s", err)
	}
}
