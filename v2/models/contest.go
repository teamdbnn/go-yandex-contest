package models

import (
	"github.com/go-openapi/strfmt"
)

type Contest struct {
	ID              *int64                  `json:"id"`
	Compilers       []string                `json:"compilers"`
	Duration        *Duration               `json:"duration,omitempty"`
	Enabled         bool                    `json:"enabled,omitempty"`
	EndTime         strfmt.DateTime         `json:"endTime,omitempty"`
	Finished        bool                    `json:"finished,omitempty"`
	Infinite        bool                    `json:"infinite,omitempty"`
	Monitor         *MonitorConfiguration   `json:"monitor,omitempty"`
	MonitorPlugin   string                  `json:"monitorPlugin,omitempty"`
	Name            string                  `json:"name,omitempty"`
	Owner           *User                   `json:"owner,omitempty"`
	ProblemSetID    string                  `json:"problemSetId,omitempty"`
	StartTime       strfmt.DateTime         `json:"startTime,omitempty"`
	TestingSettings *ContestTestingSettings `json:"testingSettings,omitempty"`
	TimeControlType string                  `json:"timeControlType,omitempty"`
	TimeLimited     bool                    `json:"timeLimited,omitempty"`
}
