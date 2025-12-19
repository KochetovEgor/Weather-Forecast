package search

import (
	"WeatherForecast/database"
	"encoding/json"
	"errors"
	"io"
	"iter"
	"net/http"
	"net/url"
)

// URL for weather forecast
const BaseUrl = "https://api.open-meteo.com/v1/forecast"

// URL query parameters for weather API
var Params = [...]string{
	"latitude",
	"longitude",
	"daily",
	"start_date",
	"end_date"}

// Type for decode errors from weather API
type Error struct {
	IsError bool `json:"error"`
	Reason  string
}

func (err Error) Error() string {
	return err.Reason
}

// Type for API's response JSON unmarshaling
type Weather struct {
	Daily Daily
}

// Result of JSON demarshaling. Contains various data about weather
type Daily struct {
	Time    []string
	TempMin []float64 `json:"temperature_2m_min"`
	TempMax []float64 `json:"temperature_2m_max"`
}

func (d *Daily) All() iter.Seq[database.Day] {
	return func(yield func(database.Day) bool) {
		n := min(len(d.Time), len(d.TempMin), len(d.TempMax))
		for i := 0; i < n; i++ {
			if !yield(database.Day{
				Time:    d.Time[i],
				TempMin: d.TempMin[i],
				TempMax: d.TempMax[i],
			}) {
				return
			}
		}
	}
}

// Receives URL query parameters for weather API.
// Returns type Daily that contains forecast data and any error,
// occured during execution, except io.EOF
func Forecast(val url.Values) (daily *Daily, err error) {
	u, err := url.Parse(BaseUrl)
	if err != nil {
		err = errors.New("url.Parse error: " + err.Error())
		return
	}
	u.RawQuery = val.Encode()
	r, err := http.Get(u.String())
	if err != nil {
		err = errors.New("http.Get error: " + err.Error())
		return
	}
	defer r.Body.Close()
	if r.StatusCode != http.StatusOK {
		var e Error
		err = json.NewDecoder(r.Body).Decode(&e)
		if err != nil {
			err = errors.New("JSON decoding for BAD REQUEST error: " + err.Error())
			return
		}
		err = errors.New("BAD REQUEST error: " + e.Error())
		return
	}
	var W Weather
	err = json.NewDecoder(r.Body).Decode(&W)
	if err != nil && err != io.EOF {
		err = errors.New("JSON decoding error: " + err.Error())
		return
	}
	daily = &W.Daily
	err = nil
	return
}
