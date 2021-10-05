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

// KParameter KParameter
//
// swagger:model KParameter
type KParameter struct {

	// annotations
	// Required: true
	Annotations []Annotation `json:"annotations"`

	// index
	// Required: true
	Index *int32 `json:"index"`

	// kind
	// Required: true
	// Enum: [INSTANCE EXTENSION_RECEIVER VALUE]
	Kind *string `json:"kind"`

	// name
	Name string `json:"name,omitempty"`

	// optional
	// Required: true
	Optional *bool `json:"optional"`

	// type
	// Required: true
	Type *KType `json:"type"`

	// vararg
	// Required: true
	Vararg *bool `json:"vararg"`
}

// Validate validates this k parameter
func (m *KParameter) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAnnotations(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIndex(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateKind(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOptional(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVararg(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *KParameter) validateAnnotations(formats strfmt.Registry) error {

	if err := validate.Required("annotations", "body", m.Annotations); err != nil {
		return err
	}

	return nil
}

func (m *KParameter) validateIndex(formats strfmt.Registry) error {

	if err := validate.Required("index", "body", m.Index); err != nil {
		return err
	}

	return nil
}

var kParameterTypeKindPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["INSTANCE","EXTENSION_RECEIVER","VALUE"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		kParameterTypeKindPropEnum = append(kParameterTypeKindPropEnum, v)
	}
}

const (

	// KParameterKindINSTANCE captures enum value "INSTANCE"
	KParameterKindINSTANCE string = "INSTANCE"

	// KParameterKindEXTENSIONRECEIVER captures enum value "EXTENSION_RECEIVER"
	KParameterKindEXTENSIONRECEIVER string = "EXTENSION_RECEIVER"

	// KParameterKindVALUE captures enum value "VALUE"
	KParameterKindVALUE string = "VALUE"
)

// prop value enum
func (m *KParameter) validateKindEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, kParameterTypeKindPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *KParameter) validateKind(formats strfmt.Registry) error {

	if err := validate.Required("kind", "body", m.Kind); err != nil {
		return err
	}

	// value enum
	if err := m.validateKindEnum("kind", "body", *m.Kind); err != nil {
		return err
	}

	return nil
}

func (m *KParameter) validateOptional(formats strfmt.Registry) error {

	if err := validate.Required("optional", "body", m.Optional); err != nil {
		return err
	}

	return nil
}

func (m *KParameter) validateType(formats strfmt.Registry) error {

	if err := validate.Required("type", "body", m.Type); err != nil {
		return err
	}

	if m.Type != nil {
		if err := m.Type.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("type")
			}
			return err
		}
	}

	return nil
}

func (m *KParameter) validateVararg(formats strfmt.Registry) error {

	if err := validate.Required("vararg", "body", m.Vararg); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this k parameter based on the context it is used
func (m *KParameter) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *KParameter) contextValidateType(ctx context.Context, formats strfmt.Registry) error {

	if m.Type != nil {
		if err := m.Type.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("type")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *KParameter) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *KParameter) UnmarshalBinary(b []byte) error {
	var res KParameter
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
