package usercommands

import "github.com/phxenix-w/gotestbot/internal/inits"

type Pong struct{}

func (c *Pong) Invokes() []string {
	return []string{"pong", "pp"}
}

func (c *Pong) Description() string {
	return "Basic Pong Command for testing purposes. Equivalent of ping. Check out that file for more info."
}

func (c *Pong) AdminPermission() bool {
	return false
}

func (c *Pong) Exec(ctx *inits.Context) error {
	_, err := ctx.Session.ChannelMessageSend(ctx.Message.ChannelID,
		"Ping!")

	if err != nil {
		return err
	}

	return nil
}
