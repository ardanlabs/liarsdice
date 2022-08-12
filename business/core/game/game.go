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

/*
	We could choose a random person to start in the Start API.
*/

// Represents the different game statues.
const (
	StatusRoundOver = "roundover"
	StatusPlaying   = "playing"
	StatusOpen      = "open"
	StatusOver      = "over"
)

// minNumberPlayers represents the minimum number of players required
// to play a game.
const minNumberPlayers = 2

// =============================================================================

// Banker represents the ability to manage money for the game. Deposits and
// Withdrawls happen outside of game play.
type Banker interface {
	Balance(ctx context.Context, account string) (*big.Int, error)
	Reconcile(ctx context.Context, winningAccount string, losingAccounts []string, ante uint, gameFee uint) error
}

// =============================================================================

// StatusInfo represents a copy of the game status.
type StatusInfo struct {
	Status        string
	Round         int
	CurrentPlayer string
	CupsOrder     []string
}

// =============================================================================

// Cup represents an individual cup being held by a player.
type Cup struct {
	Account string
	Outs    int
	Dice    []int
	order   int
}

// Claim represents a claim of dice by a player.
type Claim struct {
	Account string
	Number  int
	Suite   int
}

// =============================================================================

// Game represents a single game that is being played.
type Game struct {
	id            string
	banker        Banker
	status        string
	lastOutAcct   string
	lastWinAcct   string
	currentPlayer string
	currentCup    int
	round         int
	cups          map[string]Cup
	cupsOrder     []string
	claims        []Claim
	mu            sync.RWMutex
}

// New creates a new game.
func New(banker Banker) *Game {
	return &Game{
		id:     uuid.NewString(),
		banker: banker,
		status: StatusOpen,
		round:  1,
		cups:   make(map[string]Cup),
	}
}

// Status returns a copy of the game status.
func (g *Game) Status() StatusInfo {
	g.mu.RLock()
	defer g.mu.RUnlock()

	cupsOrder := make([]string, len(g.cupsOrder))
	copy(cupsOrder, g.cupsOrder)

	return StatusInfo{
		Status:        g.status,
		Round:         g.round,
		CurrentPlayer: g.currentPlayer,
		CupsOrder:     cupsOrder,
	}
}

// AddAccount adds a player to the game. If the account already exists, the
// function will return an error.
func (g *Game) AddAccount(account string) error {
	g.mu.Lock()
	defer g.mu.Unlock()

	if g.status != StatusOpen {
		return errors.New("can't add a new account to the game")
	}

	if account == "" {
		return errors.New("invalid account information")
	}

	if _, exists := g.cups[account]; exists {
		return fmt.Errorf("account [%s] is already in the game", account)
	}

	g.cups[account] = Cup{
		Account: account,
		Dice:    make([]int, 5),
		order:   len(g.cupsOrder),
	}

	g.cupsOrder = append(g.cupsOrder, account)

	return nil
}

// OutAccount will apply the specified number of outs to the account.
func (g *Game) OutAccount(account string, outs int) error {
	g.mu.Lock()
	defer g.mu.Unlock()

	if g.status != StatusPlaying {
		return errors.New("game is not being played")
	}

	if account == "" {
		return errors.New("invalid account information")
	}

	if outs <= 0 || outs > 3 {
		return errors.New("invalid out amount")
	}

	acc, exists := g.cups[account]
	if !exists {
		return fmt.Errorf("account [%s] does not exist in the game", account)
	}

	acc.Outs = outs
	g.cups[account] = acc

	return nil
}

// RollDice will generate 5 new random integers for the players cup.
func (g *Game) RollDice(account string) error {
	g.mu.Lock()
	defer g.mu.Unlock()

	if g.status != StatusPlaying {
		return errors.New("game is not being played")
	}

	if account == "" {
		return errors.New("invalid account information")
	}

	cup, exists := g.cups[account]
	if !exists {
		return fmt.Errorf("account [%s] does not exist in the game", account)
	}

	for i := range cup.Dice {
		cup.Dice[i] = rand.Intn(6) + 1
	}

	return nil
}

