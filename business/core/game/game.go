// Package game represents all the game play for liar's dice.
package game

import (
	"context"
	"errors"
	"fmt"
	"math/big"
	"math/rand"
	"sync"
	"time"

	"github.com/google/uuid"
)

/*
	We could choose a random person to start in the Start API.
*/

// Represents the different game status.
const (
	StatusPlaying   = "playing"
	StatusRoundOver = "roundover"
	StatusGameOver  = "gameover"
)

// minNumberPlayers represents the minimum number of players required
// to play a game.
const minNumberPlayers = 2

// =============================================================================

// Banker represents the ability to manage money for the game. Deposits and
// Withdrawls happen outside of game play.
type Banker interface {
	Balance(ctx context.Context, account string) (*big.Int, error)
	Reconcile(ctx context.Context, winningAccount string, losingAccounts []string, ante int64, gameFee int64) error
}

// =============================================================================

// Status represents a copy of the game status.
type Status struct {
	Status        string
	LastOutAcct   string
	LastWinAcct   string
	CurrentPlayer string
	CurrentCup    int
	Round         int
	Cups          map[string]Cup
	CupsOrder     []string
	Claims        []Claim
}

// Cup represents an individual cup being held by a player.
type Cup struct {
	OrderIdx int
	Account  string
	Outs     int
	Dice     []int
}

// Claim represents a claim of dice by a player.
type Claim struct {
	Account string
	Number  int
	Suite   int
}

// Game represents a single game that is being played.
type Game struct {
	id            string
	banker        Banker
	mu            sync.RWMutex
	status        string
	lastOutAcct   string
	lastWinAcct   string
	currentPlayer string
	currentCup    int
	round         int
	ante          int64
	cups          map[string]Cup
	cupsOrder     []string
	claims        []Claim
}

// New creates a new game.
func New(banker Banker, ante int64) *Game {
	rand.Seed(time.Now().UTC().UnixNano())

	return &Game{
		id:     uuid.NewString(),
		banker: banker,
		status: StatusGameOver,
		round:  1,
		ante:   ante,
		cups:   make(map[string]Cup),
	}
}

// AddAccount adds a player to the game. If the account already exists, the
// function will return an error.
func (g *Game) AddAccount(ctx context.Context, account string) error {
	g.mu.Lock()
	defer g.mu.Unlock()

	if account == "" {
		return errors.New("account provided is empty")
	}

	if g.status != StatusGameOver {
		return fmt.Errorf("game status is required to be over: status[%s]", g.status)
	}

	if _, exists := g.cups[account]; exists {
		return fmt.Errorf("account [%s] is already in the game", account)
	}

	balance, err := g.banker.Balance(ctx, account)
	if err != nil {
		return fmt.Errorf("unable to retrieve account[%s] balance", account)
	}

	ante := big.NewInt(int64(g.ante))

	// If comparison is negative, the player has no balance.
	if balance.Cmp(ante) < 0 {
		return fmt.Errorf("account [%s] does not have enough balance to play", account)
	}

	g.cups[account] = Cup{
		OrderIdx: len(g.cupsOrder),
		Account:  account,
		Outs:     0,
		Dice:     make([]int, 5),
	}

	g.cupsOrder = append(g.cupsOrder, account)

	return nil
}

// StartPlay changes the status to Playing to allow the game to begin.
func (g *Game) StartPlay() error {
	g.mu.Lock()
	defer g.mu.Unlock()

	if g.status != StatusGameOver {
		return fmt.Errorf("game status is required to be over: status[%s]", g.status)
	}

	if len(g.cups) < minNumberPlayers {
		return errors.New("not enough players to start the game")
	}

	g.status = StatusPlaying

	return nil
}

// ApplyOut will apply the specified number of outs to the account.
// If an account is out, it will check the number of active accounts, and end
// the round if there is only 1 left.
func (g *Game) ApplyOut(account string, outs int) error {
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

	// After 3 outs, an account is out of the game.
	// We need to check if there is only 1 account left, end the round.
	if outs == 3 {
		g.cupsOrder[cup.OrderIdx] = ""

		// Look for active players.
		var activePlayers int

		for _, v := range g.cupsOrder {
			if v != "" {
				activePlayers++
			}
		}

		if activePlayers == 1 {
			g.status = StatusRoundOver
		}
	}

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

		if number < lastClaim.Number {
			return fmt.Errorf("claim number must be greater or equal to the last claim number: number[%d] last[%d]", number, lastClaim.Number)
		}

		if number == lastClaim.Number && suite <= lastClaim.Suite {
			return fmt.Errorf("claim suite must be greater than the last claim suite: suite[%d] last[%d]", suite, lastClaim.Suite)
		}
	}

	// Add the claim to the list.
	c := Claim{
		Account: account,
		Number:  number,
		Suite:   suite,
	}
	g.claims = append(g.claims, c)

	g.NextTurn(account)

	return nil
}

// NextTurn determines which account make the next move.
func (g *Game) NextTurn(account string) {
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

	if len(g.claims) == 0 {
		return "", "", errors.New("no claims have been made yet")
	}

	// Validate that the account who is making the claim is the account that
	// should be making this claim.
	currentAccount := g.cupsOrder[g.currentCup]
	if currentAccount != account {
		return "", "", fmt.Errorf("account [%s] can't call liar now", account)
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

// NextRound updates the game state for players who are out and determining
// which player goes next. The function returns the number of players left
// in the game.
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
			g.cupsOrder[cup.OrderIdx] = ""
			continue
		}

		leftToPlay++
	}

	// If there is only 1 player left we have a winner.
	if leftToPlay == 1 {
		g.status = StatusGameOver
		return 1, nil
	}

	// Figure out who starts the next round.
	// The person who was last out should start the round unless they are out.
	if g.cups[g.lastOutAcct].Outs != 3 {
		g.currentPlayer = g.lastOutAcct
	} else {
		g.currentPlayer = g.lastWinAcct
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
func (g *Game) PlayerBalance(ctx context.Context, account string) (*big.Int, error) {
	if account == "" {
		return nil, errors.New("account provided is empty")
	}

	return g.banker.Balance(ctx, account)
}

// Reconcile calculates the game pot and make the transfer to the winner.
func (g *Game) Reconcile(ctx context.Context) error {
	var winner string
	var losers []string

	// Find the winner.
	// cupsOrders holds the remaining (winner) account, the others are "".
	for _, cup := range g.cupsOrder {
		if _, found := g.cups[cup]; found {
			winner = cup
		}
	}

	// Find the losers.
	for _, cup := range g.cups {
		if winner != cup.Account {
			losers = append(losers, cup.Account)
		}
	}

	return g.banker.Reconcile(ctx, winner, losers, g.ante, 10)
}

// Info returns a copy of the game status.
func (g *Game) Info() Status {
	g.mu.RLock()
	defer g.mu.RUnlock()

	cups := make(map[string]Cup)
	for k, v := range g.cups {
		cups[k] = v
	}

	cupsOrder := make([]string, len(g.cupsOrder))
	copy(cupsOrder, g.cupsOrder)

	claims := make([]Claim, len(g.claims))
	copy(claims, g.claims)

	return Status{
		Status:        g.status,
		LastOutAcct:   g.lastOutAcct,
		LastWinAcct:   g.lastWinAcct,
		CurrentPlayer: g.currentPlayer,
		CurrentCup:    g.currentCup,
		Round:         g.round,
		Cups:          cups,
		CupsOrder:     cupsOrder,
		Claims:        claims,
	}
}
