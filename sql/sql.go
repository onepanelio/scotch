package sql

import (
	"errors"
	"reflect"
	"regexp"
	"strings"

	"github.com/jmoiron/sqlx"
)

func isValidStatement(statementType string, query string) bool {
	r, _ := regexp.Compile("(i?)" + statementType)

	loc := r.FindStringIndex(strings.TrimSpace(query))

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

type Rows struct {
	*sqlx.Rows
}

type DB struct {
	*sqlx.DB
}

func MustConnect(driverName, dataSourceName string) *DB {
	db := sqlx.MustConnect(driverName, dataSourceName)

	return &DB{DB: db}
}

func (db *DB) Unsafe() *DB {
	return &DB{DB: db.DB.Unsafe()}
}

func (db *DB) Get(dest interface{}, query string, args ...interface{}) error {
	if !isValidStatement("SELECT", query) {
		return errors.New("Not a valid SELECT statement.")
	}

	err := db.DB.Get(dest, query, args...)

	if err != nil {
		return err
	}

	execHook("PostGet", dest)

	return nil
}

func (db *DB) Insert(query string, entity interface{}) (*Rows, error) {
	if !isValidStatement("INSERT", query) {
		return nil, errors.New("Not a valid INSERT statement.")
	}

	execHook("PreInsert", entity)

	rows, err := db.NamedQuery(query, entity)

	if err != nil {
		return nil, err
	}

	execHook("PostInsert", entity)

	return &Rows{Rows: rows}, nil
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

func (db *DB) Delete(query string, args ...interface{}) (*Rows, error) {
	if !isValidStatement("DELETE", query) {
		return nil, errors.New("Not a valid DELETE statement.")
	}

	rows, err := db.Queryx(query, args...)

	if err != nil {
		return nil, err
	}

	return &Rows{Rows: rows}, nil
}

func (db *DB) MustBegin() *Tx {
	tx, err := db.Beginx()

	if err != nil {
		panic(err)
	}

	return tx
}

func (db *DB) Beginx() (*Tx, error) {
	tx, err := db.DB.Beginx()

	if err != nil {
		return nil, err
	}

	return &Tx{Tx: tx}, err
}

type Tx struct {
	*sqlx.Tx
}

func (tx *Tx) Insert(query string, entity interface{}) (*Rows, error) {
	if !isValidStatement("INSERT", query) {
		return nil, errors.New("Not a valid INSERT statement.")
	}

	execHook("PreInsert", entity)

	rows, err := tx.NamedQuery(query, entity)

	if err != nil {
		return nil, err
	}

	execHook("PostInsert", entity)

	return &Rows{Rows: rows}, nil
}

func (tx *Tx) Update(query string, entity interface{}) error {
	if !isValidStatement("UPDATE", query) {
		return errors.New("Not a valid UPDATE statement.")
	}

	execHook("PreUpdate", entity)

	_, err := tx.NamedExec(query, entity)

	if err != nil {
		return err
	}

	execHook("PostUpdate", entity)

	return nil
}

func (tx *Tx) Delete(query string, args ...interface{}) (*Rows, error) {
	if !isValidStatement("DELETE", query) {
		return nil, errors.New("Not a valid DELETE statement.")
	}

	rows, err := tx.Queryx(query, args...)

	if err != nil {
		return nil, err
	}

	return &Rows{Rows: rows}, nil
}
