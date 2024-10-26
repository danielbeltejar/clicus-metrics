package handlers

import (
	"context"
	"net/http"
	"url-service/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Redirect handles redirection to the original URL and logs the click
func Redirect(w http.ResponseWriter, r *http.Request) {
	shortID := r.URL.Query().Get("id")
	var url models.URL
	err := urlCollection.FindOne(context.TODO(), bson.M{"short_id": shortID}).Decode(&url)
	if err != nil {
		http.Error(w, "URL not found", http.StatusNotFound)
		return
	}

	// Increment click count
	_, err = urlCollection.UpdateOne(
		context.TODO(),
		bson.M{"short_id": shortID},
		bson.M{"$inc": bson.M{"clicks": 1}},
		options.Update().SetUpsert(true),
	)
	if err != nil {
		http.Error(w, "Error updating click count", http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, url.OriginalURL, http.StatusFound)
}
