package sheets

import (
	"WeatherForecast/search"
	"html/template"
	"io"
)

var forecastTMPL = template.Must(template.ParseFiles("./sheets/templates/forecast.html"))

// Writes html template "forecast.html" to wr
func GetForecastPage(wr io.Writer, d *search.Daily) error {
	err := forecastTMPL.Execute(wr, d)
	return err
}
