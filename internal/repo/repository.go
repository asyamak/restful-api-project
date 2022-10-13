package repo

import (
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/jmoiron/sqlx"
)

type Users struct {
	id                       int
	name, surname, interests string
}
type User struct {
	db *sqlx.DB
}

func NewUser(db *sqlx.DB) *User {
	return &User{
		db: db,
	}
}

func (u *User) Create(name, surname, interests string) {
	query := `INSERT INTO users (name,surname,interests) VALUES($1,$2,$3)`
	_, err := u.db.Exec(query, name, surname, interests)
	if err != nil {
		log.Printf("error insert user: %v", err)
	}
}

func (u *User) Read(surname string) {
	query := `SELECT id,name,surname  FROM user WHERE surname=$1`
	persons := Users{}
	if err := u.db.Select(&persons, query, surname); err != nil {
		log.Printf("error on read user: %v ", err)
	}
	fmt.Println(persons)
	// if err != nil {
	// 	log.Printf("error on read user ")
	// }
	// fmt.Scanf(row, surname)
}

func (u *User) Update(name, toname string) {
	query := `UPDATE users SET name=$1 WHERE surname=$2`
	_, err := u.db.Exec(query, toname, name)
	if err != nil {
		log.Printf("error update user: %v", err)
	}
}

func (u *User) UpdateInterests(name, interests string) {
	query := `UPDATE users SET name=$1 WHERE surname=$2`
	_, err := u.db.Exec(query, interests, name)
	if err != nil {
		log.Printf("error update user: %v", err)
	}
}

func (u *User) Delete(name string) {
	query := `DELETE FROM users WHERE name=$1`
	_, err := u.db.Exec(query, name)
	if err != nil {
		log.Printf("error remove user: %v", err)
	}
}
