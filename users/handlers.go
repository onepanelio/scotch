package users

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/context"
	"github.com/gorilla/mux"

	"github.com/rushtehrani/scotch/lib/response"
)

func Get(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	fmt.Println(context.Get(r, "User"))

	id, _ := strconv.ParseInt(params["id"], 0, 64)

	u, err := GetUser(id)

	if u == (User{}) {
		response.Error(w, 404)
		return
	}

	if err != nil {
		response.Error(w, 500)
		return
	}

	response.JSON(w, u)
}

func Create(w http.ResponseWriter, r *http.Request) {
	var (
		u   User
		err error
	)

	err = json.NewDecoder(r.Body).Decode(&u)

	if err != nil {
		response.Error(w, 400)
		return
	}

	err = u.Save()

	if err != nil {
		response.Error(w, 500)
		return
	}

	response.JSON(w, u)
}

func Update(w http.ResponseWriter, r *http.Request) {
	var err error

	params := mux.Vars(r)

	id, _ := strconv.ParseInt(params["id"], 0, 64)

	u, err := GetUser(id)

	if u == (User{}) {
		response.Error(w, 404)
		return
	}

	if err != nil {
		response.Error(w, 500)
		return
	}

	err = json.NewDecoder(r.Body).Decode(&u)

	if err != nil {
		response.Error(w, 400)
		return
	}

	u.ID = id
	err = u.Save()

	if err != nil {
		response.Error(w, 500)
		return
	}

	response.JSON(w, u)
}
