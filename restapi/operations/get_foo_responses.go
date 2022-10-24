// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// GetFooOKCode is the HTTP code returned for type GetFooOK
const GetFooOKCode int = 200

/*
GetFooOK Successful response

swagger:response getFooOK
*/
type GetFooOK struct {

	/*
	  In: Body
	*/
	Payload string `json:"body,omitempty"`
}

// NewGetFooOK creates GetFooOK with default headers values
func NewGetFooOK() *GetFooOK {

	return &GetFooOK{}
}

// WithPayload adds the payload to the get foo o k response
func (o *GetFooOK) WithPayload(payload string) *GetFooOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get foo o k response
func (o *GetFooOK) SetPayload(payload string) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetFooOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}
