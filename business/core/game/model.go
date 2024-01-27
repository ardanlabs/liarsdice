package game

import (
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
	Status          string
	PlayerLastOut   common.Address
	PlayerLastWin   common.Address
	PlayerTurn      common.Address
	Round           int
	Cups            map[common.Address]Cup
	ExistingPlayers []common.Address
	Bets            []Bet
	Balances        []string
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
