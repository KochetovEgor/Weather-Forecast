package sheets

import (
	"WeatherForecast/search"
	"html/template"
	"io"
)

var indexTMPL = template.Must(template.ParseFiles("./sheets/templates/index.html"))

var forecastTMPL = template.Must(template.ParseFiles("./sheets/templates/forecast.html"))

func GetMainPage(wr io.Writer, citiesNames []string) error {
	err := indexTMPL.Execute(wr, citiesNames)
	return err
}

// Writes html template "forecast.html" to wr
func GetForecastPage(wr io.Writer, d *search.Daily) error {
	err := forecastTMPL.Execute(wr, d)
	return err
}
