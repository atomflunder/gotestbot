package voice

import (
	"github.com/phxenix-w/gotestbot/internal/inits"
	"github.com/phxenix-w/gotestbot/internal/utils"
)

type Stop struct{}

func (c *Stop) Invokes() []string {
	return []string{"stop", "s", "disconnect", "dc"}
}

func (c *Stop) Description() string {
	return "Stops playing an audio file."
}

func (c *Stop) AdminPermission() bool {
	return false
}

//just disconnects and sends a message
func (c *Stop) Exec(ctx *inits.Context) error {
	vs, err := utils.UserVoiceState(ctx.Session, ctx.Session.State.User.ID)
	if err != nil {
		return err
	}

	ch, err := ctx.Session.Channel(vs.ChannelID)
	if err != nil {
		return err
	}

	ctx.Session.ChannelVoiceJoin(ctx.Message.GuildID, "", false, false)

	ctx.Session.ChannelMessageSend(ctx.Message.ChannelID,
		"Disconnected from voice channel `"+ch.Name+"`!")

	return nil
}
