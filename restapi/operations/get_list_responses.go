// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/xopoww/wishes/restapi/apimodels"
)

// GetListOKCode is the HTTP code returned for type GetListOK
const GetListOKCode int = 200

/*
GetListOK Success

swagger:response getListOK
*/
type GetListOK struct {

	/*
	  In: Body
	*/
	Payload *apimodels.List `json:"body,omitempty"`
}

// NewGetListOK creates GetListOK with default headers values
func NewGetListOK() *GetListOK {

	return &GetListOK{}
}

// WithPayload adds the payload to the get list o k response
func (o *GetListOK) WithPayload(payload *apimodels.List) *GetListOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get list o k response
func (o *GetListOK) SetPayload(payload *apimodels.List) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetListOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetListForbiddenCode is the HTTP code returned for type GetListForbidden
const GetListForbiddenCode int = 403

/*
GetListForbidden Access denied

swagger:response getListForbidden
*/
type GetListForbidden struct {
}

// NewGetListForbidden creates GetListForbidden with default headers values
func NewGetListForbidden() *GetListForbidden {

	return &GetListForbidden{}
}

// WriteResponse to the client
func (o *GetListForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(403)
}

// GetListNotFoundCode is the HTTP code returned for type GetListNotFound
const GetListNotFoundCode int = 404

/*
GetListNotFound List not found

swagger:response getListNotFound
*/
type GetListNotFound struct {
}

// NewGetListNotFound creates GetListNotFound with default headers values
func NewGetListNotFound() *GetListNotFound {

	return &GetListNotFound{}
}

// WriteResponse to the client
func (o *GetListNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(404)
}

// GetListInternalServerErrorCode is the HTTP code returned for type GetListInternalServerError
const GetListInternalServerErrorCode int = 500

/*
GetListInternalServerError Server error

swagger:response getListInternalServerError
*/
type GetListInternalServerError struct {
}

// NewGetListInternalServerError creates GetListInternalServerError with default headers values
func NewGetListInternalServerError() *GetListInternalServerError {

	return &GetListInternalServerError{}
}

// WriteResponse to the client
func (o *GetListInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(500)
}
