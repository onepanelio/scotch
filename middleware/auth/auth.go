package auth

import (
	"io/ioutil"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/context"

	"github.com/rushtehrani/scotch/cfg"
	"github.com/rushtehrani/scotch/lib/response"
)

var (
	privateKey []byte
	publicKey  []byte
)

type AuthHandler struct {
	successHandler http.Handler
	failureHandler http.Handler
}

func defaultFailureHandler(w http.ResponseWriter, r *http.Request) {
	response.Error(w, 401)
}

func New(h http.Handler) http.Handler {
	var err error

	privateKey, err = ioutil.ReadFile(cfg.Get("auth.privateKeyPath"))

	if err != nil {
		panic(err)
	}

	publicKey, err = ioutil.ReadFile(cfg.Get("auth.publicKeyPath"))

	if err != nil {
		panic(err)
	}

	return &AuthHandler{successHandler: h, failureHandler: http.HandlerFunc(defaultFailureHandler)}
}

func (a *AuthHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	token, err := jwt.ParseFromRequest(r, func(token *jwt.Token) (interface{}, error) {
		return publicKey, nil
	})

	if err != nil || !token.Valid {
		a.failureHandler.ServeHTTP(w, r)
		return
	}

	context.Set(r, "User", 1)
	a.successHandler.ServeHTTP(w, r)
}
