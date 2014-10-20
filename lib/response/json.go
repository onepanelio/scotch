package response

import (
	"encoding/json"
	"net/http"
)

func JSON(w http.ResponseWriter, v interface{}) {
	js, err := json.Marshal(v)

	if err != nil {
		Error(w, 500)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}
