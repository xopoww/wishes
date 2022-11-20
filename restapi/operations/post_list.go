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

	"github.com/xopoww/wishes/restapi/apimodels"
)

// PostListHandlerFunc turns a function with the right signature into a post list handler
type PostListHandlerFunc func(PostListParams, *apimodels.Principal) middleware.Responder

// Handle executing the request and returning a response
func (fn PostListHandlerFunc) Handle(params PostListParams, principal *apimodels.Principal) middleware.Responder {
	return fn(params, principal)
}

// PostListHandler interface for that can handle valid post list params
type PostListHandler interface {
	Handle(PostListParams, *apimodels.Principal) middleware.Responder
}

// NewPostList creates a new http.Handler for the post list operation
func NewPostList(ctx *middleware.Context, handler PostListHandler) *PostList {
	return &PostList{Context: ctx, Handler: handler}
}

/*
	PostList swagger:route POST /lists Lists postList

Create new list
*/
type PostList struct {
	Context *middleware.Context
	Handler PostListHandler
}

func (o *PostList) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewPostListParams()
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

// PostListBody post list body
//
// swagger:model PostListBody
type PostListBody struct {
	apimodels.List

	// items
	Items []*apimodels.ListItem `json:"items"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (o *PostListBody) UnmarshalJSON(raw []byte) error {
	// PostListParamsBodyAO0
	var postListParamsBodyAO0 apimodels.List
	if err := swag.ReadJSON(raw, &postListParamsBodyAO0); err != nil {
		return err
	}
	o.List = postListParamsBodyAO0

	// PostListParamsBodyAO1
	var dataPostListParamsBodyAO1 struct {
		Items []*apimodels.ListItem `json:"items"`
	}
	if err := swag.ReadJSON(raw, &dataPostListParamsBodyAO1); err != nil {
		return err
	}

	o.Items = dataPostListParamsBodyAO1.Items

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (o PostListBody) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	postListParamsBodyAO0, err := swag.WriteJSON(o.List)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, postListParamsBodyAO0)
	var dataPostListParamsBodyAO1 struct {
		Items []*apimodels.ListItem `json:"items"`
	}

	dataPostListParamsBodyAO1.Items = o.Items

	jsonDataPostListParamsBodyAO1, errPostListParamsBodyAO1 := swag.WriteJSON(dataPostListParamsBodyAO1)
	if errPostListParamsBodyAO1 != nil {
		return nil, errPostListParamsBodyAO1
	}
	_parts = append(_parts, jsonDataPostListParamsBodyAO1)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this post list body
func (o *PostListBody) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with apimodels.List
	if err := o.List.Validate(formats); err != nil {
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

func (o *PostListBody) validateItems(formats strfmt.Registry) error {

	if swag.IsZero(o.Items) { // not required
		return nil
	}

	for i := 0; i < len(o.Items); i++ {
		if swag.IsZero(o.Items[i]) { // not required
			continue
		}

		if o.Items[i] != nil {
			if err := o.Items[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("list" + "." + "items" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("list" + "." + "items" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this post list body based on the context it is used
func (o *PostListBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with apimodels.List
	if err := o.List.ContextValidate(ctx, formats); err != nil {
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

func (o *PostListBody) contextValidateItems(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.Items); i++ {

		if o.Items[i] != nil {
			if err := o.Items[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("list" + "." + "items" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("list" + "." + "items" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *PostListBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostListBody) UnmarshalBinary(b []byte) error {
	var res PostListBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// PostListCreatedBody post list created body
//
// swagger:model PostListCreatedBody
type PostListCreatedBody struct {
	apimodels.Revision

	apimodels.ID
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (o *PostListCreatedBody) UnmarshalJSON(raw []byte) error {
	// PostListCreatedBodyAO0
	var postListCreatedBodyAO0 apimodels.Revision
	if err := swag.ReadJSON(raw, &postListCreatedBodyAO0); err != nil {
		return err
	}
	o.Revision = postListCreatedBodyAO0

	// PostListCreatedBodyAO1
	var postListCreatedBodyAO1 apimodels.ID
	if err := swag.ReadJSON(raw, &postListCreatedBodyAO1); err != nil {
		return err
	}
	o.ID = postListCreatedBodyAO1

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (o PostListCreatedBody) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	postListCreatedBodyAO0, err := swag.WriteJSON(o.Revision)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, postListCreatedBodyAO0)

	postListCreatedBodyAO1, err := swag.WriteJSON(o.ID)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, postListCreatedBodyAO1)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this post list created body
func (o *PostListCreatedBody) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with apimodels.Revision
	if err := o.Revision.Validate(formats); err != nil {
		res = append(res, err)
	}
	// validation for a type composition with apimodels.ID
	if err := o.ID.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validate this post list created body based on the context it is used
func (o *PostListCreatedBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with apimodels.Revision
	if err := o.Revision.ContextValidate(ctx, formats); err != nil {
		res = append(res, err)
	}
	// validation for a type composition with apimodels.ID
	if err := o.ID.ContextValidate(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (o *PostListCreatedBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostListCreatedBody) UnmarshalBinary(b []byte) error {
	var res PostListCreatedBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
