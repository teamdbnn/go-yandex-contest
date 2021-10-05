package models

type ContestTestingSettings struct {
	Contest            *Contest                   `json:"contest,omitempty"`
	ID                 *int64                     `json:"id"`
	PrecompileCheckers []*AdditionalSolutionCheck `json:"precompileCheckers"`
}
