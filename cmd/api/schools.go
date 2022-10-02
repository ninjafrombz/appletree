// Filename:
package main

import (
	"fmt"
	"net/http"
	"time"

	"appletree.desireamagwula.net/internals/data"
)

// CreateSchoolHandler for the POST /v1/schools" endpoint

func (app *application) createSchoolHandler(w http.ResponseWriter, r *http.Request) {
	// Our target decode destination fmt.Fprintln(w, "create a new school..")
	var input struct {
		Name    string   `json:"name"`
		Level   string   `json:"level"`
		Contact string   `json:"contact"`
		Phone   string   `json:"phone"`
		Email   string   `json:"email"`
		Website string   `json:"website"`
		Address string   `json:"address"`
		Mode    []string `json:"mode"`
	}

	// Initialize a new json.Decoder instance
	err := app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}
	// Display the request
	fmt.Fprintf(w, "%+v\n", input)

}

func (app *application) showSchoolHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	//Create a new instance of the school struct containing the ID we extracted
	// from our URL and some sample data

	school := data.School{
		ID:        id,
		CreatedAt: time.Now(),
		Name:      "Apple Tree",
		Level:     "High School",
		Contact:   "Anna Smith",
		Phone:     "654-1651",
		Address:   "14 Apple Steet",
		Mode:      []string{"blended", "online"},
		Version:   1,
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"school": school}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
