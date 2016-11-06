package tasks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"errors"
	"net/url"
	"strings"

	"github.com/go-openapi/swag"
)

// UploadTaskFileURL generates an URL for the upload task file operation
type UploadTaskFileURL struct {
	ID int64

	// avoid unkeyed usage
	_ struct{}
}

// Build a url path and query string
func (o *UploadTaskFileURL) Build() (*url.URL, error) {
	var result url.URL

	var path = "/tasks/{id}/files"

	id := swag.FormatInt64(o.ID)
	if id != "" {
		path = strings.Replace(path, "{id}", id, -1)
	} else {
		return nil, errors.New("ID is required on UploadTaskFileURL")
	}
	result.Path = path

	return &result, nil
}

// Must is a helper function to panic when the url builder returns an error
func (o *UploadTaskFileURL) Must(u *url.URL, err error) *url.URL {
	if err != nil {
		panic(err)
	}
	if u == nil {
		panic("url can't be nil")
	}
	return u
}

// String returns the string representation of the path with query string
func (o *UploadTaskFileURL) String() string {
	return o.Must(o.Build()).String()
}

// BuildFull builds a full url with scheme, host, path and query string
func (o *UploadTaskFileURL) BuildFull(scheme, host string) (*url.URL, error) {
	if scheme == "" {
		return nil, errors.New("scheme is required for a full url on UploadTaskFileURL")
	}
	if host == "" {
		return nil, errors.New("host is required for a full url on UploadTaskFileURL")
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
func (o *UploadTaskFileURL) StringFull(scheme, host string) string {
	return o.Must(o.BuildFull(scheme, host)).String()
}
