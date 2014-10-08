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

	id, _ := strconv.ParseUint(params["id"], 0, 64)

	u, err := GetUser(id)
	
	if u == (User{}) {
		response.Error(w, 404)
	} 
		
	if err != nil {
		response.Error(w, 500)
		return
	}

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

	id, _ := strconv.ParseUint(params["id"], 0, 64)
	
	json.NewDecoder(r.Body).Decode(&u);
	
	u.ID = id
	u.Save()

	response.JSON(w, u)
}