// Claim accpts a claim from an account, but validates the claim is valid first.
// If the claim is valid, it is added to the list of claims for the game. Then
// the next player is determined and set.
func (g *Game) Claim(claim Claim) error {
	g.mu.Lock()
	defer g.mu.Unlock()

	if g.status != StatusPlaying {
		return errors.New("game is not being played")
	}

	if claim.Account == "" {
		return errors.New("invalid account information")
	}

	// Get the account who is supposed to make the next move.
	nextMove := g.cupsOrder[g.currentCup]

	// Validate the specified account matches.
	if nextMove != claim.Account {
		return fmt.Errorf("account [%s] can't make a claim now", claim.Account)
	}

	// If this is not the first claim, we need to validate the claim is valid.
	if len(g.claims) > 0 {
		lastClaim := g.claims[len(g.claims)-1]

		if claim.Number < lastClaim.Number {
			return errors.New("claim number must be greater or equal to the last claim")
		}

		if claim.Number == lastClaim.Number && claim.Suite <= lastClaim.Suite {
			return errors.New("claim suite must be greater than the last claim suite")
		}
	}

	// Add the claim to the list.
	g.claims = append(g.claims, claim)

	// Figure out who goes next.
	l := len(g.cupsOrder)
	for i := 0; i < l; i++ {

		// Circle back to the beginning of the slice if we reached the end.
		g.currentCup++
		if g.currentCup == l {
			g.currentCup = 0
		}

		// If the account information for this index is not empty, this
		// player is still in the game and the next player to make a claim.
		if g.cupsOrder[g.currentCup] != "" {
			break
		}
	}

	return nil
}

// CallLiar checks the last claim that was made and determines the winner and
// loser of the current round.
func (g *Game) CallLiar(account string) (winningAcct string, losingAcct string, err error) {
	g.mu.Lock()
	defer g.mu.Unlock()

	if g.status != StatusPlaying {
		return "", "", errors.New("game is not being played")
	}

	if account == "" {
		return "", "", errors.New("invalid account information")
	}

	// Get the account who is supposed to make the next move.
	nextMove := g.cupsOrder[g.currentCup]

	// Validate the specified account matches.
	if nextMove != account {
		return "", "", fmt.Errorf("account [%s] can't call lair now", account)
	}

	// This call ends the round, not allowing any more claims to be made.
	g.status = StatusRoundOver

	// Hold the sum of all the dice values.
	dice := make([]int, 7)
	for _, player := range g.cups {
		for _, suite := range player.Dice {
			dice[suite]++
		}
	}

	// Capture the last claim that was made.
	lastClaim := g.claims[len(g.claims)-1]

	// Identify the winner and the loser.
	switch {
	case dice[lastClaim.Suite] < lastClaim.Number:

		// The account who made the last claim lost.
		cup := g.cups[lastClaim.Account]
		cup.Outs++
		g.cups[lastClaim.Account] = cup

		g.lastOutAcct = cup.Account
		g.lastWinAcct = account

	default:

		// The account who called liar lost.
		cup := g.cups[account]
		cup.Outs++
		g.cups[account] = cup

		g.lastOutAcct = account
		g.lastWinAcct = lastClaim.Account
	}

	return g.lastWinAcct, g.lastOutAcct, nil
}

// NewRound checks for player's out count, reset players dice and game claims.
func (g *Game) NewRound() (int, error) {
	g.mu.Lock()
	defer g.mu.Unlock()

	if g.status != StatusRoundOver {
		return 0, errors.New("current round is not over")
	}

	// If an account has three outs, remove their account from game play.
	var leftToPlay int
	for _, cup := range g.cups {
		if cup.Outs == 3 {
			g.cupsOrder[cup.order] = ""
			continue
		}

		leftToPlay++
	}

	// If there is only 1 player left we have a winner.
	if leftToPlay == 1 {
		g.status = StatusOver
		return 1, nil
	}

	// Figure out who starts the next round.
	// The person who was last out should start the round unless they are out.
	if g.cups[g.lastOutAcct].Outs != 3 {
		g.currentPlayer = g.lastOutAcct
	} else {
		g.currentPlayer = g.lastWinAcct
	}

	// Roll new dice for each active player.
	for _, cup := range g.cups {
		if cup.Outs != 3 {
			for i := range cup.Dice {
				cup.Dice[i] = rand.Intn(6) + 1
			}
		}
	}

	// Reset the game state.
	g.claims = []Claim{}
	g.status = StatusPlaying
	g.round++

	// Return the number of players for this round.
	return leftToPlay, nil
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
