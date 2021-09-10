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
	role_id := replacer.Replace(RoleMention)
	return role_id
}

func RoleSearch(RoleInput string, ctx *inits.Context) string {
	//getting every role
	roles, _ := ctx.Session.GuildRoles(ctx.Message.GuildID)

	//the name and IDs correspond so we can search for the name and use that to get the index of IDs
	var role_names []string
	var role_ids []string

	for _, role := range roles {
		role_names = append(role_names, role.Name)
		role_ids = append(role_ids, role.ID)
	}

	//uses fuzzy sort to find a good match
	match := fuzzy.Find(RoleInput, role_names)

	//uses the method mentioned above
	role_id := role_ids[match[0].Index]

	return role_id
}

//this converts the input to an int, if its true it will try to use the role ID directly. uses the function above too
//the "oversight" here is that if a role has a name only with ints and you search for it, it will not work
//need to fix this sometime, but for now this is the best i came up with
func ReturnRoleFromInput(role_input string, ctx *inits.Context) (*discordgo.Role, error) {
	var role *discordgo.Role
	var role_id string
	//we try the input out, if its an int it it tries to use it as an ID
	if _, err := strconv.Atoi(role_input); err == nil {
		role, err = ctx.Session.State.Role(ctx.Message.GuildID, role_input)
		if err != nil {
			return nil, err
		}
		//if this is not true, it searches for the closest matching role, and then gives it out
	} else {
		role_id = RoleSearch(role_input, ctx)
		role, err = ctx.Session.State.Role(ctx.Message.GuildID, role_id)
		if err != nil {
			return nil, err
		}
	}
	return role, nil
}

func GetTopRole(member *discordgo.Member, ctx *inits.Context) (*discordgo.Role, error) {
	//getting the member roles
	user_roles := member.Roles
	//this whole block is for searching the top role of the member.
	//im sure this can be done better but i certainly couldnt come up with something better at 1am
	//first we get every role ID and their position
	var role_positions []int
	var role_ids []string
	for x := range member.Roles {
		role, err := ctx.Session.State.Role(ctx.Message.GuildID, user_roles[x])
		if err != nil {
			return nil, err
		}
		role_positions = append(role_positions, role.Position)
		role_ids = append(role_ids, role.ID)
	}
	//then we search for the highest position value
	max_role := role_positions[0]
	for _, value := range role_positions {
		if value > max_role {
			max_role = value
		}
	}
	//then we search the index of the highest position role in the original slice
	n := 0
	for x := range role_positions {
		if max_role == role_positions[x] {
			break
		}
		n += 1
	}
	//and then we use that index to get the according role ID in the other slice and finally get the role that is highest
	top_role, err := ctx.Session.State.Role(ctx.Message.GuildID, role_ids[n])
	if err != nil {
		return nil, err
	}

	return top_role, nil

}
