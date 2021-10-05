// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// JSONNode JsonNode
//
// swagger:model JsonNode
type JSONNode struct {

	// array
	Array bool `json:"array,omitempty"`

	// big decimal
	BigDecimal bool `json:"bigDecimal,omitempty"`

	// big integer
	BigInteger bool `json:"bigInteger,omitempty"`

	// binary
	Binary bool `json:"binary,omitempty"`

	// boolean
	Boolean bool `json:"boolean,omitempty"`

	// container node
	ContainerNode bool `json:"containerNode,omitempty"`

	// double
	Double bool `json:"double,omitempty"`

	// empty
	Empty bool `json:"empty,omitempty"`

	// float
	Float bool `json:"float,omitempty"`

	// floating point number
	FloatingPointNumber bool `json:"floatingPointNumber,omitempty"`

	// int
	Int bool `json:"int,omitempty"`

	// integral number
	IntegralNumber bool `json:"integralNumber,omitempty"`

	// long
	Long bool `json:"long,omitempty"`

	// missing node
	MissingNode bool `json:"missingNode,omitempty"`

	// node type
	// Enum: [ARRAY BINARY BOOLEAN MISSING NULL NUMBER OBJECT POJO STRING]
	NodeType string `json:"nodeType,omitempty"`

	// null
	Null bool `json:"null,omitempty"`

	// number
	Number bool `json:"number,omitempty"`

	// object
	Object bool `json:"object,omitempty"`

	// pojo
	Pojo bool `json:"pojo,omitempty"`

	// short
	Short bool `json:"short,omitempty"`

	// textual
	Textual bool `json:"textual,omitempty"`

	// value node
	ValueNode bool `json:"valueNode,omitempty"`
}

// Validate validates this Json node
func (m *JSONNode) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateNodeType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var jsonNodeTypeNodeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["ARRAY","BINARY","BOOLEAN","MISSING","NULL","NUMBER","OBJECT","POJO","STRING"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		jsonNodeTypeNodeTypePropEnum = append(jsonNodeTypeNodeTypePropEnum, v)
	}
}

const (

	// JSONNodeNodeTypeARRAY captures enum value "ARRAY"
	JSONNodeNodeTypeARRAY string = "ARRAY"

	// JSONNodeNodeTypeBINARY captures enum value "BINARY"
	JSONNodeNodeTypeBINARY string = "BINARY"

	// JSONNodeNodeTypeBOOLEAN captures enum value "BOOLEAN"
	JSONNodeNodeTypeBOOLEAN string = "BOOLEAN"

	// JSONNodeNodeTypeMISSING captures enum value "MISSING"
	JSONNodeNodeTypeMISSING string = "MISSING"

	// JSONNodeNodeTypeNULL captures enum value "NULL"
	JSONNodeNodeTypeNULL string = "NULL"

	// JSONNodeNodeTypeNUMBER captures enum value "NUMBER"
	JSONNodeNodeTypeNUMBER string = "NUMBER"

	// JSONNodeNodeTypeOBJECT captures enum value "OBJECT"
	JSONNodeNodeTypeOBJECT string = "OBJECT"

	// JSONNodeNodeTypePOJO captures enum value "POJO"
	JSONNodeNodeTypePOJO string = "POJO"

	// JSONNodeNodeTypeSTRING captures enum value "STRING"
	JSONNodeNodeTypeSTRING string = "STRING"
)

// prop value enum
func (m *JSONNode) validateNodeTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, jsonNodeTypeNodeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *JSONNode) validateNodeType(formats strfmt.Registry) error {
	if swag.IsZero(m.NodeType) { // not required
		return nil
	}

	// value enum
	if err := m.validateNodeTypeEnum("nodeType", "body", m.NodeType); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this Json node based on context it is used
func (m *JSONNode) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *JSONNode) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *JSONNode) UnmarshalBinary(b []byte) error {
	var res JSONNode
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
