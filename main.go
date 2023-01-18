package main

import (
	"fmt"
	"net/http"
	"server/database"
	"server/pkg/connection"
	"server/routes"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {
	// err := godotenv.Load()
	// if err != nil {
	// 	log.Fatal("Error loading .env file")
	// }

	connection.Database()
	database.RunMigration()

	r := mux.NewRouter()

	routes.RouteInit(r.PathPrefix("/api/v1").Subrouter())

	var AllowedHeaders = handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	var AllowedMethods = handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS", "PATCH", "DELETE"})
	var AllowedOrigins = handlers.AllowedOrigins([]string{"*"})

	var port = "5000"
	fmt.Println("Server running localhost:" + port)
	http.ListenAndServe("localhost:"+port, handlers.CORS(AllowedHeaders, AllowedMethods, AllowedOrigins)(r))
}
