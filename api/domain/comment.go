package domain

import "time"

type Comments []Comment

type Comment struct {
	ID        int       `json:"id"`
	Text      string    `json:"text"`
	UserID    int       `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
