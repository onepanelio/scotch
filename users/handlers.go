package users

import (
	"encoding/json"
	"net/http"
	"strconv"

	//"github.com/gorilla/context"
	"github.com/gorilla/mux"

	"github.com/rushtehrani/scotch/lib/response"
	"github.com/rushtehrani/scotch/models"
)

func Get(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	id, _ := strconv.ParseInt(params["id"], 0, 64)

	u, err := models.GetUser(id)

	if u == (models.User{}) {
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
	var u models.User

	err := json.NewDecoder(r.Body).Decode(&u)

	if err != nil {
		response.Error(w, 400)
		return
	}

	errs := u.Save()

	if errs != nil {
		response.JSON(w, errs, 422)
		return
	}

	response.JSON(w, u)
}

func Update(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	id, _ := strconv.ParseInt(params["id"], 0, 64)

	u, err := models.GetUser(id)

	if u == (models.User{}) {
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
	errs := u.Save()

	if errs != nil {
		response.JSON(w, errs, 422)
		return
	}

	response.JSON(w, u)
}
