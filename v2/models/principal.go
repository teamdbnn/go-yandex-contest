package models

import (
	"github.com/go-openapi/strfmt"
)

// Principal Principal
//
// swagger:model Principal
type Principal struct {

	// checked Id
	CheckedID int64 `json:"checkedId,omitempty"`

	// checked raw external Id
	CheckedRawExternalID string `json:"checkedRawExternalId,omitempty"`

	// created at
	// Required: true
	// Format: date-time
	CreatedAt *strfmt.DateTime `json:"createdAt"`

	// deleted
	// Required: true
	Deleted *bool `json:"deleted"`

	// id
	// Required: true
	ID *IDPrincipal `json:"id"`

	// modified at
	// Required: true
	// Format: date-time
	ModifiedAt *strfmt.DateTime `json:"modifiedAt"`

	// not deleted
	// Required: true
	NotDeleted *bool `json:"notDeleted"`

	// raw external Id
	RawExternalID string `json:"rawExternalId,omitempty"`

	// state
	// Required: true
	State *State `json:"state"`

	// title
	// Required: true
	Title *PrincipalTitle `json:"title"`
}
