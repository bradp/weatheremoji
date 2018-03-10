package WeatherEmoji

import (
	"github.com/adlio/darksky"
	"log"
)

// WeatherEmoji does nothing
func WeatherEmoji() {
}

// EmojiRequest requests an emoji for a location
func EmojiRequest(key string, lat float64, long float64) string {
	client := darksky.NewClient(key)
	forecast, err := client.GetForecast(lat, long, darksky.Defaults)
	if err != nil {
		log.Fatal("Forecast request error", err)
	}

	Emoji := ""

	switch forecast.Currently.Icon {
	case "clear-day":
		Emoji = "☀️"
	case "partly-cloudy-day":
	case "partly-cloudy-night":
		Emoji = "🌤️"
	case "cloudy":
		Emoji = "☁️"
	case "rain":
		Emoji = "🌧️"
	case "snow":
	case "sleet":
		Emoji = "❄️"
	case "wind":
		Emoji = "🌬️"
	case "fog":
		Emoji = "🌫️"
	default:
		Emoji = "☀️"
	}

	return Emoji
}
