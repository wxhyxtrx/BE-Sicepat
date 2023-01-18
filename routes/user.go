package routes

import (
	"server/handler"
	"server/pkg/connection"
	"server/pkg/middleware"
	"server/repositories"

	"github.com/gorilla/mux"
)

func UserRoutes(r *mux.Router) {
	userReposetory := repositories.RepositoryUser(connection.DB)
	h := handler.HandlerUser(userReposetory)

	r.HandleFunc("/users", middleware.Auth(h.FindAllUser)).Methods("GET")
	r.HandleFunc("/user/{id}", middleware.Auth(h.GetUser)).Methods("GET")
	r.HandleFunc("/user", h.CreateUser).Methods("POST")
	r.HandleFunc("/user/{id}", middleware.Auth(h.UpdateUser)).Methods("PATCH")
	r.HandleFunc("/user/{id}", middleware.Auth(h.DeleteUser)).Methods("DELETE")
}
