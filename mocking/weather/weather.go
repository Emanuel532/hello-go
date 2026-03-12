package main

type WeatherProvider interface {
	fetchTemperature(string) (float64, error)
}

const hotWeather string = "t-shirt & sunglasses"
const coldWeather string = "wear a coat"

func WeatherWardrobe(city string, weather WeatherProvider) string {
	temperature, err := weather.fetchTemperature(city)
	if err != nil {
		return "error fetching weather data"
	}
	if temperature > 25 {
		return hotWeather
	} else {
		return coldWeather
	}
}
func main() {
	d := NewWttrWeatherProvider()
	println(WeatherWardrobe("Timisoara", d))
}
