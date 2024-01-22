package gamegrp

import "github.com/ethereum/go-ethereum/common"

// Status represents the game status.
type Status struct {
	Status          string           `json:"status"`
	AnteUSD         float64          `json:"anteUSD"`
	PlayerLastOut   common.Address   `json:"lastOut"`
	PlayerLastWin   common.Address   `json:"lastWin"`
	PlayerTurn      common.Address   `json:"currentID"`
	Round           int              `json:"round"`
	Cups            []Cup            `json:"cups"`
	ExistingPlayers []common.Address `json:"playerOrder"`
	Bets            []Bet            `json:"bets"`
	Balances        []string         `json:"balances"`
}

// Bet represents the bet response.
type Bet struct {
	Player common.Address `json:"account"`
	Number int            `json:"number"`
	Suit   int            `json:"suit"`
}

// Cup represents the cup response.
type Cup struct {
	Player  common.Address `json:"account"`
	Dice    []int          `json:"dice"`
	LastBet Bet            `json:"lastBet"`
	Outs    int            `json:"outs"`
}
