package main

import "testing"

type DummyWeatherProvider struct {
	temperature float64
}

func (d *DummyWeatherProvider) setTemperature(temp float64) {
	d.temperature = temp
}

func (d DummyWeatherProvider) fetchTemperature(string) (float64, error) {
	return d.temperature, nil
}

func TestWeatherWardrobe(t *testing.T) {
	t.Run("test high temp weather", func(t *testing.T) {
		dummyWeatherProvider := DummyWeatherProvider{}
		dummyWeatherProvider.setTemperature(29.4)

		got := WeatherWardrobe("Timisoara", dummyWeatherProvider)
		want := hotWeather

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
	t.Run("test cold weather", func(t *testing.T) {
		dummyWeatherProvider := DummyWeatherProvider{}
		dummyWeatherProvider.setTemperature(9.7)

		got := WeatherWardrobe("Arad", dummyWeatherProvider)
		want := coldWeather

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
}
