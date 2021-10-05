package models

type ContestSettings struct {
	ContestName string     `json:"contestName,omitempty"`
	ContestType string     `json:"contestType,omitempty"`
	Duration    string     `json:"duration,omitempty"`
	Languages   []Language `json:"languages"`
}
