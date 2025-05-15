package sqlite

import (
	"database/sql"

	"imageService/library/sql/noop"
)

func New() *sql.DB {
	return noop.New()
}
