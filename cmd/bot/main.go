package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/phxenix-w/gotestbot/internal/commands"
	"github.com/phxenix-w/gotestbot/internal/config"
	"github.com/phxenix-w/gotestbot/internal/events"
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
	dg, err := discordgo.New("Bot " + tokencfg.Token)
	if err != nil {
		fmt.Println("Error creating Discord session:", err)
		return
	}

	//just get every intent
	dg.Identify.Intents = discordgo.IntentsAll

	err = dg.Open()
	if err != nil {
		fmt.Println("Error opening connection:", err)
	}

	//registering all events below
	registerEvents(dg)
	registerCommands(dg, prefixcfg)

	//log in message for the console, User.String() gets you the classic username#discriminator
	fmt.Println("Bot is now running! Logged in as", dg.State.User.String(), "\nPress CTRL+C to exit.")

	//the code for running the process, keeping it open and stopping it
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, syscall.SIGTERM)
	<-sc

	dg.Close()

}

//registers all events in the internal/events folder
func registerEvents(dg *discordgo.Session) {
	dg.AddHandler(events.NewMessageHandler().Handler)
}

//registers the commands in the internal/commands folder
func registerCommands(s *discordgo.Session, prefix *config.PrefixConfig) {
	cmdHandler := commands.NewCommandHandler(prefix.Prefix)
	//generic error message telling you why the command failed
	cmdHandler.OnError = func(err error, ctx *commands.Context) {
		ctx.Session.ChannelMessageSend(ctx.Message.ChannelID, fmt.Sprintf("Command Execution failed! \nReason:`%s`", err.Error()))
	}

	cmdHandler.RegisterCommand(&commands.CmdPing{})
	cmdHandler.RegisterCommand(&commands.CmdPong{})

	cmdHandler.RegisterMiddleware(&commands.MwPermissions{})

	s.AddHandler(cmdHandler.HandleMessage)
}
