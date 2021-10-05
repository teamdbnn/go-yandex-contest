package models

type KClassObject struct {
	Abstract         *bool              `json:"abstract"`
	Annotations      []Annotation       `json:"annotations"`
	Companion        *bool              `json:"companion"`
	Constructors     []*KFunctionObject `json:"constructors"`
	Data             *bool              `json:"data"`
	Final            *bool              `json:"final"`
	Inner            *bool              `json:"inner"`
	Members          []*KCallableObject `json:"members"`
	NestedClasses    []*KClassObject    `json:"nestedClasses"`
	ObjectInstance   interface{}        `json:"objectInstance,omitempty"`
	Open             *bool              `json:"open"`
	QualifiedName    string             `json:"qualifiedName,omitempty"`
	Sealed           *bool              `json:"sealed"`
	SealedSubclasses []*KClassObject    `json:"sealedSubclasses"`
	SimpleName       string             `json:"simpleName,omitempty"`
	Supertypes       []*KType           `json:"supertypes"`
	TypeParameters   []*KTypeParameter  `json:"typeParameters"`
	Visibility       string             `json:"visibility,omitempty"`
}
