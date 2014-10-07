package db

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var db *sqlx.DB

func Connect() {
    var err error
    
    db, err = sqlx.Connect("postgres", "user=postgres password=scotch dbname=scotch sslmode=disable")
	
	if err != nil {
		fmt.Println(err)
	}
}

func Get(dest interface{}, query string, args ...interface{}) error {
	return db.Get(dest, query, args...)
}

func MustBegin() *sqlx.Tx {
	return db.MustBegin()
}