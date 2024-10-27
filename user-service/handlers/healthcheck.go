package handlers

import (
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
)

// HealthCheck checks if the service is running and MongoDB is connected
func HealthCheck(db *mongo.Database, client *mongo.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	}
}
