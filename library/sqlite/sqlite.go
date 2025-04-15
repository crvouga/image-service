package sqlite

import (
	"database/sql"

	_ "modernc.org/sqlite"
)

func New() *sql.DB {
	db, err := sql.Open("sqlite", ":memory:")

	if err != nil {
		panic(err)
	}

	return db
}
