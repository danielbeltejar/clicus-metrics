package models

type DashboardData struct {
	ShortID     string   `json:"short_id"`
	OriginalURL string   `json:"original_url"`
	Clicks      int      `json:"clicks"`
	Tags        []string `json:"tags"`
}
