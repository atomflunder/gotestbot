package info

import (
	"fmt"
	"strings"

	"github.com/phxenix-w/gotestbot/internal/inits"
	"github.com/phxenix-w/gotestbot/internal/utils"
)

type Listrole struct{}

func (c *Listrole) Invokes() []string {
	return []string{"listrole", "listroles", "lrole"}
}

func (c *Listrole) Description() string {
	return "Gets you every user with a certain role."
}

func (c *Listrole) AdminPermission() bool {
	return false
}

func (c *Listrole) Exec(ctx *inits.Context) error {
	//getting the args and role as usual
	if len(ctx.Args) < 1 {
		ctx.Session.ChannelMessageSend(ctx.Message.ChannelID,
			"You must provide a role!")
		return nil
	}

	r := utils.GetArgs(ctx.Args, 0)

	roleID := utils.RoleMentionToID(r)

	role, err := utils.RoleFromInput(roleID, ctx)
	if err != nil {
		return err
	}

	roleMembers, err := utils.GetRoleMembers(role, ctx)
	if err != nil {
		return err
	}

	//we only want the user name from this list so we have to make a new list here
	var userList []string
	for m := range roleMembers {
		userList = append(userList, roleMembers[m].User.String())
	}

	//check if the user list is not too large, chose 60 as a good cutoff. primary goal is not to send more than 2000 chars in one message
	var joinedList string
	if len(userList) > 60 {
		joinedList = "Too many users to list!"
	} else {
		joinedList = strings.Join(userList, ", ")
	}

	_, err = ctx.Session.ChannelMessageSend(ctx.Message.ChannelID,
		"Users with the "+role.Name+" role ("+fmt.Sprint(len(roleMembers))+"):\n`"+joinedList+"`")

	if err != nil {
		return err
	}

	return nil
}
