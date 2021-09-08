package utils

import "strings"

func GetArgs(args []string, number int) string {
	return strings.Join(args[number:], " ")
}
