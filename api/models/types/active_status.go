package types

//swagger: responce ActiveStatus
type ActiveStatus struct {
	GameId   int `json:"game_id"`
	Round    int `json:"round"`
	PlayerId int `json:"player_id"`
	Throw    int `json:"throw"`
}

//swagger: responce Extra
type Extra struct {
	Count  int    `json:"count"`
	Player int    `json:"player"`
	Type   string `json:"type"`
}
