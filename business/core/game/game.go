// Package game represents all the game play for liar's dice.
package game

import (
	"context"
	"errors"
	"fmt"
	"math/big"
	"math/rand"
	"sync"

	"github.com/google/uuid"
)

// Represents the different game statues.
const (
	StatusRoundOver = "roundover"
	StatusPlaying   = "playing"
	StatusOpen      = "open"
)

// MinNumberPlayers represents the minimum number of players required
// to play a game.
const MinNumberPlayers = 2

// =============================================================================

// Banker represents the ability to manage money for the game. Deposits and
// Withdrawls happen outside of game play.
type Banker interface {
	Balance(ctx context.Context, account string) (*big.Int, error)
	Reconcile(ctx context.Context, winningAccount string, losingAccounts []string, ante uint, gameFee uint) error
}

// =============================================================================

// Cup represents an individual cup being held by a player.
type Cup struct {
	Account string
	Outs    uint8
	Dice    []int
}

// Claim represents a claim of dice on the table.
type Claim struct {
	Account string
	Number  int
	Suite   int
}

// =============================================================================

// Game represents a single game that is being played.
type Game struct {
	ID            string
	Status        string
	banker        Banker
	lastOutAcct   string
	lastWinAcct   string
	CurrentPlayer int
	Round         int
	Cups          map[string]Cup
	Claims        []Claim

	mu sync.RWMutex
}

// New creates a new game.
func New(banker Banker) *Game {
	return &Game{
		ID:     uuid.NewString(),
		Status: StatusOpen,
		banker: banker,
		Cups:   make(map[string]Cup),
	}
}

// Start will check if the current game can be started and update its status.
func (g *Game) Start() error {
	g.mu.Lock()
	defer g.mu.Unlock()

	if g.Status != StatusOpen {
		return errors.New("game cannot be started")
	}

	if len(g.Cups) < MinNumberPlayers {
		return errors.New("not enough players to start game")
	}

	g.Round = 1
	g.Status = StatusPlaying

	return nil
}

// AddAccount adds a player to the game. If the account already exists, the
// function will return an error.
func (g *Game) AddAccount(account string) error {
	if account == "" {
		return errors.New("invalid account information")
	}

	g.mu.Lock()
	defer g.mu.Unlock()

	if _, exists := g.Cups[account]; exists {
		return fmt.Errorf("player [%s] is already in the game", account)
	}

	g.Cups[account] = Cup{Account: account}

	return nil
}

// RemoveAccount removes a player from the game. If the account does not exist,
// the function will return an error.
func (g *Game) RemoveAccount(account string) error {
	if account == "" {
		return errors.New("invalid account information")
	}

	g.mu.Lock()
	defer g.mu.Unlock()

	if _, exists := g.Cups[account]; !exists {
		return fmt.Errorf("player [%s] does not exist in the game", account)
	}

	delete(g.Cups, account)

	return nil
}

// RollDice will generate 5 new random integers for the players cup.
func (g *Game) RollDice(account string) error {
	if g.Status != StatusPlaying {
		return errors.New("game is not started")
	}

	if account == "" {
		return errors.New("invalid account information")
	}

	g.mu.Lock()
	defer g.mu.Unlock()

	cup, exists := g.Cups[account]
	if !exists {
		return fmt.Errorf("player [%s] does not exist in the game", account)
	}

	for i := range cup.Dice {
		cup.Dice[i] = rand.Intn(6) + 1
	}

	return nil
}

