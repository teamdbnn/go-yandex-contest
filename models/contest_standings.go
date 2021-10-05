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
)

// ContestStandings ContestStandings
//
// swagger:model ContestStandings
type ContestStandings struct {

	// rows
	Rows []*Row `json:"rows"`

	// titles
	Titles []*RowTitle `json:"titles"`
}

// Validate validates this contest standings
func (m *ContestStandings) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateRows(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTitles(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ContestStandings) validateRows(formats strfmt.Registry) error {
	if swag.IsZero(m.Rows) { // not required
		return nil
	}

	for i := 0; i < len(m.Rows); i++ {
		if swag.IsZero(m.Rows[i]) { // not required
			continue
		}

		if m.Rows[i] != nil {
			if err := m.Rows[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("rows" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ContestStandings) validateTitles(formats strfmt.Registry) error {
	if swag.IsZero(m.Titles) { // not required
		return nil
	}

	for i := 0; i < len(m.Titles); i++ {
		if swag.IsZero(m.Titles[i]) { // not required
			continue
		}

		if m.Titles[i] != nil {
			if err := m.Titles[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("titles" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this contest standings based on the context it is used
func (m *ContestStandings) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateRows(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTitles(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ContestStandings) contextValidateRows(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Rows); i++ {

		if m.Rows[i] != nil {
			if err := m.Rows[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("rows" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ContestStandings) contextValidateTitles(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Titles); i++ {

		if m.Titles[i] != nil {
			if err := m.Titles[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("titles" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ContestStandings) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ContestStandings) UnmarshalBinary(b []byte) error {
	var res ContestStandings
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}