package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/phxenix-w/gotestbot/internal/config"
	"github.com/phxenix-w/gotestbot/internal/register"
)

func main() {
	//location of config files
	const tokenFile = "./internal/config/token.json"
	const prefixFile = "./internal/config/prefix.json"

	//gets the token from the file
	tokencfg, err := config.GetToken(tokenFile)
	if err != nil {
		fmt.Println("Error getting the token:", err)
		return
	}

	prefixcfg, err := config.GetPrefix(prefixFile)
	if err != nil {
		fmt.Println("Error getting the prefix:", err)
		return
	}

	//creates a new bot instance with the token above
	s, err := discordgo.New("Bot " + tokencfg.Token)
	if err != nil {
		fmt.Println("Error creating Discord session:", err)
		return
	}

	//just get every intent
	s.Identify.Intents = discordgo.IntentsAll

	err = s.Open()
	if err != nil {
		fmt.Println("Error opening connection:", err)
	}

	//registering all events and commands in the register folder
	register.RegisterEvents(s)
	register.RegisterCommands(s, prefixcfg)

	//log in message for the console, User.String() gets you the classic username#discriminator
	fmt.Println("Bot is now running! Logged in as", s.State.User.String(), "\nPress CTRL+C to exit.")

	//the code for running the process, keeping it open and stopping it
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

	s.Close()

}
