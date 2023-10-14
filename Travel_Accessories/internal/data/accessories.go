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
	Color     string    `json:"color"`
	Material  string    `json:"material"`
	Price     float64   `json:"price"`
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

func ValidateAccessory(v *validator.Validator, accessories *Accessories) {
	v.Check(accessories.Title != "", "title", "must be provided")
	v.Check(len(accessories.Title) <= 500, "title", "must not be more than 500 bytes long")
	v.Check(accessories.Year != 0, "year", "must be provided")
	v.Check(accessories.Year >= 1888, "year", "must be greater than 1888")
	v.Check(accessories.Year <= int32(time.Now().Year()), "year", "must not be in the future")
	v.Check(accessories.Runtime != 0, "runtime", "must be provided")
	v.Check(accessories.Runtime > 0, "runtime", "must be a positive integer")
	v.Check(accessories.Color != "", "color", "must be provided")
	v.Check(accessories.Material != "", "material", "must be provided")
	v.Check(accessories.Price >= 0, "price", "must be bigger than 0")
}
