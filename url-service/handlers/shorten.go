package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"time"
	"url-service/models"
	"url-service/utils"

	"go.mongodb.org/mongo-driver/mongo"
)

var urlCollection *mongo.Collection

func InitHandlers(db *mongo.Database) {
	urlCollection = db.Collection("urls")
}

// ShortenURL creates a new short URL and stores it in MongoDB
func ShortenURL(w http.ResponseWriter, r *http.Request) {
	var req struct {
		OriginalURL string   `json:"original_url"`
		Tags        []string `json:"tags"`
	}
	json.NewDecoder(r.Body).Decode(&req)

	shortID := utils.GenerateShortID()
	url := models.URL{
		ID:          shortID,
		OriginalURL: req.OriginalURL,
		ShortID:     shortID,
		Tags:        req.Tags,
		CreatedAt:   time.Now(),
	}

	_, err := urlCollection.InsertOne(context.TODO(), url)
	if err != nil {
		http.Error(w, "Error creating short URL", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"short_url": shortID})
}
