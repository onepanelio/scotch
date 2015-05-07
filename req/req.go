package req

import (
	"encoding/json"
	"net/http"
	"net/url"

	"github.com/gorilla/mux"
)

func Query(r *http.Request) url.Values {
	return r.URL.Query()
}

func Params(r *http.Request) map[string]string {
	return mux.Vars(r)
}

func JSON(r *http.Request, v interface{}) error {
	return json.NewDecoder(r.Body).Decode(v)
}
