package info

import (
	"fmt"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/phxenix-w/gotestbot/internal/inits"
	"github.com/phxenix-w/gotestbot/internal/utils"
)

type Roleinfo struct{}

func (c *Roleinfo) Invokes() []string {
	return []string{"roleinfo", "role", "rinfo"}
}

func (c *Roleinfo) Description() string {
	return "Gets you an embed with user info about a Role."
}

func (c *Roleinfo) AdminPermission() bool {
	return false
}

func (c *Roleinfo) Exec(ctx *inits.Context) error {
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

	//getting the role created at date:
	roleCreatedAt, err := discordgo.SnowflakeTimestamp(role.ID)
	if err != nil {
		return err
	}

	//getting role members
	roleMembers, err := utils.GetRoleMembers(role, ctx)
	if err != nil {
		return err
	}

	embed := &discordgo.MessageEmbed{
		Title:     "Roleinfo of " + role.Name,
		Color:     role.Color,
		Timestamp: time.Now().Format(time.RFC3339),

		Fields: []*discordgo.MessageEmbedField{
			{
				Name:   "Name:",
				Value:  role.Mention(),
				Inline: true,
			},
			{
				Name:   "Users with role:",
				Value:  fmt.Sprint(len(roleMembers)),
				Inline: true,
			},
			{
				Name:   "Created at:",
				Value:  utils.GetDiscordTimeStamp(&roleCreatedAt, "F"),
				Inline: true,
			},
			{
				Name:   "Mentionable:",
				Value:  fmt.Sprint(role.Mentionable),
				Inline: true,
			},
			{
				Name:   "Displayed separately:",
				Value:  fmt.Sprint(role.Hoist),
				Inline: true,
			},
			{
				Name: "Color:",
				//converts from int to hex, cause thats what discord uses
				Value:  fmt.Sprintf("#%X", role.Color),
				Inline: true,
			},
		},
	}

	_, err = ctx.Session.ChannelMessageSendEmbed(ctx.Message.ChannelID, embed)
	if err != nil {
		return err
	}

	return nil
}
