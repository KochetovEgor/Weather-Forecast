package search

import (
	"encoding/json"
	"net/http"
	"net/url"
)

const baseUrl = "https://api.open-meteo.com/v1/forecast"

type Error struct {
	IsError bool `json:"error"`
	Reason  string
}

func (err Error) Error() string {
	return err.Reason
}

type Weather struct {
	Daily Daily
}

type Daily struct {
	Time    []string
	TempMin []float64 `json:"temperature_2m_min"`
	TempMax []float64 `json:"temperature_2m_max"`
}

func Forecast(val url.Values) (daily *Daily, err error) {
	u, err := url.Parse(baseUrl)
	if err != nil {
		return
	}
	u.RawQuery = val.Encode()
	r, err := http.Get(u.String())
	if err != nil {
		return
	}
	defer r.Body.Close()
	if r.StatusCode != http.StatusOK {
		var e Error
		err = json.NewDecoder(r.Body).Decode(&e)
		if err != nil {
			return
		}
		err = e
		return
	}
	var W Weather
	err = json.NewDecoder(r.Body).Decode(&W)
	if err != nil {
		return
	}
	daily = &W.Daily
	return
}
