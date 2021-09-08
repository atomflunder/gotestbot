package usercommands

import "github.com/phxenix-w/gotestbot/internal/inits"

//the first command registered with the things required from the command.go file
type Ping struct{}

//the command names
func (c *Ping) Invokes() []string {
	return []string{"ping", "p"}
}

//description
func (c *Ping) Description() string {
	return "Basic Ping Command for testing purposes."
}

//if it uses admin rights
func (c *Ping) AdminPermission() bool {
	return false
}

//the command body
func (c *Ping) Exec(ctx *inits.Context) error {
	_, err := ctx.Session.ChannelMessageSend(ctx.Message.ChannelID,
		"Pong!")

	//if theres an error, we return that
	if err != nil {
		return err
	}

	//otherwise just return nil
	return nil
}
