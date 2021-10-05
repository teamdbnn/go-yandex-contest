package models

type ParticipantStatus struct {
	ContestDuration       string   `json:"contestDuration,omitempty"`
	ContestFinishTime     string   `json:"contestFinishTime,omitempty"`
	ContestInfinite       bool     `json:"contestInfinite,omitempty"`
	ContestStartTime      string   `json:"contestStartTime,omitempty"`
	ContestState          string   `json:"contestState,omitempty"`
	ParticipantFinishTime string   `json:"participantFinishTime,omitempty"`
	ParticipantLeftTime   string   `json:"participantLeftTime,omitempty"`
	ParticipantName       string   `json:"participantName,omitempty"`
	ParticipantStartTime  string   `json:"participantStartTime,omitempty"`
	Roles                 []string `json:"roles"`
}
