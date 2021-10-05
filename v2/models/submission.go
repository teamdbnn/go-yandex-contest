package models

type Submission struct {
	Author       *string `json:"author"`
	Compiler     *string `json:"compiler"`
	ID           *int64  `json:"id"`
	Memory       *int64  `json:"memory"`
	ProblemAlias *string `json:"problemAlias"`
	ProblemID    *string `json:"problemId"`
	Test         *int32  `json:"test"`
	Time         *int64  `json:"time"`
	Verdict      *string `json:"verdict"`
}
