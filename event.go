package common

import "time"

// OfflineEvent represents an "event" such as a restaurant opening or festival
type OfflineEvent struct {
	Date    time.Time `json:"date"`
	Title   string    `json:"title"`
	Details string    `json:"details"`
	URL     string    `json:"url"`
}
