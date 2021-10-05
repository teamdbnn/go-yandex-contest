package models

type SimplifiedFullRunReport struct {
	CompileLog           *string                          `json:"compileLog"`
	Compiler             *string                          `json:"compiler"`
	ContestID            *int64                           `json:"contestId"`
	ContestName          *string                          `json:"contestName"`
	FinalScore           *string                          `json:"finalScore"`
	IP                   string                           `json:"ip,omitempty"`
	MaxMemoryUsage       *int64                           `json:"maxMemoryUsage"`
	MaxTimeUsage         *int64                           `json:"maxTimeUsage"`
	Participant          *Participant                     `json:"participant"`
	PostprocessorMessage string                           `json:"postprocessorMessage,omitempty"`
	PrecompileChecks     []*AdditionalSolutionCheckResult `json:"precompileChecks"`
	PreliminaryScore     *string                          `json:"preliminaryScore"`
	ProblemAlias         *string                          `json:"problemAlias"`
	ProblemID            *string                          `json:"problemId"`
	RunID                *int64                           `json:"runId"`
	Status               string                           `json:"status,omitempty"`
	SubmissionTime       *string                          `json:"submissionTime"`
	TestFileType         *string                          `json:"testFileType"`
	Tests                []*TestLog                       `json:"tests"`
	TimeFromStart        *int64                           `json:"timeFromStart"`
	Verdict              *string                          `json:"verdict"`
}
