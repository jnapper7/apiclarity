// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/apiclarity/apiclarity/api/server/models"
)

// PostAPIInventoryReviewIDApprovedReviewOKCode is the HTTP code returned for type PostAPIInventoryReviewIDApprovedReviewOK
const PostAPIInventoryReviewIDApprovedReviewOKCode int = 200

/*PostAPIInventoryReviewIDApprovedReviewOK Success

swagger:response postApiInventoryReviewIdApprovedReviewOK
*/
type PostAPIInventoryReviewIDApprovedReviewOK struct {

	/*
	  In: Body
	*/
	Payload interface{} `json:"body,omitempty"`
}

// NewPostAPIInventoryReviewIDApprovedReviewOK creates PostAPIInventoryReviewIDApprovedReviewOK with default headers values
func NewPostAPIInventoryReviewIDApprovedReviewOK() *PostAPIInventoryReviewIDApprovedReviewOK {

	return &PostAPIInventoryReviewIDApprovedReviewOK{}
}

// WithPayload adds the payload to the post Api inventory review Id approved review o k response
func (o *PostAPIInventoryReviewIDApprovedReviewOK) WithPayload(payload interface{}) *PostAPIInventoryReviewIDApprovedReviewOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post Api inventory review Id approved review o k response
func (o *PostAPIInventoryReviewIDApprovedReviewOK) SetPayload(payload interface{}) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostAPIInventoryReviewIDApprovedReviewOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

/*PostAPIInventoryReviewIDApprovedReviewDefault unknown error

swagger:response postApiInventoryReviewIdApprovedReviewDefault
*/
type PostAPIInventoryReviewIDApprovedReviewDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.APIResponse `json:"body,omitempty"`
}

// NewPostAPIInventoryReviewIDApprovedReviewDefault creates PostAPIInventoryReviewIDApprovedReviewDefault with default headers values
func NewPostAPIInventoryReviewIDApprovedReviewDefault(code int) *PostAPIInventoryReviewIDApprovedReviewDefault {
	if code <= 0 {
		code = 500
	}

	return &PostAPIInventoryReviewIDApprovedReviewDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the post API inventory review ID approved review default response
func (o *PostAPIInventoryReviewIDApprovedReviewDefault) WithStatusCode(code int) *PostAPIInventoryReviewIDApprovedReviewDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the post API inventory review ID approved review default response
func (o *PostAPIInventoryReviewIDApprovedReviewDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the post API inventory review ID approved review default response
func (o *PostAPIInventoryReviewIDApprovedReviewDefault) WithPayload(payload *models.APIResponse) *PostAPIInventoryReviewIDApprovedReviewDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post API inventory review ID approved review default response
func (o *PostAPIInventoryReviewIDApprovedReviewDefault) SetPayload(payload *models.APIResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostAPIInventoryReviewIDApprovedReviewDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}