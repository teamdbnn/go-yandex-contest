// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Team Team
//
// swagger:model Team
type Team struct {

	// id
	// Required: true
	ID *int64 `json:"id"`

	// invited
	Invited []*User `json:"invited"`

	// name
	Name string `json:"name,omitempty"`

	// teammates
	Teammates []*User `json:"teammates"`
}

// Validate validates this team
func (m *Team) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateInvited(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTeammates(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Team) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *Team) validateInvited(formats strfmt.Registry) error {
	if swag.IsZero(m.Invited) { // not required
		return nil
	}

	for i := 0; i < len(m.Invited); i++ {
		if swag.IsZero(m.Invited[i]) { // not required
			continue
		}

		if m.Invited[i] != nil {
			if err := m.Invited[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("invited" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Team) validateTeammates(formats strfmt.Registry) error {
	if swag.IsZero(m.Teammates) { // not required
		return nil
	}

	for i := 0; i < len(m.Teammates); i++ {
		if swag.IsZero(m.Teammates[i]) { // not required
			continue
		}

		if m.Teammates[i] != nil {
			if err := m.Teammates[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("teammates" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this team based on the context it is used
func (m *Team) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateInvited(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTeammates(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Team) contextValidateInvited(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Invited); i++ {

		if m.Invited[i] != nil {
			if err := m.Invited[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("invited" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Team) contextValidateTeammates(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Teammates); i++ {

		if m.Teammates[i] != nil {
			if err := m.Teammates[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("teammates" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *Team) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Team) UnmarshalBinary(b []byte) error {
	var res Team
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
