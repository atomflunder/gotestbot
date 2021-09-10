package admin

import (
	"github.com/phxenix-w/gotestbot/internal/inits"
	"github.com/phxenix-w/gotestbot/internal/utils"
)

type Unban struct{}

func (c *Unban) Invokes() []string {
	return []string{"unban", "ub", "unbanned"}
}

func (c *Unban) Description() string {
	return "Ban command. Bans mentioned user and provides a reason."
}

func (c *Unban) AdminPermission() bool {
	return true
}

//this command is very similar to ban, just without getting any reason for the unban
func (c *Unban) Exec(ctx *inits.Context) error {
	if len(ctx.Args) < 1 {
		ctx.Session.ChannelMessageSend(ctx.Message.ChannelID,
			"You must provide a user to unban!")
		return nil
	}

	userID := utils.UserMentionToID(ctx.Args[0])

	err := ctx.Session.GuildBanDelete(ctx.Message.GuildID, userID)

	if err != nil {
		return err
	}

	_, err = ctx.Session.ChannelMessageSend(ctx.Message.ChannelID, "Unbanned <@!"+userID+">.")

	if err != nil {
		return err
	}

	return nil
}
