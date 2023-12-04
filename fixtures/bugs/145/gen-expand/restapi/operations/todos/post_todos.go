// Code generated by go-swagger; DO NOT EDIT.

package todos

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"context"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// PostTodosHandlerFunc turns a function with the right signature into a post todos handler
type PostTodosHandlerFunc func(PostTodosParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn PostTodosHandlerFunc) Handle(params PostTodosParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// PostTodosHandler interface for that can handle valid post todos params
type PostTodosHandler interface {
	Handle(PostTodosParams, interface{}) middleware.Responder
}

// NewPostTodos creates a new http.Handler for the post todos operation
func NewPostTodos(ctx *middleware.Context, handler PostTodosHandler) *PostTodos {
	return &PostTodos{Context: ctx, Handler: handler}
}

/*
	PostTodos swagger:route POST /todos Todos postTodos

# Create Todo

This creates a Todo object.

Testing `inline code`.
*/
type PostTodos struct {
	Context *middleware.Context
	Handler PostTodosHandler
}

func (o *PostTodos) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewPostTodosParams()
	uprinc, aCtx, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	if aCtx != nil {
		*r = *aCtx
	}
	var principal interface{}
	if uprinc != nil {
		principal = uprinc.(interface{}) // this is really a interface{}, I promise
	}

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params, principal) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}

// PostTodosBody Todo Partial
//
// swagger:model PostTodosBody
type PostTodosBody struct {

	// completed
	// Required: true
	Completed *bool `json:"completed"`

	// name
	// Required: true
	Name *string `json:"name"`
}

