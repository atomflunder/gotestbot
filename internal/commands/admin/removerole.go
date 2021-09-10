package admin

import (
	"github.com/phxenix-w/gotestbot/internal/inits"
	"github.com/phxenix-w/gotestbot/internal/utils"
)

type Removerole struct{}

func (c *Removerole) Invokes() []string {
	return []string{"removerole", "rr", "remover"}
}

func (c *Removerole) Description() string {
	return "Removes a role from a user."
}

func (c *Removerole) AdminPermission() bool {
	return true
}

//this is basically the same command as addrole.go. most of this is just copied
//please go there for a detailed explanation for everything in the comments.
func (c *Removerole) Exec(ctx *inits.Context) error {
	if len(ctx.Args) < 2 {
		ctx.Session.ChannelMessageSend(ctx.Message.ChannelID,
			"You must provide a user and a role!")
		return nil
	}

	userID := utils.UserMentionToID(ctx.Args[0])

	user, err := ctx.Session.GuildMember(ctx.Message.GuildID, userID)
	if err != nil {
		return err
	}

	r := utils.GetArgs(ctx.Args, 1)

	roleID := utils.RoleMentionToID(r)

	//gets the role from the utils function
	role, err := utils.ReturnRoleFromInput(roleID, ctx)
	if err != nil {
		return err
	}

	err = ctx.Session.GuildMemberRoleRemove(ctx.Message.GuildID, userID, role.ID)
	if err != nil {
		return err
	}

	ctx.Session.ChannelMessageSend(ctx.Message.ChannelID,
		"Removed the "+role.Name+" role from "+user.Mention()+".")

	return nil
}
