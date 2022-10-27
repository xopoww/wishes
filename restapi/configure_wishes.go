// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"

	"github.com/xopoww/wishes/internal/auth"
	"github.com/xopoww/wishes/internal/db"
	"github.com/xopoww/wishes/internal/handlers"
	"github.com/xopoww/wishes/internal/log"
	"github.com/xopoww/wishes/internal/meta"
	"github.com/xopoww/wishes/restapi/operations"

	"github.com/rs/zerolog/hlog"
)

//go:generate swagger generate server --target ../../wishes --name Wishes --spec ../api/wishes.yml --principal models.Principal

func configureFlags(api *operations.WishesAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.WishesAPI) http.Handler {
	l := log.Logger()
	log.WithTraces(l)

	//TODO: move somewhere else
	if err := db.Connect("devdata/db.sqlite3"); err != nil {
		l.Fatal().Err(err).Msg("connect failed")
	}
	if db.CheckUser("test") != db.ErrNameTaken {
		hash, err := auth.HashPassword("test")
		if err != nil {
			l.Fatal().Err(err).Msg("hash test pwd failed")
		}
		_, err = db.AddUser("test", hash)
		if err != nil {
			l.Fatal().Err(err).Msg("add test user failed")
		}
	}

	l.Debug().
		Str("build version", meta.BuildVersion).
		Time("build date", meta.BuildDate).
		Send()

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

	// Set your custom authorizer if needed. Default one is security.Authorized()
	// Expected interface runtime.Authorizer
	//
	// Example:
	// api.APIAuthorizer = security.Authorized()

	ht := log.Handlers(l)

	// Applies when the "x-token" header is set
	api.KeySecurityAuth = handlers.KeySecurityAuth(ht)

	api.LoginHandler = handlers.Login(ht)

	api.GetUserHandler = handlers.GetUser(ht)
	api.PatchUserHandler = handlers.PatchUser(ht)
	api.PostUserHandler = handlers.PostUser(ht)

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
	l := log.Logger()

	wrapped := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if r := recover(); r != nil {
				l.Error().
					Interface("reason", r).
					Msg("recovered in global middleware")
			}
		}()

		handler.ServeHTTP(w, r)
	})

	return hlog.AccessHandler(func(r *http.Request, status, size int, duration time.Duration) {
		l.Debug().
			Str("method", r.Method).
			Str("path", r.URL.Path).
			Str("raddr", r.RemoteAddr).
			Int("status", status).
			Dur("latency", duration).
			Msg("request processed")

	})(wrapped)
}
