package dat

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

func execHook(hookType string, arg interface{}) {
	method := reflect.ValueOf(arg).MethodByName(hookType)

	if method != (reflect.Value{}) {
		method.Call([]reflect.Value{})
	}
}

func In(query string, args ...interface{}) (string, []interface{}, error) {
	return sqlx.In(query, args)
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

func (db *DB) Insert(query string, arg interface{}) error {
	if !isValidStatement("INSERT", query) {
		return errors.New("Not a valid INSERT statement.")
	}

	execHook("PreInsert", arg)

	rows, err := db.NamedQuery(query, arg)

	if err != nil {
		return err
	}

	defer rows.Close()

	for rows.Next() {
		err = rows.StructScan(arg)

		if err != nil {
			return err
		}
	}

	if err = rows.Err(); err != nil {
		return err
	}

	execHook("PostInsert", arg)

	return nil
}

func (db *DB) Update(query string, arg interface{}) error {
	if !isValidStatement("UPDATE", query) {
		return errors.New("Not a valid UPDATE statement.")
	}

	execHook("PreUpdate", arg)

	rows, err := db.NamedQuery(query, arg)
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.StructScan(arg)

		if err != nil {
			return err
		}
	}

	if err = rows.Err(); err != nil {
		return err
	}

	execHook("PostUpdate", arg)

	return nil
}

func (db *DB) Delete(query string, args ...interface{}) error {
	if !isValidStatement("DELETE", query) {
		return errors.New("Not a valid DELETE statement.")
	}

	_, err := db.Exec(query, args...)

	if err != nil {
		return err
	}

	return nil
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

func (tx *Tx) Insert(query string, arg interface{}) error {
	if !isValidStatement("INSERT", query) {
		return errors.New("Not a valid INSERT statement.")
	}

	execHook("PreInsert", arg)

	rows, err := tx.NamedQuery(query, arg)

	if err != nil {
		return err
	}

	defer rows.Close()

	for rows.Next() {
		err = rows.StructScan(arg)

		if err != nil {
			return err
		}
	}

	if err = rows.Err(); err != nil {
		return err
	}

	execHook("PostInsert", arg)

	return nil
}

func (tx *Tx) Update(query string, arg interface{}) error {
	if !isValidStatement("UPDATE", query) {
		return errors.New("Not a valid UPDATE statement.")
	}

	execHook("PreUpdate", arg)

	rows, err := tx.NamedQuery(query, arg)
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.StructScan(arg)

		if err != nil {
			return err
		}
	}

	if err = rows.Err(); err != nil {
		return err
	}

	execHook("PostUpdate", arg)

	return nil
}

func (tx *Tx) Delete(query string, args ...interface{}) error {
	if !isValidStatement("DELETE", query) {
		return errors.New("Not a valid DELETE statement.")
	}

	_, err := tx.Exec(query, args...)

	if err != nil {
		return err
	}

	return nil
}
