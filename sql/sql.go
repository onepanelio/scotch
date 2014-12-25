package sql

import (
	"errors"
	"reflect"
	"regexp"

	"github.com/jmoiron/sqlx"
)

type DB struct {
	*sqlx.DB
}

func MustConnect(driverName, dataSourceName string) *DB {
	db := sqlx.MustConnect(driverName, dataSourceName)

	return &DB{DB: db}
}

func (db *DB) Insert(query string, entity interface{}, entityPrimaryKey interface{}) error {
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

func (db *DB) Update(query string, entity interface{}) error {
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
