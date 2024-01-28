package game

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

// Represents the different game status.
const (
	StatusNewGame    = "newgame"
	StatusPlaying    = "playing"
	StatusRoundOver  = "roundover"
	StatusGameOver   = "gameover"
	StatusReconciled = "reconciled"
)

// minNumberPlayers represents the minimum number of players required
// to play a game.
const minNumberPlayers = 2

// State represents a copy of the game state.
type State struct {
	GameID          string
	Round           int
	Status          string
	PlayerLastOut   common.Address
	PlayerLastWin   common.Address
	PlayerTurn      common.Address
	ExistingPlayers []common.Address
	Cups            map[common.Address]Cup
	Bets            []Bet
	Balances        []BalanceFmt
}

// Bet represents a bet of dice made by a player.
type Bet struct {
	Player common.Address
	Number int
	Suit   int
}

// Cup represents an individual cup being held by a player.
type Cup struct {
	Player   common.Address
	OrderIdx int
	Outs     int
	Dice     []int
}

// Balance represents an individual balance for a player.
type Balance struct {
	Player common.Address
	Amount *big.Float
}

// BalanceFmt represents an individual formatted balance for a player.
type BalanceFmt struct {
	Player common.Address
	Amount string
}
