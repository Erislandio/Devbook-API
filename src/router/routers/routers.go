package routers

import (
	"net/http"

	"github.com/gorilla/mux"
)

// Route struct - estruturas de rotas do sistema
type Route struct {
	URI    string
	Method string
	Func   func(http.ResponseWriter, *http.Request)
	Auth   bool
}

// Configure cria todas as rotas disponiveis na api
func Configure(r *mux.Router) *mux.Router {
	routers := usersRoute

	for _, router := range routers {
		r.HandleFunc(router.URI, router.Func).Methods(router.Method)
	}

	return r
}
