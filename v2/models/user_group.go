package models

type UserGroup struct {
	Global     bool              `json:"global,omitempty"`
	ID         *int64            `json:"id"`
	Name       string            `json:"name,omitempty"`
	Owner      *User             `json:"owner,omitempty"`
	System     bool              `json:"system,omitempty"`
	UserAccess map[string]string `json:"userAccess,omitempty"`
	Users      []*User           `json:"users"`
}
