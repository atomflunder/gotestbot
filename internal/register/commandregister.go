package register

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
	"github.com/phxenix-w/gotestbot/internal/commands/admin"
	"github.com/phxenix-w/gotestbot/internal/commands/usercommands"
	"github.com/phxenix-w/gotestbot/internal/config"
	"github.com/phxenix-w/gotestbot/internal/inits"
)

//registers the commands in the internal/commands folder. remember to register every command here
func RegisterCommands(s *discordgo.Session, prefix *config.PrefixConfig) {
	cmdHandler := inits.NewCommandHandler(prefix.Prefix)
	//generic error message telling you why the command failed
	cmdHandler.OnError = func(err error, ctx *inits.Context) {
		ctx.Session.ChannelMessageSend(ctx.Message.ChannelID, fmt.Sprintf("Command Execution failed! \nReason:`%s`", err.Error()))
	}

	//first:
	//we register all commands, first up: usercommands
	cmdHandler.RegisterCommand(&usercommands.Ping{})
	cmdHandler.RegisterCommand(&usercommands.Pong{})

	//here are the admin commands
	cmdHandler.RegisterCommand(&admin.Ban{})
	cmdHandler.RegisterCommand(&admin.Clear{})

	//and here:
	//all of our permissions
	cmdHandler.RegisterMiddleware(&inits.MwPermissions{})

	//finally:
	//all of our listeners
	s.AddHandler(cmdHandler.HandleMessage)
}
