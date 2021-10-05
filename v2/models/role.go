package models

import (
	"github.com/go-openapi/strfmt"
)

type Role struct {
	CheckedID            int64            `json:"checkedId,omitempty"`
	CheckedRawExternalID string           `json:"checkedRawExternalId,omitempty"`
	CreatedAt            *strfmt.DateTime `json:"createdAt"`
	Deleted              *bool            `json:"deleted"`
	ID                   *IDRole          `json:"id"`
	ModifiedAt           *strfmt.DateTime `json:"modifiedAt"`
	NotDeleted           *bool            `json:"notDeleted"`
	PrincipalID          *IDPrincipal     `json:"principalId"`
	RawExternalID        string           `json:"rawExternalId,omitempty"`
	ResourceID           *IDResource      `json:"resourceId"`
	Roles                []*IDRoleMeta    `json:"roles"`
	State                *State           `json:"state"`
}
