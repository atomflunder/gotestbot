package admin

import (
	"github.com/phxenix-w/gotestbot/internal/inits"
	"github.com/phxenix-w/gotestbot/internal/utils"
)

type Addrole struct{}

func (c *Addrole) Invokes() []string {
	return []string{"addrole", "ar", "addr"}
}

func (c *Addrole) Description() string {
	return "Adds a role to a user."
}

func (c *Addrole) AdminPermission() bool {
	return true
}

func (c *Addrole) Exec(ctx *inits.Context) error {
	//if there are less than 2 values (user/role) this returns immediately
	if len(ctx.Args) < 2 {
		ctx.Session.ChannelMessageSend(ctx.Message.ChannelID,
			"You must provide a user and a role!")
		return nil
	}

	//converts the mention to the id
	userID := utils.UserMentionToID(ctx.Args[0])

	//this gets the member, just makes sure its valid
	user, err := ctx.Session.GuildMember(ctx.Message.GuildID, userID)
	if err != nil {
		return err
	}

	//gets the rest of the args, which has to be the role searched for. starts at 1 cause 0 is the user
	r := utils.GetArgs(ctx.Args, 1)

	roleID := utils.RoleMentionToID(r)

	//gets the role from the utils function
	role, err := utils.RoleFromInput(roleID, ctx)
	if err != nil {
		return err
	}

	err = ctx.Session.GuildMemberRoleAdd(ctx.Message.GuildID, userID, role.ID)
	if err != nil {
		return err
	}

	//final confirmation message
	ctx.Session.ChannelMessageSend(ctx.Message.ChannelID,
		"Added the "+role.Name+" role to "+user.Mention()+".")

	return nil
}
