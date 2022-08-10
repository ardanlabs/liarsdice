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
	CurrentPlayer string
	Round         int
	currentCup    int
	Cups          map[string]Cup
	CupsOrder     []string
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

	g.Cups[account] = Cup{
		Account: account,
		Dice:    make([]int, 5),
	}

	g.CupsOrder = append(g.CupsOrder, account)

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

	for i, acc := range g.CupsOrder {
		if acc == account {
			g.CupsOrder[i] = ""
		}
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

// Claim checks if the claim is valid and made by the correct player before
// adding it to the list of claims for the current game.
func (g *Game) Claim(account string, claim Claim) error {
	if account == "" {
		return errors.New("invalid account information")
	}

	// Get the current cup account.
	currentAccount := g.CupsOrder[g.currentCup]

	// Validate if it is the player's turn.
	if g.Cups[currentAccount].Account != account {
		return fmt.Errorf("player [%s] can't make a claim now", account)
	}

	// Validate this player have rolled the dice.
	if g.Cups[account].Dice == nil {
		return fmt.Errorf("player [%s] didn't roll dice yet", account)
	}

	// If this is not the first claim, validate it against the previous claim.
	if len(g.Claims) != 0 {
		lastClaim := g.Claims[len(g.Claims)-1]

		if claim.Number < lastClaim.Number {
			return errors.New("claim number must be greater or equal to the last claim")
		}

		if claim.Number == lastClaim.Number && claim.Suite <= lastClaim.Suite {
			return errors.New("claim suite must be greater that the last claim")
		}
	}

	// Specify who made the claim.
	claim.Account = account

	g.Claims = append(g.Claims, claim)

	if err := g.nextCup(); err != nil {
		return err
	}

	g.Round++

	return nil
}

// CallLiar checks all the claims made so far in the round and defines a winner
// and a loser.
func (g *Game) CallLiar(account string) (string, string, error) {
	if account == "" {
		return "", "", errors.New("invalid account information")
	}

	if g.Status != StatusPlaying {
		return "", "", errors.New("cannot call liar when game is not playable")
	}

	// Get the current cup account.
	currentAccount := g.CupsOrder[g.currentCup]

	// Validate if it is the player's turn.
	if g.Cups[currentAccount].Account != account {
		return "", "", fmt.Errorf("player [%s] can't make a claim now", account)
	}

	// Hold the sum of all the dice values.
	dice := make([]int, 7)
	for _, player := range g.Cups {
		for _, suite := range player.Dice {
			dice[suite]++
		}
	}

	// This call ends the round, not allowing any more claims to be made.
	g.Status = StatusRoundOver

	lastClaim := g.Claims[len(g.Claims)-1]

	// Find player who called the last claim.
	lastClaimPlayer := g.Cups[lastClaim.Account]

	// Find player who called liar.
	callPlayer := g.Cups[account]

	// The player who made the last claim loses the round.
	if dice[lastClaim.Suite] < lastClaim.Number {
		lastClaimPlayer.Outs++

		g.lastOutAcct = lastClaimPlayer.Account
		g.lastWinAcct = callPlayer.Account

		// Update the player data because of the outs count.
		g.Cups[lastClaim.Account] = lastClaimPlayer

		return g.lastWinAcct, g.lastOutAcct, nil
	}

	// The player who made the call loses the round.
	callPlayer.Outs++
	g.Cups[account] = callPlayer

	g.lastOutAcct = callPlayer.Account
	g.lastWinAcct = lastClaimPlayer.Account

	return g.lastWinAcct, g.lastOutAcct, nil
}

// NewRound checks for player's out count, reset players dice and game claims.
func (g *Game) NewRound() (int, error) {

	// Check the round is over.
	if g.Status != StatusRoundOver {
		return 0, errors.New("current round is not over")
	}

	// Figure out which players are left in the game from the close of
	// the previous round.
	for account, player := range g.Cups {
		if player.Outs == 3 {
			delete(g.Cups, account)
		}
	}

	// If there is only 1 player left we have a winner.
	if len(g.Cups) == 1 {
		g.Status = StatusOpen
		return 1, nil
	}

	// Figure out who starts the round. The person who was last out should
	// start the round.
	var found bool
	if g.lastOutAcct != "" {
		g.CurrentPlayer = g.lastOutAcct
	}

	// If the person who was last out is no longer in the game, then the
	// player who won the last round starts.
	if !found {
		g.CurrentPlayer = g.lastWinAcct
	}

	// Reset players dice.
	for account, player := range g.Cups {
		player.Dice = make([]int, 5)
		g.Cups[account] = player
	}

	// Reset the claims to start over.
	g.Claims = make([]Claim, 0)

	g.Status = StatusPlaying

	g.Round++

	// Return the number of players for this round.
	return len(g.Cups), nil
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

//==============================================================================

// nextCup will look for the next available player.
func (g *Game) nextCup() error {
	var control int

	for control <= len(g.CupsOrder) {
		control++
		g.currentCup = (g.currentCup + 1) % len(g.CupsOrder)
		if g.CupsOrder[g.currentCup] != "" {
			return nil
		}
	}

	return errors.New("no player available")
}
