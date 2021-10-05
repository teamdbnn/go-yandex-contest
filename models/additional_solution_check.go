// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"bytes"
	"context"
	"encoding/json"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// AdditionalSolutionCheck AdditionalSolutionCheck
//
// swagger:model AdditionalSolutionCheck
type AdditionalSolutionCheck struct {
	checkerFilesField FileSystemPackage

	// compiler Id
	CompilerID string `json:"compilerId,omitempty"`

	// main filename
	MainFilename string `json:"mainFilename,omitempty"`
}

// CheckerFiles gets the checker files of this base type
func (m *AdditionalSolutionCheck) CheckerFiles() FileSystemPackage {
	return m.checkerFilesField
}

// SetCheckerFiles sets the checker files of this base type
func (m *AdditionalSolutionCheck) SetCheckerFiles(val FileSystemPackage) {
	m.checkerFilesField = val
}

// UnmarshalJSON unmarshals this object with a polymorphic type from a JSON structure
func (m *AdditionalSolutionCheck) UnmarshalJSON(raw []byte) error {
	var data struct {
		CheckerFiles json.RawMessage `json:"checkerFiles,omitempty"`

		CompilerID string `json:"compilerId,omitempty"`

		MainFilename string `json:"mainFilename,omitempty"`
	}
	buf := bytes.NewBuffer(raw)
	dec := json.NewDecoder(buf)
	dec.UseNumber()

	if err := dec.Decode(&data); err != nil {
		return err
	}

	var propCheckerFiles FileSystemPackage
	if string(data.CheckerFiles) != "null" {
		checkerFiles, err := UnmarshalFileSystemPackage(bytes.NewBuffer(data.CheckerFiles), runtime.JSONConsumer())
		if err != nil && err != io.EOF {
			return err
		}
		propCheckerFiles = checkerFiles
	}

	var result AdditionalSolutionCheck

	// checkerFiles
	result.checkerFilesField = propCheckerFiles

	// compilerId
	result.CompilerID = data.CompilerID

	// mainFilename
	result.MainFilename = data.MainFilename

	*m = result

	return nil
}

// MarshalJSON marshals this object with a polymorphic type to a JSON structure
func (m AdditionalSolutionCheck) MarshalJSON() ([]byte, error) {
	var b1, b2, b3 []byte
	var err error
	b1, err = json.Marshal(struct {
		CompilerID string `json:"compilerId,omitempty"`

		MainFilename string `json:"mainFilename,omitempty"`
	}{

		CompilerID: m.CompilerID,

		MainFilename: m.MainFilename,
	})
	if err != nil {
		return nil, err
	}
	b2, err = json.Marshal(struct {
		CheckerFiles FileSystemPackage `json:"checkerFiles,omitempty"`
	}{

		CheckerFiles: m.checkerFilesField,
	})
	if err != nil {
		return nil, err
	}

	return swag.ConcatJSON(b1, b2, b3), nil
}

// Validate validates this additional solution check
func (m *AdditionalSolutionCheck) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCheckerFiles(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AdditionalSolutionCheck) validateCheckerFiles(formats strfmt.Registry) error {
	if swag.IsZero(m.CheckerFiles()) { // not required
		return nil
	}

	if err := m.CheckerFiles().Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("checkerFiles")
		}
		return err
	}

	return nil
}

// ContextValidate validate this additional solution check based on the context it is used
func (m *AdditionalSolutionCheck) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCheckerFiles(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AdditionalSolutionCheck) contextValidateCheckerFiles(ctx context.Context, formats strfmt.Registry) error {

	if err := m.CheckerFiles().ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("checkerFiles")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AdditionalSolutionCheck) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AdditionalSolutionCheck) UnmarshalBinary(b []byte) error {
	var res AdditionalSolutionCheck
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}