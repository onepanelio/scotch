package response

import (
	"encoding/json"
	"net/http"
)

func JSON(w http.ResponseWriter, val interface{}) {
	js, err := json.Marshal(val)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}
