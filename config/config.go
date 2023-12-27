package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

type Config struct {
	BotToken string
	WeatherToken string
}

func New() (config Config) {
	err := godotenv.Load()

	if(err != nil) {
		log.Fatal(err)
	}

	return Config {
		BotToken: "Bot " + os.Getenv("DISCORD_BOT_TOKEN"),
		WeatherToken: os.Getenv("WEATHER_API_KEY"),
	}
}
