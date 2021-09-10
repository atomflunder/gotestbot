package utils

import (
	"strings"

	"github.com/phxenix-w/gotestbot/internal/inits"
)

//the input is the user mention as a string, and the output is the ID without the junk as a string
func UserMentionToID(UserMention string) string {
	replacer := strings.NewReplacer("<", "", "@", "", "!", "", ">", "")
	userID := replacer.Replace(UserMention)
	return userID
}

func UserStatusAndActivity(userID string, ctx *inits.Context) (string, string, error) {
	//getting the presence for the status/activity
	userPresence, err := ctx.Session.State.Presence(ctx.Message.GuildID, userID)
	if err != nil {
		return "", "", err
	}
	userStatus := string(userPresence.Status)
	//defining the activity in case the user didnt set one
	userActivity := "None"
	if len(userPresence.Activities) > 0 {
		userActivity = userPresence.Activities[0].Name
	}

	return userStatus, userActivity, nil

}
