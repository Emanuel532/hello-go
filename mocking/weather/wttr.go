package main

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"
)

type CurrentCondition struct {
	FeelsLikeC string `json:"FeelsLikeC"`
	FeelsLikeF string `json:"FeelsLikeF"`
}

type WeatherData struct {
	CurrentCondition []CurrentCondition `json:"current_condition"`
}

type WttrWeatherProvider struct {
	query_params string
	base_url     string
}

func NewWttrWeatherProvider() *WttrWeatherProvider {
	w := WttrWeatherProvider{}
	w.base_url = "http://wttr.in/"
	w.query_params = "?format=j2"
	return &w
}

func (w WttrWeatherProvider) fetchTemperature(city string) (float64, error) {
	request_url := w.base_url + city + w.query_params
	println(request_url)
	resp, err := http.Get(request_url)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return 0, err
	}
	var payload WeatherData
	err = json.Unmarshal(body, &payload)
	if err != nil {
		return 0, err
	}
	return strconv.ParseFloat(payload.CurrentCondition[0].FeelsLikeC, 1)
}
