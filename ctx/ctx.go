package ctx

import (
	"net/http"

	"github.com/gorilla/context"
	sq "github.com/lann/squirrel"

	"github.com/rushtehrani/scotch/sql"
)

var (
	DB *sql.DB
	Q  sq.StatementBuilderType
)

func Set(r *http.Request, key, v interface{}) {
	context.Set(r, key, v)
}

func Get(r *http.Request, key interface{}) interface{} {
	return context.Get(r, key)
}
