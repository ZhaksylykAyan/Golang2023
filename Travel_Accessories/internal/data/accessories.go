package data

import (
	"Travel_Accessories/internal/validator"
	"context"
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
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	// Use QueryRowContext() and pass the context as the first argument.
	return a.DB.QueryRowContext(ctx, query, args...).Scan(&accessories.ID, &accessories.CreatedAt, &accessories.Version)
}

func (a AccessoriesModel) Get(id int64) (*Accessories, error) {
	if id < 1 {
		return nil, ErrRecordNotFound
	}
	query := `
SELECT id, created_at, title, year, runtime, material, version, color, price
FROM Accessories
WHERE id = $1`

	var accessory Accessories
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err := a.DB.QueryRowContext(ctx, query, id).Scan(
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
WHERE id = $5 AND version = $6
RETURNING version`

	args := []interface{}{
		accessory.Title,
		accessory.Year,
		accessory.Runtime,
		accessory.Color,
		accessory.ID,
		accessory.Version,
	}
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err := a.DB.QueryRowContext(ctx, query, args...).Scan(&accessory.Version)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return ErrEditConflict
		default:
			return err
		}
	}
	return nil

}

func (a AccessoriesModel) Delete(id int64) error {
	if id < 1 {
		return ErrRecordNotFound
	}
	query := `
DELETE FROM accessories
WHERE id = $1`
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	result, err := a.DB.ExecContext(ctx, query, id)
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

func (a AccessoriesModel) GetAll(title string, filters Filters) ([]*Accessories, Metadata, error) {
	query := fmt.Sprintf(`
SELECT count(*) OVER(), id, created_at, title, year, runtime, material, version, color, price
FROM accessories
WHERE (to_tsvector('simple', title) @@ plainto_tsquery('simple', $1) OR $1 = '')
ORDER BY %s %s, id ASC
LIMIT $3 OFFSET $4`, filters.sortColumn(), filters.sortDirection())
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	args := []interface{}{title, filters.limit(), filters.offset()}
	rows, err := a.DB.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, Metadata{}, err
	}
	defer rows.Close()
	totalRecords := 0
	accessories := []*Accessories{}
	for rows.Next() {
		var accessory Accessories
		err := rows.Scan(
			&totalRecords,
			&accessory.ID,
			&accessory.CreatedAt,
			&accessory.Title,
			&accessory.Year,
			&accessory.Runtime,
			&accessory.Material,
			&accessory.Version,
			&accessory.Color,
			&accessory.Price,
		)
		if err != nil {
			return nil, Metadata{}, err
		}
		accessories = append(accessories, &accessory)
	}
	if err = rows.Err(); err != nil {
		return nil, Metadata{}, err
	}
	metadata := calculateMetadata(totalRecords, filters.Page, filters.PageSize)
	return accessories, metadata, nil
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
