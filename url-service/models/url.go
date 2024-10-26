package models

import "time"

type URL struct {
	ID          string    `bson:"_id,omitempty" json:"id"`
	OriginalURL string    `bson:"original_url" json:"original_url"`
	ShortID     string    `bson:"short_id" json:"short_id"`
	Tags        []string  `bson:"tags,omitempty" json:"tags"`
	CreatedAt   time.Time `bson:"created_at" json:"created_at"`
	Clicks      int       `bson:"clicks" json:"clicks"`
}
