package routes

import (
	"server/handler"
	"server/pkg/connection"
	"server/pkg/middleware"
	"server/repositories"

	"github.com/gorilla/mux"
)

func AuthRoutes(r *mux.Router) {
	userRepository := repositories.RepositoryUser(connection.DB)
	h := handler.HandlerAuth(userRepository)

	r.HandleFunc("/register", h.Register).Methods("POST")
	r.HandleFunc("/login", h.Login).Methods("POST")
	r.HandleFunc("/checkauth", middleware.Auth(h.CheckAuth)).Methods("GET")
}
