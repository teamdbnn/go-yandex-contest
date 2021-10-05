// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"bytes"
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// FilePackage FilePackage
//
// swagger:model FilePackage
type FilePackage struct {
	nameField string

	FilePackageAllOf1
}

// Name gets the name of this subtype
func (m *FilePackage) Name() string {
	return m.nameField
}

// SetName sets the name of this subtype
func (m *FilePackage) SetName(val string) {
	m.nameField = val
}

// UnmarshalJSON unmarshals this object with a polymorphic type from a JSON structure
func (m *FilePackage) UnmarshalJSON(raw []byte) error {
	var data struct {
		FilePackageAllOf1
	}
	buf := bytes.NewBuffer(raw)
	dec := json.NewDecoder(buf)
	dec.UseNumber()

	if err := dec.Decode(&data); err != nil {
		return err
	}

	var base struct {
		/* Just the base type fields. Used for unmashalling polymorphic types.*/

		Name string `json:"name,omitempty"`
	}
	buf = bytes.NewBuffer(raw)
	dec = json.NewDecoder(buf)
	dec.UseNumber()

	if err := dec.Decode(&base); err != nil {
		return err
	}

	var result FilePackage

	result.nameField = base.Name

	result.FilePackageAllOf1 = data.FilePackageAllOf1

	*m = result

	return nil
}

// MarshalJSON marshals this object with a polymorphic type to a JSON structure
func (m FilePackage) MarshalJSON() ([]byte, error) {
	var b1, b2, b3 []byte
	var err error
	b1, err = json.Marshal(struct {
		FilePackageAllOf1
	}{

		FilePackageAllOf1: m.FilePackageAllOf1,
	})
	if err != nil {
		return nil, err
	}
	b2, err = json.Marshal(struct {
		Name string `json:"name,omitempty"`
	}{

		Name: m.Name(),
	})
	if err != nil {
		return nil, err
	}

	return swag.ConcatJSON(b1, b2, b3), nil
}

// Validate validates this file package
func (m *FilePackage) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with FilePackageAllOf1
	if err := m.FilePackageAllOf1.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validate this file package based on the context it is used
func (m *FilePackage) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with FilePackageAllOf1
	if err := m.FilePackageAllOf1.ContextValidate(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *FilePackage) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FilePackage) UnmarshalBinary(b []byte) error {
	var res FilePackage
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}