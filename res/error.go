package res

import "net/http"

func Error(w http.ResponseWriter, code int, m ...string) {
	message := http.StatusText(code)

	if len(m) > 0 {
		message = m[0]
	}

	http.Error(w, message, code)
}
