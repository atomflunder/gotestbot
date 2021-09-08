package admin

import (
	"github.com/phxenix-w/gotestbot/internal/inits"
	"github.com/phxenix-w/gotestbot/internal/utils"
)

type Ban struct{}

func (c *Ban) Invokes() []string {
	return []string{"ban", "b", "banned"}
}

func (c *Ban) Description() string {
	return "Ban command. Bans mentioned user and provides a reason."
}

func (c *Ban) AdminPermission() bool {
	return true
}

func (c *Ban) Exec(ctx *inits.Context) error {
	//if there are less than 2 values (user/reason) this returns immediately
	if len(ctx.Args) < 2 {
		ctx.Session.ChannelMessageSend(ctx.Message.ChannelID,
			"You must provide a user to ban and a reason!")
		return nil
	}

	//converts the mention to the id
	user_id := utils.UserMentionToID(ctx.Args[0])

	//gets the rest of the args, which has to be the reason. starts at 1 cause 0 is the user
	reason := utils.GetArgs(ctx.Args, 1)

	//actually bans the user, the last arg here is the amount of days, 0 is infinite
	err := ctx.Session.GuildBanCreateWithReason(ctx.Message.GuildID, user_id, reason, 0)

	if err != nil {
		return err
	}

	//confirmation message
	_, err = ctx.Session.ChannelMessageSend(ctx.Message.ChannelID, "Banned <@!"+user_id+"> with reason: `"+reason+"`.")

	if err != nil {
		return err
	}

	return nil
}
