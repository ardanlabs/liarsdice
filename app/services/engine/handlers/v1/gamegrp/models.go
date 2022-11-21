package gamegrp

import "github.com/ethereum/go-ethereum/common"

// Status represents the game status.
type Status struct {
	Status        string           `json:"status"`
	AnteUSD       float64          `json:"anteUSD"`
	LastOutAcctID common.Address   `json:"lastOut"`
	LastWinAcctID common.Address   `json:"lastWin"`
	CurrentAcctID common.Address   `json:"currentID"`
	Round         int              `json:"round"`
	Cups          []Cup            `json:"cups"`
	CupsOrder     []common.Address `json:"playerOrder"`
	Bets          []Bet            `json:"bets"`
	Balances      []string         `json:"balances"`
}

// Bet represents the bet response.
type Bet struct {
	AccountID common.Address `json:"account"`
	Number    int            `json:"number"`
	Suite     int            `json:"suite"`
}

// Cup represents the cup response.
type Cup struct {
	AccountID common.Address `json:"account"`
	Dice      []int          `json:"dice"`
	LastBet   Bet            `json:"lastBet"`
	Outs      int            `json:"outs"`
}
