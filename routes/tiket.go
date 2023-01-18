package routes

import (
	"server/handler"
	"server/pkg/connection"
	"server/pkg/middleware"
	"server/repositories"

	"github.com/gorilla/mux"
)

func TiketRoutes(r *mux.Router) {
	tiketRepository := repositories.RepositoryTiket(connection.DB)
	h := handler.HandlerTiket(tiketRepository)

	r.HandleFunc("/tiket", middleware.Auth(h.CreateTiket)).Methods("POST")
}
