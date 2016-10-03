package ctx

import (
	"context"
	"net/http"
)

var (
	CurrentUser interface{}
)

func Set(r *http.Request, key, v interface{}) {
	r = r.WithContext(context.WithValue(r.Context(), key, v))
}

func Get(r *http.Request, key interface{}) interface{} {
	return r.Context().Value(key)
}
