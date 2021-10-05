package models

type KCallableObject struct {
	Abstract       *bool             `json:"abstract"`
	Annotations    []Annotation      `json:"annotations"`
	Final          *bool             `json:"final"`
	Name           *string           `json:"name"`
	Open           *bool             `json:"open"`
	Parameters     []*KParameter     `json:"parameters"`
	ReturnType     *KType            `json:"returnType"`
	Suspend        *bool             `json:"suspend"`
	TypeParameters []*KTypeParameter `json:"typeParameters"`
	Visibility     string            `json:"visibility,omitempty"`
}
