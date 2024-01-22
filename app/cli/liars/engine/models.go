package engine

import "github.com/ethereum/go-ethereum/common"

// ErrorResponse is the form used for API responses from failures in the API.
type ErrorResponse struct {
	Error  string            `json:"error"`
	Fields map[string]string `json:"fields,omitempty"`
}

// Config represents the configuration of the game engine.
type Config struct {
	Network    string         `json:"network"`
	ChainID    int            `json:"chainId"`
	ContractID common.Address `json:"contractId"`
}

// Token contains the user game token and public address of the player.
type Token struct {
	Token   string         `json:"token"`
	Address common.Address `json:"address"`
}

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
	Suit      int            `json:"suit"`
}

// Cup represents the cup response.
type Cup struct {
	AccountID common.Address `json:"account"`
	Dice      []int          `json:"dice"`
	LastBet   Bet            `json:"lastBet"`
	Outs      int            `json:"outs"`
}
