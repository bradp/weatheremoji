package main

import (
	"github.com/shawntoffel/darksky"
	"log"
)

func main(key string, lat float, long float) {
	client := darksky.New(key)
	request := darksky.ForecastRequest{}

	request.Latitude = lat
	request.Longitude = long

	request.Options = darksky.ForecastRequestOptions{Exclude: "hourly,minutely"}

	forecast, err := client.Forecast(request)
	if err != nil {
		log.Fatal("Forecast request error", err)
	}

	switch forecast.Currently.Icon {
	case "clear-day":
		return "☀️"
	case "partly-cloudy-day":
	case "partly-cloudy-night":
		return "🌤️"
	case "cloudy":
		return "☁️"
	case "rain":
		return "🌧️"
	case "snow":
	case "sleet":
		return "❄️"
	case "wind":
		return "🌬️"
	case "fog":
		return "🌫️"
	default:
		return "☀️"
	}
}
