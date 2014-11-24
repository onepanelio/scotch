package models

import (
	"github.com/jmoiron/sqlx"

	"github.com/go-validator/validator"
	"github.com/rushtehrani/scotch/db"
)

type User struct {
	ID    int64  `json:"id,omitempty" db:"id"`
	Name  string `json:"name,omitempty" db:"name" validate:"nonzero"`
	Email string `json:"email,omitempty" db:"email"`
}

func GetUser(ID int64) (User, error) {
	u := User{}

	err := db.Get(&u, "SELECT * FROM users WHERE id = $1 LIMIT 1", ID)

	return u, err
}

func (u *User) Save() map[string][]error {
	var (
		err  error
		errs map[string][]error
		rows *sqlx.Rows
	)

	if isValid, errs := validator.Validate(u); !isValid {
		return errs
	}

	if u.ID > 0 {
		_, err = db.NamedExec("UPDATE users SET name = :name, email = :email WHERE id = :id", u)
	} else {
		rows, err = db.NamedQuery("INSERT INTO users (name, email) VALUES (:name, :email) RETURNING id", u)

		if rows.Next() {
			rows.Scan(&u.ID)
		}
	}

	if err != nil {
		errs["database"][0] = err
	}

	return errs
}
