package voice

import (
	"github.com/phxenix-w/gotestbot/internal/inits"
	"github.com/phxenix-w/gotestbot/internal/utils"
)

type Play struct{}

func (c *Play) Invokes() []string {
	return []string{"play", "p", "aufsicht", "aufsichtsperson"}
}

func (c *Play) Description() string {
	return "Plays an audio file."
}

func (c *Play) AdminPermission() bool {
	return false
}

//TODO: Actually play the music file lol
func (c *Play) Exec(ctx *inits.Context) error {
	vs, err := utils.UserVoiceState(ctx.Session, ctx.Message.Author.ID)

	if err != nil {
		ctx.Session.ChannelMessage(ctx.Message.ChannelID,
			"Could not find your voice channel!")
		return err
	}

	ch, err := ctx.Session.Channel(vs.ChannelID)
	if err != nil {
		return err
	}

	ctx.Session.ChannelVoiceJoin(ctx.Message.GuildID, vs.ChannelID, false, false)

	ctx.Session.ChannelMessageSend(ctx.Message.ChannelID,
		"Connected to voice channel `"+ch.Name+"`!")

	return nil
}
