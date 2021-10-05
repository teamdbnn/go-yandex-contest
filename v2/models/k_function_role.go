package models

type KFunctionRole struct {
	Abstract       *bool             `json:"abstract"`
	Annotations    []Annotation      `json:"annotations"`
	External       *bool             `json:"external"`
	Final          *bool             `json:"final"`
	Infix          *bool             `json:"infix"`
	Inline         *bool             `json:"inline"`
	Name           *string           `json:"name"`
	Open           *bool             `json:"open"`
	Operator       *bool             `json:"operator"`
	Parameters     []*KParameter     `json:"parameters"`
	ReturnType     *KType            `json:"returnType"`
	Suspend        *bool             `json:"suspend"`
	TypeParameters []*KTypeParameter `json:"typeParameters"`
	Visibility     string            `json:"visibility,omitempty"`
}
