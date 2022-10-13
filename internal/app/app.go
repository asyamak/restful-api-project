package app

import (
	"log"

	"restapi/config"
	"restapi/internal/repo"
	"restapi/pkg/database"
)

type App struct {
	Config *config.Config
}

func NewApp(c *config.Config) *App {
	return &App{
		Config: c,
	}
}

func (a *App) Run() {
	db, err := database.InitDB(a.Config)
	if err != nil {
		log.Panicf("error initiate databse: %v", err)
	}
	if err := database.CreateTables(db); err != nil {
		log.Printf("error create tables: %v", err)
	}
	user := repo.NewUser(db)

	user.Create("Asya", "M", "music")
	user.Create("Rahat", "R", "cooking")
	user.Create("Oleg", "D", "running")
	user.Create("Meruert", "T", "cola")
	user.Delete("Oleg")
	user.UpdateInterests("Asya", "golang")
	// user.Read("T")
	user.Update("R", "Rakhat")
}
