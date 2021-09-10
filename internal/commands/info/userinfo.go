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
	//first we get the userID. if the author does not mention any user, it will use them
	userID := ""
	if len(ctx.Args) < 1 {
		userID = ctx.Message.Author.ID
	} else {
		userID = utils.UserMentionToID(ctx.Args[0])
	}

	//getting the member
	member, err := ctx.Session.GuildMember(ctx.Message.GuildID, userID)
	if err != nil {
		return err
	}

	//prepping some info for the embed
	userColor := ctx.Session.State.UserColor(userID, ctx.Message.ChannelID)

	//here we get the timestamps
	userCreatedAt, err := discordgo.SnowflakeTimestamp(userID)
	if err != nil {
		return err
	}
	//getting time.Time stamp for the join date, need to convert that first sadly
	userJoinedAt, err := time.Parse(time.RFC3339, string(member.JoinedAt))
	if err != nil {
		return err
	}

	//getting the user status/activity
	userStatus, userActivity, err := utils.UserStatusAndActivity(userID, ctx)
	if err != nil {
		return err
	}

	//getting the member roles
	userRoles := member.Roles
	//getting the top role
	topRole, err := utils.GetTopRole(member, ctx)
	if err != nil {
		return err
	}

	//this here gets all guild members. this caps at 1000 guild members so if the bot gets invited to a server with more than this
	//we would need to rewrite this using pagination or something. right now its not *that* important
	guildMembers, err := ctx.Session.GuildMembers(ctx.Message.GuildID, "", 1000)
	if err != nil {
		return err
	}

	//thankfully guild members are already ordered by join date (from my testing at least),
	//so we just need to search for our user ID and then we already got the rank. needs to start at 1.
	rank := 1
	for x := range guildMembers {
		if guildMembers[x].User.ID == member.User.ID {
			break
		}
		rank += 1
	}

	//the actual embed itself
	embed := &discordgo.MessageEmbed{
		//first the basic stuff
		Title: "Userinfo of " + member.User.String(),
		Color: userColor,
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
				Value:  topRole.Mention(),
				Inline: true,
			},
			{
				Name:   "Number of Roles",
				Value:  fmt.Sprint(len(userRoles)),
				Inline: true,
			},
			{
				Name:   "Joined Server on:",
				Value:  utils.GetDiscordTimeStamp(&userJoinedAt, "F"),
				Inline: true,
			},
			{
				Name:   "Join Rank:",
				Value:  fmt.Sprintf("%d/%d", rank, len(guildMembers)),
				Inline: true,
			},
			{
				Name:   "Joined Discord on:",
				Value:  utils.GetDiscordTimeStamp(&userCreatedAt, "F"),
				Inline: true,
			},
			{
				Name:   "Online status:",
				Value:  userStatus,
				Inline: true,
			},
			{
				Name:   "Activity:",
				Value:  userActivity,
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
