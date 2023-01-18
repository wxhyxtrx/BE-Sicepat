package routes

import (
	"server/handler"
	"server/pkg/connection"
	"server/pkg/middleware"
	"server/repositories"

	"github.com/gorilla/mux"
)

func ChatRoutes(r *mux.Router) {
	chatRepository := repositories.RepositoryChat(connection.DB)
	h := handler.HandlerChat(chatRepository)

	r.HandleFunc("/chat", middleware.Auth(h.CreateChat)).Methods("POST")
}
