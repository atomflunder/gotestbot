package admin

import (
	"github.com/phxenix-w/gotestbot/internal/inits"
	"github.com/phxenix-w/gotestbot/internal/utils"
)

type Rename struct{}

func (c *Rename) Invokes() []string {
	return []string{"rename", "name", "rn"}
}

func (c *Rename) Description() string {
	return "Renames a given user."
}

func (c *Rename) AdminPermission() bool {
	return true
}

func (c *Rename) Exec(ctx *inits.Context) error {
	//an empty nickname is allowed, this will just reset their name then
	if len(ctx.Args) < 1 {
		ctx.Session.ChannelMessageSend(ctx.Message.ChannelID,
			"You must provide a user to rename.")
		return nil
	}

	user_id := utils.UserMentionToID(ctx.Args[0])

	new_name := utils.GetArgs(ctx.Args, 1)

	//sets their new nickname, as mentioned above if new_name is empty this will just reset it
	err := ctx.Session.GuildMemberNickname(ctx.Message.GuildID, user_id, new_name)

	if err != nil {
		return err
	}

	//a different confirmation message for if the name got reset and changed, its nicer that way
	if len(new_name) > 0 {
		_, err = ctx.Session.ChannelMessageSend(ctx.Message.ChannelID,
			"Renamed <@!"+user_id+"> to: "+new_name)
	} else {
		_, err = ctx.Session.ChannelMessageSend(ctx.Message.ChannelID,
			"Reset the nickname of <@!"+user_id+">!")
	}

	if err != nil {
		return err
	}

	return nil
}
