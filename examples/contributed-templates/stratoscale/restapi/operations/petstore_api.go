// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/loads"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/runtime/security"
	"github.com/go-openapi/spec"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/go-swagger/go-swagger/examples/contributed-templates/stratoscale/restapi/operations/pet"
	"github.com/go-swagger/go-swagger/examples/contributed-templates/stratoscale/restapi/operations/store"
)

// NewPetstoreAPI creates a new Petstore instance
func NewPetstoreAPI(spec *loads.Document) *PetstoreAPI {
	return &PetstoreAPI{
		handlers:            make(map[string]map[string]http.Handler),
		formats:             strfmt.Default,
		defaultConsumes:     "application/json",
		defaultProduces:     "application/json",
		customConsumers:     make(map[string]runtime.Consumer),
		customProducers:     make(map[string]runtime.Producer),
		PreServerShutdown:   func() {},
		ServerShutdown:      func() {},
		spec:                spec,
		ServeError:          errors.ServeError,
		BasicAuthenticator:  security.BasicAuth,
		APIKeyAuthenticator: security.APIKeyAuth,
		BearerAuthenticator: security.BearerAuth,

		JSONConsumer: runtime.JSONConsumer(),

		JSONProducer: runtime.JSONProducer(),

		StoreInventoryGetHandler: store.InventoryGetHandlerFunc(func(params store.InventoryGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation store.InventoryGet has not yet been implemented")
		}),
		StoreOrderCreateHandler: store.OrderCreateHandlerFunc(func(params store.OrderCreateParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation store.OrderCreate has not yet been implemented")
		}),
		StoreOrderDeleteHandler: store.OrderDeleteHandlerFunc(func(params store.OrderDeleteParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation store.OrderDelete has not yet been implemented")
		}),
		StoreOrderGetHandler: store.OrderGetHandlerFunc(func(params store.OrderGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation store.OrderGet has not yet been implemented")
		}),
		PetPetCreateHandler: pet.PetCreateHandlerFunc(func(params pet.PetCreateParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation pet.PetCreate has not yet been implemented")
		}),
		PetPetDeleteHandler: pet.PetDeleteHandlerFunc(func(params pet.PetDeleteParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation pet.PetDelete has not yet been implemented")
		}),
		PetPetGetHandler: pet.PetGetHandlerFunc(func(params pet.PetGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation pet.PetGet has not yet been implemented")
		}),
		PetPetListHandler: pet.PetListHandlerFunc(func(params pet.PetListParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation pet.PetList has not yet been implemented")
		}),
		PetPetUpdateHandler: pet.PetUpdateHandlerFunc(func(params pet.PetUpdateParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation pet.PetUpdate has not yet been implemented")
		}),
		PetPetUploadImageHandler: pet.PetUploadImageHandlerFunc(func(params pet.PetUploadImageParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation pet.PetUploadImage has not yet been implemented")
		}),

		// Applies when the "X-Auth-Roles" header is set
		RolesAuth: func(token string) (interface{}, error) {
			return nil, errors.NotImplemented("api key auth (roles) X-Auth-Roles from header param [X-Auth-Roles] has not yet been implemented")
		},
		// default authorizer is authorized meaning no requests are blocked
		APIAuthorizer: security.Authorized(),
	}
}

