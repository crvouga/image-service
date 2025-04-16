package noop

import (
	"database/sql"
	"database/sql/driver"
)

// Driver implements database/sql/driver.Driver interface
type Driver struct{}

// Open returns a new connection to the database
func (d *Driver) Open(name string) (driver.Conn, error) {
	return &conn{}, nil
}

// conn implements driver.Conn interface
type conn struct{}

// Prepare returns a prepared statement
func (c *conn) Prepare(query string) (driver.Stmt, error) {
	return &stmt{}, nil
}

// Close closes the connection
func (c *conn) Close() error {
	return nil
}

// Begin starts a transaction
func (c *conn) Begin() (driver.Tx, error) {
	return &tx{}, nil
}

// stmt implements driver.Stmt interface
type stmt struct{}

// Close closes the statement
func (s *stmt) Close() error {
	return nil
}

// NumInput returns the number of placeholder parameters
func (s *stmt) NumInput() int {
	return -1 // -1 means we don't know how many parameters
}

// Exec executes a query that doesn't return rows
func (s *stmt) Exec(args []driver.Value) (driver.Result, error) {
	return &result{}, nil
}

// Query executes a query that returns rows
func (s *stmt) Query(args []driver.Value) (driver.Rows, error) {
	return &rows{}, nil
}

// result implements driver.Result interface
type result struct{}

// LastInsertId returns the database's auto-generated ID after an insert
func (r *result) LastInsertId() (int64, error) {
	return 0, nil
}

// RowsAffected returns the number of rows affected by the query
func (r *result) RowsAffected() (int64, error) {
	return 0, nil
}

// rows implements driver.Rows interface
type rows struct {
	columns []string
}

// Columns returns the names of the columns
func (r *rows) Columns() []string {
	return []string{}
}

// Close closes the rows iterator
func (r *rows) Close() error {
	return nil
}

// Next is called to populate the next row of data
func (r *rows) Next(dest []driver.Value) error {
	return nil
}

// tx implements driver.Tx interface
type tx struct{}

// Commit commits the transaction
func (t *tx) Commit() error {
	return nil
}

// Rollback rolls back the transaction
func (t *tx) Rollback() error {
	return nil
}

// Register registers the driver
func Register() {
	sql.Register("noop", &Driver{})
}

// New creates a new no-op database connection
func New() *sql.DB {
	Register()
	db, err := sql.Open("noop", "")
	if err != nil {
		panic(err) // This should never happen with our no-op driver
	}
	return db
}
