package databasesql

import "github.com/jmoiron/sqlx"

type DB interface {
	sqlx.ExtContext
	sqlx.PreparerContext
}
