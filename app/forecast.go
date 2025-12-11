package app

import (
	"WeatherForecast/search"
	"fmt"
	"net/http"
)

func forecastPost(w http.ResponseWriter, r *http.Request) {
	val, err := search.CreateValues(r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Bad parameters: %v\n", err)
		return
	}
	daily, err := search.Forecast(val)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, err, "help")
		return
	}
	fmt.Fprintln(w, daily)
}

func forecastGet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "method GET still in progress...")
}

func forecast(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		forecastGet(w, r)
	case http.MethodPost:
		forecastPost(w, r)
	default:
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "Unknown HTTP method")
	}
}
