package res

import (
	"encoding/json"
	"net/http"
)

func JSON(w http.ResponseWriter, v interface{}, code ...int) {
	js, err := json.Marshal(v)

	if err != nil {
		Error(w, 500)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	if len(code) > 0 {
		w.WriteHeader(code[0])
	}

	w.Write(js)
}

func Error(w http.ResponseWriter, code int, m ...string) {
	message := http.StatusText(code)

	if len(m) > 0 {
		message = m[0]
	}

	http.Error(w, message, code)
}
