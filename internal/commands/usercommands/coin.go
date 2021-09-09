package usercommands

import (
	"math/rand"

	"github.com/phxenix-w/gotestbot/internal/inits"
)

type Coin struct{}

func (c *Coin) Invokes() []string {
	return []string{"coin", "c", "coinflip", "flip"}
}

func (c *Coin) Description() string {
	return "Flips a coin and tells you the result."
}

func (c *Coin) AdminPermission() bool {
	return false
}

func (c *Coin) Exec(ctx *inits.Context) error {
	//this command does not take any args so its fairly simple
	//first we get the array and then a random number between 0-1
	coin := [2]string{"Heads", "Tails"}
	n := rand.Intn(2)

	//then we send the result
	_, err := ctx.Session.ChannelMessageSend(ctx.Message.ChannelID,
		"Coinflip: \n"+coin[n]+"!")

	if err != nil {
		return err
	}

	return nil
}
