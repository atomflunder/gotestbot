package utils

import "strings"

//the input is the user mention as a string, and the output is the ID without the junk as a string
func UserMentionToID(UserMention string) string {
	replacer := strings.NewReplacer("<", "", "@", "", "!", "", ">", "")
	user_id := replacer.Replace(UserMention)
	return user_id
}
