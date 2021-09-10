package events

import (
	"fmt"
	"time"

	"github.com/bwmarrin/discordgo"
)

type ReadyHandler struct{}

func NewReadyHandler() *ReadyHandler {
	return &ReadyHandler{}
}

//the on ready event, fires when the bot logs in
func (h *ReadyHandler) Handler(s *discordgo.Session, e *discordgo.Ready) {
	//log in message for the console, User.String() gets you the classic username#discriminator
	fmt.Println("Bot is now running! Logged in as", s.State.User.String(), "\nPress CTRL+C to exit.")

	//updates the status every 5 minutes
	statuses := []string{
		"Cycling statuses",
		"is pretty cool",
		"right?",
	}
	n := 0
	for {
		err := s.UpdateGameStatus(0, statuses[n])
		if err != nil {
			fmt.Println("Error updating the status:", err)
		}
		fmt.Println("Set status to: ", statuses[n])
		time.Sleep(5 * time.Minute)
		n += 1
		//loops back around
		if n >= len(statuses) {
			n = 0
		}
	}
}
