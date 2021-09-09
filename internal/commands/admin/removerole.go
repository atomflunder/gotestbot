package admin

import (
	"strconv"

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

	user_id := utils.UserMentionToID(ctx.Args[0])

	user, err := ctx.Session.GuildMember(ctx.Message.GuildID, user_id)
	if err != nil {
		return err
	}

	role := utils.GetArgs(ctx.Args, 1)

	role_id := utils.RoleMentionToID(role)

	if _, err := strconv.Atoi(role_id); err == nil {
		err := ctx.Session.GuildMemberRoleRemove(ctx.Message.GuildID, user_id, role_id)
		if err != nil {
			return err
		}
	} else {
		role_id = utils.RoleSearch(role, ctx)
		err := ctx.Session.GuildMemberRoleRemove(ctx.Message.GuildID, user_id, role_id)
		if err != nil {
			return err
		}
	}

	final_role, err := ctx.Session.State.Role(ctx.Message.GuildID, role_id)
	if err != nil {
		return err
	}

	ctx.Session.ChannelMessageSend(ctx.Message.ChannelID,
		"Removed the "+final_role.Name+" role from "+user.Mention()+".")

	return nil
}
