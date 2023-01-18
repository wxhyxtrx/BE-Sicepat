package routes

import (
	"server/handler"
	"server/pkg/connection"
	"server/pkg/middleware"
	"server/repositories"

	"github.com/gorilla/mux"
)

func OrderRoutes(r *mux.Router) {
	orderRepository := repositories.RepositoryOrder(connection.DB)
	h := handler.HandlerOrder(orderRepository)

	r.HandleFunc("/order", middleware.Auth(h.CreateOrder)).Methods("POST")
}
