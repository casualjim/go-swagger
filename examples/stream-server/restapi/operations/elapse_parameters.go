// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// NewElapseParams creates a new ElapseParams object
// no default values defined in spec.
func NewElapseParams() ElapseParams {

	return ElapseParams{}
}

// ElapseParams contains all the bound params for the elapse operation
// typically these are obtained from a http.Request
//
// swagger:parameters elapse
type ElapseParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*How many seconds to count down
	  Required: true
	  Maximum: 30
	  Minimum: 2
	  In: path
	*/
	Length int64
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewElapseParams() beforehand.
func (o *ElapseParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	rLength, rhkLength, _ := route.Params.GetOK("length")
	if err := o.bindLength(rLength, rhkLength, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindLength binds and validates parameter Length from path.
func (o *ElapseParams) bindLength(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	value, err := swag.ConvertInt64(raw)
	if err != nil {
		return errors.InvalidType("length", "path", "int64", raw)
	}
	o.Length = value

	if err := o.validateLength(formats); err != nil {
		return err
	}

	return nil
}

// validateLength carries on validations for parameter Length
func (o *ElapseParams) validateLength(formats strfmt.Registry) error {

	if err := validate.MinimumInt("length", "path", int64(o.Length), 2, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("length", "path", int64(o.Length), 30, false); err != nil {
		return err
	}

	return nil
}
