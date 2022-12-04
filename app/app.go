package app

import (
	"go_mysql/domain"
	"net/http"

	"github.com/gorilla/mux"
)

var (
	r = mux.NewRouter().StrictSlash(true)
)

func StartApp() {
	Routers()
	domain.Connect()
	domain.Migrate()
	http.ListenAndServe(":8080", r)

}
