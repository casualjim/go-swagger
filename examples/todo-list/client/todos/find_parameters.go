package todos

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewFindParams creates a new FindParams object
// with the default values initialized.
func NewFindParams() *FindParams {
	var ()
	return &FindParams{}
}

/*FindParams contains all the parameters to send to the API endpoint
for the find operation typically these are written to a http.Request
*/
type FindParams struct {

	/*XRateLimit*/
	XRateLimit int32
	/*Limit*/
	Limit int32
	/*Tags*/
	Tags []int32
}

// WithXRateLimit adds the xRateLimit to the find params
func (o *FindParams) WithXRateLimit(XRateLimit int32) *FindParams {
	o.XRateLimit = XRateLimit
	return o
}

// WithLimit adds the limit to the find params
func (o *FindParams) WithLimit(Limit int32) *FindParams {
	o.Limit = Limit
	return o
}

// WithTags adds the tags to the find params
func (o *FindParams) WithTags(Tags []int32) *FindParams {
	o.Tags = Tags
	return o
}

// WriteToRequest writes these params to a swagger request
func (o *FindParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	var res []error

	// header param X-Rate-Limit
	if err := r.SetHeaderParam("X-Rate-Limit", swag.FormatInt32(o.XRateLimit)); err != nil {
		return err
	}

	// form param limit
	frLimit := o.Limit
	fLimit := swag.FormatInt32(frLimit)
	if err := r.SetFormParam("limit", fLimit); err != nil {
		return err
	}

	var valuesTags []string
	for _, v := range o.Tags {
		valuesTags = append(valuesTags, swag.FormatInt32(v))
	}

	joinedTags := swag.JoinByFormat(valuesTags, "multi")
	// form array param tags
	if err := r.SetFormParam("tags", joinedTags...); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
