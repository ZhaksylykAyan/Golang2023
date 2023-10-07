package main

import (
	"Travel_Accessories/internal/data"
	"fmt"
	"net/http"
	"time"
)

func (app *application) createAccessoriesHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "create a new accessory")
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
	}
	err = app.writeJSON(w, http.StatusOK, envelope{"accessories": accessories}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
