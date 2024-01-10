package config

import (
	"github.com/joho/godotenv"
	"os"
)

type Config struct {
	BotToken string
	WeatherToken string
}

func new() (config Config) {
	godotenv.Load()

	return Config {
		BotToken: "Bot " + os.Getenv("DISCORD_BOT_TOKEN"),
		WeatherToken: os.Getenv("WEATHER_API_KEY"),
	}
}


var AppConfig Config = new()
