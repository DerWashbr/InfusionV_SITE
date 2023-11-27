package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Streamer struct {
	gorm.Model
	ID           uuid.UUID  `gorm:"column:id"`
	Username     string     `gorm:"index:idx_username,unique;column:user_name"`
	Charname     string     `gorm:"column:charname"`
	Fraktion     string     `gorm:"column:fraktion"`
	ShortStory   string     `gorm:"column:shortStory"`
	Title        *string    `gorm:"column:title"`
	Followers    int        `gorm:"column:followers"`
	IsLive       bool       `gorm:"column:isLive"`
	Viewers      *int       `gorm:"column:viewers"`
	Active       bool       `gorm:"column:active"`
	DataFetched  bool       `gorm:"column:dataFetched"`
	ProfileImg   string     `gorm:"column:profileImg"`
	ThumbnailImg *string    `gorm:"column:thumbnailImg"`
	LastStreamed *time.Time `gorm:"column:startedAt"`
}

func (Streamer) TableName() string {
	return "twitch"
}

type StreamerDetails struct {
	Username   string `json:"user_name"`
	ProfileImg string `json:"profile_img"`
	Followers  int    `json:"followers"`
	Fraktion   string `json:"fraktion"`
	Charname   string `json:"charname"`
	ShortStory string `json:"short_story"`
}
