package sqlite

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

func New() *sql.DB {
	db, err := sql.Open("sqlite3", ":memory:")

	if err != nil {
		panic(err)
	}

	return db
}
