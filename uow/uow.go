package uow

import (
	"database/sql"
	"imageresizerservice/uow/uowInMemory"
	"imageresizerservice/uow/uowSql"
)

type Uow struct {
	InMemory uowInMemory.UowInMemory
	Sql      uowSql.UowSql
}

func Begin(db *sql.DB) (*Uow, error) {
	sql, err := uowSql.Begin(db)

	if err != nil {
		return nil, err
	}

	inMemory := uowInMemory.Begin()

	uow := Uow{
		InMemory: inMemory,
		Sql:      *sql,
	}
	return &uow, nil
}

func (uow *Uow) Commit() error {
	return uow.InMemory.Commit()
}

func (uow *Uow) Rollback() error {
	return uow.InMemory.Rollback()
}
