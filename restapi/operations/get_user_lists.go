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

	"github.com/xopoww/wishes/models"
)

// GetUserListsHandlerFunc turns a function with the right signature into a get user lists handler
type GetUserListsHandlerFunc func(GetUserListsParams, *models.Principal) middleware.Responder

// Handle executing the request and returning a response
func (fn GetUserListsHandlerFunc) Handle(params GetUserListsParams, principal *models.Principal) middleware.Responder {
	return fn(params, principal)
}

// GetUserListsHandler interface for that can handle valid get user lists params
type GetUserListsHandler interface {
	Handle(GetUserListsParams, *models.Principal) middleware.Responder
}

// NewGetUserLists creates a new http.Handler for the get user lists operation
func NewGetUserLists(ctx *middleware.Context, handler GetUserListsHandler) *GetUserLists {
	return &GetUserLists{Context: ctx, Handler: handler}
}

/*
	GetUserLists swagger:route GET /users/{id}/lists getUserLists

Get user list ids (visible by client)
*/
type GetUserLists struct {
	Context *middleware.Context
	Handler GetUserListsHandler
}

func (o *GetUserLists) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewGetUserListsParams()
	uprinc, aCtx, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	if aCtx != nil {
		*r = *aCtx
	}
	var principal *models.Principal
	if uprinc != nil {
		principal = uprinc.(*models.Principal) // this is really a models.Principal, I promise
	}

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params, principal) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}

// GetUserListsInternalServerErrorBody get user lists internal server error body
//
// swagger:model GetUserListsInternalServerErrorBody
type GetUserListsInternalServerErrorBody struct {

	// error
	// Required: true
	Error *string `json:"error"`
}

// Validate validates this get user lists internal server error body
func (o *GetUserListsInternalServerErrorBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateError(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetUserListsInternalServerErrorBody) validateError(formats strfmt.Registry) error {

	if err := validate.Required("getUserListsInternalServerError"+"."+"error", "body", o.Error); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this get user lists internal server error body based on context it is used
func (o *GetUserListsInternalServerErrorBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetUserListsInternalServerErrorBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetUserListsInternalServerErrorBody) UnmarshalBinary(b []byte) error {
	var res GetUserListsInternalServerErrorBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
