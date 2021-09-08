package commands

type CmdPong struct{}

func (c *CmdPong) Invokes() []string {
	return []string{"pong", "pp"}
}

func (c *CmdPong) Description() string {
	return "Basic Pong Command for testing purposes. Equivalent of ping. Check out that file for more info."
}

func (c *CmdPong) AdminPermission() bool {
	return false
}

func (c *CmdPong) Exec(ctx *Context) error {
	_, err := ctx.Session.ChannelMessageSend(ctx.Message.ChannelID,
		"Ping!")

	if err != nil {
		return err
	}

	return nil
}
