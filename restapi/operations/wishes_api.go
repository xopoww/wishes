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

	"github.com/xopoww/wishes/models"
)

// NewWishesAPI creates a new Wishes instance
func NewWishesAPI(spec *loads.Document) *WishesAPI {
	return &WishesAPI{
		handlers:            make(map[string]map[string]http.Handler),
		formats:             strfmt.Default,
		defaultConsumes:     "application/json",
		defaultProduces:     "application/json",
		customConsumers:     make(map[string]runtime.Consumer),
		customProducers:     make(map[string]runtime.Producer),
		PreServerShutdown:   func() {},
		ServerShutdown:      func() {},
		spec:                spec,
		useSwaggerUI:        false,
		ServeError:          errors.ServeError,
		BasicAuthenticator:  security.BasicAuth,
		APIKeyAuthenticator: security.APIKeyAuth,
		BearerAuthenticator: security.BearerAuth,

		JSONConsumer: runtime.JSONConsumer(),

		JSONProducer: runtime.JSONProducer(),

		DeleteListHandler: DeleteListHandlerFunc(func(params DeleteListParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation DeleteList has not yet been implemented")
		}),
		GetListHandler: GetListHandlerFunc(func(params GetListParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation GetList has not yet been implemented")
		}),
		GetUserHandler: GetUserHandlerFunc(func(params GetUserParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation GetUser has not yet been implemented")
		}),
		LoginHandler: LoginHandlerFunc(func(params LoginParams) middleware.Responder {
			return middleware.NotImplemented("operation Login has not yet been implemented")
		}),
		PatchListHandler: PatchListHandlerFunc(func(params PatchListParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation PatchList has not yet been implemented")
		}),
		PatchUserHandler: PatchUserHandlerFunc(func(params PatchUserParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation PatchUser has not yet been implemented")
		}),
		PostListHandler: PostListHandlerFunc(func(params PostListParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation PostList has not yet been implemented")
		}),
		PostUserHandler: PostUserHandlerFunc(func(params PostUserParams) middleware.Responder {
			return middleware.NotImplemented("operation PostUser has not yet been implemented")
		}),

		// Applies when the "x-token" header is set
		KeySecurityAuth: func(token string) (*models.Principal, error) {
			return nil, errors.NotImplemented("api key auth (KeySecurity) x-token from header param [x-token] has not yet been implemented")
		},
		// default authorizer is authorized meaning no requests are blocked
		APIAuthorizer: security.Authorized(),
	}
}

