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

// GetListTokenHandlerFunc turns a function with the right signature into a get list token handler
type GetListTokenHandlerFunc func(GetListTokenParams, *apimodels.Principal) middleware.Responder

// Handle executing the request and returning a response
func (fn GetListTokenHandlerFunc) Handle(params GetListTokenParams, principal *apimodels.Principal) middleware.Responder {
	return fn(params, principal)
}

// GetListTokenHandler interface for that can handle valid get list token params
type GetListTokenHandler interface {
	Handle(GetListTokenParams, *apimodels.Principal) middleware.Responder
}

// NewGetListToken creates a new http.Handler for the get list token operation
func NewGetListToken(ctx *middleware.Context, handler GetListTokenHandler) *GetListToken {
	return &GetListToken{Context: ctx, Handler: handler}
}

/*
	GetListToken swagger:route GET /lists/{id}/token getListToken

Get access token for a list
*/
type GetListToken struct {
	Context *middleware.Context
	Handler GetListTokenHandler
}

func (o *GetListToken) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewGetListTokenParams()
	uprinc, aCtx, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	if aCtx != nil {
		*r = *aCtx
	}
	var principal *apimodels.Principal
	if uprinc != nil {
		principal = uprinc.(*apimodels.Principal) // this is really a apimodels.Principal, I promise
	}

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params, principal) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}

// GetListTokenInternalServerErrorBody get list token internal server error body
//
// swagger:model GetListTokenInternalServerErrorBody
type GetListTokenInternalServerErrorBody struct {

	// error
	// Required: true
	Error *string `json:"error"`
}

// Validate validates this get list token internal server error body
func (o *GetListTokenInternalServerErrorBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateError(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetListTokenInternalServerErrorBody) validateError(formats strfmt.Registry) error {

	if err := validate.Required("getListTokenInternalServerError"+"."+"error", "body", o.Error); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this get list token internal server error body based on context it is used
func (o *GetListTokenInternalServerErrorBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetListTokenInternalServerErrorBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetListTokenInternalServerErrorBody) UnmarshalBinary(b []byte) error {
	var res GetListTokenInternalServerErrorBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// GetListTokenOKBody get list token o k body
//
// swagger:model GetListTokenOKBody
type GetListTokenOKBody struct {

	// token
	Token string `json:"token,omitempty"`
}

// Validate validates this get list token o k body
func (o *GetListTokenOKBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this get list token o k body based on context it is used
func (o *GetListTokenOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetListTokenOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetListTokenOKBody) UnmarshalBinary(b []byte) error {
	var res GetListTokenOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}