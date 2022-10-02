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
	fmt.Fprintln(w, "create a new school..")
}

func (app *application) showSchoolHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	//Create a new instance of the school struct containing the ID we extracted 
	// from our URL and some sample data 

	school := data.School {
		ID: id,
		CreatedAt: time.Now(),
		Name: "Apple Tree",
		Level: "High School",
		Contact: "Anna Smith",
		Phone: "654-1651",
		Address: "14 Apple Steet",
		Mode: []string{"blended", "online"},
		Version: 1,
		
	}

	err = app.writeJSON(w, http.StatusOK, school, nil)
	if err != nil {
		app.logger.Println(err)
		http.Error(w, "The Server encountered a problem and could not process your request",http.StatusInternalServerError)
	}
}
