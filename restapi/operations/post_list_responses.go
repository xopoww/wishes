// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// PostListCreatedCode is the HTTP code returned for type PostListCreated
const PostListCreatedCode int = 201

/*
PostListCreated Success

swagger:response postListCreated
*/
type PostListCreated struct {

	/*
	  In: Body
	*/
	Payload *PostListCreatedBody `json:"body,omitempty"`
}

// NewPostListCreated creates PostListCreated with default headers values
func NewPostListCreated() *PostListCreated {

	return &PostListCreated{}
}

// WithPayload adds the payload to the post list created response
func (o *PostListCreated) WithPayload(payload *PostListCreatedBody) *PostListCreated {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post list created response
func (o *PostListCreated) SetPayload(payload *PostListCreatedBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostListCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(201)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PostListInternalServerErrorCode is the HTTP code returned for type PostListInternalServerError
const PostListInternalServerErrorCode int = 500

/*
PostListInternalServerError Server error

swagger:response postListInternalServerError
*/
type PostListInternalServerError struct {
}

// NewPostListInternalServerError creates PostListInternalServerError with default headers values
func NewPostListInternalServerError() *PostListInternalServerError {

	return &PostListInternalServerError{}
}

// WriteResponse to the client
func (o *PostListInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(500)
}
