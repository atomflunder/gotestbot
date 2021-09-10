package utils

import (
	"fmt"
	"time"
)

func GetDiscordTimeStamp(ts *time.Time, style string) string {
	timestamp := "<t:" + fmt.Sprint(ts.Unix()) + ":" + style + ">"

	return timestamp
}
