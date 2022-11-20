// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"

	"github.com/xopoww/wishes/restapi/apimodels"
)

// GetUserListsHandlerFunc turns a function with the right signature into a get user lists handler
type GetUserListsHandlerFunc func(GetUserListsParams, *apimodels.Principal) middleware.Responder

// Handle executing the request and returning a response
func (fn GetUserListsHandlerFunc) Handle(params GetUserListsParams, principal *apimodels.Principal) middleware.Responder {
	return fn(params, principal)
}

// GetUserListsHandler interface for that can handle valid get user lists params
type GetUserListsHandler interface {
	Handle(GetUserListsParams, *apimodels.Principal) middleware.Responder
}

// NewGetUserLists creates a new http.Handler for the get user lists operation
func NewGetUserLists(ctx *middleware.Context, handler GetUserListsHandler) *GetUserLists {
	return &GetUserLists{Context: ctx, Handler: handler}
}

/*
	GetUserLists swagger:route GET /lists Lists getUserLists

Get user list IDs (visible by client)
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
