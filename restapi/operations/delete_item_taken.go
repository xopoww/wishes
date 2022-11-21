// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"

	"github.com/xopoww/wishes/restapi/apimodels"
)

// DeleteItemTakenHandlerFunc turns a function with the right signature into a delete item taken handler
type DeleteItemTakenHandlerFunc func(DeleteItemTakenParams, *apimodels.Principal) middleware.Responder

// Handle executing the request and returning a response
func (fn DeleteItemTakenHandlerFunc) Handle(params DeleteItemTakenParams, principal *apimodels.Principal) middleware.Responder {
	return fn(params, principal)
}

// DeleteItemTakenHandler interface for that can handle valid delete item taken params
type DeleteItemTakenHandler interface {
	Handle(DeleteItemTakenParams, *apimodels.Principal) middleware.Responder
}

// NewDeleteItemTaken creates a new http.Handler for the delete item taken operation
func NewDeleteItemTaken(ctx *middleware.Context, handler DeleteItemTakenHandler) *DeleteItemTaken {
	return &DeleteItemTaken{Context: ctx, Handler: handler}
}

/*
	DeleteItemTaken swagger:route DELETE /lists/{id}/items/{item_id}/taken_by Items deleteItemTaken

Unmark previously taken item
*/
type DeleteItemTaken struct {
	Context *middleware.Context
	Handler DeleteItemTakenHandler
}

func (o *DeleteItemTaken) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewDeleteItemTakenParams()
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

// DeleteItemTakenConflictBody delete item taken conflict body
//
// swagger:model DeleteItemTakenConflictBody
type DeleteItemTakenConflictBody struct {

	// reason
	// Required: true
	// Enum: [outdated revision not taken]
	Reason *string `json:"reason"`
}

// Validate validates this delete item taken conflict body
func (o *DeleteItemTakenConflictBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateReason(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var deleteItemTakenConflictBodyTypeReasonPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["outdated revision","not taken"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		deleteItemTakenConflictBodyTypeReasonPropEnum = append(deleteItemTakenConflictBodyTypeReasonPropEnum, v)
	}
}

const (

	// DeleteItemTakenConflictBodyReasonOutdatedRevision captures enum value "outdated revision"
	DeleteItemTakenConflictBodyReasonOutdatedRevision string = "outdated revision"

	// DeleteItemTakenConflictBodyReasonNotTaken captures enum value "not taken"
	DeleteItemTakenConflictBodyReasonNotTaken string = "not taken"
)

// prop value enum
func (o *DeleteItemTakenConflictBody) validateReasonEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, deleteItemTakenConflictBodyTypeReasonPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *DeleteItemTakenConflictBody) validateReason(formats strfmt.Registry) error {

	if err := validate.Required("deleteItemTakenConflict"+"."+"reason", "body", o.Reason); err != nil {
		return err
	}

	// value enum
	if err := o.validateReasonEnum("deleteItemTakenConflict"+"."+"reason", "body", *o.Reason); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this delete item taken conflict body based on context it is used
func (o *DeleteItemTakenConflictBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *DeleteItemTakenConflictBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *DeleteItemTakenConflictBody) UnmarshalBinary(b []byte) error {
	var res DeleteItemTakenConflictBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}