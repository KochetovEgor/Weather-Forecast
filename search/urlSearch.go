package search

import (
	"encoding/json"
	"net/http"
	"net/url"
)

type weatherResponse struct {
	Daily Daily
}

type Daily struct {
	Time    []string
	TempMin []float64 `json:"temperature_2m_min"`
	TempMax []float64 `json:"temperature_2m_max"`
}

type City struct {
	Latitude  string
	Longitude string
}

func SearchTemp(baseUrl string, city City, startDate, endDate string) (daily *Daily, err error) {
	u, err := url.Parse(baseUrl)
	if err != nil {
		return
	}
	q := u.Query()
	q.Set("latitude", city.Latitude)
	q.Set("longitude", city.Longitude)
	q.Set("daily", "temperature_2m_min")
	q.Add("daily", "temperature_2m_max")
	q.Set("start_date", startDate)
	q.Set("end_date", endDate)

	u.RawQuery = q.Encode()
	r, err := http.Get(u.String())
	if err != nil {
		return
	}
	defer r.Body.Close()
	var wR weatherResponse
	err = json.NewDecoder(r.Body).Decode(&wR)
	if err != nil {
		return
	}
	daily = &wR.Daily
	return
}
