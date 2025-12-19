package app

import (
	"WeatherForecast/search"
	"WeatherForecast/sheets"
	"fmt"
	"net/http"
)

func (a *App) forecastPost(w http.ResponseWriter, r *http.Request) {
	val, err := a.createValues(r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Bad parameters: %v\n", err)
		return
	}
	daily, err := search.Forecast(val)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, err)
		return
	}
	err = sheets.GetForecastPage(w, daily)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Invalid response from weather API: %v\n", err)
		return
	}
}

func (a *App) forecastGet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "method GET still in progress...")
}

func (a *App) forecast(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		a.forecastGet(w, r)
	case http.MethodPost:
		a.forecastPost(w, r)
	default:
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "Unknown HTTP method")
	}
}

func (a *App) addCity(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		city, err := a.parseCity(r)
		if err != errEmptyCity {
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				fmt.Fprintf(w, "bad request: %v\n", err)
				return
			}
			err = a.db.AddCity(city)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				fmt.Fprintf(w, "Can't add city to DB: %v\n", err)
				return
			}
		}
	}

	cities, err := a.db.GetCities()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Can't get cities from DB: %v\n", err)
		return
	}
	err = sheets.GetAddCityPage(w, cities)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Can't execute template: %v\n", err)
		return
	}
}
