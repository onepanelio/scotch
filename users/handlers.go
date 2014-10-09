package users

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"github.com/rushtehrani/scotch/lib/response"
)

func Get(w http.ResponseWriter, r *http.Request) {
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

	response.JSON(w, u)
}

func Create(w http.ResponseWriter, r *http.Request) {
	var u User

	json.NewDecoder(r.Body).Decode(&u)

	err := u.Save()

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

	json.NewDecoder(r.Body).Decode(&u)

	u.ID = id
	err = u.Save()

	if err != nil {
		response.Error(w, 500)
		return
	}

	response.JSON(w, u)
}
