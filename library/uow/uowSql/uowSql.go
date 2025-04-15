package uowSql

import (
	"database/sql"
)

type UowSql struct {
	tx *sql.Tx
}

func Begin(db *sql.DB) (*UowSql, error) {
	tx, err := db.Begin()
	if err != nil {
		return nil, err
	}
	return &UowSql{tx: tx}, nil
}

func (uow *UowSql) Rollback() error {
	return uow.tx.Rollback()
}

func (uow *UowSql) Commit() error {
	return uow.tx.Commit()
}

func (uow *UowSql) GetTx() *sql.Tx {
	return uow.tx
}
