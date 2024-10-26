// main.go
package main

import (
	"context"
	"dashboard-service/handlers"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	initMongoDB()
	defer mongoClient.Disconnect(context.Background())

	// Initialize handlers with the URL collection
	handlers.InitHandlers(urlCollection)

	r := mux.NewRouter()
	r.HandleFunc("/dashboard", handlers.GetDashboardData).Methods("GET")

	log.Println("Dashboard service running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
