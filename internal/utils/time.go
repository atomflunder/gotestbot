package utils

import (
	"fmt"
	"time"
)

func GetDiscordTimeStamp(timestamp *time.Time, style string) string {
	ts := "<t:" + fmt.Sprint(timestamp.Unix()) + ":" + style + ">"

	return ts
}