// Validate validates this post todos body
func (o *PostTodosBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateCompleted(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *PostTodosBody) validateCompleted(formats strfmt.Registry) error {

	if err := validate.Required("body"+"."+"completed", "body", o.Completed); err != nil {
		return err
	}

	return nil
}

func (o *PostTodosBody) validateName(formats strfmt.Registry) error {

	if err := validate.Required("body"+"."+"name", "body", o.Name); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this post todos body based on context it is used
func (o *PostTodosBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *PostTodosBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostTodosBody) UnmarshalBinary(b []byte) error {
	var res PostTodosBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// PostTodosCreatedBody Todo Full
//
// swagger:model PostTodosCreatedBody
type PostTodosCreatedBody struct {

	// completed
	// Required: true
	Completed *bool `json:"completed"`

	// name
	// Required: true
	Name *string `json:"name"`

	// completed at
	// Format: date-time
	CompletedAt strfmt.DateTime `json:"completed_at,omitempty"`

	// created at
	// Format: date-time
	CreatedAt strfmt.DateTime `json:"created_at,omitempty"`

	// id
	// Required: true
	// Maximum: 1e+06
	// Minimum: 0
	ID *int64 `json:"id"`

	// updated at
	// Format: date-time
	UpdatedAt strfmt.DateTime `json:"updated_at,omitempty"`

	// user
	// Required: true
	User *PostTodosCreatedBodyPostTodosCreatedBodyAO1User `json:"user"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (o *PostTodosCreatedBody) UnmarshalJSON(raw []byte) error {
	// PostTodosCreatedBodyAO0
	var dataPostTodosCreatedBodyAO0 struct {
		Completed *bool `json:"completed"`

		Name *string `json:"name"`
	}
	if err := swag.ReadJSON(raw, &dataPostTodosCreatedBodyAO0); err != nil {
		return err
	}

	o.Completed = dataPostTodosCreatedBodyAO0.Completed

	o.Name = dataPostTodosCreatedBodyAO0.Name

	// PostTodosCreatedBodyAO1
	var dataPostTodosCreatedBodyAO1 struct {
		CompletedAt strfmt.DateTime `json:"completed_at,omitempty"`

		CreatedAt strfmt.DateTime `json:"created_at,omitempty"`

		ID *int64 `json:"id"`

		UpdatedAt strfmt.DateTime `json:"updated_at,omitempty"`

		User *PostTodosCreatedBodyPostTodosCreatedBodyAO1User `json:"user"`
	}
	if err := swag.ReadJSON(raw, &dataPostTodosCreatedBodyAO1); err != nil {
		return err
	}

	o.CompletedAt = dataPostTodosCreatedBodyAO1.CompletedAt

	o.CreatedAt = dataPostTodosCreatedBodyAO1.CreatedAt

	o.ID = dataPostTodosCreatedBodyAO1.ID

	o.UpdatedAt = dataPostTodosCreatedBodyAO1.UpdatedAt

	o.User = dataPostTodosCreatedBodyAO1.User

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (o PostTodosCreatedBody) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	var dataPostTodosCreatedBodyAO0 struct {
		Completed *bool `json:"completed"`

		Name *string `json:"name"`
	}

	dataPostTodosCreatedBodyAO0.Completed = o.Completed

	dataPostTodosCreatedBodyAO0.Name = o.Name

	jsonDataPostTodosCreatedBodyAO0, errPostTodosCreatedBodyAO0 := swag.WriteJSON(dataPostTodosCreatedBodyAO0)
	if errPostTodosCreatedBodyAO0 != nil {
		return nil, errPostTodosCreatedBodyAO0
	}
	_parts = append(_parts, jsonDataPostTodosCreatedBodyAO0)
	var dataPostTodosCreatedBodyAO1 struct {
		CompletedAt strfmt.DateTime `json:"completed_at,omitempty"`

		CreatedAt strfmt.DateTime `json:"created_at,omitempty"`

		ID *int64 `json:"id"`

		UpdatedAt strfmt.DateTime `json:"updated_at,omitempty"`

		User *PostTodosCreatedBodyPostTodosCreatedBodyAO1User `json:"user"`
	}

	dataPostTodosCreatedBodyAO1.CompletedAt = o.CompletedAt

	dataPostTodosCreatedBodyAO1.CreatedAt = o.CreatedAt

	dataPostTodosCreatedBodyAO1.ID = o.ID

	dataPostTodosCreatedBodyAO1.UpdatedAt = o.UpdatedAt

	dataPostTodosCreatedBodyAO1.User = o.User

	jsonDataPostTodosCreatedBodyAO1, errPostTodosCreatedBodyAO1 := swag.WriteJSON(dataPostTodosCreatedBodyAO1)
	if errPostTodosCreatedBodyAO1 != nil {
		return nil, errPostTodosCreatedBodyAO1
	}
	_parts = append(_parts, jsonDataPostTodosCreatedBodyAO1)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this post todos created body
func (o *PostTodosCreatedBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateCompleted(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateCompletedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateCreatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateUpdatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateUser(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *PostTodosCreatedBody) validateCompleted(formats strfmt.Registry) error {

	if err := validate.Required("postTodosCreated"+"."+"completed", "body", o.Completed); err != nil {
		return err
	}

	return nil
}

func (o *PostTodosCreatedBody) validateName(formats strfmt.Registry) error {

	if err := validate.Required("postTodosCreated"+"."+"name", "body", o.Name); err != nil {
		return err
	}

	return nil
}

func (o *PostTodosCreatedBody) validateCompletedAt(formats strfmt.Registry) error {

	if swag.IsZero(o.CompletedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("postTodosCreated"+"."+"completed_at", "body", "date-time", o.CompletedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (o *PostTodosCreatedBody) validateCreatedAt(formats strfmt.Registry) error {

	if swag.IsZero(o.CreatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("postTodosCreated"+"."+"created_at", "body", "date-time", o.CreatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (o *PostTodosCreatedBody) validateID(formats strfmt.Registry) error {

	if err := validate.Required("postTodosCreated"+"."+"id", "body", o.ID); err != nil {
		return err
	}

	if err := validate.MinimumInt("postTodosCreated"+"."+"id", "body", *o.ID, 0, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("postTodosCreated"+"."+"id", "body", *o.ID, 1e+06, false); err != nil {
		return err
	}

	return nil
}

func (o *PostTodosCreatedBody) validateUpdatedAt(formats strfmt.Registry) error {

	if swag.IsZero(o.UpdatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("postTodosCreated"+"."+"updated_at", "body", "date-time", o.UpdatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (o *PostTodosCreatedBody) validateUser(formats strfmt.Registry) error {

	if err := validate.Required("postTodosCreated"+"."+"user", "body", o.User); err != nil {
		return err
	}

	if o.User != nil {
		if err := o.User.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("postTodosCreated" + "." + "user")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("postTodosCreated" + "." + "user")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this post todos created body based on the context it is used
func (o *PostTodosCreatedBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateUser(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *PostTodosCreatedBody) contextValidateUser(ctx context.Context, formats strfmt.Registry) error {

	if o.User != nil {

		if err := o.User.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("postTodosCreated" + "." + "user")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("postTodosCreated" + "." + "user")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *PostTodosCreatedBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostTodosCreatedBody) UnmarshalBinary(b []byte) error {
	var res PostTodosCreatedBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// PostTodosCreatedBodyPostTodosCreatedBodyAO1User User
//
// swagger:model PostTodosCreatedBodyPostTodosCreatedBodyAO1User
type PostTodosCreatedBodyPostTodosCreatedBodyAO1User struct {

	// age
	// Required: true
	// Maximum: 150
	// Minimum: 0
	Age *float64 `json:"age"`

	// error
	Error *PostTodosCreatedBodyPostTodosCreatedBodyAO1UserError `json:"error,omitempty"`

	// The user's full name.
	// Required: true
	Name *string `json:"name"`
}

// Validate validates this post todos created body post todos created body a o1 user
func (o *PostTodosCreatedBodyPostTodosCreatedBodyAO1User) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateAge(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateError(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *PostTodosCreatedBodyPostTodosCreatedBodyAO1User) validateAge(formats strfmt.Registry) error {

	if err := validate.Required("postTodosCreated"+"."+"user"+"."+"age", "body", o.Age); err != nil {
		return err
	}

	if err := validate.Minimum("postTodosCreated"+"."+"user"+"."+"age", "body", *o.Age, 0, false); err != nil {
		return err
	}

	if err := validate.Maximum("postTodosCreated"+"."+"user"+"."+"age", "body", *o.Age, 150, false); err != nil {
		return err
	}

	return nil
}

func (o *PostTodosCreatedBodyPostTodosCreatedBodyAO1User) validateError(formats strfmt.Registry) error {
	if swag.IsZero(o.Error) { // not required
		return nil
	}

	if o.Error != nil {
		if err := o.Error.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("postTodosCreated" + "." + "user" + "." + "error")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("postTodosCreated" + "." + "user" + "." + "error")
			}
			return err
		}
	}

	return nil
}

func (o *PostTodosCreatedBodyPostTodosCreatedBodyAO1User) validateName(formats strfmt.Registry) error {

	if err := validate.Required("postTodosCreated"+"."+"user"+"."+"name", "body", o.Name); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this post todos created body post todos created body a o1 user based on the context it is used
func (o *PostTodosCreatedBodyPostTodosCreatedBodyAO1User) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateError(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *PostTodosCreatedBodyPostTodosCreatedBodyAO1User) contextValidateError(ctx context.Context, formats strfmt.Registry) error {

	if o.Error != nil {

		if swag.IsZero(o.Error) { // not required
			return nil
		}

		if err := o.Error.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("postTodosCreated" + "." + "user" + "." + "error")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("postTodosCreated" + "." + "user" + "." + "error")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *PostTodosCreatedBodyPostTodosCreatedBodyAO1User) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostTodosCreatedBodyPostTodosCreatedBodyAO1User) UnmarshalBinary(b []byte) error {
	var res PostTodosCreatedBodyPostTodosCreatedBodyAO1User
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// PostTodosCreatedBodyPostTodosCreatedBodyAO1UserError Error Response
//
// swagger:model PostTodosCreatedBodyPostTodosCreatedBodyAO1UserError
type PostTodosCreatedBodyPostTodosCreatedBodyAO1UserError struct {

	// error
	// Required: true
	Error *string `json:"error"`

	// status
	// Required: true
	Status *string `json:"status"`
}

// Validate validates this post todos created body post todos created body a o1 user error
func (o *PostTodosCreatedBodyPostTodosCreatedBodyAO1UserError) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateError(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *PostTodosCreatedBodyPostTodosCreatedBodyAO1UserError) validateError(formats strfmt.Registry) error {

	if err := validate.Required("postTodosCreated"+"."+"user"+"."+"error"+"."+"error", "body", o.Error); err != nil {
		return err
	}

	return nil
}

func (o *PostTodosCreatedBodyPostTodosCreatedBodyAO1UserError) validateStatus(formats strfmt.Registry) error {

	if err := validate.Required("postTodosCreated"+"."+"user"+"."+"error"+"."+"status", "body", o.Status); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this post todos created body post todos created body a o1 user error based on context it is used
func (o *PostTodosCreatedBodyPostTodosCreatedBodyAO1UserError) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *PostTodosCreatedBodyPostTodosCreatedBodyAO1UserError) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostTodosCreatedBodyPostTodosCreatedBodyAO1UserError) UnmarshalBinary(b []byte) error {
	var res PostTodosCreatedBodyPostTodosCreatedBodyAO1UserError
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// PostTodosInternalServerErrorBody Error Response
//
// swagger:model PostTodosInternalServerErrorBody
type PostTodosInternalServerErrorBody struct {

	// error
	// Required: true
	Error *string `json:"error"`

	// status
	// Required: true
	Status *string `json:"status"`
}

// Validate validates this post todos internal server error body
func (o *PostTodosInternalServerErrorBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateError(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *PostTodosInternalServerErrorBody) validateError(formats strfmt.Registry) error {

	if err := validate.Required("postTodosInternalServerError"+"."+"error", "body", o.Error); err != nil {
		return err
	}

	return nil
}

func (o *PostTodosInternalServerErrorBody) validateStatus(formats strfmt.Registry) error {

	if err := validate.Required("postTodosInternalServerError"+"."+"status", "body", o.Status); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this post todos internal server error body based on context it is used
func (o *PostTodosInternalServerErrorBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *PostTodosInternalServerErrorBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostTodosInternalServerErrorBody) UnmarshalBinary(b []byte) error {
	var res PostTodosInternalServerErrorBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// PostTodosUnauthorizedBody Error Response
//
// swagger:model PostTodosUnauthorizedBody
type PostTodosUnauthorizedBody struct {

	// error
	// Required: true
	Error *string `json:"error"`

	// status
	// Required: true
	Status *string `json:"status"`
}

// Validate validates this post todos unauthorized body
func (o *PostTodosUnauthorizedBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateError(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *PostTodosUnauthorizedBody) validateError(formats strfmt.Registry) error {

	if err := validate.Required("postTodosUnauthorized"+"."+"error", "body", o.Error); err != nil {
		return err
	}

	return nil
}

func (o *PostTodosUnauthorizedBody) validateStatus(formats strfmt.Registry) error {

	if err := validate.Required("postTodosUnauthorized"+"."+"status", "body", o.Status); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this post todos unauthorized body based on context it is used
func (o *PostTodosUnauthorizedBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *PostTodosUnauthorizedBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostTodosUnauthorizedBody) UnmarshalBinary(b []byte) error {
	var res PostTodosUnauthorizedBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
