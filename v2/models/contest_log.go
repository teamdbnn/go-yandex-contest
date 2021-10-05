package models

import (
	"github.com/go-openapi/strfmt"
)

type ContestLog struct {
	Events         []*Event         `json:"events"`
	GenerationTime strfmt.DateTime  `json:"generationTime,omitempty"`
	Problems       []Problem        `json:"problems"`
	Settings       *ContestSettings `json:"settings,omitempty"`
	Users          []*User          `json:"users"`
}
