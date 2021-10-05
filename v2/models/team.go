package models

type Team struct {
	ID        *int64  `json:"id"`
	Invited   []*User `json:"invited"`
	Name      string  `json:"name,omitempty"`
	Teammates []*User `json:"teammates"`
}
