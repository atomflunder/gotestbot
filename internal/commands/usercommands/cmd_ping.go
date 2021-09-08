package usercommands

import "github.com/phxenix-w/gotestbot/internal/inits"

//the first command registered with the things required from the command.go file
type CmdPing struct{}

//the command names
func (c *CmdPing) Invokes() []string {
	return []string{"ping", "p"}
}

//description
func (c *CmdPing) Description() string {
	return "Basic Ping Command for testing purposes."
}

//if it uses admin rights
func (c *CmdPing) AdminPermission() bool {
	return false
}

//the command body
func (c *CmdPing) Exec(ctx *inits.Context) error {
	_, err := ctx.Session.ChannelMessageSend(ctx.Message.ChannelID,
		"Pong!")

	//if theres an error, we return that
	if err != nil {
		return err
	}

	//otherwise just return nil
	return nil
}
