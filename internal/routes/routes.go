package routes

import (
	"github.com/gorilla/mux"
)

type Router struct {
	MuxRouter *mux.Router
}

func ConfigureRouter() Router {

	router := Router{MuxRouter: mux.NewRouter()}
	UserRoutes(router.MuxRouter)

	return router
}
