package models

type GrantResponse struct {
	Roles []*ExtendedRole `json:"roles"`
}
