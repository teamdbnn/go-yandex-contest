package models

type KTypeProjection struct {
	Type     *KType `json:"type,omitempty"`
	Variance string `json:"variance,omitempty"`
}