/*PetstoreAPI This is a sample server Petstore server.  You can find out more about Swagger at [http://swagger.io](http://swagger.io) or on [irc.freenode.net, #swagger](http://swagger.io/irc/).  For this sample, you can use the api key `special-key` to test the authorization filters. */
type PetstoreAPI struct {
	spec            *loads.Document
	context         *middleware.Context
	handlers        map[string]map[string]http.Handler
	formats         strfmt.Registry
	customConsumers map[string]runtime.Consumer
	customProducers map[string]runtime.Producer
	defaultConsumes string
	defaultProduces string
	Middleware      func(middleware.Builder) http.Handler

	// BasicAuthenticator generates a runtime.Authenticator from the supplied basic auth function.
	// It has a default implementation in the security package, however you can replace it for your particular usage.
	BasicAuthenticator func(security.UserPassAuthentication) runtime.Authenticator
	// APIKeyAuthenticator generates a runtime.Authenticator from the supplied token auth function.
	// It has a default implementation in the security package, however you can replace it for your particular usage.
	APIKeyAuthenticator func(string, string, security.TokenAuthentication) runtime.Authenticator
	// BearerAuthenticator generates a runtime.Authenticator from the supplied bearer token auth function.
	// It has a default implementation in the security package, however you can replace it for your particular usage.
	BearerAuthenticator func(string, security.ScopedTokenAuthentication) runtime.Authenticator

	// JSONConsumer registers a consumer for the following mime types:
	//   - application/json
	JSONConsumer runtime.Consumer

	// JSONProducer registers a producer for the following mime types:
	//   - application/json
	JSONProducer runtime.Producer

	// RolesAuth registers a function that takes a token and returns a principal
	// it performs authentication based on an api key X-Auth-Roles provided in the header
	RolesAuth func(string) (interface{}, error)

	// APIAuthorizer provides access control (ACL/RBAC/ABAC) by providing access to the request and authenticated principal
	APIAuthorizer runtime.Authorizer

	// StoreInventoryGetHandler sets the operation handler for the inventory get operation
	StoreInventoryGetHandler store.InventoryGetHandler
	// StoreOrderCreateHandler sets the operation handler for the order create operation
	StoreOrderCreateHandler store.OrderCreateHandler
	// StoreOrderDeleteHandler sets the operation handler for the order delete operation
	StoreOrderDeleteHandler store.OrderDeleteHandler
	// StoreOrderGetHandler sets the operation handler for the order get operation
	StoreOrderGetHandler store.OrderGetHandler
	// PetPetCreateHandler sets the operation handler for the pet create operation
	PetPetCreateHandler pet.PetCreateHandler
	// PetPetDeleteHandler sets the operation handler for the pet delete operation
	PetPetDeleteHandler pet.PetDeleteHandler
	// PetPetGetHandler sets the operation handler for the pet get operation
	PetPetGetHandler pet.PetGetHandler
	// PetPetListHandler sets the operation handler for the pet list operation
	PetPetListHandler pet.PetListHandler
	// PetPetUpdateHandler sets the operation handler for the pet update operation
	PetPetUpdateHandler pet.PetUpdateHandler
	// PetPetUploadImageHandler sets the operation handler for the pet upload image operation
	PetPetUploadImageHandler pet.PetUploadImageHandler
	// ServeError is called when an error is received, there is a default handler
	// but you can set your own with this
	ServeError func(http.ResponseWriter, *http.Request, error)

	// PreServerShutdown is called before the HTTP(S) server is shutdown
	// This allows for custom functions to get executed before the HTTP(S) server stops accepting traffic
	PreServerShutdown func()

	// ServerShutdown is called when the HTTP(S) server is shut down and done
	// handling all active connections and does not accept connections any more
	ServerShutdown func()

	// Custom command line argument groups with their descriptions
	CommandLineOptionsGroups []swag.CommandLineOptionsGroup

	// User defined logger function.
	Logger func(string, ...interface{})
}

// SetDefaultProduces sets the default produces media type
func (o *PetstoreAPI) SetDefaultProduces(mediaType string) {
	o.defaultProduces = mediaType
}

// SetDefaultConsumes returns the default consumes media type
func (o *PetstoreAPI) SetDefaultConsumes(mediaType string) {
	o.defaultConsumes = mediaType
}

// SetSpec sets a spec that will be served for the clients.
func (o *PetstoreAPI) SetSpec(spec *loads.Document) {
	o.spec = spec
}

// DefaultProduces returns the default produces media type
func (o *PetstoreAPI) DefaultProduces() string {
	return o.defaultProduces
}

// DefaultConsumes returns the default consumes media type
func (o *PetstoreAPI) DefaultConsumes() string {
	return o.defaultConsumes
}

// Formats returns the registered string formats
func (o *PetstoreAPI) Formats() strfmt.Registry {
	return o.formats
}

// RegisterFormat registers a custom format validator
func (o *PetstoreAPI) RegisterFormat(name string, format strfmt.Format, validator strfmt.Validator) {
	o.formats.Add(name, format, validator)
}

// Validate validates the registrations in the PetstoreAPI
func (o *PetstoreAPI) Validate() error {
	var unregistered []string

	if o.JSONConsumer == nil {
		unregistered = append(unregistered, "JSONConsumer")
	}

	if o.JSONProducer == nil {
		unregistered = append(unregistered, "JSONProducer")
	}

	if o.RolesAuth == nil {
		unregistered = append(unregistered, "XAuthRolesAuth")
	}

	if o.StoreInventoryGetHandler == nil {
		unregistered = append(unregistered, "store.InventoryGetHandler")
	}
	if o.StoreOrderCreateHandler == nil {
		unregistered = append(unregistered, "store.OrderCreateHandler")
	}
	if o.StoreOrderDeleteHandler == nil {
		unregistered = append(unregistered, "store.OrderDeleteHandler")
	}
	if o.StoreOrderGetHandler == nil {
		unregistered = append(unregistered, "store.OrderGetHandler")
	}
	if o.PetPetCreateHandler == nil {
		unregistered = append(unregistered, "pet.PetCreateHandler")
	}
	if o.PetPetDeleteHandler == nil {
		unregistered = append(unregistered, "pet.PetDeleteHandler")
	}
	if o.PetPetGetHandler == nil {
		unregistered = append(unregistered, "pet.PetGetHandler")
	}
	if o.PetPetListHandler == nil {
		unregistered = append(unregistered, "pet.PetListHandler")
	}
	if o.PetPetUpdateHandler == nil {
		unregistered = append(unregistered, "pet.PetUpdateHandler")
	}
	if o.PetPetUploadImageHandler == nil {
		unregistered = append(unregistered, "pet.PetUploadImageHandler")
	}

	if len(unregistered) > 0 {
		return fmt.Errorf("missing registration: %s", strings.Join(unregistered, ", "))
	}

	return nil
}

