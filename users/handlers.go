package users

import (
	"net/http"
	"strconv"
	"encoding/json"

	"github.com/gorilla/mux"

	"github.com/rushtehrani/scotch/lib/response"
)

func Get(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	ID, _ := strconv.ParseUint(params["id"], 0, 64)

	u := GetUser(ID)

	response.JSON(w, u)
}


func Create(w http.ResponseWriter, r *http.Request) {
	var u User
	
	json.NewDecoder(r.Body).Decode(&u);

	u.Save()

	response.JSON(w, u)
}

func Update(w http.ResponseWriter, r *http.Request) {
	var u User
	
	params := mux.Vars(r)

	ID, _ := strconv.ParseUint(params["id"], 0, 64)
	
	json.NewDecoder(r.Body).Decode(&u);
	
	u.ID = ID
	u.Save()

	response.JSON(w, u)
}