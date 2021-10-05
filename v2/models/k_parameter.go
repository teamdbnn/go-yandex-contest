package models

type KParameter struct {
	Annotations []Annotation `json:"annotations"`
	Index       *int32       `json:"index"`
	Kind        *string      `json:"kind"`
	Name        string       `json:"name,omitempty"`
	Optional    *bool        `json:"optional"`
	Type        *KType       `json:"type"`
	Vararg      *bool        `json:"vararg"`
}
