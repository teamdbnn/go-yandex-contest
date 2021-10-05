package models

type RoleMeta struct {
	CheckedID            int64       `json:"checkedId,omitempty"`
	CheckedRawExternalID string      `json:"checkedRawExternalId,omitempty"`
	Content              *JSONNode   `json:"content,omitempty"`
	ID                   *IDRoleMeta `json:"id"`
	RawExternalID        string      `json:"rawExternalId,omitempty"`
}
