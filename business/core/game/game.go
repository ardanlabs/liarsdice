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

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/google/uuid"
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

// =============================================================================

// Converter represents the ability to convert USD to GWei for game play.
type Converter interface {
	USD2GWei(amountUSD *big.Float) *big.Float
}

// Banker represents the ability to manage money for the game. Deposits and
// Withdrawls happen outside of game play.
type Banker interface {
	AccountBalance(ctx context.Context, account string) (GWei *big.Float, err error)
	Reconcile(ctx context.Context, winningAccount string, losingAccounts []string, anteGWei *big.Float, gameFeeGWei *big.Float) (*types.Transaction, *types.Receipt, error)
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
	converter     Converter
	banker        Banker
	mu            sync.RWMutex
	owner         string
	status        string
	lastOutAcct   string
	lastWinAcct   string
	currentPlayer string
	currentCup    int
	round         int
	anteUSD       float64
	cups          map[string]Cup
	cupsOrder     []string
	claims        []Claim
}

// New creates a new game.
func New(ctx context.Context, converter Converter, banker Banker, owner string, anteUSD float64) (*Game, error) {
	balance, err := banker.AccountBalance(ctx, owner)
	if err != nil {
		return nil, fmt.Errorf("unable to retrieve account[%s] balance", owner)
	}

	// If comparison is negative, the player has no balance.
	anteGwei := converter.USD2GWei(big.NewFloat(anteUSD))
	if balance.Cmp(anteGwei) < 0 {
		return nil, fmt.Errorf("account [%s] does not have enough balance to play", owner)
	}

	rand.Seed(time.Now().UTC().UnixNano())

	g := Game{
		id:        uuid.NewString(),
		converter: converter,
		banker:    banker,
		owner:     owner,
		status:    StatusNewGame,
		round:     1,
		anteUSD:   anteUSD,
		cups:      make(map[string]Cup),
	}

	if err := g.AddAccount(ctx, owner); err != nil {
		return nil, errors.New("unable to add owner to the game")
	}

	return &g, nil
}

// AddAccount adds a player to the game. If the account already exists, the
// function will return an error.
func (g *Game) AddAccount(ctx context.Context, account string) error {
	g.mu.Lock()
	defer g.mu.Unlock()

	if account == "" {
		return errors.New("account provided is empty")
	}

	if g.status != StatusNewGame {
		return fmt.Errorf("game status is required to be over: status[%s]", g.status)
	}

	if _, exists := g.cups[account]; exists {
		return fmt.Errorf("account [%s] is already in the game", account)
	}

	balance, err := g.banker.AccountBalance(ctx, account)
	if err != nil {
		return fmt.Errorf("unable to retrieve account[%s] balance", account)
	}

	ante := big.NewFloat(float64(g.anteUSD))

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

// StartGame changes the status to Playing to allow the game to begin.
func (g *Game) StartGame(owner string) error {
	g.mu.Lock()
	defer g.mu.Unlock()

	if g.status != StatusNewGame {
		return fmt.Errorf("game status is required to be over: status[%s]", g.status)
	}

	if g.owner != owner {
		return fmt.Errorf("only the game owner can start the game: owner[%s] caller[%s]", g.owner, owner)
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

// RollDice will generate 5 new random integers for the players cup. The caller
// can specific the dice if they choose.
func (g *Game) RollDice(account string, manualRole ...int) error {
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

	if manualRole == nil || len(manualRole) < 5 {
		for i := range cup.Dice {
			cup.Dice[i] = rand.Intn(6) + 1
		}
	} else {
		for i := range cup.Dice {
			cup.Dice[i] = manualRole[i]
		}
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

	// Move the turn to the next player.
	g.nextTurn(account)

	return nil
}

// NextTurn determines which account make the next move.
func (g *Game) NextTurn(account string) error {
	g.mu.Lock()
	defer g.mu.Unlock()

	if g.status != StatusPlaying {
		return fmt.Errorf("game status is required to be playing: status[%s]", g.status)
	}

	g.nextTurn(account)

	return nil
}

// nextTurn determines which account make the next move.
func (g *Game) nextTurn(account string) {
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

// Reconcile calculates the game pot and make the transfer to the winner.
func (g *Game) Reconcile(ctx context.Context, winningAccount string) (*types.Transaction, *types.Receipt, error) {
	g.mu.RLock()
	defer g.mu.RUnlock()

	if g.status != StatusGameOver {
		return nil, nil, fmt.Errorf("game status is required to be gameover: status[%s]", g.status)
	}

	// Find the winner.
	// cupsOrders holds the remaining (winner) account, the others are "".
	var winner string
	for _, cup := range g.cupsOrder {
		if _, found := g.cups[cup]; found {
			winner = cup
		}
	}

	if winner != winningAccount {
		return nil, nil, fmt.Errorf("only winning account can reconcile the game, winner[%s]", winner)
	}

	// Find the losers.
	var losers []string
	for _, cup := range g.cups {
		if winner != cup.Account {
			losers = append(losers, cup.Account)
		}
	}

	// Convert the anti and game fee from USD to Wei.
	antiGWei := g.converter.USD2GWei(big.NewFloat(g.anteUSD))
	gameFeeGWei := g.converter.USD2GWei(big.NewFloat(g.anteUSD))

	// Perform the reconcile against the bank.
	tx, receipt, err := g.banker.Reconcile(ctx, winner, losers, antiGWei, gameFeeGWei)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to reconcile the game: %w", err)
	}

	g.status = StatusReconciled

	return tx, receipt, nil
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
