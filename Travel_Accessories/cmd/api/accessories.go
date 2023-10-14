package main

import (
	"Travel_Accessories/internal/data"
	"Travel_Accessories/internal/validator"
	"fmt"
	"net/http"
	"time"
)

func (app *application) createAccessoriesHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Title    string       `json:"title"`
		Year     int32        `json:"year"`
		Runtime  data.Runtime `json:"runtime"`
		Color    string       `json:"color"`
		Material string       `json:"material"`
		Price    float64      `json:"price"`
	}
	err := app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}
	// Copy the values from the input struct to a new Movie struct.
	accessory := &data.Accessories{
		Title:    input.Title,
		Year:     input.Year,
		Runtime:  input.Runtime,
		Color:    input.Color,
		Material: input.Material,
		Price:    input.Price,
		//Genres: input.Genres,
	}
	// Initialize a new Validator.
	v := validator.New()
	// Call the ValidateMovie() function and return a response containing the errors if
	// any of the checks fail.
	if data.ValidateAccessory(v, accessory); !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	fmt.Fprintf(w, "%+v\n", input)
}
func (app *application) showAccessoriesHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}
	accessories := data.Accessories{
		ID:        id,
		CreatedAt: time.Now(),
		Title:     "Suitcases",
		Runtime:   102,
		Version:   1,
		Year:      1996,
		Color:     "blue",
		Material:  "Aluminium",
		Price:     12776,
	}
	err = app.writeJSON(w, http.StatusOK, envelope{"accessories": accessories}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
