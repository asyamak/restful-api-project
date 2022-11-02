package user

import (
	"log"

	"restapi/db"

	"github.com/jmoiron/sqlx"
	_ "github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type UserDb struct {
	Db   *sqlx.DB
	Id   int `json:"id" db:"id"`
	Data `json:"data" db:"data"`
}

type Data struct {
	FirstName string `json:"first_name,omitempty"`
	LastName  string `json:"last_name,omitempty"`
	Interests string `json:"interests,omitempty"`
}

func NewUserDb(db *sqlx.DB) *UserDb {
	return &UserDb{
		Db: db,
	}
}

func (u *UserDb) Create(data string) error {
	query := `INSERT INTO user1 (data) VALUES($1) `
	_, err := db.DB.Exec(query, data)
	if err != nil {
		log.Printf("error insert user: %v", err)
		return err
	}
	return nil
}

func (u *UserDb) Read() (string, error) {
	query := `SELECT data FROM user1 WHERE id=$1`
	var str string
	err := db.DB.QueryRow(query, u.Id).Scan(&str)
	if err != nil {
		log.Printf("error on read user: %v ", err)
		return " ", err
	}
	return str, nil
}

func (u *UserDb) Update(data string) error {
	query := `UPDATE user1 SET data=$1 WHERE id=$2`
	_, err := db.DB.Exec(query, data, u.Id)
	if err != nil {
		log.Printf("error in update method user: %v", err)
	}
	return nil
}

func (u *UserDb) Delete() error {
	query := `DELETE FROM user1 WHERE id=$1`
	_, err := db.DB.Exec(query, u.Id)
	if err != nil {
		log.Printf("error remove user: %v", err)
		return err
	}
	return nil
}

// func (u *UserDb) UpdateInterests(name, interests string) error {
// 	query := `UPDATE users SET name=$1 WHERE interests=$2`
// 	_, err := db.DB.Exec(query, name, interests)
// 	if err != nil {
// 		log.Printf("error update user: %v", err)
// 		return err
// 	}
// 	return nil
// }
