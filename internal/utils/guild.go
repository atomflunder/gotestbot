package utils

import "github.com/bwmarrin/discordgo"

//gets you the server boosters within the supplied list
func GuildBoosters(m []*discordgo.Member) (boosters []*discordgo.Member) {
	for i := range m {
		if m[i].PremiumSince != "" {
			boosters = append(boosters, m[i])
		}
	}
	return boosters
}

//gets you the bot accounts within the supplied list
func BotMembers(m []*discordgo.Member) (bots []*discordgo.Member) {
	for i := range m {
		if m[i].User.Bot {
			bots = append(bots, m[i])
		}
	}
	return bots
}
