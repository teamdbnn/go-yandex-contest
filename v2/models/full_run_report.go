package models

type FullRunReport struct {
	CheckerLog       []*TestLog                       `json:"checkerLog"`
	CompileLog       *string                          `json:"compileLog"`
	Compiler         *string                          `json:"compiler"`
	ContestID        *int64                           `json:"contestId"`
	ContestName      *string                          `json:"contestName"`
	Diff             *string                          `json:"diff"`
	FinalScore       *string                          `json:"finalScore"`
	IP               string                           `json:"ip,omitempty"`
	MaxMemoryUsage   *int64                           `json:"maxMemoryUsage"`
	MaxTimeUsage     *int64                           `json:"maxTimeUsage"`
	Participant      *Participant                     `json:"participant"`
	PrecompileChecks []*AdditionalSolutionCheckResult `json:"precompileChecks"`
	PreliminaryScore *string                          `json:"preliminaryScore"`
	ProblemAlias     *string                          `json:"problemAlias"`
	ProblemID        *string                          `json:"problemId"`
	ProblemName      *string                          `json:"problemName"`
	RunID            *int64                           `json:"runId"`
	Source           *string                          `json:"source"`
	Status           string                           `json:"status,omitempty"`
	SubmissionTime   *string                          `json:"submissionTime"`
	TestFileType     *string                          `json:"testFileType"`
	TimeFromStart    *int64                           `json:"timeFromStart"`
	Verdict          *string                          `json:"verdict"`
}
