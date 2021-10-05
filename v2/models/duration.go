package models

type Duration struct {
	Millis          int64 `json:"millis,omitempty"`
	StandardDays    int64 `json:"standardDays,omitempty"`
	StandardHours   int64 `json:"standardHours,omitempty"`
	StandardMinutes int64 `json:"standardMinutes,omitempty"`
	StandardSeconds int64 `json:"standardSeconds,omitempty"`
}
