// Filename:
package main

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

func (app *application) readIDParam(r *http.Request) (int64, error) {
	params := httprouter.ParamsFromContext(r.Context())

	id, err := strconv.ParseInt(params.ByName("id"), 10, 64)
	if err != nil || id < 1 {
		return 0, errors.New("Invalid id parameter")
		
	}

	return id, nil
	//Display the school id

}


func (app *application) writeJSON(w http.ResponseWriter, status int, data interface{}, headers http.Header) (error) {
	js, err := json.MarshalIndent(data, "", "\t")
	if err != nil {
		return err
	}
	js = append(js, '\n')
	// Add any of the headers 
	for key, value := range headers {
		w.Header()[key] = value 
	}

	// SPecify that we will serve our responses using JSON 
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	//Write the byte slice containing the json response body
	w.Write(js)
	return nil

}
// func (app *application) healthcheckHandler(w http.ResponseWriter, r *http.Request) {
// 	js := `{"status":"available", "environment": %q, "version": %q}`
// 	js = fmt.Sprintf(js, app.config.env, version)
// 	// Specify that we will serve our responses using JSON 
// 	w.Header().Set("Content-Type", "application/json")
// 	// write the json as the HTTP 
// 	w.Write([]byte(js)) 
// }