/*WishesAPI the wishes API */
type WishesAPI struct {
	spec            *loads.Document
	context         *middleware.Context
	handlers        map[string]map[string]http.Handler
	formats         strfmt.Registry
	customConsumers map[string]runtime.Consumer
	customProducers map[string]runtime.Producer
	defaultConsumes string
	defaultProduces string
	Middleware      func(middleware.Builder) http.Handler
	useSwaggerUI    bool

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

	// KeySecurityAuth registers a function that takes a token and returns a principal
	// it performs authentication based on an api key x-token provided in the header
	KeySecurityAuth func(string) (*models.Principal, error)

	// APIAuthorizer provides access control (ACL/RBAC/ABAC) by providing access to the request and authenticated principal
	APIAuthorizer runtime.Authorizer

	// DeleteListHandler sets the operation handler for the delete list operation
	DeleteListHandler DeleteListHandler
	// GetListHandler sets the operation handler for the get list operation
	GetListHandler GetListHandler
	// GetUserHandler sets the operation handler for the get user operation
	GetUserHandler GetUserHandler
	// LoginHandler sets the operation handler for the login operation
	LoginHandler LoginHandler
	// PatchListHandler sets the operation handler for the patch list operation
	PatchListHandler PatchListHandler
	// PatchUserHandler sets the operation handler for the patch user operation
	PatchUserHandler PatchUserHandler
	// PostListHandler sets the operation handler for the post list operation
	PostListHandler PostListHandler
	// PostUserHandler sets the operation handler for the post user operation
	PostUserHandler PostUserHandler

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

// UseRedoc for documentation at /docs
func (o *WishesAPI) UseRedoc() {
	o.useSwaggerUI = false
}

// UseSwaggerUI for documentation at /docs
func (o *WishesAPI) UseSwaggerUI() {
	o.useSwaggerUI = true
}

// SetDefaultProduces sets the default produces media type
func (o *WishesAPI) SetDefaultProduces(mediaType string) {
	o.defaultProduces = mediaType
}

// SetDefaultConsumes returns the default consumes media type
func (o *WishesAPI) SetDefaultConsumes(mediaType string) {
	o.defaultConsumes = mediaType
}

// SetSpec sets a spec that will be served for the clients.
func (o *WishesAPI) SetSpec(spec *loads.Document) {
	o.spec = spec
}

// DefaultProduces returns the default produces media type
func (o *WishesAPI) DefaultProduces() string {
	return o.defaultProduces
}

// DefaultConsumes returns the default consumes media type
func (o *WishesAPI) DefaultConsumes() string {
	return o.defaultConsumes
}

// Formats returns the registered string formats
func (o *WishesAPI) Formats() strfmt.Registry {
	return o.formats
}

// RegisterFormat registers a custom format validator
func (o *WishesAPI) RegisterFormat(name string, format strfmt.Format, validator strfmt.Validator) {
	o.formats.Add(name, format, validator)
}

// Validate validates the registrations in the WishesAPI
func (o *WishesAPI) Validate() error {
	var unregistered []string

	if o.JSONConsumer == nil {
		unregistered = append(unregistered, "JSONConsumer")
	}

	if o.JSONProducer == nil {
		unregistered = append(unregistered, "JSONProducer")
	}

	if o.KeySecurityAuth == nil {
		unregistered = append(unregistered, "XTokenAuth")
	}

	if o.DeleteListHandler == nil {
		unregistered = append(unregistered, "DeleteListHandler")
	}
	if o.GetListHandler == nil {
		unregistered = append(unregistered, "GetListHandler")
	}
	if o.GetUserHandler == nil {
		unregistered = append(unregistered, "GetUserHandler")
	}
	if o.LoginHandler == nil {
		unregistered = append(unregistered, "LoginHandler")
	}
	if o.PatchListHandler == nil {
		unregistered = append(unregistered, "PatchListHandler")
	}
	if o.PatchUserHandler == nil {
		unregistered = append(unregistered, "PatchUserHandler")
	}
	if o.PostListHandler == nil {
		unregistered = append(unregistered, "PostListHandler")
	}
	if o.PostUserHandler == nil {
		unregistered = append(unregistered, "PostUserHandler")
	}

	if len(unregistered) > 0 {
		return fmt.Errorf("missing registration: %s", strings.Join(unregistered, ", "))
	}

	return nil
}

// ServeErrorFor gets a error handler for a given operation id
func (o *WishesAPI) ServeErrorFor(operationID string) func(http.ResponseWriter, *http.Request, error) {
	return o.ServeError
}

// AuthenticatorsFor gets the authenticators for the specified security schemes
func (o *WishesAPI) AuthenticatorsFor(schemes map[string]spec.SecurityScheme) map[string]runtime.Authenticator {
	result := make(map[string]runtime.Authenticator)
	for name := range schemes {
		switch name {
		case "KeySecurity":
			scheme := schemes[name]
			result[name] = o.APIKeyAuthenticator(scheme.Name, scheme.In, func(token string) (interface{}, error) {
				return o.KeySecurityAuth(token)
			})

		}
	}
	return result
}

// Authorizer returns the registered authorizer
func (o *WishesAPI) Authorizer() runtime.Authorizer {
	return o.APIAuthorizer
}

// ConsumersFor gets the consumers for the specified media types.
// MIME type parameters are ignored here.
func (o *WishesAPI) ConsumersFor(mediaTypes []string) map[string]runtime.Consumer {
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
func (o *WishesAPI) ProducersFor(mediaTypes []string) map[string]runtime.Producer {
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
func (o *WishesAPI) HandlerFor(method, path string) (http.Handler, bool) {
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

// Context returns the middleware context for the wishes API
func (o *WishesAPI) Context() *middleware.Context {
	if o.context == nil {
		o.context = middleware.NewRoutableContext(o.spec, o, nil)
	}

	return o.context
}

func (o *WishesAPI) initHandlerCache() {
	o.Context() // don't care about the result, just that the initialization happened
	if o.handlers == nil {
		o.handlers = make(map[string]map[string]http.Handler)
	}

	if o.handlers["DELETE"] == nil {
		o.handlers["DELETE"] = make(map[string]http.Handler)
	}
	o.handlers["DELETE"]["/list"] = NewDeleteList(o.context, o.DeleteListHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/list"] = NewGetList(o.context, o.GetListHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/user"] = NewGetUser(o.context, o.GetUserHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/login"] = NewLogin(o.context, o.LoginHandler)
	if o.handlers["PATCH"] == nil {
		o.handlers["PATCH"] = make(map[string]http.Handler)
	}
	o.handlers["PATCH"]["/list"] = NewPatchList(o.context, o.PatchListHandler)
	if o.handlers["PATCH"] == nil {
		o.handlers["PATCH"] = make(map[string]http.Handler)
	}
	o.handlers["PATCH"]["/user"] = NewPatchUser(o.context, o.PatchUserHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/list"] = NewPostList(o.context, o.PostListHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/user"] = NewPostUser(o.context, o.PostUserHandler)
}

// Serve creates a http handler to serve the API over HTTP
// can be used directly in http.ListenAndServe(":8000", api.Serve(nil))
func (o *WishesAPI) Serve(builder middleware.Builder) http.Handler {
	o.Init()

	if o.Middleware != nil {
		return o.Middleware(builder)
	}
	if o.useSwaggerUI {
		return o.context.APIHandlerSwaggerUI(builder)
	}
	return o.context.APIHandler(builder)
}

// Init allows you to just initialize the handler cache, you can then recompose the middleware as you see fit
func (o *WishesAPI) Init() {
	if len(o.handlers) == 0 {
		o.initHandlerCache()
	}
}

// RegisterConsumer allows you to add (or override) a consumer for a media type.
func (o *WishesAPI) RegisterConsumer(mediaType string, consumer runtime.Consumer) {
	o.customConsumers[mediaType] = consumer
}

// RegisterProducer allows you to add (or override) a producer for a media type.
func (o *WishesAPI) RegisterProducer(mediaType string, producer runtime.Producer) {
	o.customProducers[mediaType] = producer
}

// AddMiddlewareFor adds a http middleware to existing handler
func (o *WishesAPI) AddMiddlewareFor(method, path string, builder middleware.Builder) {
	um := strings.ToUpper(method)
	if path == "/" {
		path = ""
	}
	o.Init()
	if h, ok := o.handlers[um][path]; ok {
		o.handlers[method][path] = builder(h)
	}
}
