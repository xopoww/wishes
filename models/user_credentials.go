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

// UserCredentials user credentials
//
// swagger:model UserCredentials
type UserCredentials struct {

	// password
	// Required: true
	// Format: password
	Password *strfmt.Password `json:"password"`

	// username
	// Required: true
	Username *UserName `json:"username"`
}

// Validate validates this user credentials
func (m *UserCredentials) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePassword(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUsername(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UserCredentials) validatePassword(formats strfmt.Registry) error {

	if err := validate.Required("password", "body", m.Password); err != nil {
		return err
	}

	if err := validate.FormatOf("password", "body", "password", m.Password.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *UserCredentials) validateUsername(formats strfmt.Registry) error {

	if err := validate.Required("username", "body", m.Username); err != nil {
		return err
	}

	if err := validate.Required("username", "body", m.Username); err != nil {
		return err
	}

	if m.Username != nil {
		if err := m.Username.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("username")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("username")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this user credentials based on the context it is used
func (m *UserCredentials) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateUsername(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UserCredentials) contextValidateUsername(ctx context.Context, formats strfmt.Registry) error {

	if m.Username != nil {
		if err := m.Username.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("username")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("username")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *UserCredentials) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UserCredentials) UnmarshalBinary(b []byte) error {
	var res UserCredentials
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}