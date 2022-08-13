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

	orderIdx int
}

// =============================================================================

// Claim represents a claim of dice by a player.
type claim struct {
	account string
	number  int
	suite   int
}

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
	claims        []claim
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

// =============================================================================

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

// =============================================================================

// AddAccount adds a player to the game. If the account already exists, the
// function will return an error.
func (g *Game) AddAccount(account string) error {
	g.mu.Lock()
	defer g.mu.Unlock()

	if account == "" {
		return errors.New("account provided is empty")
	}

	if g.status != StatusOpen {
		return fmt.Errorf("game status is required to be open: status[%s]", g.status)
	}

	if _, exists := g.cups[account]; exists {
		return fmt.Errorf("account [%s] is already in the game", account)
	}

	g.cups[account] = Cup{
		Account:  account,
		Outs:     0,
		Dice:     make([]int, 5),
		orderIdx: len(g.cupsOrder),
	}

	g.cupsOrder = append(g.cupsOrder, account)

	return nil
}

// OutAccount will apply the specified number of outs to the account.
func (g *Game) OutAccount(account string, outs int) error {
	g.mu.Lock()
	defer g.mu.Unlock()

	if account == "" {
		return errors.New("account provided is empty")
	}

	if outs < 0 || outs > 3 {
		return errors.New("invalid out value")
	}

	if g.status != StatusPlaying {
		return fmt.Errorf("game status is required to be playing: status[%s]", g.status)
	}

	cup, exists := g.cups[account]
	if !exists {
		return fmt.Errorf("account [%s] does not exist in the game", account)
	}

	cup.Outs = outs
	g.cups[account] = cup

	return nil
}

// RollDice will generate 5 new random integers for the players cup.
func (g *Game) RollDice(account string) error {
	g.mu.Lock()
	defer g.mu.Unlock()

	if account == "" {
		return errors.New("account provided is empty")
	}

	if g.status != StatusPlaying {
		return fmt.Errorf("game status is required to be playing: status[%s]", g.status)
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

// Claim accepts a claim from an account, but validates the claim is valid first.
// If the claim is valid, it's added to the list of claims for the game. Then
// the next player is determined and set.
func (g *Game) Claim(account string, number int, suite int) error {
	g.mu.Lock()
	defer g.mu.Unlock()

	if account == "" {
		return errors.New("account provided is empty")
	}

	if g.status != StatusPlaying {
		return fmt.Errorf("game status is required to be playing: status[%s]", g.status)
	}

	// Validate that the account who is making the claim is the account that
	// should be making this claim.
	currentAccount := g.cupsOrder[g.currentCup]
	if currentAccount != account {
		return fmt.Errorf("account [%s] can't make a claim now", account)
	}

	// If this is not the first claim, we need to validate the claim is valid.
	if len(g.claims) > 0 {
		lastClaim := g.claims[len(g.claims)-1]

		if number < lastClaim.number {
			return fmt.Errorf("claim number must be greater or equal to the last claim: number[%d] last[%d]", number, lastClaim.number)
		}

		if number == lastClaim.number && suite <= lastClaim.suite {
			return fmt.Errorf("claim suite must be greater than the last claim suite: suite[%d] last[%d]", suite, lastClaim.suite)
		}
	}

	// Add the claim to the list.
	c := claim{
		account: account,
		number:  number,
		suite:   suite,
	}
	g.claims = append(g.claims, c)

	return nil
}

// NextTurn determines which account make the next move.
func (g *Game) NextTurn() {
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
}

// CallLiar checks the last claim that was made and determines the winner and
// loser of the current round.
func (g *Game) CallLiar(account string) (winningAcct string, losingAcct string, err error) {
	g.mu.Lock()
	defer g.mu.Unlock()

	if account == "" {
		return "", "", errors.New("account provided is empty")
	}

	if g.status != StatusPlaying {
		return "", "", fmt.Errorf("game status is required to be playing: status[%s]", g.status)
	}

	// Validate that the account who is making the claim is the account that
	// should be making this claim.
	currentAccount := g.cupsOrder[g.currentCup]
	if currentAccount != account {
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
	case dice[lastClaim.suite] < lastClaim.number:

		// The account who made the last claim lost.
		cup := g.cups[lastClaim.account]
		cup.Outs++
		g.cups[lastClaim.account] = cup

		g.lastOutAcct = cup.Account
		g.lastWinAcct = account

	default:

		// The account who called liar lost.
		cup := g.cups[account]
		cup.Outs++
		g.cups[account] = cup

		g.lastOutAcct = account
		g.lastWinAcct = lastClaim.account
	}

	return g.lastWinAcct, g.lastOutAcct, nil
}

// NextRound updates the game state for players who are out and determining
// which player goes next.
func (g *Game) NextRound() (int, error) {
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
