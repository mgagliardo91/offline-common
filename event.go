package common

import (
	"time"

	"github.com/mgagliardo91/go-utils"
)

type OfflineEvent struct {
	Address       string   `json:"address"`
	Category      string   `json:"category"`
	Description   string   `json:"description"`
	EndDateTime   int64    `json:"endDateTime"`
	EventURL      string   `json:"eventUrl"`
	EventID       string   `json:"eventId"`
	ID            string   `json:"id"`
	ImageURL      string   `json:"imageUrl"`
	Latitude      float64  `json:"lat"`
	Longitude     float64  `json:"lng"`
	OfflineURL    string   `json:"offlineUrl"`
	Price         float64  `json:"price"`
	ReferralURLs  []string `json:"referralUrls"`
	StartDateTime int64    `json:"startDateTime"`
	Tags          []string `json:"tags"`
	Teaser        string   `json:"teaser"`
	Title         string   `json:"title"`
}

func (o *OfflineEvent) GetStartDateTime() time.Time {
	return utils.MilisToTime(o.StartDateTime)
}

func (o *OfflineEvent) GetEndDateTime() time.Time {
	return utils.MilisToTime(o.EndDateTime)
}
