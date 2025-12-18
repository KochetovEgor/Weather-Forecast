package app

import (
	"WeatherForecast/database"
	"fmt"
	"log"
	"net/http"
)

type App struct {
	db *database.DataBase
}

func New(db *database.DataBase) *App {
	return &App{db: db}
}

func (a *App) mainPage(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./pages/index.html")
}

func (a *App) Run() error {
	err := a.db.CreateTables()
	if err != nil {
		return fmt.Errorf("error creating tables: %v", err)
	}
	log.Println("Tables Created")

	err = a.db.InitTables()
	if err != nil {
		return fmt.Errorf("error initialising tables: %v", err)
	}
	log.Println("Tables initialized")

	http.HandleFunc("/", a.mainPage)
	http.HandleFunc("/forecast", a.forecast)
	err = http.ListenAndServe("localhost:8000", nil)
	if err != nil {
		return fmt.Errorf("server error: %v", err)
	}
	return nil
}
