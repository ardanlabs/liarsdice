// Package game represents all the game play for liar's dice.
package game

import (
	"errors"

	"github.com/google/uuid"
)

const (
	StatusEmpty = "empty"
	StatusNew   = "new"
	StatusPlay  = "playing"
)

// Player represents someone in the system.
type Player struct {
	UserID string
	Dice   []int
}

// Bet represents a single bet by a player.
type Bet struct {
	Player Player
	Number int
	Suite  int
}

// Game represents an instance of game play.
type Game struct {
	ID      string
	Price   int
	Round   int
	Next    int
	Players []*Player
	LastBet Bet
	LastOut *Player
	Outs    map[*Player]uint8
	Bets    [][]Bet
}

// Table represents a place players can play a game.
type Table struct {
	ID      string
	Name    string
	Price   int
	Status  string
	Players map[string]*Player
	Game    *Game
}

// NewTable constructs a table for players to use.
func NewTable(name string, price int) *Table {
	t := Table{
		ID:     uuid.NewString(),
		Name:   name,
		Price:  price,
		Status: StatusEmpty,
	}

	return &t
}

// AddPlayer adds a player to the table who can play in any future games.
func (t *Table) AddPlayer(userID string) error {
	if _, exists := t.Players[userID]; exists {
		return errors.New("player already on table")
	}

	t.Players[userID] = &Player{
		UserID: userID,
	}

	return nil
}

// RemovePlayer removes a player from the table so they can't play in
// any future games.
func (t *Table) RemovePlayer(userID string) error {
	if _, exists := t.Players[userID]; !exists {
		return errors.New("player doesn't exist on table")
	}

	delete(t.Players, userID)
	return nil
}

// StartNewGame creates a new game for the table.
func (t *Table) StartNewGame() error {
	if t.Status == StatusPlay {
		return errors.New("table is in the middle of a game")
	}

	players := make([]*Player, len(t.Players))
	outs := make(map[*Player]uint8)
	for _, player := range t.Players {
		players = append(players, player)
		outs[player] = 0
	}

	t.Status = StatusNew
	t.Game = &Game{
		ID:      uuid.NewString(),
		Price:   t.Price,
		Players: players,
		Outs:    outs,
		Bets:    make([][]Bet, 1),
	}

	return nil
}

// NewRound starts a new round of play with players who are not out. It returns
// the number of players left. If only 1 player is left, the game is over.
func (t *Table) NewRound() int {
	var players []*Player
	for player, outs := range t.Game.Outs {
		if outs != 3 {
			players = append(players, player)
		}
	}
	t.Game.Players = players

	if len(players) == 1 {
		return 1
	}

	var found bool
	for i, player := range t.Game.Players {
		if player.UserID == t.Game.LastOut.UserID {
			t.Game.Next = i
			found = true
		}
	}

	if !found {
		t.Game.Next--
		if t.Game.Next < 0 {
			t.Game.Next = len(t.Game.Players) - 1
		}
	}

	return len(players)
}

// NextTurn returns the next player who's turn it is to make a bet
func (t *Table) NextTurn() *Player {
	return t.Game.Players[t.Game.Next]
}

// MakeBet allows the specified player to make the next bet.
func (t *Table) MakeBet(p *Player, bet Bet) error {
	if p.UserID != t.Game.Players[t.Game.Next].UserID {
		return errors.New("wrong player making bet")
	}

	if bet.Number < t.Game.LastBet.Number {
		return errors.New("bet number must be greater or equal to the last bet")
	}

	if bet.Number == t.Game.LastBet.Number && bet.Suite <= t.Game.LastBet.Suite {
		return errors.New("bet suite must be greater that the last bet")
	}

	t.Game.LastBet = bet
	t.Game.Bets[t.Game.Round] = append(t.Game.Bets[t.Game.Round], bet)

	t.Game.Next++
	if t.Game.Next == len(t.Game.Players) {
		t.Game.Next = 0
	}

	return nil
}

// CallLiar ends the round and checks to see who won the round.
func (t *Table) CallLiar(p *Player) error {
	if p.UserID != t.Game.Players[t.Game.Next].UserID {
		return errors.New("wrong player calling lair")
	}

	return nil
}

// =============================================================================

// TableHistory maintains a history of every game played at a table.
type TableHistory struct {
	ID      string
	TableID string
	GameID  string
	Winner  *Player
	Losers  []*Player
	Payout  int
}
