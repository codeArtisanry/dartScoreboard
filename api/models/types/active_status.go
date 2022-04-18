package types

type ActiveStatus struct {
	GameId   int `json:"game_id"`
	Round    int `json:"round"`
	PlayerId int `json:"player_id"`
	Throw    int `json:"throw"`
}
