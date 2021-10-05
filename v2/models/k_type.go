package models

type KType struct {
	Annotations    []Annotation       `json:"annotations"`
	Arguments      []*KTypeProjection `json:"arguments"`
	Classifier     KClassifier        `json:"classifier,omitempty"`
	MarkedNullable *bool              `json:"markedNullable"`
}
