package types

type Score struct {
	Score int `json:"score"`
}

type ResScore struct {
	Score       int  `json:"score"`
	TotalScore  int  `json:"total_score"`
	FoundWinner bool `json:"found_winner"`
}
