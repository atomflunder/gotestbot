package info

import (
	"fmt"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/phxenix-w/gotestbot/internal/inits"
	"github.com/phxenix-w/gotestbot/internal/utils"
)

type Serverinfo struct{}

func (c *Serverinfo) Invokes() []string {
	return []string{"serverinfo", "server", "sinfo"}
}

func (c *Serverinfo) Description() string {
	return "Gets you an embed with server information."
}

func (c *Serverinfo) AdminPermission() bool {
	return false
}

func (c *Serverinfo) Exec(ctx *inits.Context) error {
	//gets the guild object
	guild, err := ctx.Session.State.Guild(ctx.Message.GuildID)
	if err != nil {
		return err
	}

	err = ctx.Session.RequestGuildMembers(guild.ID, "", 1000, true)
	if err != nil {
		return err
	}

	//getting the owner
	guildOwner, err := ctx.Session.GuildMember(guild.ID, guild.OwnerID)
	if err != nil {
		return err
	}

	//getting the top role from the owner for the embed color
	guildTopRole, err := utils.GetTopRole(guildOwner, ctx)
	if err != nil {
		return err
	}

	//gets the timestamp
	guildCreatedAt, err := discordgo.SnowflakeTimestamp(guild.ID)
	if err != nil {
		return err
	}

	//counts the channels in the guild
	var textChannels []*discordgo.Channel
	var voiceChannels []*discordgo.Channel
	for c := range guild.Channels {
		if guild.Channels[c].Type == 0 {
			textChannels = append(textChannels, guild.Channels[c])
		} else if guild.Channels[c].Type == 2 {
			voiceChannels = append(voiceChannels, guild.Channels[c])
		}
	}

	//counts the invites to the guild
	guildInvites, err := ctx.Session.GuildInvites(guild.ID)
	if err != nil {
		return err
	}

	//getting the amount of individual server boosters
	var guildBoosters []*discordgo.Member
	for m := range guild.Members {
		if guild.Members[m].PremiumSince != "" {
			guildBoosters = append(guildBoosters, guild.Members[m])
		}
	}

	//getting the amount of bot members
	var botMembers []*discordgo.Member
	for b := range guild.Members {
		if guild.Members[b].User.Bot {
			botMembers = append(botMembers, guild.Members[b])
		}
	}

	//the embed itself, quite full
	embed := &discordgo.MessageEmbed{
		Title: "Serverinfo of " + guild.Name + " (" + guild.ID + ")",
		Color: guildTopRole.Color,
		Thumbnail: &discordgo.MessageEmbedThumbnail{
			URL: "https://cdn.discordapp.com/icons/" + guild.ID + "/" + guild.Icon + ".png",
		},
		Timestamp: time.Now().Format(time.RFC3339),

		Fields: []*discordgo.MessageEmbedField{
			{
				Name:   "Created on:",
				Value:  utils.GetDiscordTimeStamp(&guildCreatedAt, "F"),
				Inline: true,
			},
			{
				Name:   "Owner:",
				Value:  guildOwner.Mention(),
				Inline: true,
			},
			{
				Name:   "Server Region:",
				Value:  guild.Region,
				Inline: true,
			},
			{
				Name:   "Members:",
				Value:  fmt.Sprint(len(guild.Members)) + " (Bots: " + fmt.Sprint(len(botMembers)) + ")",
				Inline: true,
			},
			{
				Name:   "Boosts:",
				Value:  fmt.Sprint(guild.PremiumSubscriptionCount) + " (Boosters: " + fmt.Sprint(len(guildBoosters)) + ")",
				Inline: true,
			},
			{
				Name:   "Active Invites:",
				Value:  fmt.Sprint(len(guildInvites)),
				Inline: true,
			},
			{
				Name:   "Roles:",
				Value:  fmt.Sprint(len(guild.Roles)),
				Inline: true,
			},
			{
				Name:   "Emojis:",
				Value:  fmt.Sprint(len(guild.Emojis)),
				Inline: true,
			},
			{
				Name:   "Stickers:",
				Value:  "Yet to be implemented in discordgo. Sorry :(",
				Inline: true,
			},
			{
				Name:   "Text Channels:",
				Value:  fmt.Sprint(len(textChannels)),
				Inline: true,
			},
			{
				Name:   "Voice Channels:",
				Value:  fmt.Sprint(len(voiceChannels)),
				Inline: true,
			},
			{
				Name:   "Active Threads:",
				Value:  "Yet to be implemented in discordgo. Sorry :(",
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
