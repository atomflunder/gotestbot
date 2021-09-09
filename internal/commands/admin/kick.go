package admin

import (
	"github.com/phxenix-w/gotestbot/internal/inits"
	"github.com/phxenix-w/gotestbot/internal/utils"
)

type Kick struct{}

func (c *Kick) Invokes() []string {
	return []string{"kick", "k", "kicked"}
}

func (c *Kick) Description() string {
	return "Kicks a user."
}

func (c *Kick) AdminPermission() bool {
	return true
}

//the kick command functions pretty similarly to the ban command.
func (c *Kick) Exec(ctx *inits.Context) error {
	if len(ctx.Args) < 2 {
		ctx.Session.ChannelMessageSend(ctx.Message.ChannelID,
			"You must provide a user to kick and a reason!")
		return nil
	}

	user_id := utils.UserMentionToID(ctx.Args[0])

	reason := utils.GetArgs(ctx.Args, 1)

	//the kick command. no clue why its named member delete
	err := ctx.Session.GuildMemberDeleteWithReason(ctx.Message.GuildID, user_id, reason)
	if err != nil {
		return err
	}

	_, err = ctx.Session.ChannelMessageSend(ctx.Message.ChannelID,
		"Kicked <@!"+user_id+"> from this server!")

	if err != nil {
		return err
	}

	return nil

}
