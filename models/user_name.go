// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
)

// UserName user name
//
// swagger:model UserName
type UserName string

// Validate validates this user name
func (m UserName) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this user name based on context it is used
func (m UserName) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
