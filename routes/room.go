package routes

import (
	"server/handler"
	"server/pkg/connection"
	"server/pkg/middleware"
	"server/repositories"

	"github.com/gorilla/mux"
)

func RoomRoutes(r *mux.Router) {
	roomRepository := repositories.RepositoryRoom(connection.DB)
	h := handler.HandlerRoom(roomRepository)

	r.HandleFunc("/room", middleware.Auth(h.CreateRoom)).Methods("POST")
	r.HandleFunc("/room/{id}", middleware.Auth(h.GetRoom)).Methods("GET")
}
