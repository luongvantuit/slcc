package interfaces

type RoleSet struct {
	OrgId string   `json:"orgId"`
	Roles []string `json:"roles"`
}
