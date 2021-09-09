package utils

import (
	"strings"

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
