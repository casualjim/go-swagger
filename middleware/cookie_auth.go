package middleware

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// CookieAuthenticator authenticates requests using cookies
type CookieAuthenticator struct {
	Name string
}

// NewCookieAuthenticator creates a new cookie authenticator
func NewCookieAuthenticator(name string) *CookieAuthenticator {
	return &CookieAuthenticator{Name: name}
}

// Authenticate implements the runtime.Authenticator interface
func (c *CookieAuthenticator) Authenticate(params interface{}) (bool, interface{}, error) {
	req, ok := params.(*http.Request)
	if !ok {
		return false, nil, nil
	}

	cookie, err := req.Cookie(c.Name)
	if err != nil {
		return false, nil, nil
	}

	return true, cookie.Value, nil
}