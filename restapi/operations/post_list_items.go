// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"context"
	"net/http"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"

	"github.com/xopoww/wishes/restapi/apimodels"
)

// PostListItemsHandlerFunc turns a function with the right signature into a post list items handler
type PostListItemsHandlerFunc func(PostListItemsParams, *apimodels.Principal) middleware.Responder

// Handle executing the request and returning a response
func (fn PostListItemsHandlerFunc) Handle(params PostListItemsParams, principal *apimodels.Principal) middleware.Responder {
	return fn(params, principal)
}

// PostListItemsHandler interface for that can handle valid post list items params
type PostListItemsHandler interface {
	Handle(PostListItemsParams, *apimodels.Principal) middleware.Responder
}

// NewPostListItems creates a new http.Handler for the post list items operation
func NewPostListItems(ctx *middleware.Context, handler PostListItemsHandler) *PostListItems {
	return &PostListItems{Context: ctx, Handler: handler}
}

/*
	PostListItems swagger:route POST /lists/{id}/items postListItems

Add items to existing list
*/
type PostListItems struct {
	Context *middleware.Context
	Handler PostListItemsHandler
}

func (o *PostListItems) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewPostListItemsParams()
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

// PostListItemsBody post list items body
//
// swagger:model PostListItemsBody
type PostListItemsBody struct {
	apimodels.Revision

	// items
	// Required: true
	Items []*apimodels.ListItem `json:"items"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (o *PostListItemsBody) UnmarshalJSON(raw []byte) error {
	// PostListItemsParamsBodyAO0
	var postListItemsParamsBodyAO0 apimodels.Revision
	if err := swag.ReadJSON(raw, &postListItemsParamsBodyAO0); err != nil {
		return err
	}
	o.Revision = postListItemsParamsBodyAO0

	// PostListItemsParamsBodyAO1
	var dataPostListItemsParamsBodyAO1 struct {
		Items []*apimodels.ListItem `json:"items"`
	}
	if err := swag.ReadJSON(raw, &dataPostListItemsParamsBodyAO1); err != nil {
		return err
	}

	o.Items = dataPostListItemsParamsBodyAO1.Items

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (o PostListItemsBody) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	postListItemsParamsBodyAO0, err := swag.WriteJSON(o.Revision)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, postListItemsParamsBodyAO0)
	var dataPostListItemsParamsBodyAO1 struct {
		Items []*apimodels.ListItem `json:"items"`
	}

	dataPostListItemsParamsBodyAO1.Items = o.Items

	jsonDataPostListItemsParamsBodyAO1, errPostListItemsParamsBodyAO1 := swag.WriteJSON(dataPostListItemsParamsBodyAO1)
	if errPostListItemsParamsBodyAO1 != nil {
		return nil, errPostListItemsParamsBodyAO1
	}
	_parts = append(_parts, jsonDataPostListItemsParamsBodyAO1)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this post list items body
func (o *PostListItemsBody) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with apimodels.Revision
	if err := o.Revision.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateItems(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *PostListItemsBody) validateItems(formats strfmt.Registry) error {

	if err := validate.Required("items"+"."+"items", "body", o.Items); err != nil {
		return err
	}

	for i := 0; i < len(o.Items); i++ {
		if swag.IsZero(o.Items[i]) { // not required
			continue
		}

		if o.Items[i] != nil {
			if err := o.Items[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("items" + "." + "items" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("items" + "." + "items" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this post list items body based on the context it is used
func (o *PostListItemsBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with apimodels.Revision
	if err := o.Revision.ContextValidate(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateItems(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *PostListItemsBody) contextValidateItems(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.Items); i++ {

		if o.Items[i] != nil {
			if err := o.Items[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("items" + "." + "items" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("items" + "." + "items" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *PostListItemsBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostListItemsBody) UnmarshalBinary(b []byte) error {
	var res PostListItemsBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// PostListItemsInternalServerErrorBody post list items internal server error body
//
// swagger:model PostListItemsInternalServerErrorBody
type PostListItemsInternalServerErrorBody struct {

	// error
	// Required: true
	Error *string `json:"error"`
}

// Validate validates this post list items internal server error body
func (o *PostListItemsInternalServerErrorBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateError(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *PostListItemsInternalServerErrorBody) validateError(formats strfmt.Registry) error {

	if err := validate.Required("postListItemsInternalServerError"+"."+"error", "body", o.Error); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this post list items internal server error body based on context it is used
func (o *PostListItemsInternalServerErrorBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *PostListItemsInternalServerErrorBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostListItemsInternalServerErrorBody) UnmarshalBinary(b []byte) error {
	var res PostListItemsInternalServerErrorBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}