package model

import "time"

type Notification struct {
	ID        int64     `json:"id" db:"id"`
	UserID    int64     `json:"user_id" db:"user_id"`
	Title     string    `json:"title" db:"title"`
	Body      string    `json:"body" db:"body"`
	Type      string    `json:"type" db:"type"`
	Read      bool      `json:"read" db:"read"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}

type NotificationList struct {
	Items  []Notification `json:"items"`
	Total  int64          `json:"total"`
	Unread int64          `json:"unread"`
}
