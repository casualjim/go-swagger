// Package client contains a client to send http requests
// to a swagger API. This implementation is untyped
package client

import (
	"fmt"
	"net/http"

	"github.com/casualjim/go-swagger/httpkit"
	"github.com/casualjim/go-swagger/spec"
	"github.com/casualjim/go-swagger/strfmt"
)

// Runtime represents an API client that uses the transport
// to make http requests based on a swagger specification.
type Runtime struct {
	DefaultMediaType string
	Consumers        map[string]httpkit.Consumer
	Producers        map[string]httpkit.Producer
	Transport        http.Transport
	Spec             *spec.Document
	Host             string
	BasePath         string
	client           *http.Client
	Formats          strfmt.Registry
}

// New creates a new default runtime for a swagger api client.
func New(swaggerSpec *spec.Document) *Runtime {
	var rt Runtime
	rt.DefaultMediaType = "application/json"
	rt.Consumers = map[string]httpkit.Consumer{
		"application/json": httpkit.JSONConsumer(),
	}
	rt.Producers = map[string]httpkit.Producer{
		"application/json": httpkit.JSONProducer(),
	}
	return &rt
}

// APIError wraps an error model and captures the status code
type APIError struct {
	OperationName string
	Value         interface{}
	Code          int
}

func (a *APIError) Error() string {
	return fmt.Sprintf("%s (status %d): %+v ", a.OperationName, a.Code, a.Value)
}

// Submit a request and when there is a body on success it will turn that into the result
// all other things are turned into an api error for swagger which retains the status code
func (r *Runtime) Submit(request *Request, result interface{}) error {
	//req, _ := http.NewRequest(request.Method, request.Path, body)
	//req.Header.Add(httpkit.HeaderContentType, producerMediaType) // use selected producer mime type
	//req.Header.Add(httpkit.HeaderAccept, consumerMediaType)      // use selected consumer mime type
	//if body != nil && body.Len() > 0 {
	//req.Header.Add("Content-Length", fmt.Sprintf("%d", body.Len()))
	//}

	//res, err := r.client.Do(req) // make requests, by default follows 10 redirects before failing
	//if err != nil {
	//return err
	//}

	//sc := res.StatusCode / 100 // read the response
	//switch sc {
	//case 2:
	//if res.StatusCode == 200 { // only 200 should parse the response body in the result
	//cons, ok := r.Consumers[consumerMediaType]
	//if ok {
	//if err := cons.Consume(res.Body, result); err != nil {
	//return err
	//}
	//} else {
	//return &APIError{
	//OperationName: request.Operation.ID,
	//Value:         fmt.Sprintf("no consumer for %q", consumerMediaType),
	//Code:          res.StatusCode,
	//}
	//}
	//}
	//return nil

	//case 4, 5:
	//// this is an error, check for default model and use that
	//cons, ok := r.Consumers[consumerMediaType]
	//if ok {
	//var eres interface{}
	//if err := cons.Consume(res.Body, &eres); err != nil {
	//return &APIError{OperationName: request.Operation.ID, Value: err, Code: res.StatusCode}
	//}
	//return &APIError{OperationName: request.Operation.ID, Value: eres, Code: res.StatusCode}
	//}
	//return fmt.Errorf("%s: no consumer for %q (status %d)", request.Operation.ID, consumerMediaType, res.StatusCode)
	//}

	return nil
}
