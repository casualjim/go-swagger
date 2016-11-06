package pet

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"errors"
	"net/url"
	"strings"

	"github.com/go-openapi/swag"
)

// GetPetByIDURL generates an URL for the get pet by Id operation
type GetPetByIDURL struct {
	PetID int64

	// avoid unkeyed usage
	_ struct{}
}

// Build a url path and query string
func (o *GetPetByIDURL) Build() (*url.URL, error) {
	var result url.URL

	var _path = "/pets/{petId}"

	petID := swag.FormatInt64(o.PetID)
	if petID != "" {
		_path = strings.Replace(_path, "{petId}", petID, -1)
	} else {
		return nil, errors.New("PetID is required on GetPetByIDURL")
	}
	result.Path = _path

	return &result, nil
}

// Must is a helper function to panic when the url builder returns an error
func (o *GetPetByIDURL) Must(u *url.URL, err error) *url.URL {
	if err != nil {
		panic(err)
	}
	if u == nil {
		panic("url can't be nil")
	}
	return u
}

// String returns the string representation of the path with query string
func (o *GetPetByIDURL) String() string {
	return o.Must(o.Build()).String()
}

// BuildFull builds a full url with scheme, host, path and query string
func (o *GetPetByIDURL) BuildFull(scheme, host string) (*url.URL, error) {
	if scheme == "" {
		return nil, errors.New("scheme is required for a full url on GetPetByIDURL")
	}
	if host == "" {
		return nil, errors.New("host is required for a full url on GetPetByIDURL")
	}

	base, err := o.Build()
	if err != nil {
		return nil, err
	}

	base.Scheme = scheme
	base.Host = host
	return base, nil
}

// StringFull returns the string representation of a complete url
func (o *GetPetByIDURL) StringFull(scheme, host string) string {
	return o.Must(o.BuildFull(scheme, host)).String()
}
