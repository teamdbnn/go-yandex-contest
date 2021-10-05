package models

type AdditionalSolutionCheckResult struct {
	CheckerName  string `json:"checkerName,omitempty"`
	ErrorMessage string `json:"errorMessage,omitempty"`
	Message      string `json:"message,omitempty"`
	Verdict      string `json:"verdict,omitempty"`
}
