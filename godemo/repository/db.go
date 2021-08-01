package repository

import (
	"context"
	"database/sql"
)

type DB interface {
	Select(dest interface{}, query string, args ...interface{}) error
	Get(dest interface{}, query string, args ...interface{}) error
	ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error)
	Begin() (*sql.Tx, error)
}
