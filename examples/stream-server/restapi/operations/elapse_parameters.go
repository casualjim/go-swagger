// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"

	strfmt "github.com/go-openapi/strfmt"
)

// NewElapseParams creates a new ElapseParams object
// with the default values initialized.
func NewElapseParams() ElapseParams {
	var ()
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
// for simple values it will use straight method calls
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

func (o *ElapseParams) bindLength(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

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

func (o *ElapseParams) validateLength(formats strfmt.Registry) error {

	if err := validate.MinimumInt("length", "path", int64(o.Length), 2, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("length", "path", int64(o.Length), 30, false); err != nil {
		return err
	}

	return nil
}
