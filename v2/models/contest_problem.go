package models

// ContestProblem ContestProblem
//
// swagger:model ContestProblem
type ContestProblem struct {

	// alias
	// Required: true
	Alias *string `json:"alias"`

	// compilers
	// Required: true
	Compilers []string `json:"compilers"`

	// id
	// Required: true
	ID *string `json:"id"`

	// limits
	// Required: true
	Limits []*CompilerLimit `json:"limits"`

	// name
	// Required: true
	Name *string `json:"name"`

	// statements
	// Required: true
	Statements []*Statement `json:"statements"`
}
