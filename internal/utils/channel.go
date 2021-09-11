package utils

import "github.com/bwmarrin/discordgo"

//makes it easier to read instead of getting these ints back and looking them up each time
func CheckChannelType(ch *discordgo.Channel) string {
	if ch.Type == 0 {
		return "Text"
	} else if ch.Type == 1 {
		return "DM"
	} else if ch.Type == 2 {
		return "Voice"
	} else if ch.Type == 3 {
		return "Group DM"
	} else if ch.Type == 4 {
		return "Category"
	} else if ch.Type == 5 {
		return "News"
	} else if ch.Type == 6 {
		return "Store"
	} else {
		return "Invalid"
	}
}
