package common

import (
	"sync"
	"time"
)

// OfflineEvent represents an "event" such as a restaurant opening or festival
type OfflineEvent struct {
	OfflineDate  time.Time `json:"offlineDate"`
	OfflineURL   string    `json:"offlineUrl"`
	ImageURL     string    `json:"imageUrl"`
	Title        string    `json:"title"`
	Teaser       string    `json:"teaser"`
	Description  string    `json:"description"`
	LocationRaw  string    `json:"locationRaw"`
	DateRaw      string    `json:"dateRaw"`
	TimeRaw      string    `json:"timeRaw"`
	PriceRaw     string    `json:"priceRaw"`
	EventURL     string    `json:"eventUrl"`
	ReferralURLs []string  `json:"referralUrls"`
	mux          *sync.Mutex
}

func NewOfflineEvent() *OfflineEvent {
	return &OfflineEvent{
		mux:          &sync.Mutex{},
		ReferralURLs: make([]string, 0),
	}
}

type UpdateFunc func()

func (event *OfflineEvent) LockAndUpdate(update UpdateFunc) {
	event.mux.Lock()
	defer event.mux.Unlock()
	update()
}
