package main

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"time"
)

func main() {
	bot, err := discordgo.New("Bot OTE4NDU0MTE0NjgwMTkzMDM2.YbHfDA.3V0vrsmcOjgK6jOyT8KFnIeWVFE")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	bot.LogLevel = discordgo.LogInformational

	err = bot.Open()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("Bot running!")

	for {
		fmt.Println("Doing the nasty...")

		err := bot.GuildMemberNickname("770229888195493888", "533331780397301790", "Potato LauraPleaseFuckMe")
		if err != nil {
			fmt.Println(err.Error())
			continue
		}

		time.Sleep(time.Second * 10)
	}
}
