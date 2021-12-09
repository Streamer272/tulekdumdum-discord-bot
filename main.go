package main

import (
	"flag"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"time"
)

func main() {
	token := flag.String("token", "", "Bot token")
	nickname := flag.String("nickname", "LauraFuckMePls", "Tulek's new nickname")
	sleepInterval := flag.Int("sleep", 10, "Sleep interval in seconds")
	flag.Parse()

	if *token == "" {
		fmt.Errorf("Please provide a token\n")
		return
	}

	if len(*nickname) > 25 {
		fmt.Errorf("Nickname is too long\n")
		return
	}

	bot, err := discordgo.New("Bot " + *token)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	bot.LogLevel = discordgo.LogWarning

	err = bot.Open()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("Bot running!")

	for {
		fmt.Println("Doing the nasty...")

		err := bot.GuildMemberNickname("770229888195493888", "533331780397301790", "Potato "+*nickname)
		if err != nil {
			fmt.Println(err.Error())
			continue
		}

		time.Sleep(time.Second * time.Duration(*sleepInterval))
	}
}
