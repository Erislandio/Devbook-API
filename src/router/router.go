package router

import (
	"api/src/router/routers"

	"github.com/gorilla/mux"
)

// Init inicializa as rotas da api
func Init() *mux.Router {
	router := mux.NewRouter()
	return routers.Configure(router)
}
