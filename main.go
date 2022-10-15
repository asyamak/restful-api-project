package main

import (
	"restapi/db"
	"restapi/pkg/user"
)

func main() {
	DB := db.InitDB()
	db.CreateTables(DB)
	u := user.NewUserDb(DB)
	// u.Delete("Murzagazina")
	// u.Create("A", "B", "C")
	u.Create("L", "H", "iG")
	u.Create("S", "Q", "C")
	// u.Create("Asya", "M", "golang")
	// _, err := u.Read("M")
	// if err != nil {
	// 	log.Printf("error read value from the db: %v", err)
	// }
	// fmt.Println(a)
	// u.Update("A", "Asya")
	// u.Delete("O")
	// u.UpdateInterests("J", "H")
	// u.Delete("Asya")
}
