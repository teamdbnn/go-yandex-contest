// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// IDPrincipal Id«Principal»
//
// swagger:model Id«Principal»
type IDPrincipal struct {

	// checked Id
	// Required: true
	CheckedID *int64 `json:"checkedId"`

	// checked raw external Id
	// Required: true
	CheckedRawExternalID *string `json:"checkedRawExternalId"`

	// external Id
	// Required: true
	ExternalID *ExternalID `json:"externalId"`

	// id
	ID int64 `json:"id,omitempty"`

	// model
	// Required: true
	Model *KClassPrincipal `json:"model"`

	// new
	// Required: true
	New *bool `json:"new"`

	// raw external Id
	RawExternalID string `json:"rawExternalId,omitempty"`

	// uid
	// Format: uuid
	UID strfmt.UUID `json:"uid,omitempty"`
}

// Validate validates this Id principal
func (m *IDPrincipal) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCheckedID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCheckedRawExternalID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateExternalID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateModel(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNew(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IDPrincipal) validateCheckedID(formats strfmt.Registry) error {

	if err := validate.Required("checkedId", "body", m.CheckedID); err != nil {
		return err
	}

	return nil
}

func (m *IDPrincipal) validateCheckedRawExternalID(formats strfmt.Registry) error {

	if err := validate.Required("checkedRawExternalId", "body", m.CheckedRawExternalID); err != nil {
		return err
	}

	return nil
}

func (m *IDPrincipal) validateExternalID(formats strfmt.Registry) error {

	if err := validate.Required("externalId", "body", m.ExternalID); err != nil {
		return err
	}

	if m.ExternalID != nil {
		if err := m.ExternalID.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("externalId")
			}
			return err
		}
	}

	return nil
}

func (m *IDPrincipal) validateModel(formats strfmt.Registry) error {

	if err := validate.Required("model", "body", m.Model); err != nil {
		return err
	}

	if m.Model != nil {
		if err := m.Model.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("model")
			}
			return err
		}
	}

	return nil
}

func (m *IDPrincipal) validateNew(formats strfmt.Registry) error {

	if err := validate.Required("new", "body", m.New); err != nil {
		return err
	}

	return nil
}

func (m *IDPrincipal) validateUID(formats strfmt.Registry) error {
	if swag.IsZero(m.UID) { // not required
		return nil
	}

	if err := validate.FormatOf("uid", "body", "uuid", m.UID.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this Id principal based on the context it is used
func (m *IDPrincipal) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateExternalID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateModel(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IDPrincipal) contextValidateExternalID(ctx context.Context, formats strfmt.Registry) error {

	if m.ExternalID != nil {
		if err := m.ExternalID.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("externalId")
			}
			return err
		}
	}

	return nil
}

func (m *IDPrincipal) contextValidateModel(ctx context.Context, formats strfmt.Registry) error {

	if m.Model != nil {
		if err := m.Model.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("model")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *IDPrincipal) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IDPrincipal) UnmarshalBinary(b []byte) error {
	var res IDPrincipal
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
