package models

type Event struct {
	AbsoluteTime int64 `json:"absoluteTime,omitempty"`
	ContestTime  int64 `json:"contestTime,omitempty"`
}
