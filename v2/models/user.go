package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

// User User
//
// swagger:model User
type User struct {

	// access level
	AccessLevel *AccessLevel `json:"accessLevel,omitempty"`

	// creation time
	CreationTime []int64 `json:"creationTime"`

	// creator
	Creator *User `json:"creator,omitempty"`

	// creator Id
	CreatorID []int64 `json:"creatorId"`

	// displayed name
	DisplayedName string `json:"displayedName,omitempty"`

	// guest
	Guest bool `json:"guest,omitempty"`

	// id
	// Required: true
	ID *int64 `json:"id"`

	// login
	Login string `json:"login,omitempty"`

	// password hash
	PasswordHash string `json:"passwordHash,omitempty"`

	// salt
	Salt string `json:"salt,omitempty"`

	// students group
	StudentsGroup *UserGroup `json:"studentsGroup,omitempty"`

	// students invites group
	StudentsInvitesGroup *UserGroup `json:"studentsInvitesGroup,omitempty"`

	// uid
	UID string `json:"uid,omitempty"`

	// uid as long
	UIDAsLong int64 `json:"uidAsLong,omitempty"`

	// user type
	// Enum: [PASSPORT GUEST SHAD INTERN OPENCUP LTI INTERNAL]
	UserType string `json:"userType,omitempty"`
}