// CallLiar checks all the claims made so far in the round and defines a winner
// and a loser.
func (g *Game) CallLiar(wallet string) (string, string, error) {
	if wallet == "" {
		return "", "", errors.New("invalid wallet address")
	}

	// Validate if it is the player's turn..
	if g.Players[g.CurrentPlayer].Wallet != wallet {
		return "", "", fmt.Errorf("player [%s] can't make a claim now", wallet)
	}

	// Hold the sum of all the dice values.
	dice := make([]int, 7)
	for _, player := range g.Players {
		for _, suite := range player.Dice {
			dice[suite]++
		}
	}

	// Find player who called a liar.
	callPlayer, _, found := g.findPlayer(wallet)
	if !found {
		return "", "", fmt.Errorf("player [%s] was not found", wallet)
	}

	g.Status = STATUSROUNDOVER

	lastClaim := g.claims[len(g.claims)-1]

	// If the last Claim is incorrect, the player who made it, gets an out.
	if dice[lastClaim.Suite] < lastClaim.Number {
		lastClaim.player.Outs++
		g.lastOut = lastClaim.player
		g.lastWin = callPlayer

		return wallet, g.lastOut.Wallet, nil
	}

	callPlayer.Outs++
	g.lastOut = callPlayer
	g.lastWin = lastClaim.player

	return lastClaim.player.Wallet, wallet, nil
}

// NewRound checks for players out count, reset players dice and game claims.
func (g *Game) NewRound() (int, error) {

	// Check the round is over.
	if g.Status != STATUSROUNDOVER {
		return 0, errors.New("current round is not over")
	}

	g.Round++

	// Figure out which players are left in the game from the close of
	// the previous round.
	for _, player := range g.Players {
		if player.Outs == 3 {
			g.RemovePlayer(player.Wallet)
		}
	}

	// If there is only 1 player left we have a winner.
	if len(g.Players) == 1 {
		g.Status = STATUSOPEN
		return 1, nil
	}

	// Figure out who starts the round. The person who was last out should
	// start the round.
	_, i, found := g.findPlayer(g.lastOut.Wallet)
	if found {
		g.CurrentPlayer = i
	}

	// If the person who was last out is no longer in the game, then the
	// player who won the last round starts.
	if !found {
		_, i, found := g.findPlayer(g.lastWin.Wallet)
		if found {
			g.CurrentPlayer = i
		}
	}

	// Reset players dice.
	for i := range g.Players {
		g.Players[i].Dice = []int{}
	}

	// Reset the claims to start over.
	g.claims = []Claim{}

	g.Status = STATUSPLAYING

	// Return the number of players for this round.
	return len(g.Players), nil
}

// Claim checks if the claim is valid and made by the correct player before
// adding it to the list of claims for the current game.
func (g *Game) Claim(wallet string, claim Claim) error {
	if wallet == "" {
		return errors.New("invalid wallet address")
	}

	// Validate if it is the player's turn.
	if g.Players[g.CurrentPlayer].Wallet != wallet {
		return fmt.Errorf("player [%s] can't make a claim now", wallet)
	}

	// Validate this player have rolled the dice.
	if g.Players[g.CurrentPlayer].Dice == nil {
		return fmt.Errorf("player [%s] didn't roll dice yet", wallet)
	}

	// If this is not the first claim, validate it against the previous claim.
	if len(g.claims) != 0 {
		lastClaim := g.claims[len(g.claims)-1]

		if claim.Number < lastClaim.Number {
			return errors.New("claim number must be greater or equal to the last claim")
		}

		if claim.Number == lastClaim.Number && claim.Suite <= lastClaim.Suite {
			return errors.New("claim suite must be greater that the last claim")
		}
	}

	player, _, found := g.findPlayer(wallet)
	if !found {
		return fmt.Errorf("player [%s] was not found", wallet)
	}

	// Specify who made the claim.
	claim.player = player

	g.claims = append(g.claims, claim)

	// Update the CurrentPlayer index.
	g.CurrentPlayer++
	g.CurrentPlayer = g.CurrentPlayer % len(g.Players)

	g.Round++

	return nil
}

// PlayerBalance returns the player's balance, by calling the banks contract
// method.
func (g *Game) PlayerBalance(ctx context.Context, wallet string) (*big.Int, error) {
	if wallet == "" {
		return nil, errors.New("invalid wallet address")
	}

	return g.banker.Balance(ctx, wallet)
}

// Reconcile calculates the game pot and make the transfer to the winner.
func (g *Game) Reconcile(ctx context.Context, winner string, losers []string, ante uint, gameFee uint) error {
	return nil
}
