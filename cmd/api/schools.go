// Filename:
package main

import (
	"fmt"
	"net/http"
	"time"

	"appletree.desireamagwula.net/internals/data"
	"appletree.desireamagwula.net/internals/validator"
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

	// Copy the values from the input struct to a new school struct 
	school := &data.School{
		Name: input.Name,
		Level: input.Level,
		Contact: input.Contact,
		Phone: input.Phone,
		Email: input.Email,
		Website: input.Website,
		Address: input.Address,
		Mode: input.Mode,
	}

	//Initialize a new validator instance 
	v := validator.New()

	// Check the map to determine if there were any validation errors 
	if data.ValidateSchool(v, school); !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}
	// CReate a school
	err = app.models.Schools.Insert(school)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}

	// CReate a location header for the newly created 
	headers := make(http.Header)
	headers.Set("Location", fmt.Sprintf("/v1/schools/%d", school.ID))
	//Write the JSON response with 201 - Created status code with the body
	// being the school data and the header being the headers map

	err = app.writeJSON(w, http.StatusCreated, envelope{"school": school}, headers)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		
	}

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
