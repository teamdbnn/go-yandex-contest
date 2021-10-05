package models

type ContestProblem struct {
	ID         *string          `json:"id"`
	Name       *string          `json:"name"`
	Alias      *string          `json:"alias"`
	Compilers  []string         `json:"compilers"`
	Statements []*Statement     `json:"statements"`
	Limits     []*CompilerLimit `json:"limits"`
}
