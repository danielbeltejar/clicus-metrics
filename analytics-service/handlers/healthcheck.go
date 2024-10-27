package handlers

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"net/http"
	"time"
)

// HealthCheck checks if the service is running and MongoDB is connected
func HealthCheck(db *mongo.Database) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx, cancel := context.WithTimeout(r.Context(), 2*time.Second)
		defer cancel()

		// Ping MongoDB to ensure the connection is alive
		if err := db.Client().Ping(ctx, nil); err != nil {
			log.Println("MongoDB connection error:", err)
			http.Error(w, "Service unavailable", http.StatusServiceUnavailable)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	}
}
