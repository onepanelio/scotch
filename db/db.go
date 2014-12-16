package db

import (
	"reflect"
	"database/sql"
	"regexp"
	"errors"

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

func Insert(query string, entity interface{}, entityPrimaryKey interface{}) error {
	
	if !isValidStatement("INSERT", query) {
		return errors.New("Not a valid INSERT statement.")
	}
	
	execHook("PreInsert", entity)
	
	rows, err := db.NamedQuery(query, entity) 
	
	if err != nil {
		return err
	}

	if rows.Next() {
		rows.Scan(entityPrimaryKey)
	}
	
	execHook("PostInsert", entity)
	
	return nil
}

func Update(query string, entity interface{}) error {
	
	if !isValidStatement("UPDATE", query) {
		return errors.New("Not a valid UPDATE statement.")
	}
	
	execHook("PreUpdate", entity)
	
	_, err := db.NamedExec(query, entity)
	
	if err != nil {
		return err
	}
	
	execHook("PostUpdate", entity)
	
	return nil
}

func isValidStatement(statementType string, query string) bool {
	r, _ := regexp.Compile("(i?)" + statementType)
	
	loc := r.FindStringIndex(query)
	
	if loc == nil || loc[0] != 0 {
		return false
	}
	
	return true
}

func execHook(hookType string, entity interface{}) {
	method := reflect.ValueOf(entity).MethodByName(hookType)
	
	if method != (reflect.Value{}) {
		method.Call([]reflect.Value{})
	}
}
