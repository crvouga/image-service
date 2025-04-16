package sqlite

import (
	"database/sql"

	"imageresizerservice/library/sql/noop"
)

func New() *sql.DB {
	return noop.New()
}
