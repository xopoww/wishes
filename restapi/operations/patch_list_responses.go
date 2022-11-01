// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// PatchListOKCode is the HTTP code returned for type PatchListOK
const PatchListOKCode int = 200

/*
PatchListOK Success

swagger:response patchListOK
*/
type PatchListOK struct {
}

// NewPatchListOK creates PatchListOK with default headers values
func NewPatchListOK() *PatchListOK {

	return &PatchListOK{}
}

// WriteResponse to the client
func (o *PatchListOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(200)
}

// PatchListForbiddenCode is the HTTP code returned for type PatchListForbidden
const PatchListForbiddenCode int = 403

/*
PatchListForbidden Access denied

swagger:response patchListForbidden
*/
type PatchListForbidden struct {
}

// NewPatchListForbidden creates PatchListForbidden with default headers values
func NewPatchListForbidden() *PatchListForbidden {

	return &PatchListForbidden{}
}

// WriteResponse to the client
func (o *PatchListForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(403)
}

// PatchListNotFoundCode is the HTTP code returned for type PatchListNotFound
const PatchListNotFoundCode int = 404

/*
PatchListNotFound List not found

swagger:response patchListNotFound
*/
type PatchListNotFound struct {
}

// NewPatchListNotFound creates PatchListNotFound with default headers values
func NewPatchListNotFound() *PatchListNotFound {

	return &PatchListNotFound{}
}

// WriteResponse to the client
func (o *PatchListNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(404)
}

// PatchListInternalServerErrorCode is the HTTP code returned for type PatchListInternalServerError
const PatchListInternalServerErrorCode int = 500

/*
PatchListInternalServerError Server error

swagger:response patchListInternalServerError
*/
type PatchListInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *PatchListInternalServerErrorBody `json:"body,omitempty"`
}

// NewPatchListInternalServerError creates PatchListInternalServerError with default headers values
func NewPatchListInternalServerError() *PatchListInternalServerError {

	return &PatchListInternalServerError{}
}

// WithPayload adds the payload to the patch list internal server error response
func (o *PatchListInternalServerError) WithPayload(payload *PatchListInternalServerErrorBody) *PatchListInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the patch list internal server error response
func (o *PatchListInternalServerError) SetPayload(payload *PatchListInternalServerErrorBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PatchListInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