// ServeErrorFor gets a error handler for a given operation id
func (o *PetstoreAPI) ServeErrorFor(operationID string) func(http.ResponseWriter, *http.Request, error) {
	return o.ServeError
}

// AuthenticatorsFor gets the authenticators for the specified security schemes
func (o *PetstoreAPI) AuthenticatorsFor(schemes map[string]spec.SecurityScheme) map[string]runtime.Authenticator {
	result := make(map[string]runtime.Authenticator)
	for name := range schemes {
		switch name {
		case "roles":
			scheme := schemes[name]
			result[name] = o.APIKeyAuthenticator(scheme.Name, scheme.In, o.RolesAuth)

		}
	}
	return result
}

// Authorizer returns the registered authorizer
func (o *PetstoreAPI) Authorizer() runtime.Authorizer {
	return o.APIAuthorizer
}

// ConsumersFor gets the consumers for the specified media types.
// MIME type parameters are ignored here.
func (o *PetstoreAPI) ConsumersFor(mediaTypes []string) map[string]runtime.Consumer {
	result := make(map[string]runtime.Consumer, len(mediaTypes))
	for _, mt := range mediaTypes {
		switch mt {
		case "application/json":
			result["application/json"] = o.JSONConsumer
		}

		if c, ok := o.customConsumers[mt]; ok {
			result[mt] = c
		}
	}
	return result
}

// ProducersFor gets the producers for the specified media types.
// MIME type parameters are ignored here.
func (o *PetstoreAPI) ProducersFor(mediaTypes []string) map[string]runtime.Producer {
	result := make(map[string]runtime.Producer, len(mediaTypes))
	for _, mt := range mediaTypes {
		switch mt {
		case "application/json":
			result["application/json"] = o.JSONProducer
		}

		if p, ok := o.customProducers[mt]; ok {
			result[mt] = p
		}
	}
	return result
}

// HandlerFor gets a http.Handler for the provided operation method and path
func (o *PetstoreAPI) HandlerFor(method, path string) (http.Handler, bool) {
	if o.handlers == nil {
		return nil, false
	}
	um := strings.ToUpper(method)
	if _, ok := o.handlers[um]; !ok {
		return nil, false
	}
	if path == "/" {
		path = ""
	}
	h, ok := o.handlers[um][path]
	return h, ok
}

// Context returns the middleware context for the petstore API
func (o *PetstoreAPI) Context() *middleware.Context {
	if o.context == nil {
		o.context = middleware.NewRoutableContext(o.spec, o, nil)
	}

	return o.context
}

func (o *PetstoreAPI) initHandlerCache() {
	o.Context() // don't care about the result, just that the initialization happened
	if o.handlers == nil {
		o.handlers = make(map[string]map[string]http.Handler)
	}

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/store/inventory"] = store.NewInventoryGet(o.context, o.StoreInventoryGetHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/store/order"] = store.NewOrderCreate(o.context, o.StoreOrderCreateHandler)
	if o.handlers["DELETE"] == nil {
		o.handlers["DELETE"] = make(map[string]http.Handler)
	}
	o.handlers["DELETE"]["/store/order/{orderId}"] = store.NewOrderDelete(o.context, o.StoreOrderDeleteHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/store/order/{orderId}"] = store.NewOrderGet(o.context, o.StoreOrderGetHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/pet"] = pet.NewPetCreate(o.context, o.PetPetCreateHandler)
	if o.handlers["DELETE"] == nil {
		o.handlers["DELETE"] = make(map[string]http.Handler)
	}
	o.handlers["DELETE"]["/pet/{petId}"] = pet.NewPetDelete(o.context, o.PetPetDeleteHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/pet/{petId}"] = pet.NewPetGet(o.context, o.PetPetGetHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/pet"] = pet.NewPetList(o.context, o.PetPetListHandler)
	if o.handlers["PUT"] == nil {
		o.handlers["PUT"] = make(map[string]http.Handler)
	}
	o.handlers["PUT"]["/pet"] = pet.NewPetUpdate(o.context, o.PetPetUpdateHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/pet/{petId}/image"] = pet.NewPetUploadImage(o.context, o.PetPetUploadImageHandler)
}

// Serve creates a http handler to serve the API over HTTP
// can be used directly in http.ListenAndServe(":8000", api.Serve(nil))
func (o *PetstoreAPI) Serve(builder middleware.Builder) http.Handler {
	o.Init()

	if o.Middleware != nil {
		return o.Middleware(builder)
	}
	return o.context.APIHandler(builder)
}

// Init allows you to just initialize the handler cache, you can then recompose the middleware as you see fit
func (o *PetstoreAPI) Init() {
	if len(o.handlers) == 0 {
		o.initHandlerCache()
	}
}

// RegisterConsumer allows you to add (or override) a consumer for a media type.
func (o *PetstoreAPI) RegisterConsumer(mediaType string, consumer runtime.Consumer) {
	o.customConsumers[mediaType] = consumer
}

// RegisterProducer allows you to add (or override) a producer for a media type.
func (o *PetstoreAPI) RegisterProducer(mediaType string, producer runtime.Producer) {
	o.customProducers[mediaType] = producer
}
