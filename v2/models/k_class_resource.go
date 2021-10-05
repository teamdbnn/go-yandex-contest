package models

// KClassResource KClass«Resource»
type KClassResource struct {

	// abstract
	// Required: true
	Abstract *bool `json:"abstract"`

	// annotations
	// Required: true
	Annotations []Annotation `json:"annotations"`

	// companion
	// Required: true
	Companion *bool `json:"companion"`

	// constructors
	// Required: true
	Constructors []*KFunctionResource `json:"constructors"`

	// data
	// Required: true
	Data *bool `json:"data"`

	// final
	// Required: true
	Final *bool `json:"final"`

	// inner
	// Required: true
	Inner *bool `json:"inner"`

	// members
	// Required: true
	Members []*KCallableObject `json:"members"`

	// nested classes
	// Required: true
	NestedClasses []*KClassObject `json:"nestedClasses"`

	// object instance
	ObjectInstance *Resource `json:"objectInstance,omitempty"`

	// open
	// Required: true
	Open *bool `json:"open"`

	// qualified name
	QualifiedName string `json:"qualifiedName,omitempty"`

	// sealed
	// Required: true
	Sealed *bool `json:"sealed"`

	// sealed subclasses
	// Required: true
	SealedSubclasses []*KClassResource `json:"sealedSubclasses"`

	// simple name
	SimpleName string `json:"simpleName,omitempty"`

	// supertypes
	// Required: true
	Supertypes []*KType `json:"supertypes"`

	// type parameters
	// Required: true
	TypeParameters []*KTypeParameter `json:"typeParameters"`

	// visibility
	// Enum: [PUBLIC PROTECTED INTERNAL PRIVATE]
	Visibility string `json:"visibility,omitempty"`
}
