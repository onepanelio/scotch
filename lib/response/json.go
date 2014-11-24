package response

import (
	"encoding/json"
	"net/http"
	"reflect"
)

func JSON(w http.ResponseWriter, v interface{}, code ...int) {
	if reflect.TypeOf(v).String() == "map[string][]error" {
		v = makeJSONFriendly(v.(map[string][]error))
	}

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

func makeJSONFriendly(errs map[string][]error) map[string][]string {
	ferr := make(map[string][]string)

	for k, v := range errs {
		for _, err := range v {
			ferr[k] = append(ferr[k], err.Error())
		}
	}

	return ferr
}
