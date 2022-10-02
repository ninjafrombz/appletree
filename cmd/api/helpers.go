// Filename:
package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
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

// Define a new type named envelope 
type envelope map[string]interface{}

func (app *application) writeJSON(w http.ResponseWriter, status int, data envelope, headers http.Header) (error) {
	js, err := json.MarshalIndent(data, "", "\t")
	if err != nil {
		return err
	}
	js = append(js, '\n')
	// Add any of the headers 
	for key, value := range headers {
		w.Header()[key] = value 
	}
	// Specify that we will serve our responses using JSON 
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	//Write the byte slice containing the json response body
	w.Write(js)
	return nil
}

func (app *application) readJSON(w http.ResponseWriter, r *http.Request, dst interface{}) error {
	// decode the request body into the target destination 
	err := json.NewDecoder(r.Body).Decode(dst)
	// Check for a bad request
	if err != nil {
		var syntaxError *json.SyntaxError
		var unmarshalTypeError *json.UnmarshalTypeError
		var invalidUnmarshalError *json.InvalidUnmarshalError
		//Switch to check for the errors
		switch {
			//Check for syntax Errors
		case errors.As(err, &syntaxError):
			return fmt.Errorf("body contains badly formed JSON (at character %d)", syntaxError.Offset)
		case errors.Is(err, io.ErrUnexpectedEOF):
			return errors.New("body contains badly formed JSON")
		// Check for wrong types passed by the client 
		case errors.As(err, &unmarshalTypeError):
			if unmarshalTypeError.Field != "" {
				return fmt.Errorf("body contains incorrect JSON type for field %q", unmarshalTypeError.Field)
			}
			return fmt.Errorf("body contains incorrect JSON type (at character %d)", unmarshalTypeError.Offset)
			//Empty body
		case errors.Is(err, io.EOF):
			return errors.New("body must not be empty")
			// Pass non-nil pointer
		case errors.As(err, &invalidUnmarshalError):
			panic(err)
			//default
		default:
			return err
		}
		//default
	
	}
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
