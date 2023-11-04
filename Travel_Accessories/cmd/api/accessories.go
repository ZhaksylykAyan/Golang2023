package main

import (
	"Travel_Accessories/internal/data"
	"Travel_Accessories/internal/validator"
	"errors"
	"fmt"
	"net/http"
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
	accessory := &data.Accessories{
		Title:    input.Title,
		Year:     input.Year,
		Runtime:  input.Runtime,
		Color:    input.Color,
		Material: input.Material,
		Price:    input.Price,
	}
	v := validator.New()
	if data.ValidateAccessory(v, accessory); !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	err = app.models.Accessories.Insert(accessory)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}
	headers := make(http.Header)
	headers.Set("Location", fmt.Sprintf("/v1/accessory/%d", accessory.ID))
	err = app.writeJSON(w, http.StatusCreated, envelope{"accessory": accessory}, headers)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}

	fmt.Fprintf(w, "%+v\n", input)
}
func (app *application) showAccessoriesHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	accessories, err := app.models.Accessories.Get(id)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrRecordNotFound):
			app.notFoundResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}
	err = app.writeJSON(w, http.StatusOK, envelope{"accessories": accessories}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
func (app *application) updateAccessoryHandler(w http.ResponseWriter, r *http.Request) {
	// Extract the movie ID from the URL.
	id, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}
	accessory, err := app.models.Accessories.Get(id)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrRecordNotFound):
			app.notFoundResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}
	var input struct {
		Title   string       `json:"title"`
		Year    int32        `json:"year"`
		Runtime data.Runtime `json:"runtime"`
		Genres  []string     `json:"genres"`
	}
	err = app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}
	accessory.Title = input.Title
	accessory.Year = input.Year
	accessory.Runtime = input.Runtime
	v := validator.New()
	if data.ValidateAccessory(v, accessory); !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}
	err = app.models.Accessories.Update(accessory)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}
	err = app.writeJSON(w, http.StatusOK, envelope{"accessory": accessory}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) deleteAccessoryHandler(w http.ResponseWriter, r *http.Request) {
	// Extract the movie ID from the URL.
	id, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}
	// Delete the movie from the database, sending a 404 Not Found response to the
	// client if there isn't a matching record.
	err = app.models.Accessories.Delete(id)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrRecordNotFound):
			app.notFoundResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}
	// Return a 200 OK status code along with a success message.
	err = app.writeJSON(w, http.StatusOK, envelope{"message": "movie successfully deleted"}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
