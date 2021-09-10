package utils

import "strings"

func GetArgs(args []string, n int) string {
	return strings.Join(args[n:], " ")
}
