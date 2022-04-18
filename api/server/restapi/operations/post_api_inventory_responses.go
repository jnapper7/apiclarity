// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/apiclarity/apiclarity/api/server/models"
)

// PostAPIInventoryOKCode is the HTTP code returned for type PostAPIInventoryOK
const PostAPIInventoryOKCode int = 200

/*PostAPIInventoryOK Success

swagger:response postApiInventoryOK
*/
type PostAPIInventoryOK struct {

	/*
	  In: Body
	*/
	Payload *models.APIInfo `json:"body,omitempty"`
}

// NewPostAPIInventoryOK creates PostAPIInventoryOK with default headers values
func NewPostAPIInventoryOK() *PostAPIInventoryOK {

	return &PostAPIInventoryOK{}
}

// WithPayload adds the payload to the post Api inventory o k response
func (o *PostAPIInventoryOK) WithPayload(payload *models.APIInfo) *PostAPIInventoryOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post Api inventory o k response
func (o *PostAPIInventoryOK) SetPayload(payload *models.APIInfo) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostAPIInventoryOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*PostAPIInventoryDefault unknown error

swagger:response postApiInventoryDefault
*/
type PostAPIInventoryDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.APIResponse `json:"body,omitempty"`
}

// NewPostAPIInventoryDefault creates PostAPIInventoryDefault with default headers values
func NewPostAPIInventoryDefault(code int) *PostAPIInventoryDefault {
	if code <= 0 {
		code = 500
	}

	return &PostAPIInventoryDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the post API inventory default response
func (o *PostAPIInventoryDefault) WithStatusCode(code int) *PostAPIInventoryDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the post API inventory default response
func (o *PostAPIInventoryDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the post API inventory default response
func (o *PostAPIInventoryDefault) WithPayload(payload *models.APIResponse) *PostAPIInventoryDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post API inventory default response
func (o *PostAPIInventoryDefault) SetPayload(payload *models.APIResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostAPIInventoryDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}