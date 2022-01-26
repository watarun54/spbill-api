package domain

import "time"

type Papers []Paper

type Paper struct {
	ID        int       `json:"id"`
	Text      string    `json:"text"`
	URL       string    `json:"url" gorm:"not null;size:255"`
	UserID    int       `json:"user_id"`
	IsDeleted *bool     `json:"is_deleted"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
