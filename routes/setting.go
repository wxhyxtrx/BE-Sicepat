package routes

import (
	"server/handler"
	"server/pkg/connection"
	"server/pkg/middleware"
	"server/repositories"

	"github.com/gorilla/mux"
)

func SettingRoutes(r *mux.Router) {
	settingRepository := repositories.RepositorySetting(connection.DB)
	h := handler.HandlerSetting(settingRepository)

	r.HandleFunc("/setting", middleware.Auth(h.CreateSetting)).Methods("POST")
}
