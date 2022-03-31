package models

import "time"

type User struct {
	Id        int        `json:"id"`
	FirstName string     `json:"firstName"`
	LastName  string     `json:"lastName"`
	Email     string     `json:"email"`
	AvatarURL string     `json:"avatarUrl"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}

type Game struct {
	Id           int        `json:"id"`
	Name         string     `json:"gameName"`
	Type         string     `json:"gameType"`
	CreaterEmail string     `json:"createrEmail"`
	CreatedAt    *time.Time `json:"created_at,omitempty"`
	UpdatedAt    *time.Time `json:"updated_at,omitempty"`
}

type Score struct {
	Id         int        `json:"id"`
	Round      int        `json:"round"`
	Point      int        `json:"point"`
	PlayerName string     `json:"playerName"`
	Throw      int        `json:"throw"`
	CreatedAt  *time.Time `json:"created_at,omitempty"`
	UpdatedAt  *time.Time `json:"updated_at,omitempty"`
}

type Player struct {
	UserId    int        `json:"userId"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}
