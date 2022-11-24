// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"encoding/base64"
	"net/http"
	"os"
	"runtime/debug"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"

	"github.com/xopoww/wishes/internal/controllers/handlers"
	"github.com/xopoww/wishes/internal/log"
	"github.com/xopoww/wishes/internal/meta"
	"github.com/xopoww/wishes/internal/oauth/yandex"
	"github.com/xopoww/wishes/internal/repository/sqlite"
	"github.com/xopoww/wishes/internal/service"
	"github.com/xopoww/wishes/restapi/operations"

	"github.com/rs/zerolog/hlog"
)

//go:generate go run ./clean.go --quiet
//go:generate swagger generate server --quiet --target ../../wishes --name Wishes --spec ../api/wishes.yml --principal apimodels.Principal -m restapi/apimodels --skip-tag-packages

func configureFlags(api *operations.WishesAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.WishesAPI) http.Handler {
	l := log.Logger()

	//TODO: move somewhere else
	dbs, exists := os.LookupEnv("WISHES_DBS")
	if !exists {
		l.Fatal().Msg("WISHES_DBS is not set")
	}
	repo, err := sqlite.NewRepository(dbs, log.Sqlite(l))
	if err != nil {
		l.Fatal().Err(err).Msg("connect to repository failed")
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
	api.UseRedoc()

	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()

	// Set your custom authorizer if needed. Default one is security.Authorized()
	// Expected interface runtime.Authorizer
	//
	// Example:
	// api.APIAuthorizer = security.Authorized()

	listSecretStr, exists := os.LookupEnv("WISHES_LIST_SECRET")
	if !exists {
		l.Fatal().Msg("WISHES_LIST_SECRET is not set")
	}
	listSecret, err := base64.RawStdEncoding.DecodeString(listSecretStr)
	if err != nil {
		l.Fatal().Err(err).Msg("decode WISHES_LIST_SECRET failed")
	}
	serv := service.NewService(repo, service.NewListTokenProvider(listSecret))

	if cid, exists := os.LookupEnv("WISHES_OAUTH_YANDEX_CLIENT_ID"); exists {
		serv.AddOAuthProvider("yandex", yandex.NewOAuthProvider(log.YandexOAuth(l), cid))
		l.Debug().Str("provider", "yandex").Msg("added oauth provider")
	}

	controller := handlers.NewApiController(log.Handlers(l), serv)

	// Applies when the "x-token" header is set
	api.KeySecurityAuth = controller.KeySecurityAuth()

	api.LoginHandler = controller.Login()

	api.GetUserHandler = controller.GetUser()
	api.PatchUserHandler = controller.PatchUser()
	api.RegisterHandler = controller.Register()

	api.GetUserListsHandler = controller.GetUserLists()

	api.GetListHandler = controller.GetList()
	api.GetListItemsHandler = controller.GetListItems()
	api.PostListHandler = controller.PostList()
	api.PatchListHandler = controller.PatchList()
	api.DeleteListHandler = controller.DeleteList()
	api.GetListTokenHandler = controller.GetListToken()
	api.PostListItemsHandler = controller.PostListItems()
	api.DeleteListItemsHandler = controller.DeleteListItems()
	api.PostItemTakenHandler = controller.PostItemTaken()
	api.DeleteItemTakenHandler = controller.DeleteItemTaken()

	api.OAuthRegisterHandler = controller.OAuthRegister()
	api.OAuthLoginHandler = controller.OAuthLogin()

	api.PreServerShutdown = func() {}

	api.ServerShutdown = func() {
		if err := repo.Close(); err != nil {
			l.Error().Err(err).Msg("db disconnect failed")
		}
	}

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
				l.Printf("stack trace: %s", string(debug.Stack()))
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
