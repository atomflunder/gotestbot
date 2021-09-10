package utils

import (
	"strings"

	"github.com/phxenix-w/gotestbot/internal/inits"
)

//the input is the user mention as a string, and the output is the ID without the junk as a string
func UserMentionToID(UserMention string) string {
	replacer := strings.NewReplacer("<", "", "@", "", "!", "", ">", "")
	user_id := replacer.Replace(UserMention)
	return user_id
}

func UserStatusAndActivity(user_id string, ctx *inits.Context) (string, string, error) {
	//getting the presence for the status/activity
	user_presence, err := ctx.Session.State.Presence(ctx.Message.GuildID, user_id)
	if err != nil {
		return "", "", err
	}
	user_status := string(user_presence.Status)
	//defining the activity in case the user didnt set one
	user_activity := "None"
	if len(user_presence.Activities) > 0 {
		user_activity = user_presence.Activities[0].Name
	}

	return user_status, user_activity, nil

}
