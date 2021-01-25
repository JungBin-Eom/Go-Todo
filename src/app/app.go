package app

import (
	"net/http"

	"github.com/gorilla/mux"
)

func MakeNewHandler() http.Handler {
	r := mux.NewRouter()

	return r
}
