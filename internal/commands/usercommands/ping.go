package usercommands

import (
	"strings"

	"github.com/phxenix-w/gotestbot/internal/inits"
)

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
	//gets the amount of ms, without the digits after the dot, also converts it to a string all in one go
	p := strings.Split(ctx.Session.HeartbeatLatency().String(), ".")[0]

	_, err := ctx.Session.ChannelMessageSend(ctx.Message.ChannelID,
		"My ping is: "+p+"ms")

	//if theres an error, we return that
	if err != nil {
		return err
	}

	//otherwise just return nil
	return nil
}
