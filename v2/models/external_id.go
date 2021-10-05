package models

type ExternalID struct {
	ClientID *string `json:"clientId"`
	Empty    *bool   `json:"empty"`
	ID       string  `json:"id,omitempty"`
}
