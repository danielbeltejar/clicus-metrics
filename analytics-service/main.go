package main

import (
	"analytics-service/handlers"
	"context"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"net/http"
)

// @title Analytics Service API
// @version 1.0
// @description API for logging and retrieving URL click analytics
// @termsOfService http://clicus.danielbeltejar.es/terms/
// @contact.name API Support
// @contact.url http://danielbeltejar.es/
// @contact.email hi@danielbeltejar.es/
// @license.name MIT
// @license.url http://opensource.org/licenses/MIT
// @host localhost:8084
// @BasePath /
func main() {
	initConfig()

	// Connect to MongoDB
	clientOptions := options.Client().ApplyURI(MongoURI)
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}
	defer client.Disconnect(context.Background())

	db := client.Database("clicus_metrics")

	// Initialize analytics service
	handlers.InitAnalyticsService(db)

	r := mux.NewRouter()
	r.Use(handlers.AuthMiddleware) // Use the AuthMiddleware for all routes
	r.HandleFunc("/analytics/log/{url_id}", handlers.LogClick).Methods("POST")
	r.HandleFunc("/analytics/{url_id}", handlers.GetAnalytics).Methods("GET")

	log.Println("Analytics service running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
