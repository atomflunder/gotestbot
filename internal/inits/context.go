package inits

import "github.com/bwmarrin/discordgo"

//passes in the relevant context information for our commands
type Context struct {
	Session *discordgo.Session
	Message *discordgo.Message
	Args    []string
	Handler *CommandHandler
}
