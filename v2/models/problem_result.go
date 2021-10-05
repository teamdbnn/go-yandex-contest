package models

type ProblemResult struct {
	Score           string `json:"score,omitempty"`
	Status          string `json:"status,omitempty"`
	SubmissionCount string `json:"submissionCount,omitempty"`
	SubmitDelay     int64  `json:"submitDelay,omitempty"`
}
