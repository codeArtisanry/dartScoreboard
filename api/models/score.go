package models

import "time"

type Score struct {
	Id        int        `json:"id"`
	RoundId   int        `json:"roundId"`
	PlayerID  int        `json:"playerId"`
	Throw     int        `json:"throw"`
	Score     int        `json:"score"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}
