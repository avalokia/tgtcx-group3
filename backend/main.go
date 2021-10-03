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

	router.HandleFunc("/upload-coupon", handlers.UploadCoupon).Methods(http.MethodPost)
	router.HandleFunc("/get-coupon-list", handlers.GetCouponList).Methods(http.MethodGet)
	router.HandleFunc("/update-coupon", handlers.UpdateCoupon).Methods(http.MethodPatch)
	router.HandleFunc("/set-coupon-duration", handlers.SetCouponDuration).Methods(http.MethodPatch)
	router.HandleFunc("/set-target-users", handlers.SetTargetUsers).Methods(http.MethodPatch)

	serverConfig := server.Config{
		WriteTimeout: 5 * time.Second,
		ReadTimeout:  5 * time.Second,
		Port:         8000,
	}
	server.Serve(serverConfig, router)
}
