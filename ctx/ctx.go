package ctx

import (
	"net/http"

	"github.com/gorilla/context"
)

var (
	CurrentUser interface{}
)

func Set(r *http.Request, key, v interface{}) {
	context.Set(r, key, v)
}

func Get(r *http.Request, key interface{}) interface{} {
	return context.Get(r, key)
}
