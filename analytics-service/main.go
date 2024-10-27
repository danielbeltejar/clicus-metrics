package main

import (
	"analytics-service/handlers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	InitConfig()
	db := InitMongo()
	handlers.InitAnalyticsService(db, JwtSecret)

	r := mux.NewRouter()
	r.Use(handlers.AuthMiddleware(JwtSecret)) // Pass JwtSecret to AuthMiddleware
	r.HandleFunc("/analytics/log/{url_id}", handlers.LogClick).Methods("POST")
	r.HandleFunc("/analytics/{url_id}", handlers.GetAnalytics).Methods("GET")
	r.HandleFunc("/healthz", handlers.HealthCheck(db)).Methods("GET")

	log.Println("Analytics service running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
