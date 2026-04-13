package model

import "time"

type Tag struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name"`
	Slug      string    `json:"slug"`
	Color     string    `json:"color"`
	CreatedAt time.Time `json:"created_at"`
}

type ItemTag struct {
	ItemID int64 `json:"item_id"`
	TagID  int64 `json:"tag_id"`
}
