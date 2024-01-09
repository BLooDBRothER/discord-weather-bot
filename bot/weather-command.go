package bot

import (
	"encoding/json"
	"fmt"
	"errors"
	"io"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"

	"github.com/bloodbrother/discord-weather-bot/config"
)

var BASE_URL = "https://api.openweathermap.org/data/2.5/weather"

type WeatherData struct {
	Weather []struct {
		Main string `json:"main"`
		Description string `json:"description"`
		Icon string `json:"icon"`
	}

	Main struct {
		Temp     float64 `json:"temp"`
		Humidity int     `json:"humidity"`
	} `json:"main"`

	Wind struct {
		Speed float64 `json:"speed"`
	} `json:"wind"`

	Name string `json:"name"`
}

type WeatherError struct {
	Message string `json:"message"`
}


func createWeatherEmbed(weatherData *WeatherData) ([]*discordgo.MessageEmbed) {
	title := fmt.Sprintf("Current Weather in %s", weatherData.Name)
	description := fmt.Sprintf("**%s** - %s", weatherData.Weather[0].Main, weatherData.Weather[0].Description)
	temperature := strconv.FormatFloat(weatherData.Main.Temp, 'f', 2, 64)


	return []*discordgo.MessageEmbed {{
		Type: discordgo.EmbedTypeRich,
		Title: title,
		Fields: []*discordgo.MessageEmbedField{
			{
				Name: "Condiition",
				Value: description,
				Inline: false,
			},
			{
				Name: "Temperature 🌡️",
				Value: temperature + " ℃",
				Inline: true,
			},
		},
	}}
}

func getCurrentWeatherFromCityName (city string) *discordgo.MessageSend {
	weatherUrl := fmt.Sprintf("%s?q=%s&units=metric&appid=%s", BASE_URL, city, config.AppConfig.WeatherToken)

	client := http.Client {
		Timeout: 5 * time.Second,
	}

	response, err := client.Get(weatherUrl)

	if(err != nil) {
		return &discordgo.MessageSend{
			Content: "Error while getting the weather please try again",
		}
	}

	body, _ := io.ReadAll(response.Body)
	defer response.Body.Close()

	if(response.StatusCode != 200) {
		var weatherError WeatherError
		json.Unmarshal(body, &weatherError)
		return &discordgo.MessageSend{
			Content: weatherError.Message,
		}
	}

	var weatherData WeatherData
	json.Unmarshal(body, &weatherData)

	embed := &discordgo.MessageSend{
		Embeds: createWeatherEmbed(&weatherData),
	}

	return embed
}

func parseCity (command string) (string, error) {
	processedString := strings.Split(command, " ")

	if(len(processedString) < 2 && processedString[0] != "!city") {
		return "", errors.New("Invalid command")
	}

	return strings.Join(processedString[1:], " "), nil
}
