package	main

import (
	"fmt"
	
	//main-folder-of-app/..
	"github.com/project-discord-bot/bot"
	"github.com/project-discord-bot/config"
)

func main()  {
	
	err := config.ReadConfig()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	bot.Start()

	<-make(chan struct{})
	return
}