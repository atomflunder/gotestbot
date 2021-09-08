package inits

//the basic things needed for commands
type Command interface {
	Invokes() []string
	Description() string
	AdminPermission() bool //permissions check: admin y/n is enough for now
	Exec(ctx *Context) error
}
