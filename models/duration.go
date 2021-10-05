// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// Duration Duration
//
// swagger:model Duration
type Duration struct {

	// millis
	Millis int64 `json:"millis,omitempty"`

	// standard days
	StandardDays int64 `json:"standardDays,omitempty"`

	// standard hours
	StandardHours int64 `json:"standardHours,omitempty"`

	// standard minutes
	StandardMinutes int64 `json:"standardMinutes,omitempty"`

	// standard seconds
	StandardSeconds int64 `json:"standardSeconds,omitempty"`
}

// Validate validates this duration
func (m *Duration) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this duration based on context it is used
func (m *Duration) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Duration) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Duration) UnmarshalBinary(b []byte) error {
	var res Duration
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}