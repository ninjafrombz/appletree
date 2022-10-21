// Filename

package main 

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
)


func (app *application) routes() *httprouter.Router {
	// Create
	router := httprouter.New()
	router.NotFound = http.HandlerFunc(app.notFoundResponse)
	router.MethodNotAllowed = http.HandlerFunc(app.methodNotAllowedResponse)
	router.HandlerFunc(http.MethodGet, "/v1/healthcheck", app.healthcheckHandler)
	router.HandlerFunc(http.MethodPost, "/v1/schools", app.createSchoolHandler)
	router.HandlerFunc(http.MethodGet, "/v1/schools/:id", app.showSchoolHandler)
	router.HandlerFunc(http.MethodPut, "/v1/schools/:id", app.updateSchoolHandler)

	return router
}