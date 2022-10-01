// Filename:
package main

import (
	"fmt"
	"net/http"
)

// CreateSchoolHandler for the POST /v1/schools" endpoint

func (app *application) createSchoolHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "create a new school..")
}

func (app *application) showSchoolHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		http.NotFound(w, r)
		return
	}
	// fmt.Fprintf(w, "show the details for school %d\n", id)
	fmt.Fprintf(w, "Show the details for school %d\n", id)
}
