package main

import (
	"net/http"
	"time"

	"github.com/avalokia/tgtcx/backend/database"
	"github.com/avalokia/tgtcx/backend/handlers"
	"github.com/avalokia/tgtcx/backend/server"
	"github.com/gorilla/mux"
)

func main() {
	// Init database connection
	database.InitDB()

	// Init serve HTTP
	router := mux.NewRouter()

	router.HandleFunc("/ping", handlers.Ping).Methods(http.MethodGet)

	serverConfig := server.Config{
		WriteTimeout: 5 * time.Second,
		ReadTimeout:  5 * time.Second,
		Port:         8000,
	}
	server.Serve(serverConfig, router)
}
