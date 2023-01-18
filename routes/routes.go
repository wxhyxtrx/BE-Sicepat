package routes

import (
	"github.com/gorilla/mux"
)

func RouteInit(r *mux.Router) {
	UserRoutes(r)
	AuthRoutes(r)
	RoomRoutes(r)
	ChatRoutes(r)
	SettingRoutes(r)

	TiketRoutes(r)
	OrderRoutes(r)
}
