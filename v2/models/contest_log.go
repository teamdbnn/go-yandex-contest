package models

import (
	"github.com/go-openapi/strfmt"
)

type ContestLog struct {

	// events
	Events []*Event `json:"events"`

	// generation time
	// Format: date-time
	GenerationTime strfmt.DateTime `json:"generationTime,omitempty"`

	// problems
	Problems []Problem `json:"problems"`

	// settings
	Settings *ContestSettings `json:"settings,omitempty"`

	// users
	Users []*User `json:"users"`
}
