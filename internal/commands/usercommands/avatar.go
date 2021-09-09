package usercommands

import (
	"github.com/phxenix-w/gotestbot/internal/inits"
	"github.com/phxenix-w/gotestbot/internal/utils"
)

type Avatar struct{}

func (c *Avatar) Invokes() []string {
	return []string{"avatar", "ava", "icon"}
}

func (c *Avatar) Description() string {
	return "Returns you the avatar of a user or yourself."
}

func (c *Avatar) AdminPermission() bool {
	return false
}

func (c *Avatar) Exec(ctx *inits.Context) error {
	//first we get the user_id. if the author does not mention any user, it will use them
	user_id := ""
	if len(ctx.Args) < 1 {
		user_id = ctx.Message.Author.ID
	} else {
		user_id = utils.UserMentionToID(ctx.Args[0])
	}

	//getting the member
	member, err := ctx.Session.GuildMember(ctx.Message.GuildID, user_id)
	if err != nil {
		return err
	}

	//sending the avatar url, discord will preview it automatically as an image
	_, err = ctx.Session.ChannelMessageSend(ctx.Message.ChannelID,
		member.User.AvatarURL(""))

	if err != nil {
		return err
	}

	return nil
}
