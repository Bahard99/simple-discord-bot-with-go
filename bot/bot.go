package bot

import (
	"fmt"
	//main-folder-of-app/config
	"github.com/project-discord-bot/config"

	"github.com/bwmarrin/discordgo"
)

var BotId string
var goBot *discordgo.Session

func Start()  {
	
	//create new bot session
	goBot, err := discordgo.New("Bot "+ config.Token)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	//making our bot a user using User function
	u, err := goBot.User("@me")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	BotId = u.ID

	//Add handler for bot to handle our message handler
	goBot.AddHandler(messageHandler)
	err = goBot.Open()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("Bot is running!!")
}

func messageHandler(s *discordgo.Session, m *discordgo.MessageCreate)  {
	
	//Bot shouldn't reply its own message
	if m.Author.ID == BotId {
		return
	}

	//message content is ping
	if m.Content == "ping" {
		_, _ = s.ChannelMessageSend(m.ChannelID, "pong")
	}

}