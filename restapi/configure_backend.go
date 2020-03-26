// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	"github.com/julianpitt/go-swagger-test/restapi/models"
	"github.com/julianpitt/go-swagger-test/restapi/operations"
	"github.com/julianpitt/go-swagger-test/restapi/operations/hello"
)

//go:generate swagger generate server --target ../../go-swagger-test --name Backend --spec ../swagger.yaml --model-package restapi/models

func configureFlags(api *operations.BackendAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func sayHello() (message *models.Message) {
	message = &models.Message{
		Message: "hello",
	}
	return
}

func configureAPI(api *operations.BackendAPI) http.Handler {

	api.ServeError = errors.ServeError

	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()

	api.HelloSayHelloHandler = hello.SayHelloHandlerFunc(func(params hello.SayHelloParams) middleware.Responder {
		return hello.NewSayHelloOK().WithPayload(sayHello())
	})

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
// scheme value will be set accordingly: "http", "https" or "unix"
func configureServer(s *http.Server, scheme, addr string) {
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	return handler
}
