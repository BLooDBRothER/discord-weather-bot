package bot

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"strings"

	"github.com/bwmarrin/discordgo"
	
	"github.com/bloodbrother/discord-weather-bot/config"
)

func Run() {
	app_config := config.New()

	discord, err := discordgo.New(app_config.BotToken)

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
	switch {
		case strings.Contains(message.Content, "weather"):
			discord.ChannelMessageSend(message.ChannelID, "I can help with that!")
		case strings.Contains(message.Content, "bot"):
			discord.ChannelMessageSend(message.ChannelID, "Hi there! I am a weather bot")
	}
}
