package sheets

import (
	"WeatherForecast/database"
	"WeatherForecast/search"
	"html/template"
	"io"
)

var indexTMPL = template.Must(template.ParseFiles("./sheets/templates/index.html"))

var forecastTMPL = template.Must(template.ParseFiles("./sheets/templates/forecast.html"))

var addCityTMPL = template.Must(template.ParseFiles("./sheets/templates/add_city.html"))

// Writes html template "index.html" to wr
func GetMainPage(wr io.Writer, citiesNames []string) error {
	err := indexTMPL.Execute(wr, citiesNames)
	return err
}

// Writes html template "forecast.html" to wr
func GetForecastPage(wr io.Writer, d *search.Daily) error {
	err := forecastTMPL.Execute(wr, d)
	return err
}

// Writes html template "add_city.html" to wr
func GetAddCityPage(wr io.Writer, cities []database.City) error {
	err := addCityTMPL.Execute(wr, cities)
	return err
}
