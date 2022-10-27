// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"log"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	"github.com/xopoww/wishes/internal/auth"
	"github.com/xopoww/wishes/internal/db"
	"github.com/xopoww/wishes/internal/handlers"
	"github.com/xopoww/wishes/internal/meta"
	"github.com/xopoww/wishes/models"
	"github.com/xopoww/wishes/restapi/operations"
)

//go:generate swagger generate server --target ../../wishes --name Wishes --spec ../api/wishes.yml --principal models.Principal

func configureFlags(api *operations.WishesAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.WishesAPI) http.Handler {
	//TODO: move somewhere else
	if err := db.Connect("devdata/db.sqlite3"); err != nil {
		log.Fatalf("connect: %s", err)
	}
	if db.CheckUser("test") != db.ErrNameTaken {
		hash, err := auth.HashPassword("test")
		if err != nil {
			log.Fatalf("hash test pwd: %s", err)
		}
		_, err = db.AddUser("test", hash)
		if err != nil {
			log.Fatalf("add test user: %s", err)
		}
	}

	log.Printf("build version: %s", meta.BuildVersion)
	log.Printf("build date: %s", meta.BuildDate)

	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	api.Logger = log.Printf

	api.UseSwaggerUI()
	// To continue using redoc as your UI, uncomment the following line
	// api.UseRedoc()

	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()

	// Applies when the "x-token" header is set
	api.KeySecurityAuth = func(token string) (*models.Principal, error) {
		principal, err := auth.ValidateToken(token)
		if err != nil {
			api.Logger("incorrect api token: %s (token=%s)", err, token)
			return nil, errors.New(401, "incorrect api key auth")
		}
		return principal, nil
	}

	// Set your custom authorizer if needed. Default one is security.Authorized()
	// Expected interface runtime.Authorizer
	//
	// Example:
	// api.APIAuthorizer = security.Authorized()

	api.GetFooHandler = operations.GetFooHandlerFunc(func(params operations.GetFooParams, principal *models.Principal) middleware.Responder {
		api.Logger("authenticated request from %q", *principal)
		return middleware.NotImplemented("operation operations.GetFoo has not yet been implemented")
	})
	api.LoginHandler = handlers.Login()

	api.PreServerShutdown = func() {}

	api.ServerShutdown = func() {}

	return setupGlobalMiddleware(api.Serve(setupMiddlewares))
}

// The TLS configuration before HTTPS server starts.
func configureTLS(tlsConfig *tls.Config) {
	// Make all necessary changes to the TLS configuration here.
}

// As soon as server is initialized but not run yet, this function will be called.
// If you need to modify a config, store server instance to stop it individually later, this is the place.
// This function can be called multiple times, depending on the number of serving schemes.
// scheme value will be set accordingly: "http", "https" or "unix".
func configureServer(s *http.Server, scheme, addr string) {
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation.
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics.
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if r := recover(); r != nil {
				log.Printf("recovered: %v", r)
			}
		}()

		handler.ServeHTTP(w, r)
	})
}