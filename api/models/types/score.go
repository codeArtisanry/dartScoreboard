package types

// Score score
// swagger:model score
type Score struct {
	Score int `json:"score"`
}

// swagger:response ResScore
type ResScore struct {
	// in: body
	Score       int  `json:"score"`
	TotalScore  int  `json:"total_score"`
	FoundWinner bool `json:"found_winner"`
}
