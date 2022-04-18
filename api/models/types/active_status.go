package types

type ActiveStatus struct {
	GameId   int `json:"game_id"`
	Round    int `json:"round"`
	PlayerId int `json:"player_id"`
	Throw    int `json:"throw"`
}

type NextTurn struct {
	Count  int    `json:"count"`
	Player int    `json:"player"`
	Type   string `json:"type"`
}
