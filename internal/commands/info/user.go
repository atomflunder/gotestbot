package info

import (
	"fmt"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/phxenix-w/gotestbot/internal/inits"
	"github.com/phxenix-w/gotestbot/internal/utils"
)

type Userinfo struct{}

func (c *Userinfo) Invokes() []string {
	return []string{"userinfo", "user", "uinfo"}
}

func (c *Userinfo) Description() string {
	return "Gets you an embed with user info about a mentioned user or yourself."
}

func (c *Userinfo) AdminPermission() bool {
	return false
}

func (c *Userinfo) Exec(ctx *inits.Context) error {
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

	//prepping some info for the embed
	user_color := ctx.Session.State.UserColor(user_id, ctx.Message.ChannelID)

	//here we get the timestamps
	user_created_at, err := discordgo.SnowflakeTimestamp(user_id)
	if err != nil {
		return err
	}
	//getting unix timestamp for the join date, need to convert that first sadly
	user_joined_at, err := time.Parse(time.RFC3339, string(member.JoinedAt))
	if err != nil {
		return err
	}

	//getting the presence for the status/activity
	user_presence, err := ctx.Session.State.Presence(ctx.Message.GuildID, user_id)
	if err != nil {
		return err
	}
	//defining the activity in case the user didnt set one
	user_activity := "None"
	if len(user_presence.Activities) > 0 {
		user_activity = user_presence.Activities[0].Name
	}

	//getting the member roles
	user_roles := member.Roles
	//this whole block is for searching the top role of the member.
	//im sure this can be done better but i certainly couldnt come up with something better at 1am
	//first we get every role ID and their position
	var role_positions []int
	var role_ids []string
	for x := range member.Roles {
		role, err := ctx.Session.State.Role(ctx.Message.GuildID, user_roles[x])
		if err != nil {
			return err
		}
		role_positions = append(role_positions, role.Position)
		role_ids = append(role_ids, role.ID)
	}
	//then we search for the highest position value
	max_role := role_positions[0]
	for _, value := range role_positions {
		if value > max_role {
			max_role = value
		}
	}
	//then we search the index of the highest position role in the original slice
	n := 0
	for x := range role_positions {
		if max_role == role_positions[x] {
			break
		}
		n += 1
	}
	//and then we use that index to get the according role ID in the other slice and finally get the role that is highest
	top_role, err := ctx.Session.State.Role(ctx.Message.GuildID, role_ids[n])
	if err != nil {
		return err
	}

	//this here gets all guild members. this caps at 1000 guild members so if the bot gets invited to a server with more than this
	//we would need to rewrite this using pagination or something. right now its not *that* important
	guild_members, err := ctx.Session.GuildMembers(ctx.Message.GuildID, "", 1000)
	if err != nil {
		return err
	}

	//thankfully guild members are already ordered by join date (from my testing at least),
	//so we just need to search for our user ID and then we already got the rank. needs to start at 1.
	rank := 1
	for x := range guild_members {
		if guild_members[x].User.ID == member.User.ID {
			break
		}
		rank += 1
	}

	//the actual embed itself
	embed := &discordgo.MessageEmbed{
		//first the basic stuff
		Title: "Userinfo of " + member.User.String(),
		Color: user_color,
		Thumbnail: &discordgo.MessageEmbedThumbnail{
			URL: member.User.AvatarURL(""),
		},
		//discord compatible timestamp
		Timestamp: time.Now().Format(time.RFC3339),

		//now the embed fields
		Fields: []*discordgo.MessageEmbedField{
			{
				Name:   "Name:",
				Value:  member.Mention(),
				Inline: true,
			},
			{
				Name:   "Top Role:",
				Value:  top_role.Mention(),
				Inline: true,
			},
			{
				Name:   "Number of Roles",
				Value:  fmt.Sprint(len(user_roles)),
				Inline: true,
			},
			{
				Name:   "Joined Server on:",
				Value:  "<t:" + fmt.Sprint(user_joined_at.Unix()) + ":F>",
				Inline: true,
			},
			{
				Name:   "Join Rank:",
				Value:  fmt.Sprintf("%d/%d", rank, len(guild_members)),
				Inline: true,
			},
			{
				Name:   "Joined Discord on:",
				Value:  "<t:" + fmt.Sprint(user_created_at.Unix()) + ":F>",
				Inline: true,
			},
			{
				Name:   "Online status:",
				Value:  string(user_presence.Status),
				Inline: true,
			},
			{
				Name:   "Activity:",
				Value:  user_activity,
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
