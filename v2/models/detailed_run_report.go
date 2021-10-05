package models

type DetailedRunReport struct {
	CompileLog     *string `json:"compileLog"`
	Compiler       *string `json:"compiler"`
	Diff           *string `json:"diff"`
	ProblemAlias   *string `json:"problemAlias"`
	ProblemID      *string `json:"problemId"`
	Source         *string `json:"source"`
	SubmissionTime *string `json:"submissionTime"`
	Verdict        *string `json:"verdict"`
}
