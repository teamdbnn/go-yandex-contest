package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

// ParticipantStats ParticipantStats
//
// swagger:model ParticipantStats
type ParticipantStats struct {

	// contest Id
	// Required: true
	ContestID *int64 `json:"contestId"`

	// first submission time
	FirstSubmissionTime string `json:"firstSubmissionTime,omitempty"`

	// id
	// Required: true
	ID *int64 `json:"id"`

	// login
	// Required: true
	Login *string `json:"login"`

	// name
	// Required: true
	Name *string `json:"name"`

	// runs
	// Required: true
	Runs []*FullRunReport `json:"runs"`

	// started at
	StartedAt string `json:"startedAt,omitempty"`
}