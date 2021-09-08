package inits

import "github.com/bwmarrin/discordgo"

type MwPermissions struct{}

//permissions, just the basic admin y/n check for now
func (mw *MwPermissions) Exec(ctx *Context, cmd Command) (next bool, err error) {
	//if the command does not have admin required it goes through obviously
	if !cmd.AdminPermission() {
		next = true
		return
	}

	//defer always gets executed after the last return of this Exec function, its out error message for missing permissions
	defer func() {
		if !next && err == nil {
			_, err = ctx.Session.ChannelMessageSend(ctx.Message.ChannelID, "Nice try, but you don't have the required permissions for this command!")
		}
	}()

	guild, err := ctx.Session.Guild(ctx.Message.GuildID)
	if err != nil {
		return
	}

	//checks if the command was invoked by the owner of the server whos admin too ofc
	if guild.OwnerID == ctx.Message.Author.ID {
		next = true
	}

	//makes an array of all guild roles for the permission check
	roleMap := make(map[string]*discordgo.Role)

	for _, role := range guild.Roles {
		roleMap[role.ID] = role
	}

	//the admin check, checks if you have any role with the admin permission
	for _, rID := range ctx.Message.Member.Roles {
		if role, ok := roleMap[rID]; ok && role.Permissions&discordgo.PermissionAdministrator > 0 {
			next = true
			break
		}
	}

	return

}
