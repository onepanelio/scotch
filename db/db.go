package db

import (
	"reflect"
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

func Select(dest interface{}, query string, args ...interface{}) error {
	return db.Select(dest, query, args...)
}

func NamedQuery(query string, arg interface{}) (*sqlx.Rows, error) {
	return db.NamedQuery(query, arg)
}

func NamedExec(query string, arg interface{}) (sql.Result, error) {
	return db.NamedExec(query, arg)
}


func Insert(query string, arg interface{}, primaryKey *int64) error {
	execHook("PreInsert", arg)
	
	rows, err := db.NamedQuery(query, arg) 
	
	if err != nil {
		return err
	}

	if rows.Next() {
		rows.Scan(primaryKey)
	}
	
	execHook("PostInsert", arg)
	
	return nil
}

func Update(query string, arg interface{}) error {
	execHook("PreUpdate", arg)
	
	_, err := db.NamedExec(query, arg)
	
	if err != nil {
		return err
	}
	
	execHook("PostUpdate", arg)
	
	return nil
}

func execHook(hookType string, v interface{}) {
	method := reflect.ValueOf(v).MethodByName(hookType)
	
	if method != (reflect.Value{}) {
		method.Call([]reflect.Value{})
	}
}
