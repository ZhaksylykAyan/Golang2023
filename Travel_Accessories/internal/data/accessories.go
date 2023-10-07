package data

import (
	"Travel_Accessories/internal/validator"
	"encoding/json"
	"fmt"
	"time"
)

type Accessories struct {
	ID        int64     `json:"ID"`
	CreatedAt time.Time `json:"-"`
	Title     string    `json:"title"`
	Year      int32     `json:"year"`
	Runtime   Runtime   `json:"runtime,omitempty"`
	Version   int32     `json:"version"`
}

func (a Accessories) MarshalJSON() ([]byte, error) {

	var runtime string
	if a.Runtime != 0 {
		runtime = fmt.Sprintf("%d mins", a.Runtime)
	}
	type AccessoryAlias Accessories

	aux := struct {
		AccessoryAlias
		Runtime string `json:"runtime,omitempty"`
	}{
		AccessoryAlias: AccessoryAlias(a),
		Runtime:        runtime,
	}
	return json.Marshal(aux)
}

func ValidateAccessory(v *validator.Validator, movie *Accessories) {
	v.Check(movie.Title != "", "title", "must be provided")
	v.Check(len(movie.Title) <= 500, "title", "must not be more than 500 bytes long")
	v.Check(movie.Year != 0, "year", "must be provided")
	v.Check(movie.Year >= 1888, "year", "must be greater than 1888")
	v.Check(movie.Year <= int32(time.Now().Year()), "year", "must not be in the future")
	v.Check(movie.Runtime != 0, "runtime", "must be provided")
	v.Check(movie.Runtime > 0, "runtime", "must be a positive integer")
	//v.Check(movie.Genres != nil, "genres", "must be provided")
	//v.Check(len(movie.Genres) >= 1, "genres", "must contain at least 1 genre")
	//v.Check(len(movie.Genres) <= 5, "genres", "must not contain more than 5 genres")
	//v.Check(validator.Unique(movie.Genres), "genres", "must not contain duplicate values")
}
