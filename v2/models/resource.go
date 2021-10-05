package models

import (
	"github.com/go-openapi/strfmt"
)

type Resource struct {
	CheckedID            int64            `json:"checkedId,omitempty"`
	CheckedRawExternalID string           `json:"checkedRawExternalId,omitempty"`
	CreatedAt            *strfmt.DateTime `json:"createdAt"`
	Deleted              *bool            `json:"deleted"`
	ID                   *IDResource      `json:"id"`
	ModifiedAt           *strfmt.DateTime `json:"modifiedAt"`
	NotDeleted           *bool            `json:"notDeleted"`
	RawExternalID        string           `json:"rawExternalId,omitempty"`
	State                *State           `json:"state"`
	Title                *ResourceTitle   `json:"title"`
}
