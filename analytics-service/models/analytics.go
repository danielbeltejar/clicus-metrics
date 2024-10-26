// models/analytics.go
package models

import "time"

// ClickAnalytics represents the analytics data for clicks.
type ClickAnalytics struct {
	ID        string    `json:"id" bson:"_id,omitempty"`
	URLID     string    `json:"url_id" bson:"url_id"`
	Timestamp time.Time `json:"timestamp" bson:"timestamp"`
	Tags      []string  `json:"tags" bson:"tags"`
}
