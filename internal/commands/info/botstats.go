package info

import (
	"fmt"
	"math"
	"runtime"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/phxenix-w/gotestbot/internal/inits"
	"github.com/phxenix-w/gotestbot/internal/utils"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/mem"
)

type Botstats struct{}

func (c *Botstats) Invokes() []string {
	return []string{"botstats", "stats", "bot", "bstats"}
}

func (c *Botstats) Description() string {
	return "Gives you some useful information about the bot."
}

func (c *Botstats) AdminPermission() bool {
	return false
}

func (c *Botstats) Exec(ctx *inits.Context) error {
	//getting the color of the bot
	botColor := ctx.Session.State.UserColor(ctx.Session.State.User.ID, ctx.Message.ChannelID)

	//getting all guilds the bot has access to
	botGuilds, err := ctx.Session.UserGuilds(10, "", "")
	if err != nil {
		return err
	}

	//getting the user count on those guilds
	totalMembers := 0
	for g := range botGuilds {
		guild, err := ctx.Session.State.Guild(botGuilds[g].ID)
		if err != nil {
			return err
		}
		totalMembers += len(guild.Members)
	}

	//getting the cpu percentage
	per, err := cpu.Percent(0, false)
	if err != nil {
		return err
	}

	//getting the ram usage
	v, err := mem.VirtualMemory()
	if err != nil {
		return err
	}

	//getting the uptime
	up, err := host.Uptime()
	if err != nil {
		return err
	}

	//the actual embed itself
	embed := &discordgo.MessageEmbed{
		//first the basic stuff
		Title: "Bot info about " + ctx.Session.State.User.String(),
		Color: botColor,
		Thumbnail: &discordgo.MessageEmbedThumbnail{
			URL: ctx.Session.State.User.AvatarURL(""),
		},
		//discord compatible timestamp
		Timestamp: time.Now().Format(time.RFC3339),

		//now the embed fields
		Fields: []*discordgo.MessageEmbedField{
			{
				Name:   "Name:",
				Value:  ctx.Session.State.User.Mention(),
				Inline: true,
			},
			{
				Name:   "Servers",
				Value:  fmt.Sprint(len(botGuilds)),
				Inline: true,
			},
			{
				Name:   "Total Users:",
				Value:  fmt.Sprint(totalMembers),
				Inline: true,
			},
			{
				Name:   "Bot Version:",
				Value:  utils.BotVersion,
				Inline: true,
			},
			{
				Name:   "Go Version:",
				Value:  runtime.Version(),
				Inline: true,
			},
			{
				Name:   "discordgo Version:",
				Value:  discordgo.VERSION,
				Inline: true,
			},
			{
				Name:   "CPU Usage:",
				Value:  fmt.Sprint(math.Round(per[0])) + "%", //theres only 1 value in per anyways, its the average between all cores
				Inline: true,
			},
			{
				Name:   "RAM Usage:",
				Value:  fmt.Sprint(v.UsedPercent) + "%",
				Inline: true,
			},
			{
				Name:   "System Uptime:",
				Value:  utils.SecondsToHumanTime(int(up)),
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
