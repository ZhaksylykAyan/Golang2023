package data

import (
	"database/sql"
	"errors"
)

var (
	ErrRecordNotFound = errors.New("record not found")
)

type Models struct {
	Accessories AccessoriesModel
}

func NewModels(db *sql.DB) Models {
	return Models{
		Accessories: AccessoriesModel{DB: db},
	}
}
