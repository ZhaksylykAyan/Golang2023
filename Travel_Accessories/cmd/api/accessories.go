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
		Title    *string       `json:"title"`
		Year     *int32        `json:"year"`
		Runtime  *data.Runtime `json:"runtime"`
		Color    *string       `json:"color"`
		Material *string       `json:"material"`
		Price    *float64      `json:"price"`
	}
	err = app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}
	if input.Title != nil {
		accessory.Title = *input.Title
	}
	if input.Year != nil {
		accessory.Year = *input.Year
	}
	if input.Runtime != nil {
		accessory.Runtime = *input.Runtime
	}
	if input.Color != nil {
		accessory.Color = *input.Color
	}
	if input.Material != nil {
		accessory.Material = *input.Material
	}
	if input.Price != nil {
		accessory.Price = *input.Price
	}

	v := validator.New()
	if data.ValidateAccessory(v, accessory); !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}
	err = app.models.Accessories.Update(accessory)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrEditConflict):
			app.editConflictResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"accessory": accessory}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) deleteAccessoryHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}
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
	err = app.writeJSON(w, http.StatusOK, envelope{"message": "movie successfully deleted"}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) listAccessoriesHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Title string
		//Genres   []string
		data.Filters
	}
	v := validator.New()
	qs := r.URL.Query()
	input.Title = app.readString(qs, "title", "")
	//input.Genres = app.readCSV(qs, "genres", []string{})
	input.Filters.Page = app.readInt(qs, "page", 1, v)
	input.Filters.PageSize = app.readInt(qs, "page_size", 20, v)
	input.Filters.Sort = app.readString(qs, "sort", "id")
	input.Filters.SortSafelist = []string{"id", "title", "year", "runtime", "-id", "-title", "-year", "-runtime"}
	if data.ValidateFilters(v, input.Filters); !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}
	accessories, metadata, err := app.models.Accessories.GetAll(input.Title, input.Filters)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}
	err = app.writeJSON(w, http.StatusOK, envelope{"accessories": accessories, "metadata": metadata}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}

	fmt.Fprintf(w, "%+v\n", input)
}
