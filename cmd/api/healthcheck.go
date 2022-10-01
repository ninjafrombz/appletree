// Filename : cmd/api/healthcheck.go
package main

import (
	"encoding/json"
	"net/http"
)

func (app *application) healthcheckHandler(w http.ResponseWriter, r *http.Request) {
	// Create a map to hold our healthcheck data
	data := map[string]string{
		"status":      "available",
		"environment": app.config.env,
		"version":     version,
	}

	//Convert Map to a JSON object
	js, err := json.Marshal(data)
	if err != nil {
		app.logger.Println(err)
		http.Error(w, "The server encountered a problem and could not process your request", http.StatusInternalServerError)
		return
	}
	// Add a newline for easier viewing
	js = append(js, '\n')

	w.Header().Set("Content-Type", "application/json")
	//Write the byte slice containing the json response body
	w.Write(js)

}
