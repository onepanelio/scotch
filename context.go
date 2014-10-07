package main

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type context struct {
	db *sqlx.DB
}
