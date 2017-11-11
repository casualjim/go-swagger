// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"io"
	"net/http"

	"github.com/go-openapi/runtime"
)

// ElapseOKCode is the HTTP code returned for type ElapseOK
const ElapseOKCode int = 200

/*ElapseOK Secondly update on remaining time

swagger:response elapseOK
*/
type ElapseOK struct {

	/*
	  In: Body
	*/
	Payload io.ReadCloser `json:"body,omitempty"`
}

// NewElapseOK creates ElapseOK with default headers values
func NewElapseOK() *ElapseOK {
	return &ElapseOK{}
}

// WithPayload adds the payload to the elapse o k response
func (o *ElapseOK) WithPayload(payload io.ReadCloser) *ElapseOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the elapse o k response
func (o *ElapseOK) SetPayload(payload io.ReadCloser) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ElapseOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}

}

// ElapseForbiddenCode is the HTTP code returned for type ElapseForbidden
const ElapseForbiddenCode int = 403

/*ElapseForbidden Contrived - thrown when length of 11 is chosen

swagger:response elapseForbidden
*/
type ElapseForbidden struct {
}

// NewElapseForbidden creates ElapseForbidden with default headers values
func NewElapseForbidden() *ElapseForbidden {
	return &ElapseForbidden{}
}

// WriteResponse to the client
func (o *ElapseForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(403)
}
