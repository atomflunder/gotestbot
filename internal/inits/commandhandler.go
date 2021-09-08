package inits

import (
	"strings"

	"github.com/bwmarrin/discordgo"
)

type CommandHandler struct {
	prefix string

	cmdInstances []Command
	cmdMap       map[string]Command
	middlewares  []Middleware

	OnError func(err error, ctx *Context)
}

//the new command handler we call from main.go
func NewCommandHandler(prefix string) *CommandHandler {
	return &CommandHandler{
		prefix:       prefix,
		cmdInstances: make([]Command, 0),
		cmdMap:       make(map[string]Command),
		middlewares:  make([]Middleware, 0),
		OnError:      func(error, *Context) {},
	}
}

//registers every command for us, also called from main.go
func (c *CommandHandler) RegisterCommand(cmd Command) {
	c.cmdInstances = append(c.cmdInstances, cmd)
	for _, invoke := range cmd.Invokes() {
		c.cmdMap[invoke] = cmd
	}
}

func (c *CommandHandler) RegisterMiddleware(mw Middleware) {
	c.middlewares = append(c.middlewares, mw)
}

//this here parses the messages, "converts" them into commands
func (c *CommandHandler) HandleMessage(s *discordgo.Session, e *discordgo.MessageCreate) {
	//just returns if its our bot, another bot or the message does not start with the prefix
	if e.Author.ID == s.State.User.ID || e.Author.Bot || !strings.HasPrefix(e.Content, c.prefix) {
		return
	}

	//slices off the prefix
	split := strings.Split(e.Content[len(c.prefix):], " ")
	//checks if there is anything after the prefix
	if len(split) < 1 {
		return
	}

	//this is the command being invoked (in lowercase)
	invoke := strings.ToLower(split[0])
	//these are the rest of the args after the command
	args := split[1:]

	cmd, ok := c.cmdMap[invoke]
	//if the command is not found, returns
	if !ok || cmd == nil {
		return
	}

	ctx := &Context{
		Session: s,
		Args:    args,
		Handler: c,
		Message: e.Message,
	}

	for _, mw := range c.middlewares {
		next, err := mw.Exec(ctx, cmd)
		//if theres some kind of error in our command we obv need to return
		if err != nil {
			c.OnError(err, ctx)
			return
		}
		if !next {
			return
		}
	}

	if err := cmd.Exec(ctx); err != nil {
		c.OnError(err, ctx)
	}

}
