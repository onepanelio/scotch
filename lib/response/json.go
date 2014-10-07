package response

import (
	"net/http"
	"encoding/json"
)

func JSON(w http.ResponseWriter, value interface{}) {
	js, err := json.Marshal(value)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}
