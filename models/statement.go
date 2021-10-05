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

// Statement Statement
//
// swagger:model Statement
type Statement struct {

	// path
	Path string `json:"path,omitempty"`

	// type
	// Enum: [TEX PDF BINARY OLYMP MARKDOWN]
	Type string `json:"type,omitempty"`
}

// Validate validates this statement
func (m *Statement) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var statementTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["TEX","PDF","BINARY","OLYMP","MARKDOWN"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		statementTypeTypePropEnum = append(statementTypeTypePropEnum, v)
	}
}

const (

	// StatementTypeTEX captures enum value "TEX"
	StatementTypeTEX string = "TEX"

	// StatementTypePDF captures enum value "PDF"
	StatementTypePDF string = "PDF"

	// StatementTypeBINARY captures enum value "BINARY"
	StatementTypeBINARY string = "BINARY"

	// StatementTypeOLYMP captures enum value "OLYMP"
	StatementTypeOLYMP string = "OLYMP"

	// StatementTypeMARKDOWN captures enum value "MARKDOWN"
	StatementTypeMARKDOWN string = "MARKDOWN"
)

// prop value enum
func (m *Statement) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, statementTypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *Statement) validateType(formats strfmt.Registry) error {
	if swag.IsZero(m.Type) { // not required
		return nil
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this statement based on context it is used
func (m *Statement) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Statement) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Statement) UnmarshalBinary(b []byte) error {
	var res Statement
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
