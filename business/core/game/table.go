// Package game represents all the game play for liar's dice.
package game

import (
	"errors"
	"math/rand"
	"time"

	"github.com/google/uuid"
)

/*
General account

1. Player login occurs
2. They enter the game room
	a. Automatically added to the list of players.
	b. They can see their balance.
	c. They see the current game being played.
	---
	a. Button: Connect Wallet
	b. Button: Deposit Money
	c. Button: Withdrawl Money
	d. Button: Join/Leave Game


General game play

1. New Game Setup
	a. 30 second clock starts
	b. Ask players to Join game
		1. Check they have enough money for the ante
	c. Still have time to leave the game

2. New Game Start
	a. Must have at lease 2 players
	b. Randomly select the first player to bet

3. Game Play
	a. Players in slice order makes bets
	b. Player calls the last player a liar
		1. Validate bet against all the player dice to determine winner/loser
		2. Give loser an out

4. Round Over
	a. Remove players with 3 outs from the game
	b. If: Only one player is left Close Game (6)
	c. Else: Start Next Round (5)

5. Next Round
	a. Loser of last round starts the betting
	b. If loser was eliminated, the player who won the last round starts
	c. Back to (3) Game Play

6. Close Game
	a. Provide SC the winner, losers, ante, gameFee
	b. Reconcile the accounting
	c. New Game Setup (1)


Things to Consider

1. We can't allow a player to withdrawl while they are in a game.
2. A player closing their browser or logging out doesn't matter.
*/

const (
	StatusPlay      = "playing"
	StatusRoundOver = "roundover"
)

// Player represents someone in the system.
type Player struct {
	UserID  string
	Address string
	Dice    []int
}

// RollDice changes the dice in the players hand.
func (p *Player) RollDice() {
	dice := make([]int, 6)
	for i := range dice {
		dice[i] = rand.Intn(6) + 1
	}
	p.Dice = dice
}

// Bet represents a single bet by a player.
type Bet struct {
	Player *Player
	Number int
	Suite  int
}

// Game represents an instance of game play.
type Game struct {
	ID         string
	Players    []*Player
	nextPlayer int
	Bets       []Bet
	LastOut    *Player
	LastWin    *Player
	Outs       map[*Player]uint8
}

// Table represents a place players can play a game.
type Table struct {
	ID      string
	Ante    int
	Status  string
	Players map[string]*Player
	Game    *Game
}

// NewTable constructs a table for players to use.
func NewTable(ante int) *Table {
	t := Table{
		ID:     uuid.NewString(),
		Ante:   ante,
		Status: StatusRoundOver,
	}

	rand.Seed(time.Now().Unix())

	return &t
}

// AddPlayer adds a player to the table who can play in any future games.
func (t *Table) AddPlayer(userID string) error {
	if _, exists := t.Players[userID]; exists {
		return errors.New("player already at the table")
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
		return errors.New("player doesn't exist at table")
	}

	delete(t.Players, userID)
	return nil
}

// StartGame creates a new game for the table.
func (t *Table) StartGame() error {
	if t.Game != nil {
		return errors.New("table is in the middle of a game")
	}

	// Add all the existing players at the table to this new game.
	players := make([]*Player, len(t.Players))
	outs := make(map[*Player]uint8)
	for _, player := range t.Players {
		players = append(players, player)
		outs[player] = 0
	}

	t.Status = StatusPlay
	t.Game = &Game{
		ID:      uuid.NewString(),
		Players: players,
		Bets:    []Bet{},
		Outs:    outs,
	}

	return nil
}

// CloseGame closes the game and settles the accounts.
func (t *Table) CloseGame() error {
	if t.Game == nil {
		return errors.New("table doesn't have a current game")
	}

	// Close out the accounts and paid players.

	// Check the round is over.
	// if t.Status != StatusRoundOver {

	// 	// I guess we are shutting down this game and reseting the pot
	// }

	t.Game = nil

	return nil
}

// NewRound starts a new round of play with players who are not out. It returns
// the number of players left. If only 1 player is left, the game is over.
func (t *Table) NewRound() (int, error) {

	// Check the round is over.
	if t.Status != StatusRoundOver {
		return 0, errors.New("current round is not over")
	}

	// Figure out which players are left in the game from the close of
	// the previous round.
	var players []*Player
	for player, outs := range t.Game.Outs {
		if outs != 3 {
			players = append(players, player)
		}
	}
	t.Game.Players = players

	// If there is only 1 player left we have a winner.
	if len(players) == 1 {
		return 1, nil
	}

	// Figure out who starts the round. The person who was last out should
	// start the round.
	var found bool
	for i, player := range t.Game.Players {
		if player.UserID == t.Game.LastOut.UserID {
			t.Game.nextPlayer = i
			found = true
		}
	}

	// If the person who was last out is no longer in the game, then the
	// player who won the last round starts.
	if !found {
		for i, player := range t.Game.Players {
			if player.UserID == t.Game.LastWin.UserID {
				t.Game.nextPlayer = i
			}
		}
	}

	// Reset the bets to start over.
	t.Game.Bets = []Bet{}

	// Return the number of players for this round.
	return len(players), nil
}

// NextTurn returns the next player who's turn it is to make a bet
func (t *Table) NextTurn() *Player {
	return t.Game.Players[t.Game.nextPlayer]
}

// MakeBet allows the specified player to make the next bet.
func (t *Table) MakeBet(bet Bet) error {

	// Validate this player does have the next turn.
	if bet.Player.UserID != t.Game.Players[t.Game.nextPlayer].UserID {
		return errors.New("wrong player making bet")
	}

	// If this is not the first bet, validate the bet.
	if len(t.Game.Bets) != 0 {
		lastBet := t.Game.Bets[len(t.Game.Players)-1]

		if bet.Number < lastBet.Number {
			return errors.New("bet number must be greater or equal to the last bet")
		}

		if bet.Number == lastBet.Number && bet.Suite <= lastBet.Suite {
			return errors.New("bet suite must be greater that the last bet")
		}
	}

	// Add the bet to the set of bets for this round.
	t.Game.Bets = append(t.Game.Bets, bet)

	// Increment the next player index.
	t.Game.nextPlayer++
	if t.Game.nextPlayer == len(t.Game.Players) {
		t.Game.nextPlayer = 0
	}

	return nil
}

// CallLiar ends the round and checks to see who won the round. The losing
// player is returned.
func (t *Table) CallLiar(p *Player) (winner *Player, loser *Player, err error) {

	// Validate this player does have the next turn.
	if p.UserID != t.Game.Players[t.Game.nextPlayer].UserID {
		return nil, nil, errors.New("wrong player calling lair")
	}

	// Compare the last bet to all the dice the players are holding.
	t.Status = StatusRoundOver

	// Add up the number of each number of dice players have.
	dice := make([]int, 6)
	for _, player := range t.Game.Players {
		for _, suite := range player.Dice {
			dice[suite]++
		}
	}

	// Capture the last bet.
	lastBet := t.Game.Bets[len(t.Game.Players)-1]

	// Did the person calling Liar win?
	if dice[lastBet.Suite] < lastBet.Number {
		t.Game.Outs[lastBet.Player]++
		return p, lastBet.Player, nil
	}

	// The person calling Liar lost.
	t.Game.Outs[p]++
	return lastBet.Player, p, nil
}
