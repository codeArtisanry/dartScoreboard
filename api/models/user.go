package models

type User struct {
	FirstName string `json:"firstName"`
	LastName string `json:"lastName"`
	Email   string `json:"email"`
	AvatarURL string `json:"avatarUrl"`
}

type Game struct {
	Name        string `json:"gameName"`
	Type        string `json:"gameType"`
	CreaterEmail string `json:"createrEmail"`
}

type Score struct {
	Round int `json:"round"`
	Point int `json:"point"`
	PlayerName string `json:"playerName"`
	Throw int `json:"throw"`
}

type Player struct {
	UserId int `json:"userId"`
}
