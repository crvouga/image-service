package uow

import (
	"database/sql"
)

type UowFactory struct {
	db *sql.DB
}

func NewFactory(db *sql.DB) *UowFactory {
	return &UowFactory{
		db: db,
	}
}

func (uowFactory *UowFactory) Begin() (*Uow, error) {
	return Begin(uowFactory.db)
}
