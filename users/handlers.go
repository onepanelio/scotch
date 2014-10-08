package users

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"github.com/rushtehrani/scotch/lib/response"
)

func Get(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	id, _ := strconv.ParseUint(params["id"], 0, 64)

	u := GetUser(id)

	response.JSON(w, u)
}
