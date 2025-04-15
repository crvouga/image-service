package uow

import (
	"database/sql"
)

type UowFactory struct {
	Db *sql.DB
}

func (uowFactory *UowFactory) Begin() (*Uow, error) {
	return Begin(uowFactory.Db)
}
