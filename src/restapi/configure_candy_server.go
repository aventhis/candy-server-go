// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"fmt"
	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"net/http"

	"github.com/aventhis/candy-server-go/restapi/operations"
)

//go:generate swagger generate server --target ../../src --name CandyServer --spec ../swagger.yaml --principal interface{}

func configureFlags(api *operations.CandyServerAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.CandyServerAPI) http.Handler {

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

	api.BuyCandyHandler = operations.BuyCandyHandlerFunc(func(params operations.BuyCandyParams) middleware.Responder {
		candyPrice := map[string]int64{
			"CE": 10,
			"AA": 15,
			"NT": 17,
			"DE": 21,
			"YR": 23,
		}

		//проверка на входные данные
		price, exist := candyPrice[*params.Order.CandyType]
		if !exist {
			return operations.NewBuyCandyBadRequest().WithPayload(&operations.BuyCandyBadRequestBody{
				Error: "Invalid candy type",
			})
		}

		if *params.Order.CandyCount <= 0 {
			return operations.NewBuyCandyBadRequest().WithPayload(&operations.BuyCandyBadRequestBody{
				Error: "Candy count must be greater than zero",
			})
		}

		totalCost := price * (*params.Order.CandyCount)

		if totalCost > *params.Order.Money {
			return operations.NewBuyCandyPaymentRequired().WithPayload(&operations.BuyCandyPaymentRequiredBody{
				Error: fmt.Sprintf("You need %v more money", totalCost-*params.Order.Money),
			})
		}
		change := *params.Order.Money - totalCost
		return operations.NewBuyCandyCreated().WithPayload(&operations.BuyCandyCreatedBody{
			Thanks: "Thank you",
			Change: change,
		})
	})

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
