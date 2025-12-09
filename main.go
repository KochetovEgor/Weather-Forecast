package main

import (
	"WeatherForecast/search"
	"fmt"
	"net/http"
)

var baseUrl = "https://api.open-meteo.com/v1/forecast"

func mainPage(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./pages/index.html")
}

func forecastPost(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	var city search.City
	city.Latitude = r.Form.Get("latitude")
	city.Longitude = r.Form.Get("longitude")
	startDate := r.Form.Get("start_date")
	endDate := r.Form.Get("end_date")
	daily, _ := search.SearchTemp(baseUrl, city, startDate, endDate)
	fmt.Fprintln(w, daily)
}

func main() {
	http.HandleFunc("/", mainPage)
	http.HandleFunc("/forecast", forecastPost)
	http.ListenAndServe("localhost:8000", nil)
}
