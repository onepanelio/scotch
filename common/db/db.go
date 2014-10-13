package db

import (
	"database/sql"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var db *sqlx.DB

func MustConnect(driverName, dataSourceName string) {
	db = sqlx.MustConnect(driverName, dataSourceName)
}

func Get(dest interface{}, query string, args ...interface{}) error {
	return db.Get(dest, query, args...)
}

func NamedExec(query string, arg interface{}) (sql.Result, error) {
	return db.NamedExec(query, arg)
}

func NamedQuery(query string, arg interface{}) (*sqlx.Rows, error) {
	return db.NamedQuery(query, arg)
}

func QueryRowx(query string, args ...interface{}) *sqlx.Row {
	return db.QueryRowx(query, args)
}

func MustBegin() *sqlx.Tx {
	return db.MustBegin()
}
