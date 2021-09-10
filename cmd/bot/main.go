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
	//gets the token&prefix from the file
	tokencfg, err := config.GetToken("./internal/config/token.json")
	if err != nil {
		fmt.Println("Error getting the token:", err)
		return
	}
	prefixcfg, err := config.GetPrefix("./internal/config/prefix.json")
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

	//registering all events and commands in the register folder... before opening the connection to discord
	register.RegisterEvents(s)
	register.RegisterCommands(s, prefixcfg)

	//opening up the connection
	err = s.Open()
	if err != nil {
		fmt.Println("Error opening connection:", err)
	}

	//the code for running the process, keeping it open and stopping it
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

	s.Close()

}
