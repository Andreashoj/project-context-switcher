package models

import "time"

type Project struct {
	Id        uint      `json:"id"`
	Name      string    `json:"name"`
	Path      string    `json:"path"` // TODO: Add to migration
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
