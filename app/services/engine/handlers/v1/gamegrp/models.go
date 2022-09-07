package gamegrp

// Status represents the game status.
type Status struct {
	Status        string   `json:"status"`
	AnteUSD       float64  `json:"anteUSD"`
	LastOutAcctID string   `json:"lastOut"`
	LastWinAcctID string   `json:"lastWin"`
	CurrentAcctID string   `json:"currentID"`
	Round         int      `json:"round"`
	Cups          []Cup    `json:"cups"`
	CupsOrder     []string `json:"playerOrder"`
	Bets          []Bet    `json:"bets"`
	Balances      []string `json:"balances"`
}

// Bet represents the bet response.
type Bet struct {
	AccountID string `json:"account"`
	Number    int    `json:"number"`
	Suite     int    `json:"suite"`
}

// Cup represents the cup response.
type Cup struct {
	AccountID string `json:"account"`
	Dice      []int  `json:"dice"`
	LastBet   Bet    `json:"lastBet"`
	Outs      int    `json:"outs"`
}
