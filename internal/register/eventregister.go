package register

import (
	"github.com/bwmarrin/discordgo"
	"github.com/phxenix-w/gotestbot/internal/events"
)

//registers all events in the internal/events folder
func RegisterEvents(s *discordgo.Session) {
	s.AddHandler(events.NewMessageHandler().Handler)
}
