package store

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/eatigo/go-swagger/examples/generated/models"
)

// GetOrderByIDOKCode is the HTTP code returned for type GetOrderByIDOK
const GetOrderByIDOKCode int = 200

/*GetOrderByIDOK successful operation

swagger:response getOrderByIdOK
*/
type GetOrderByIDOK struct {

	/*
	  In: Body
	*/
	Payload *models.Order `json:"body,omitempty"`
}

// NewGetOrderByIDOK creates GetOrderByIDOK with default headers values
func NewGetOrderByIDOK() *GetOrderByIDOK {
	return &GetOrderByIDOK{}
}

// WithPayload adds the payload to the get order by Id o k response
func (o *GetOrderByIDOK) WithPayload(payload *models.Order) *GetOrderByIDOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get order by Id o k response
func (o *GetOrderByIDOK) SetPayload(payload *models.Order) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetOrderByIDOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetOrderByIDBadRequestCode is the HTTP code returned for type GetOrderByIDBadRequest
const GetOrderByIDBadRequestCode int = 400

/*GetOrderByIDBadRequest Invalid ID supplied

swagger:response getOrderByIdBadRequest
*/
type GetOrderByIDBadRequest struct {
}

// NewGetOrderByIDBadRequest creates GetOrderByIDBadRequest with default headers values
func NewGetOrderByIDBadRequest() *GetOrderByIDBadRequest {
	return &GetOrderByIDBadRequest{}
}

// WriteResponse to the client
func (o *GetOrderByIDBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
}

// GetOrderByIDNotFoundCode is the HTTP code returned for type GetOrderByIDNotFound
const GetOrderByIDNotFoundCode int = 404

/*GetOrderByIDNotFound Order not found

swagger:response getOrderByIdNotFound
*/
type GetOrderByIDNotFound struct {
}

// NewGetOrderByIDNotFound creates GetOrderByIDNotFound with default headers values
func NewGetOrderByIDNotFound() *GetOrderByIDNotFound {
	return &GetOrderByIDNotFound{}
}

// WriteResponse to the client
func (o *GetOrderByIDNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
}
