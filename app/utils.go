package app

import (
	"WeatherForecast/search"
	"fmt"
	"net/http"
	"net/url"
)

// Reads forms data from http.Request (r.Form.Get) and returns
// variable of type url.Values that contains URL query parameters for weather API.
// If an error occurs during form parsing, the result of the function will be nil
func (a *App) createValues(r *http.Request) (url.Values, error) {
	val := url.Values{}
	err := r.ParseForm()
	if err != nil {
		return nil, err
	}
	for _, p := range search.Params {
		v := r.PostForm.Get(p)
		if v != "" {
			val.Add(p, v)
		}
	}
	val.Add("daily", "temperature_2m_min")
	val.Add("daily", "temperature_2m_max")
	name := r.PostForm.Get("city")
	city, err := a.db.GetCityByName(name)
	if err == nil {
		val.Add("latitude", fmt.Sprint(city.Latitude))
		val.Add("longitude", fmt.Sprint(city.Longitude))
	}
	return val, nil
}
