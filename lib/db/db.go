package db

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var DB *sqlx.DB

func Connect() {
    var err error
    
    DB, err = sqlx.Connect("postgres", "user=postgres password=scotch dbname=scotch sslmode=disable")
	
	if err != nil {
		fmt.Println(err)
	}
}
