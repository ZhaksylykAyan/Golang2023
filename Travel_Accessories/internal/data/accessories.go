package data

import (
	"Travel_Accessories/internal/validator"
	"database/sql"
	"encoding/json"
	"errors"
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

type AccessoriesModel struct {
	DB *sql.DB
}

func (a *AccessoriesModel) Insert(accessories *Accessories) error {
	query := `
		INSERT INTO accessories (title, year, runtime, color)
		VALUES ($1, $2, $3, $4)
		RETURNING id, created_at, version`
	args := []interface{}{accessories.Title, accessories.Year, accessories.Runtime, accessories.Color}
	return a.DB.QueryRow(query, args...).Scan(&accessories.ID, &accessories.CreatedAt, &accessories.Version)
}

func (a AccessoriesModel) Get(id int64) (*Accessories, error) {
	if id < 1 {
		return nil, ErrRecordNotFound
	}
	query := `
SELECT id, created_at, title, year, runtime, genres, version
FROM Accessoriess
WHERE id = $1`
	var accessory Accessories
	err := a.DB.QueryRow(query, id).Scan(
		&accessory.ID,
		&accessory.CreatedAt,
		&accessory.Title,
		&accessory.Year,
		&accessory.Runtime,
		&accessory.Version,
		&accessory.Material,
		&accessory.Color,
		&accessory.Price,
	)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, ErrRecordNotFound
		default:
			return nil, err
		}
	}
	return &accessory, nil
}
func (a AccessoriesModel) Update(accessory *Accessories) error {
	query := `
UPDATE movies
SET title = $1, year = $2, runtime = $3, color = $4, version = version + 1
WHERE id = $5
RETURNING version`
	args := []interface{}{
		accessory.Title,
		accessory.Year,
		accessory.Runtime,
		accessory.Color,
		accessory.ID,
	}
	return a.DB.QueryRow(query, args...).Scan(&accessory.Version)
}

func (a AccessoriesModel) Delete(id int64) error {
	// Return an ErrRecordNotFound error if the movie ID is less than 1.
	if id < 1 {
		return ErrRecordNotFound
	}
	// Construct the SQL query to delete the record.
	query := `
DELETE FROM accessories
WHERE id = $1`
	// Execute the SQL query using the Exec() method, passing in the id variable as
	// the value for the placeholder parameter. The Exec() method returns a sql.Result
	// object.
	result, err := a.DB.Exec(query, id)
	if err != nil {
		return err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return ErrRecordNotFound
	}
	return nil
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
