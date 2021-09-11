package utils

import "strings"

//makes it a bit more readable in my opinion
func GetArgs(args []string, n int) string {
	return strings.Join(args[n:], " ")
}
