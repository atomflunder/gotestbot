package admin

import (
	"strconv"

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
	user_id := utils.UserMentionToID(ctx.Args[0])

	//this gets the member, just makes sure its valid
	user, err := ctx.Session.GuildMember(ctx.Message.GuildID, user_id)
	if err != nil {
		return err
	}

	//gets the rest of the args, which has to be the role searched for. starts at 1 cause 0 is the user
	role := utils.GetArgs(ctx.Args, 1)

	role_id := utils.RoleMentionToID(role)

	//this converts the input to an int, if its true it will try to use the role ID directly.
	//the "oversight" here is that if a role has a name only with ints and you search for it, it will not work
	//need to fix this sometime, but for now this is the best i came up with
	if _, err := strconv.Atoi(role_id); err == nil {
		err := ctx.Session.GuildMemberRoleAdd(ctx.Message.GuildID, user_id, role_id)
		if err != nil {
			return err
		}
		//if this is not true, it searches for the closest matching role, and then gives it out
	} else {
		role_id = utils.RoleSearch(role, ctx)
		err := ctx.Session.GuildMemberRoleAdd(ctx.Message.GuildID, user_id, role_id)
		if err != nil {
			return err
		}
	}

	//this gets the role, to use the role name in the message later
	final_role, err := ctx.Session.State.Role(ctx.Message.GuildID, role_id)
	if err != nil {
		return err
	}

	//final confirmation message
	ctx.Session.ChannelMessageSend(ctx.Message.ChannelID,
		"Added the "+final_role.Name+" role to "+user.Mention()+".")

	return nil
}
