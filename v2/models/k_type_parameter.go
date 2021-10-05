package models

type KTypeParameter struct {
	Name        *string  `json:"name"`
	Reified     *bool    `json:"reified"`
	UpperBounds []*KType `json:"upperBounds"`
	Variance    *string  `json:"variance"`
}
