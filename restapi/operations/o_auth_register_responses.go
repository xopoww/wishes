// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// OAuthRegisterOKCode is the HTTP code returned for type OAuthRegisterOK
const OAuthRegisterOKCode int = 200

/*
OAuthRegisterOK Registration result

swagger:response oAuthRegisterOK
*/
type OAuthRegisterOK struct {

	/*
	  In: Body
	*/
	Payload *OAuthRegisterOKBody `json:"body,omitempty"`
}

// NewOAuthRegisterOK creates OAuthRegisterOK with default headers values
func NewOAuthRegisterOK() *OAuthRegisterOK {

	return &OAuthRegisterOK{}
}

// WithPayload adds the payload to the o auth register o k response
func (o *OAuthRegisterOK) WithPayload(payload *OAuthRegisterOKBody) *OAuthRegisterOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the o auth register o k response
func (o *OAuthRegisterOK) SetPayload(payload *OAuthRegisterOKBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *OAuthRegisterOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// OAuthRegisterInternalServerErrorCode is the HTTP code returned for type OAuthRegisterInternalServerError
const OAuthRegisterInternalServerErrorCode int = 500

/*
OAuthRegisterInternalServerError Server error

swagger:response oAuthRegisterInternalServerError
*/
type OAuthRegisterInternalServerError struct {
}

// NewOAuthRegisterInternalServerError creates OAuthRegisterInternalServerError with default headers values
func NewOAuthRegisterInternalServerError() *OAuthRegisterInternalServerError {

	return &OAuthRegisterInternalServerError{}
}

// WriteResponse to the client
func (o *OAuthRegisterInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(500)
}