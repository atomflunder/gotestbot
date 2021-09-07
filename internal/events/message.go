package events

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

type MessageHandler struct{}

func NewMessageHandler() *MessageHandler {
	return &MessageHandler{}
}

//the on message event. logs all sent messages to the console for now
func (h *MessageHandler) Handler(s *discordgo.Session, e *discordgo.MessageCreate) {
	c, err := s.Channel(e.ChannelID)
	if err != nil {
		fmt.Println("Error getting the channel:", err)
		return
	}

	fmt.Printf("%s said in %s: %s\n", e.Author, c.Name, e.Content)
}
