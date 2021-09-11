package utils

import (
	"strconv"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/phxenix-w/gotestbot/internal/inits"
	"github.com/sahilm/fuzzy"
)

func RoleMentionToID(RoleMention string) string {
	replacer := strings.NewReplacer("<", "", "@", "", "&", "", ">", "")
	roleID := replacer.Replace(RoleMention)
	return roleID
}

//this converts the input to a matching *discordgo.Role
func RoleFromInput(roleInput string, ctx *inits.Context) (*discordgo.Role, error) {
	var role *discordgo.Role
	//we try the input out, if its an int it it tries to use it as an ID
	//the "oversight" here is that if a role has a name only with ints and you search for it, it will not work
	if _, err := strconv.Atoi(roleInput); err == nil {
		role, err = ctx.Session.State.Role(ctx.Message.GuildID, roleInput)
		if err != nil {
			return nil, err
		}
		//if this is not true, it searches for the closest matching role, and then gives it out
	} else {
		roles, _ := ctx.Session.GuildRoles(ctx.Message.GuildID)

		//the name and IDs correspond so we can search for the name and use that to get the index of IDs
		var roleNames []string
		var roleIDs []string

		for _, role := range roles {
			roleNames = append(roleNames, role.Name)
			roleIDs = append(roleIDs, role.ID)
		}

		//uses fuzzy sort to find a good match
		match := fuzzy.Find(roleInput, roleNames)

		//uses the method mentioned above
		roleID := roleIDs[match[0].Index]

		role, err = ctx.Session.State.Role(ctx.Message.GuildID, roleID)
		if err != nil {
			return nil, err
		}
	}
	return role, nil
}

func GetTopRole(member *discordgo.Member, ctx *inits.Context) (*discordgo.Role, error) {
	//getting the member roles
	userRoles := member.Roles
	//this whole block is for searching the top role of the member.
	//im sure this can be done better but i certainly couldnt come up with something better at 1am
	//first we get every role ID and their position
	var rolePos []int
	var roleIDs []string
	for x := range member.Roles {
		role, err := ctx.Session.State.Role(ctx.Message.GuildID, userRoles[x])
		if err != nil {
			return nil, err
		}
		rolePos = append(rolePos, role.Position)
		roleIDs = append(roleIDs, role.ID)
	}
	//then we search for the highest position value
	maxRole := rolePos[0]
	for _, value := range rolePos {
		if value > maxRole {
			maxRole = value
		}
	}
	//then we search the index of the highest position role in the original slice
	n := 0
	for x := range rolePos {
		if maxRole == rolePos[x] {
			break
		}
		n += 1
	}
	//and then we use that index to get the according role ID in the other slice and finally get the role that is highest
	topRole, err := ctx.Session.State.Role(ctx.Message.GuildID, roleIDs[n])
	if err != nil {
		return nil, err
	}

	return topRole, nil

}
