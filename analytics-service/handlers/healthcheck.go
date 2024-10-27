package handlers

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
)

// HealthCheck checks if the service is running and MongoDB is connected.
func HealthCheck(db *mongo.Database) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := db.Client().Ping(context.Background(), nil); err != nil {
			http.Error(w, "MongoDB connection failed", http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Service is running and MongoDB is connected"))
	}
}
