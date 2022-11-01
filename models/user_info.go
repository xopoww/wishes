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

// UserInfo user info
//
// swagger:model UserInfo
type UserInfo struct {

	// fname
	// Required: true
	Fname *string `json:"fname"`

	// lname
	// Required: true
	Lname *string `json:"lname"`
}

// Validate validates this user info
func (m *UserInfo) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFname(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLname(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UserInfo) validateFname(formats strfmt.Registry) error {

	if err := validate.Required("fname", "body", m.Fname); err != nil {
		return err
	}

	return nil
}

func (m *UserInfo) validateLname(formats strfmt.Registry) error {

	if err := validate.Required("lname", "body", m.Lname); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this user info based on context it is used
func (m *UserInfo) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *UserInfo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UserInfo) UnmarshalBinary(b []byte) error {
	var res UserInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
