package handlers

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
	"time"
)

// HealthCheck checks if the service is running and MongoDB is connected
func HealthCheck(db *mongo.Database, client *mongo.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		if err := client.Ping(ctx, nil); err != nil {
			http.Error(w, "Database is not reachable", http.StatusServiceUnavailable)
			return
		}

		if err := db.Client().Database("clicus").RunCommand(ctx, bson.D{{"ping", 1}}).Err(); err != nil {
			http.Error(w, "Database is not responsive", http.StatusServiceUnavailable)
			return
		}

		// If everything is fine, return a 200 OK status
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	}
}
