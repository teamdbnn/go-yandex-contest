package models

type User struct {
	AccessLevel          *AccessLevel `json:"accessLevel,omitempty"`
	CreationTime         []int64      `json:"creationTime"`
	Creator              *User        `json:"creator,omitempty"`
	CreatorID            []int64      `json:"creatorId"`
	DisplayedName        string       `json:"displayedName,omitempty"`
	Guest                bool         `json:"guest,omitempty"`
	ID                   *int64       `json:"id"`
	Login                string       `json:"login,omitempty"`
	PasswordHash         string       `json:"passwordHash,omitempty"`
	Salt                 string       `json:"salt,omitempty"`
	StudentsGroup        *UserGroup   `json:"studentsGroup,omitempty"`
	StudentsInvitesGroup *UserGroup   `json:"studentsInvitesGroup,omitempty"`
	UID                  string       `json:"uid,omitempty"`
	UIDAsLong            int64        `json:"uidAsLong,omitempty"`
	UserType             string       `json:"userType,omitempty"`
}
