package users

import (
	"fmt"
	
	"github.com/rushtehrani/scotch/db"
)

type User struct {
	ID    uint64 `json:"id,omitempty" db:"id"`
	Name  string `json:"name,omitempty" db:"name"`
	Email string `json:"email,omitempty" db:"email"`
}

func (u *User) Save() {
	tx := db.MustBegin()

	if u.ID > 0 {
		tx.NamedExec("UPDATE users SET name = :name, email = :email WHERE id = :id", u)
	} else {
		tx.NamedExec("INSERT INTO users (name, email) VALUES (:name, :email)", u)
	}

	tx.Commit()
}

func GetUser(ID uint64) User {
	u := User{}

	err := db.Get(&u, "SELECT * FROM users WHERE id = $1 LIMIT 1", ID)

	if err != nil {
		fmt.Println(err)
	}

	return u
}
