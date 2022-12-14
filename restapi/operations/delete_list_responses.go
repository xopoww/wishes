// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// DeleteListNoContentCode is the HTTP code returned for type DeleteListNoContent
const DeleteListNoContentCode int = 204

/*
DeleteListNoContent Success

swagger:response deleteListNoContent
*/
type DeleteListNoContent struct {
}

// NewDeleteListNoContent creates DeleteListNoContent with default headers values
func NewDeleteListNoContent() *DeleteListNoContent {

	return &DeleteListNoContent{}
}

// WriteResponse to the client
func (o *DeleteListNoContent) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(204)
}

// DeleteListForbiddenCode is the HTTP code returned for type DeleteListForbidden
const DeleteListForbiddenCode int = 403

/*
DeleteListForbidden Access denied

swagger:response deleteListForbidden
*/
type DeleteListForbidden struct {
}

// NewDeleteListForbidden creates DeleteListForbidden with default headers values
func NewDeleteListForbidden() *DeleteListForbidden {

	return &DeleteListForbidden{}
}

// WriteResponse to the client
func (o *DeleteListForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(403)
}

// DeleteListNotFoundCode is the HTTP code returned for type DeleteListNotFound
const DeleteListNotFoundCode int = 404

/*
DeleteListNotFound List not found

swagger:response deleteListNotFound
*/
type DeleteListNotFound struct {
}

// NewDeleteListNotFound creates DeleteListNotFound with default headers values
func NewDeleteListNotFound() *DeleteListNotFound {

	return &DeleteListNotFound{}
}

// WriteResponse to the client
func (o *DeleteListNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(404)
}

// DeleteListInternalServerErrorCode is the HTTP code returned for type DeleteListInternalServerError
const DeleteListInternalServerErrorCode int = 500

/*
DeleteListInternalServerError Server error

swagger:response deleteListInternalServerError
*/
type DeleteListInternalServerError struct {
}

// NewDeleteListInternalServerError creates DeleteListInternalServerError with default headers values
func NewDeleteListInternalServerError() *DeleteListInternalServerError {

	return &DeleteListInternalServerError{}
}

// WriteResponse to the client
func (o *DeleteListInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(500)
}
