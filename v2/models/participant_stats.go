package models

type ParticipantStats struct {
	ContestID           *int64           `json:"contestId"`
	FirstSubmissionTime string           `json:"firstSubmissionTime,omitempty"`
	ID                  *int64           `json:"id"`
	Login               *string          `json:"login"`
	Name                *string          `json:"name"`
	Runs                []*FullRunReport `json:"runs"`
	StartedAt           string           `json:"startedAt,omitempty"`
}
