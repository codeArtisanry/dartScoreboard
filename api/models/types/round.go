package types

import "time"

type Round struct {
	Id        int        `json:"id"`
	Round     int        `json:"round"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}
