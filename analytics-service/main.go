package main

import (
	"analytics-service/handlers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	InitConfig()
	db, client := InitMongo()
	handlers.InitAnalyticsService(db, JwtSecret)

	r := mux.NewRouter()

	// Health check endpoint without authentication middleware
	r.HandleFunc("/healthz", handlers.HealthCheck(db, client)).Methods("GET")

	// Routes that require authentication
	securedRoutes := r.PathPrefix("/analytics").Subrouter()
	securedRoutes.Use(handlers.AuthMiddleware(JwtSecret)) // Apply middleware only to these routes
	securedRoutes.HandleFunc("/log/{url_id}", handlers.LogClick).Methods("POST")
	securedRoutes.HandleFunc("/{url_id}", handlers.GetAnalytics).Methods("GET")

	log.Println("Analytics service running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
