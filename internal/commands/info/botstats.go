package info

import "github.com/phxenix-w/gotestbot/internal/inits"

type Botstats struct{}

func (c *Botstats) Invokes() []string {
	return []string{"botstats", "stats", "bot", "bstats"}
}

func (c *Botstats) Description() string {
	return "Gives you some useful information about the bot."
}

func (c *Botstats) AdminPermission() bool {
	return false
}

func (c *Botstats) Exec(ctx *inits.Context) error {
	//TODO: everything here, just wanted to setup this for tomorrow
	return nil
}
