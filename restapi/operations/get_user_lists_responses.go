// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// GetUserListsOKCode is the HTTP code returned for type GetUserListsOK
const GetUserListsOKCode int = 200

/*
GetUserListsOK Success

swagger:response getUserListsOK
*/
type GetUserListsOK struct {

	/*
	  Unique: true
	  In: Body
	*/
	Payload []int64 `json:"body,omitempty"`
}

// NewGetUserListsOK creates GetUserListsOK with default headers values
func NewGetUserListsOK() *GetUserListsOK {

	return &GetUserListsOK{}
}

// WithPayload adds the payload to the get user lists o k response
func (o *GetUserListsOK) WithPayload(payload []int64) *GetUserListsOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get user lists o k response
func (o *GetUserListsOK) SetPayload(payload []int64) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetUserListsOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		// return empty array
		payload = make([]int64, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// GetUserListsNotFoundCode is the HTTP code returned for type GetUserListsNotFound
const GetUserListsNotFoundCode int = 404

/*
GetUserListsNotFound User not found

swagger:response getUserListsNotFound
*/
type GetUserListsNotFound struct {
}

// NewGetUserListsNotFound creates GetUserListsNotFound with default headers values
func NewGetUserListsNotFound() *GetUserListsNotFound {

	return &GetUserListsNotFound{}
}

// WriteResponse to the client
func (o *GetUserListsNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(404)
}

// GetUserListsInternalServerErrorCode is the HTTP code returned for type GetUserListsInternalServerError
const GetUserListsInternalServerErrorCode int = 500

/*
GetUserListsInternalServerError Server error

swagger:response getUserListsInternalServerError
*/
type GetUserListsInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *GetUserListsInternalServerErrorBody `json:"body,omitempty"`
}

// NewGetUserListsInternalServerError creates GetUserListsInternalServerError with default headers values
func NewGetUserListsInternalServerError() *GetUserListsInternalServerError {

	return &GetUserListsInternalServerError{}
}

// WithPayload adds the payload to the get user lists internal server error response
func (o *GetUserListsInternalServerError) WithPayload(payload *GetUserListsInternalServerErrorBody) *GetUserListsInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get user lists internal server error response
func (o *GetUserListsInternalServerError) SetPayload(payload *GetUserListsInternalServerErrorBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetUserListsInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}