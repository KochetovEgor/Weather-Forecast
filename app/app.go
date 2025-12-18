package app

import (
	"WeatherForecast/database"
	"fmt"
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
		return fmt.Errorf("Error creating tables: %v", err)
	}

	http.HandleFunc("/", a.mainPage)
	http.HandleFunc("/forecast", a.forecast)
	err = http.ListenAndServe("localhost:8000", nil)
	if err != nil {
		return fmt.Errorf("Server error: %v", err)
	}
	return nil
}
