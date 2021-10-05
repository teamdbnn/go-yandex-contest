package models

import (
	"github.com/go-openapi/strfmt"
)

type IDResource struct {
	CheckedID            *int64          `json:"checkedId"`
	CheckedRawExternalID *string         `json:"checkedRawExternalId"`
	ExternalID           *ExternalID     `json:"externalId"`
	ID                   int64           `json:"id,omitempty"`
	Model                *KClassResource `json:"model"`
	New                  *bool           `json:"new"`
	RawExternalID        string          `json:"rawExternalId,omitempty"`
	UID                  strfmt.UUID     `json:"uid,omitempty"`
}
