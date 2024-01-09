package bot

import (
	"fmt"
	"log"
	"os"
	"os/signal"

	"github.com/bwmarrin/discordgo"
	
	"github.com/bloodbrother/discord-weather-bot/config"
)

func Run() {
	discord, err := discordgo.New(config.AppConfig.BotToken)

	if err != nil {
		log.Fatal(err)
	}

	// Add event handler
	discord.AddHandler(newMessage)

	// Open session
	discord.Open()
	defer discord.Close()

	// Run until code is terminated
	fmt.Println("Bot running...")
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c
}


func newMessage(discord *discordgo.Session, message *discordgo.MessageCreate) {

	// Ignore bot messaage
	if message.Author.ID == discord.State.User.ID {
		return
	}

	// Respond to messages
	command := commandCondition{
		userCommand: message.Content,
	} 

	switch {
		case command.isGreetCommand():
			discord.ChannelMessageSend(message.ChannelID, "Hi there! I am a weather bot")
		case command.isCityWeatherCommand():
			
			city, parseErr := parseCity(command.userCommand)

			if(parseErr != nil) {
				discord.ChannelMessageSend(message.ChannelID, parseErr.Error())
			}

			currentWeather := getCurrentWeatherFromCityName(city)
			discord.ChannelMessageSendComplex(message.ChannelID, currentWeather)
	}
}
