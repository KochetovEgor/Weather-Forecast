package search

import (
	"net/http"
	"net/url"
)

var Params = [...]string{
	"latitude",
	"longitude",
	"daily",
	"start_date",
	"end_date"}

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
