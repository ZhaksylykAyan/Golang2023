package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func (app *application) routes() http.Handler {
	router := httprouter.New()
	router.NotFound = http.HandlerFunc(app.notFoundResponse)
	router.MethodNotAllowed = http.HandlerFunc(app.methodNotAllowedResponse)

	router.HandlerFunc(http.MethodGet, "/v1/healthcheck", app.healthcheckHandler)
	router.HandlerFunc(http.MethodGet, "/v1/accessories", app.listAccessoriesHandler)
	router.HandlerFunc(http.MethodPost, "/v1/accessories", app.createAccessoriesHandler)
	router.HandlerFunc(http.MethodGet, "/v1/accessories/:id", app.showAccessoriesHandler)
	router.HandlerFunc(http.MethodPatch, "/v1/accessories/:id", app.updateAccessoryHandler)
	router.HandlerFunc(http.MethodDelete, "/v1/accessories/:id", app.deleteAccessoryHandler)
	router.HandlerFunc(http.MethodPost, "/v1/users", app.registerUserHandler)
	router.HandlerFunc(http.MethodPut, "/v1/users/activated", app.activateUserHandler)
	router.HandlerFunc(http.MethodPost, "/v1/tokens/authentication", app.createAuthenticationTokenHandler)

	return app.recoverPanic(app.rateLimit(app.authenticate(router)))
}
