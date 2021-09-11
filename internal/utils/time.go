package utils

import (
	"fmt"
	"time"
)

//makes it easy to create a time zone sensitive datetime obj
func GetDiscordTimeStamp(timestamp *time.Time, style string) string {
	ts := "<t:" + fmt.Sprint(timestamp.Unix()) + ":" + style + ">"

	return ts
}
