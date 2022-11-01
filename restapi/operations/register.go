// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"context"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"

	"github.com/xopoww/wishes/internal/models"
)

// RegisterHandlerFunc turns a function with the right signature into a register handler
type RegisterHandlerFunc func(RegisterParams) middleware.Responder

// Handle executing the request and returning a response
func (fn RegisterHandlerFunc) Handle(params RegisterParams) middleware.Responder {
	return fn(params)
}

// RegisterHandler interface for that can handle valid register params
type RegisterHandler interface {
	Handle(RegisterParams) middleware.Responder
}

// NewRegister creates a new http.Handler for the register operation
func NewRegister(ctx *middleware.Context, handler RegisterHandler) *Register {
	return &Register{Context: ctx, Handler: handler}
}

/*
	Register swagger:route POST /users register

Register new user
*/
type Register struct {
	Context *middleware.Context
	Handler RegisterHandler
}

func (o *Register) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewRegisterParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}

// RegisterInternalServerErrorBody register internal server error body
//
// swagger:model RegisterInternalServerErrorBody
type RegisterInternalServerErrorBody struct {

	// error
	// Required: true
	Error *string `json:"error"`
}

// Validate validates this register internal server error body
func (o *RegisterInternalServerErrorBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateError(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *RegisterInternalServerErrorBody) validateError(formats strfmt.Registry) error {

	if err := validate.Required("registerInternalServerError"+"."+"error", "body", o.Error); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this register internal server error body based on context it is used
func (o *RegisterInternalServerErrorBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *RegisterInternalServerErrorBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *RegisterInternalServerErrorBody) UnmarshalBinary(b []byte) error {
	var res RegisterInternalServerErrorBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// RegisterOKBody register o k body
//
// swagger:model RegisterOKBody
type RegisterOKBody struct {

	// error
	Error string `json:"error,omitempty"`

	// ok
	// Required: true
	Ok *bool `json:"ok"`

	// user
	User *models.ID `json:"user,omitempty"`
}

// Validate validates this register o k body
func (o *RegisterOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateOk(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateUser(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *RegisterOKBody) validateOk(formats strfmt.Registry) error {

	if err := validate.Required("registerOK"+"."+"ok", "body", o.Ok); err != nil {
		return err
	}

	return nil
}

func (o *RegisterOKBody) validateUser(formats strfmt.Registry) error {
	if swag.IsZero(o.User) { // not required
		return nil
	}

	if o.User != nil {
		if err := o.User.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("registerOK" + "." + "user")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("registerOK" + "." + "user")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this register o k body based on the context it is used
func (o *RegisterOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateUser(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *RegisterOKBody) contextValidateUser(ctx context.Context, formats strfmt.Registry) error {

	if o.User != nil {
		if err := o.User.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("registerOK" + "." + "user")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("registerOK" + "." + "user")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *RegisterOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *RegisterOKBody) UnmarshalBinary(b []byte) error {
	var res RegisterOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}