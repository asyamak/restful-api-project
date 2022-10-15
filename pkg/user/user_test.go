package user

import (
	"log"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
)

func TestCRUD(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatalf("an error occured '%s' expected when opening a stub db connction", err)
	}
	sqlxDB := sqlx.NewDb(db, "sqlmock")
	storeUser := NewUserDb(sqlxDB)
	// creating user struct to compare
	u := &User{Firstname: "A", Surname: "M", Interests: "Golang"}
	mock.ExpectExec(`INSERT INTO users`).WithArgs(u.Firstname, u.Surname, u.Interests).WillReturnResult(sqlmock.NewResult(1, 1))
	// checing create method
	err = storeUser.Create(u.Firstname, u.Surname, u.Interests)
	if err != nil {
		t.Errorf("error on create method: %v", err)
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were infulfilled expectations: %s", err)
	}
	// checking read method
	mock.ExpectQuery(`SELECT`).WithArgs(u.Surname).WillReturnRows(sqlmock.NewRows([]string{"name", "surname", "interests"}).AddRow(u.Firstname, u.Surname, u.Interests))
	_, err = storeUser.Read(u.Surname)
	if err != nil {
		t.Errorf("there were error in read: %v", err)
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were infulfilled expectations: %s", err)
	}
	// checking update method
	u1 := &User{Firstname: "Rahat", Surname: "R", Interests: "Python"}
	mock.ExpectExec(`UPDATE users SET`).WithArgs(u1.Firstname, "Rakhat").WillReturnResult(sqlmock.NewResult(0, 0))
	err = storeUser.Update(u1.Firstname, "Rakhat")
	if err != nil {
		t.Errorf("error UPDATE name with givenname: %v", err)
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were infulfilled expectations: %s", err)
	}
	// read
	mock.ExpectQuery(`SELECT`).WithArgs(u1.Surname).WillReturnRows(sqlmock.NewRows([]string{"name", "surname", "interests"}).AddRow(u1.Firstname, u1.Surname, u1.Interests))
	_, err = storeUser.Read(u1.Surname)
	if err != nil {
		t.Errorf("there were error in read: %v", err)
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were infulfilled expectations: %s", err)
	}
	// checking delete method
	u2 := &User{Firstname: "Oleg", Surname: "D", Interests: "JS"}
	mock.ExpectExec(`DELETE FROM users`).WithArgs(u2.Firstname).WillReturnResult(sqlmock.NewResult(1, 1))
	err = storeUser.Delete(u2.Firstname)
	if err != nil {
		log.Printf("error in DELETE method mock: %v", err)
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were infulfilled expectations: %s", err)
	}
	// read
	mock.ExpectQuery(`SELECT`).WithArgs(u2.Surname).WillReturnRows(sqlmock.NewRows([]string{"name", "surname", "interests"}))
	_, err = storeUser.Read(u2.Surname)
	if err != nil {
		t.Errorf("there were error in read: %v", err)
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were infulfilled expectations: %s", err)
	}

	// checking update interests
	u3 := &User{Firstname: "Ernur", Surname: "A", Interests: "Java"}
	mock.ExpectExec(`UPDATE users SET`).WithArgs(u3.Firstname, "golang").WillReturnResult(sqlmock.NewResult(1, 1))
	err = storeUser.UpdateInterests(u3.Firstname, "golang")
	if err != nil {
		t.Errorf("there were error in update interest: %v", err)
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were infulfilled expectations: %s", err)
	}
	// read
	mock.ExpectQuery(`SELECT`).WithArgs(u3.Surname).WillReturnRows(sqlmock.NewRows([]string{"name", "surname", "interests"}))
	_, err = storeUser.Read(u3.Surname)
	if err != nil {
		t.Errorf("there were error in read: %v", err)
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were infulfilled expectations: %s", err)
	}
}
