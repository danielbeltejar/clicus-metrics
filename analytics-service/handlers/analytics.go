// handlers/analytics.go
package handlers

import (
	"analytics-service/models"
	"context"
	"encoding/json"
	"github.com/golang-jwt/jwt/v4"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
	"time"
)

var analyticsCollection *mongo.Collection
var jwtSecret string

// InitAnalyticsService initializes the analytics service and sets the JWT secret.
func InitAnalyticsService(db *mongo.Database, secret string) {
	analyticsCollection = db.Collection("analytics")
	jwtSecret = secret
}

// LogClick logs a click for a given URL ID.
func LogClick(w http.ResponseWriter, r *http.Request) {
	urlID := mux.Vars(r)["url_id"]
	var tags []string

	if err := json.NewDecoder(r.Body).Decode(&tags); err != nil {
		http.Error(w, "Failed to parse request body", http.StatusBadRequest)
		return
	}

	click := models.ClickAnalytics{
		URLID:     urlID,
		Timestamp: time.Now(),
		Tags:      tags,
	}

	_, err := analyticsCollection.InsertOne(context.Background(), click)
	if err != nil {
		http.Error(w, "Failed to log click", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Click logged successfully"})
}

// GetAnalytics retrieves analytics data for a given URL ID.
func GetAnalytics(w http.ResponseWriter, r *http.Request) {
	urlID := mux.Vars(r)["url_id"]

	var clicks []models.ClickAnalytics
	cursor, err := analyticsCollection.Find(context.Background(), bson.M{"url_id": urlID})
	if err != nil {
		http.Error(w, "Failed to retrieve analytics", http.StatusInternalServerError)
		return
	}
	defer cursor.Close(context.Background())

	if err := cursor.All(context.Background(), &clicks); err != nil {
		http.Error(w, "Failed to decode analytics", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(clicks)
}

// AuthMiddleware is a middleware for authenticating requests using JWT.
func AuthMiddleware(secret string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			tokenString := r.Header.Get("Authorization")
			if tokenString == "" {
				http.Error(w, "Authorization header is required", http.StatusUnauthorized)
				return
			}

			tokenString = tokenString[len("Bearer "):]

			claims := &jwt.StandardClaims{}
			token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
				return []byte(secret), nil
			})

			if err != nil || !token.Valid {
				http.Error(w, "Invalid token", http.StatusUnauthorized)
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}
