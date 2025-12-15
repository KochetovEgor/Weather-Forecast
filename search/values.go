package search

import (
	"net/http"
	"net/url"
)

// URL query parameters for weather API
var Params = [...]string{
	"latitude",
	"longitude",
	"daily",
	"start_date",
	"end_date"}

// Reads forms data from http.Request (r.Form.Get) and returns
// variable of type url.Values that contains URL query parameters for weather API.
// If an error occurs during form parsing, the result of the function will be nil
func CreateValues(r *http.Request) (url.Values, error) {
	val := url.Values{}
	err := r.ParseForm()
	if err != nil {
		return nil, err
	}
	for _, p := range Params {
		v := r.Form.Get(p)
		if v != "" {
			val.Add(p, v)
		}
	}
	val.Add("daily", "temperature_2m_min")
	val.Add("daily", "temperature_2m_max")
	return val, nil
}
