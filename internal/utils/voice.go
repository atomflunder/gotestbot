package utils

import (
	"errors"

	"github.com/bwmarrin/discordgo"
)

func UserVoiceState(s *discordgo.Session, userID string) (*discordgo.VoiceState, error) {
	for _, guild := range s.State.Guilds {
		for _, vs := range guild.VoiceStates {
			if vs.UserID == userID {
				return vs, nil
			}
		}
	}
	return nil, errors.New("could not find your voice channel")
}
