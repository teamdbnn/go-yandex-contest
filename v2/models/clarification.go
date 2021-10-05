package models

type Clarification struct {
	Message string `json:"message,omitempty"`
	Subject string `json:"subject,omitempty"`
}
