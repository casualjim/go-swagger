// Code generated by go-swagger; DO NOT EDIT.

package pet

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// PetDeleteReader is a Reader for the PetDelete structure.
type PetDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PetDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewPetDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPetDeleteBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPetDeleteNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPetDeleteNoContent creates a PetDeleteNoContent with default headers values
func NewPetDeleteNoContent() *PetDeleteNoContent {
	return &PetDeleteNoContent{}
}

/*PetDeleteNoContent handles this case with default header values.

Deleted successfully
*/
type PetDeleteNoContent struct {
}

func (o *PetDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /pet/{petId}][%d] petDeleteNoContent ", 204)
}

func (o *PetDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPetDeleteBadRequest creates a PetDeleteBadRequest with default headers values
func NewPetDeleteBadRequest() *PetDeleteBadRequest {
	return &PetDeleteBadRequest{}
}

/*PetDeleteBadRequest handles this case with default header values.

Invalid ID supplied
*/
type PetDeleteBadRequest struct {
}

func (o *PetDeleteBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /pet/{petId}][%d] petDeleteBadRequest ", 400)
}

func (o *PetDeleteBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPetDeleteNotFound creates a PetDeleteNotFound with default headers values
func NewPetDeleteNotFound() *PetDeleteNotFound {
	return &PetDeleteNotFound{}
}

/*PetDeleteNotFound handles this case with default header values.

Pet not found
*/
type PetDeleteNotFound struct {
}

func (o *PetDeleteNotFound) Error() string {
	return fmt.Sprintf("[DELETE /pet/{petId}][%d] petDeleteNotFound ", 404)
}

func (o *PetDeleteNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
