package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

/*Comment A comment for an issue.

Users can comment on issues to discuss plans for resolution etc.


swagger:model Comment
*/
type Comment struct {

	/* The content of the comment.

	This is a free text field with support for github flavored markdown.


	Required: true
	*/
	Content *string `json:"content"`

	/* The time at which this comment was created.

	This field is autogenerated when the content is posted.

	Read Only: true
	*/
	CreatedAt strfmt.DateTime `json:"createdAt,omitempty"`

	/* user

	Required: true
	*/
	User *UserCard `json:"user"`
}

// Validate validates this comment
func (m *Comment) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateContent(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateUser(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Comment) validateContent(formats strfmt.Registry) error {

	if err := validate.Required("content", "body", m.Content); err != nil {
		return err
	}

	return nil
}

func (m *Comment) validateUser(formats strfmt.Registry) error {

	if err := validate.Required("user", "body", m.User); err != nil {
		return err
	}

	if m.User != nil {

		if err := m.User.Validate(formats); err != nil {
			return err
		}
	}

	return nil
}
