package domain

import "time"

// ItemStatus represents the lifecycle state of an item.
type ItemStatus string

const (
	ItemStatusActive   ItemStatus = "active"
	ItemStatusArchived ItemStatus = "archived"
	ItemStatusDraft    ItemStatus = "draft"
)

// Item represents the core domain entity.
type Item struct {
	ID          int64      `json:"id"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	Status      ItemStatus `json:"status"`
	OwnerID     int64      `json:"owner_id"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
}

type ItemCreate struct {
	Title       string     `json:"title"`
	Description string     `json:"description"`
	Status      ItemStatus `json:"status"`
}

type ItemUpdate struct {
	Title       *string     `json:"title"`
	Description *string     `json:"description"`
	Status      *ItemStatus `json:"status"`
}
