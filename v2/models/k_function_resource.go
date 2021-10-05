package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

// KFunctionResource KFunction«Resource»
//
// swagger:model KFunction«Resource»
type KFunctionResource struct {

	// abstract
	// Required: true
	Abstract *bool `json:"abstract"`

	// annotations
	// Required: true
	Annotations []Annotation `json:"annotations"`

	// external
	// Required: true
	External *bool `json:"external"`

	// final
	// Required: true
	Final *bool `json:"final"`

	// infix
	// Required: true
	Infix *bool `json:"infix"`

	// inline
	// Required: true
	Inline *bool `json:"inline"`

	// name
	// Required: true
	Name *string `json:"name"`

	// open
	// Required: true
	Open *bool `json:"open"`

	// operator
	// Required: true
	Operator *bool `json:"operator"`

	// parameters
	// Required: true
	Parameters []*KParameter `json:"parameters"`

	// return type
	// Required: true
	ReturnType *KType `json:"returnType"`

	// suspend
	// Required: true
	Suspend *bool `json:"suspend"`

	// type parameters
	// Required: true
	TypeParameters []*KTypeParameter `json:"typeParameters"`

	// visibility
	// Enum: [PUBLIC PROTECTED INTERNAL PRIVATE]
	Visibility string `json:"visibility,omitempty"`
}
