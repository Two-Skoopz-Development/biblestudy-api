package routes

import (
	"github.com/Two-Skoopz-Development/biblestudy-api/internal/api/services"
	"github.com/gorilla/mux"
)

func UserRoutes(r *mux.Router) {
	r.HandleFunc("/user", services.UserHandler)
}
