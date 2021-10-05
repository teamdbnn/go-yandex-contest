package models

type ContestDescription struct {
	Duration  int64  `json:"duration,omitempty"`
	Name      string `json:"name,omitempty"`
	StartTime string `json:"startTime,omitempty"`
	Type      string `json:"type,omitempty"`
}
