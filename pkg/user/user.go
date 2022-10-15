package user

import (
	"log"

	//"restapi/config"

	//"github.com/jmoiron/sqlx"
	"github.com/jmoiron/sqlx"
	//_ "github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"

	//"honnef.co/go/tools/config"

	//"github.com/jmoiron/sqlx"
	_ "github.com/jmoiron/sqlx"
)

type User struct {
	Firstname string `json:"name" db:"name"`
	Surname   string `json:"surname" db:"surname"`
	Interests string `json:"interests" db:"interests" `
}

type UserDb struct {
	db *sqlx.DB
}

func NewUserDb(db *sqlx.DB) *UserDb {
	return &UserDb{
		db: db,
	}
}

func (u *UserDb) Create(name, surname, interests string) error {
	query := `INSERT INTO users (name,surname,interests) VALUES($1,$2,$3)`
	_, err := u.db.Exec(query, name, surname, interests)
	if err != nil {
		log.Printf("error insert user: %v", err)
		return err
	}
	return nil
}

func (u *UserDb) Read(surname string) (*User, error) {
	query := `SELECT * FROM users WHERE surname=:surname`
	person := &User{}
	row, err := u.db.NamedQuery(query, map[string]interface{}{"surname": surname})
	if err != nil {
		log.Printf("error on read user: %v ", err)
		return nil, err
	}
	for row.Next() {
		err := row.StructScan(&person)
		if err != nil {
			return nil, err
		}
	}
	return person, nil
}

func (u *UserDb) Update(name, toname string) error {
	query := `UPDATE users SET name=$1 WHERE name=$2`
	_, err := u.db.Exec(query, name, toname)
	if err != nil {
		log.Printf("error in update method user: %v", err)
	}
	return nil
}

func (u *UserDb) UpdateInterests(name, interests string) error {
	query := `UPDATE users SET name=$1 WHERE interests=$2`
	_, err := u.db.Exec(query, name, interests)
	if err != nil {
		log.Printf("error update user: %v", err)
		return err
	}
	return nil
}

func (u *UserDb) Delete(name string) error {
	query := `DELETE FROM users WHERE name=$1`
	_, err := u.db.Exec(query, name)
	if err != nil {
		log.Printf("error remove user: %v", err)
		return err
	}
	return nil
}
