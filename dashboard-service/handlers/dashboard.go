// handlers/dashboard.go
package handlers

import (
	"context"
	"dashboard-service/models"
	"encoding/json"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var urlCollection *mongo.Collection

// InitHandlers initializes the URL collection for the handlers
func InitHandlers(db *mongo.Database) {
	urlCollection = db.Collection("urls")
}

func GetDashboardData(w http.ResponseWriter, r *http.Request) {
	cursor, err := urlCollection.Find(context.TODO(), bson.M{})
	if err != nil {
		http.Error(w, "Error retrieving data", http.StatusInternalServerError)
		return
	}
	defer cursor.Close(context.TODO())

	var data []models.DashboardData
	for cursor.Next(context.TODO()) {
		var url models.DashboardData
		if err := cursor.Decode(&url); err != nil {
			http.Error(w, "Error decoding data", http.StatusInternalServerError)
			return
		}
		data = append(data, url)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}
