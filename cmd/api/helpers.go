// Filename:
package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"

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

func (app *application) writeJSON(w http.ResponseWriter, status int, data envelope, headers http.Header) error {
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
	maxBytes := 1_048_576
	// decode the request body into the target destination
	r.Body = http.MaxBytesReader(w, r.Body, int64(maxBytes))
	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()
	err := dec.Decode(dst)
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

		// Check for Unmappable fields
		case strings.HasPrefix(err.Error(), "json: unknown field "):
			fieldName := strings.TrimPrefix(err.Error(), "json: unknown field")
			return fmt.Errorf("body contains unknown key %s", fieldName)

		// Too large
		case err.Error() == "http: request body too large":
			return fmt.Errorf("The body must not be larger than %d bytes", maxBytes)

		// Pass non-nil pointer
		case errors.As(err, &invalidUnmarshalError):
			panic(err)
		//default
		default:
			return err
		}
	}
	// Call decode again
	err = dec.Decode(&struct{}{})
	if err != io.EOF {
		return errors.New("body must only contain a single JSON value")
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
