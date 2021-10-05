package models

type ContestDescription struct {

	// duration
	Duration int64 `json:"duration,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// start time
	StartTime string `json:"startTime,omitempty"`

	// type
	Type string `json:"type,omitempty"`
}
