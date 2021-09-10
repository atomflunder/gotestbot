package utils

import (
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
