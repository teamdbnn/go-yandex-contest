package models

import (
	"github.com/go-openapi/strfmt"
)

// Contest Contest
//
// swagger:model Contest
type Contest struct {

	// compilers
	Compilers []string `json:"compilers"`

	// duration
	Duration *Duration `json:"duration,omitempty"`

	// enabled
	Enabled bool `json:"enabled,omitempty"`

	// end time
	// Format: date-time
	EndTime strfmt.DateTime `json:"endTime,omitempty"`

	// finished
	Finished bool `json:"finished,omitempty"`

	// id
	// Required: true
	ID *int64 `json:"id"`

	// infinite
	Infinite bool `json:"infinite,omitempty"`

	// monitor
	Monitor *MonitorConfiguration `json:"monitor,omitempty"`

	// monitor plugin
	MonitorPlugin string `json:"monitorPlugin,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// owner
	Owner *User `json:"owner,omitempty"`

	// problem set Id
	ProblemSetID string `json:"problemSetId,omitempty"`

	// start time
	// Format: date-time
	StartTime strfmt.DateTime `json:"startTime,omitempty"`

	// testing settings
	TestingSettings *ContestTestingSettings `json:"testingSettings,omitempty"`

	// time control type
	// Enum: [USUAL VIRTUAL]
	TimeControlType string `json:"timeControlType,omitempty"`

	// time limited
	TimeLimited bool `json:"timeLimited,omitempty"`
}
