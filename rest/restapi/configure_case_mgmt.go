// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	"github.com/case-management-suite/api/rest/restapi/ops"
)

//go:generate swagger generate server --target ../../rest --name CaseMgmt --spec ../swagger.yml --api-package ops --principal interface{} --exclude-main

func configureFlags(api *ops.CaseMgmtAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *ops.CaseMgmtAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// api.Logger = log.Printf

	api.UseSwaggerUI()
	// To continue using redoc as your UI, uncomment the following line
	// api.UseRedoc()

	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()

	if api.CreateCaseHandler == nil {
		api.CreateCaseHandler = ops.CreateCaseHandlerFunc(func(params ops.CreateCaseParams) middleware.Responder {
			return middleware.NotImplemented("operation ops.CreateCase has not yet been implemented")
		})
	}
	if api.ExecuteActionHandler == nil {
		api.ExecuteActionHandler = ops.ExecuteActionHandlerFunc(func(params ops.ExecuteActionParams) middleware.Responder {
			return middleware.NotImplemented("operation ops.ExecuteAction has not yet been implemented")
		})
	}
	if api.FindCaseHandler == nil {
		api.FindCaseHandler = ops.FindCaseHandlerFunc(func(params ops.FindCaseParams) middleware.Responder {
			return middleware.NotImplemented("operation ops.FindCase has not yet been implemented")
		})
	}
	if api.FindCasesHandler == nil {
		api.FindCasesHandler = ops.FindCasesHandlerFunc(func(params ops.FindCasesParams) middleware.Responder {
			return middleware.NotImplemented("operation ops.FindCases has not yet been implemented")
		})
	}
	if api.GetActionRecordsHandler == nil {
		api.GetActionRecordsHandler = ops.GetActionRecordsHandlerFunc(func(params ops.GetActionRecordsParams) middleware.Responder {
			return middleware.NotImplemented("operation ops.GetActionRecords has not yet been implemented")
		})
	}

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
	return handler
}
