// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/xopoww/wishes/restapi/apimodels"
)

// GetListItemsOKCode is the HTTP code returned for type GetListItemsOK
const GetListItemsOKCode int = 200

/*
GetListItemsOK Success

swagger:response getListItemsOK
*/
type GetListItemsOK struct {

	/*
	  In: Body
	*/
	Payload *apimodels.ListItems `json:"body,omitempty"`
}

// NewGetListItemsOK creates GetListItemsOK with default headers values
func NewGetListItemsOK() *GetListItemsOK {

	return &GetListItemsOK{}
}

// WithPayload adds the payload to the get list items o k response
func (o *GetListItemsOK) WithPayload(payload *apimodels.ListItems) *GetListItemsOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get list items o k response
func (o *GetListItemsOK) SetPayload(payload *apimodels.ListItems) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetListItemsOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetListItemsForbiddenCode is the HTTP code returned for type GetListItemsForbidden
const GetListItemsForbiddenCode int = 403

/*
GetListItemsForbidden Access denied

swagger:response getListItemsForbidden
*/
type GetListItemsForbidden struct {
}

// NewGetListItemsForbidden creates GetListItemsForbidden with default headers values
func NewGetListItemsForbidden() *GetListItemsForbidden {

	return &GetListItemsForbidden{}
}

// WriteResponse to the client
func (o *GetListItemsForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(403)
}

// GetListItemsNotFoundCode is the HTTP code returned for type GetListItemsNotFound
const GetListItemsNotFoundCode int = 404

/*
GetListItemsNotFound List not found

swagger:response getListItemsNotFound
*/
type GetListItemsNotFound struct {
}

// NewGetListItemsNotFound creates GetListItemsNotFound with default headers values
func NewGetListItemsNotFound() *GetListItemsNotFound {

	return &GetListItemsNotFound{}
}

// WriteResponse to the client
func (o *GetListItemsNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(404)
}

// GetListItemsInternalServerErrorCode is the HTTP code returned for type GetListItemsInternalServerError
const GetListItemsInternalServerErrorCode int = 500

/*
GetListItemsInternalServerError Server error

swagger:response getListItemsInternalServerError
*/
type GetListItemsInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *GetListItemsInternalServerErrorBody `json:"body,omitempty"`
}

// NewGetListItemsInternalServerError creates GetListItemsInternalServerError with default headers values
func NewGetListItemsInternalServerError() *GetListItemsInternalServerError {

	return &GetListItemsInternalServerError{}
}

// WithPayload adds the payload to the get list items internal server error response
func (o *GetListItemsInternalServerError) WithPayload(payload *GetListItemsInternalServerErrorBody) *GetListItemsInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get list items internal server error response
func (o *GetListItemsInternalServerError) SetPayload(payload *GetListItemsInternalServerErrorBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetListItemsInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}