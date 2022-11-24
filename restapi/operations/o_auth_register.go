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

	"github.com/xopoww/wishes/restapi/apimodels"
)

// OAuthRegisterHandlerFunc turns a function with the right signature into a o auth register handler
type OAuthRegisterHandlerFunc func(OAuthRegisterParams) middleware.Responder

// Handle executing the request and returning a response
func (fn OAuthRegisterHandlerFunc) Handle(params OAuthRegisterParams) middleware.Responder {
	return fn(params)
}

// OAuthRegisterHandler interface for that can handle valid o auth register params
type OAuthRegisterHandler interface {
	Handle(OAuthRegisterParams) middleware.Responder
}

// NewOAuthRegister creates a new http.Handler for the o auth register operation
func NewOAuthRegister(ctx *middleware.Context, handler OAuthRegisterHandler) *OAuthRegister {
	return &OAuthRegister{Context: ctx, Handler: handler}
}

/*
	OAuthRegister swagger:route POST /oauth/register Auth oAuthRegister

Register new OAuth user
*/
type OAuthRegister struct {
	Context *middleware.Context
	Handler OAuthRegisterHandler
}

func (o *OAuthRegister) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewOAuthRegisterParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}

// OAuthRegisterBody o auth register body
//
// swagger:model OAuthRegisterBody
type OAuthRegisterBody struct {

	// username
	// Required: true
	Username *apimodels.UserName `json:"username"`

	apimodels.OAuthCredentials
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (o *OAuthRegisterBody) UnmarshalJSON(raw []byte) error {
	// OAuthRegisterParamsBodyAO0
	var dataOAuthRegisterParamsBodyAO0 struct {
		Username *apimodels.UserName `json:"username"`
	}
	if err := swag.ReadJSON(raw, &dataOAuthRegisterParamsBodyAO0); err != nil {
		return err
	}

	o.Username = dataOAuthRegisterParamsBodyAO0.Username

	// OAuthRegisterParamsBodyAO1
	var oAuthRegisterParamsBodyAO1 apimodels.OAuthCredentials
	if err := swag.ReadJSON(raw, &oAuthRegisterParamsBodyAO1); err != nil {
		return err
	}
	o.OAuthCredentials = oAuthRegisterParamsBodyAO1

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (o OAuthRegisterBody) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	var dataOAuthRegisterParamsBodyAO0 struct {
		Username *apimodels.UserName `json:"username"`
	}

	dataOAuthRegisterParamsBodyAO0.Username = o.Username

	jsonDataOAuthRegisterParamsBodyAO0, errOAuthRegisterParamsBodyAO0 := swag.WriteJSON(dataOAuthRegisterParamsBodyAO0)
	if errOAuthRegisterParamsBodyAO0 != nil {
		return nil, errOAuthRegisterParamsBodyAO0
	}
	_parts = append(_parts, jsonDataOAuthRegisterParamsBodyAO0)

	oAuthRegisterParamsBodyAO1, err := swag.WriteJSON(o.OAuthCredentials)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, oAuthRegisterParamsBodyAO1)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this o auth register body
func (o *OAuthRegisterBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateUsername(formats); err != nil {
		res = append(res, err)
	}

	// validation for a type composition with apimodels.OAuthCredentials
	if err := o.OAuthCredentials.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *OAuthRegisterBody) validateUsername(formats strfmt.Registry) error {

	if err := validate.Required("body"+"."+"username", "body", o.Username); err != nil {
		return err
	}

	if err := validate.Required("body"+"."+"username", "body", o.Username); err != nil {
		return err
	}

	if o.Username != nil {
		if err := o.Username.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("body" + "." + "username")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("body" + "." + "username")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this o auth register body based on the context it is used
func (o *OAuthRegisterBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateUsername(ctx, formats); err != nil {
		res = append(res, err)
	}

	// validation for a type composition with apimodels.OAuthCredentials
	if err := o.OAuthCredentials.ContextValidate(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *OAuthRegisterBody) contextValidateUsername(ctx context.Context, formats strfmt.Registry) error {

	if o.Username != nil {
		if err := o.Username.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("body" + "." + "username")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("body" + "." + "username")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *OAuthRegisterBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *OAuthRegisterBody) UnmarshalBinary(b []byte) error {
	var res OAuthRegisterBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// OAuthRegisterOKBody o auth register o k body
//
// swagger:model OAuthRegisterOKBody
type OAuthRegisterOKBody struct {

	// error
	Error string `json:"error,omitempty"`

	// ok
	// Required: true
	Ok *bool `json:"ok"`

	// user
	User *apimodels.ID `json:"user,omitempty"`
}

// Validate validates this o auth register o k body
func (o *OAuthRegisterOKBody) Validate(formats strfmt.Registry) error {
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

func (o *OAuthRegisterOKBody) validateOk(formats strfmt.Registry) error {

	if err := validate.Required("oAuthRegisterOK"+"."+"ok", "body", o.Ok); err != nil {
		return err
	}

	return nil
}

func (o *OAuthRegisterOKBody) validateUser(formats strfmt.Registry) error {
	if swag.IsZero(o.User) { // not required
		return nil
	}

	if o.User != nil {
		if err := o.User.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("oAuthRegisterOK" + "." + "user")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("oAuthRegisterOK" + "." + "user")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this o auth register o k body based on the context it is used
func (o *OAuthRegisterOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateUser(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *OAuthRegisterOKBody) contextValidateUser(ctx context.Context, formats strfmt.Registry) error {

	if o.User != nil {
		if err := o.User.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("oAuthRegisterOK" + "." + "user")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("oAuthRegisterOK" + "." + "user")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *OAuthRegisterOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *OAuthRegisterOKBody) UnmarshalBinary(b []byte) error {
	var res OAuthRegisterOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}