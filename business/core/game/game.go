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

	"github.com/ardanlabs/ethereum/currency"
	"github.com/ardanlabs/liarsdice/foundation/logger"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/google/uuid"
)

var ErrNotFound = errors.New("game not found")

// Storer interface declares the behaviour this package needs to persist and
// retrieve data.
type Storer interface {
	Create(ctx context.Context, g *Game) error
	InsertRound(ctx context.Context, state State) error
	QueryStateByID(ctx context.Context, gameID uuid.UUID, round int) (State, error)
}

// Banker represents the ability to manage money for the game. Deposits and
// Withdrawls happen outside of game play.
type Banker interface {
	AccountBalance(ctx context.Context, player common.Address) (GWei *big.Float, err error)
	Reconcile(ctx context.Context, winningPlayer common.Address, losingPlayers []common.Address, anteGWei *big.Float, gameFeeGWei *big.Float) (*types.Transaction, *types.Receipt, error)
}

// Game represents a single game that is being played.
type Game struct {
	log             *logger.Logger
	converter       *currency.Converter
	storer          Storer
	banker          Banker
	mu              sync.RWMutex
	id              uuid.UUID              // Unique game id.
	dateCreated     time.Time              // The time the game was created. Used to help with caching.
	round           int                    // Current round of the game.
	status          string                 // Current status of the game.
	anteUSD         float64                // The ante for joining this game.
	playerLastOut   common.Address         // The player who lost the last round.
	playerLastWin   common.Address         // The player who won the last round.
	playerTurn      int                    // The index of the player who's turn it is.
	players         []common.Address       // Game players in the order they were added.
	existingPlayers []common.Address       // The set of players still in the game.
	cups            map[common.Address]Cup // Game players with indexes cup access.
	bets            []Bet                  // History of bets for the current round.
	balancesGWei    []Balance              // The balances of the players when added.
}

// New creates a new game.
func New(ctx context.Context, log *logger.Logger, converter *currency.Converter, storer Storer, banker Banker, player common.Address, anteUSD float64) (*Game, error) {
	balance, err := banker.AccountBalance(ctx, player)
	if err != nil {
		return nil, fmt.Errorf("unable to retrieve account[%s] balance", player)
	}

	// If comparison is negative, the player has no balance.
	anteGWei := converter.USD2GWei(big.NewFloat(anteUSD))
	if balance.Cmp(anteGWei) < 0 {
		return nil, fmt.Errorf("account [%s] does not have enough balance to play, balance[%v]", player, balance)
	}

	g := Game{
		log:         log,
		converter:   converter,
		storer:      storer,
		banker:      banker,
		id:          uuid.New(),
		status:      StatusNewGame,
		round:       0,
		anteUSD:     anteUSD,
		cups:        make(map[common.Address]Cup),
		dateCreated: time.Now().UTC(),
	}

	if err := g.AddAccount(ctx, player); err != nil {
		return nil, errors.New("unable to add owner to the game")
	}

	if err := g.storer.Create(ctx, &g); err != nil {
		return nil, errors.New("unable to add the game to the db")
	}

	Tables.add(&g)

	g.log.Info(ctx, "game.new", "id", g.id, "player", player, "anteUSD", g.anteUSD, "anteGWei", anteGWei)

	return &g, nil
}

// ID returns the game id.
func (g *Game) ID() uuid.UUID {
	return g.id
}

// DateCreated returns the date/time the game was created.
func (g *Game) DateCreated() time.Time {
	return g.dateCreated
}

// Status returns the current status of the game.
func (g *Game) Status() string {
	g.mu.Lock()
	defer g.mu.Unlock()

	return g.status
}

// AddAccount adds a player to the game. If the account already exists, the
// function will return an error.
func (g *Game) AddAccount(ctx context.Context, player common.Address) error {
	g.mu.Lock()
	defer g.mu.Unlock()

	var empty common.Address
	if player == empty {
		return errors.New("account id provided is empty")
	}

	if _, exists := g.cups[player]; exists {
		return fmt.Errorf("account id [%s] is already in the game", player)
	}

	if g.status != StatusNewGame {
		return fmt.Errorf("game status is required to be over: status[%s]", g.status)
	}

	if _, exists := g.cups[player]; exists {
		return fmt.Errorf("account id [%s] is already in the game", player)
	}

	balanceGwei, err := g.banker.AccountBalance(ctx, player)
	if err != nil {
		return fmt.Errorf("unable to retrieve account id [%s] balance", player)
	}

	anteGWei := g.converter.USD2GWei(big.NewFloat(g.anteUSD))

	// If comparison is negative, the player has no balance.
	if balanceGwei.Cmp(anteGWei) < 0 {
		return fmt.Errorf("player [%s] does not have enough balance to play", player)
	}

	g.cups[player] = Cup{
		OrderIdx: len(g.players),
		Player:   player,
		Outs:     0,
		Dice:     make([]int, 5),
	}

	g.players = append(g.players, player)
	g.existingPlayers = append(g.existingPlayers, player)

	g.balancesGWei = append(g.balancesGWei, Balance{
		Player: player,
		Amount: balanceGwei,
	})

	return nil
}

// StartGame changes the status to Playing to allow the game to begin.
func (g *Game) StartGame(ctx context.Context) error {
	g.mu.Lock()
	defer g.mu.Unlock()

	if g.status != StatusNewGame {
		return fmt.Errorf("game status is required to be over: status[%s]", g.status)
	}

	if len(g.cups) < minNumberPlayers {
		return errors.New("not enough players to start the game")
	}

	g.playerTurn = rand.Intn(len(g.cups))
	g.status = StatusPlaying
	g.round = 1

	return nil
}

// ApplyOut will apply the specified number of outs to the account.
// If an account is out, it will check the number of active accounts, and end
// the round if there is only 1 left.
func (g *Game) ApplyOut(ctx context.Context, player common.Address, outs int) error {
	g.mu.Lock()
	defer g.mu.Unlock()

	var empty common.Address
	if player == empty {
		return errors.New("account provided is empty")
	}

	cup, exists := g.cups[player]
	if !exists {
		return fmt.Errorf("player [%s] does not exist in the game", player)
	}

	if outs < 0 || outs > 3 {
		return errors.New("invalid out value")
	}

	if g.status != StatusPlaying {
		return fmt.Errorf("game status is required to be playing: status[%s]", g.status)
	}

	cup.Outs = outs
	g.cups[player] = cup

	// After 3 outs, an account is out of the game.
	// We need to check if there is only 1 account left, end the round.
	if outs == 3 {
		var empty common.Address
		g.existingPlayers[cup.OrderIdx] = empty

		// Look for active players.
		var activePlayers int

		for _, v := range g.existingPlayers {
			if v != empty {
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
func (g *Game) RollDice(ctx context.Context, player common.Address, manualRole ...int) error {
	g.mu.Lock()
	defer g.mu.Unlock()

	var empty common.Address
	if player == empty {
		return errors.New("player provided is empty")
	}

	if g.status != StatusPlaying {
		return fmt.Errorf("game status is required to be playing: status[%s]", g.status)
	}

	return g.rollDice(ctx, player, manualRole...)
}

// rollDice will generate 5 new random integers for the players cup. The caller
// can specific the dice if they choose.
func (g *Game) rollDice(ctx context.Context, player common.Address, manualRole ...int) error {
	cup, exists := g.cups[player]
	if !exists {
		return fmt.Errorf("player [%s] does not exist in the game", player)
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

	g.log.Info(ctx, "game.rolldice", "id", g.id, "player", player, "dice", cup.Dice)

	return nil
}

// Bet accepts a bet from an account, but validates the bet is valid first.
// If the bet is valid, it's added to the list of bets for the game. Then
// the next player is determined and set.
func (g *Game) Bet(ctx context.Context, player common.Address, number int, suit int) error {
	g.mu.Lock()
	defer g.mu.Unlock()

	var empty common.Address
	if player == empty {
		return errors.New("account id provided is empty")
	}

	cup, exists := g.cups[player]
	if !exists {
		return fmt.Errorf("player [%s] does not exist in the game", player)
	}

	if g.status != StatusPlaying {
		return fmt.Errorf("game status is required to be playing: status[%s]", g.status)
	}

	// Validate that the account who is making the bet is the account that
	// should be making this bet.
	currentPlayer := g.existingPlayers[g.playerTurn]
	if currentPlayer != player {
		return fmt.Errorf("player [%s] can't make a bet now", player)
	}

	// If this is not the first bet, we need to validate the bet is valid.
	if len(g.bets) > 0 {
		lastBet := g.bets[len(g.bets)-1]

		if number < lastBet.Number {
			return fmt.Errorf("bet number must be greater or equal to the last bet number: number[%d] last[%d]", number, lastBet.Number)
		}

		if number == lastBet.Number && suit <= lastBet.Suit {
			return fmt.Errorf("bet suit must be greater than the last bet suit: suit[%d] last[%d]", suit, lastBet.Suit)
		}
	}

	// Add the bet to the list.
	bet := Bet{
		Player: player,
		Number: number,
		Suit:   suit,
	}
	g.bets = append(g.bets, bet)

	// Add the last bet to the cup.
	g.cups[player] = cup

	// Move the turn to the next player.
	g.nextTurn()

	return nil
}

// NextTurn determines which account makes the next move.
func (g *Game) NextTurn(ctx context.Context) error {
	g.mu.Lock()
	defer g.mu.Unlock()

	if g.status != StatusPlaying {
		return fmt.Errorf("game status is required to be playing: status[%s]", g.status)
	}

	g.nextTurn()

	return nil
}

// nextTurn determines which account makes the next move.
func (g *Game) nextTurn() {
	l := len(g.existingPlayers)

	for i := 0; i < l; i++ {

		// Circle back to the beginning of the slice if we reached the end.
		g.playerTurn++
		if g.playerTurn == l {
			g.playerTurn = 0
		}

		// If the account information for this index is not empty, this
		// player is still in the game and the next player to make a bet.
		var empty common.Address
		if g.existingPlayers[g.playerTurn] != empty {
			break
		}
	}
}

// CallLiar checks the last bet that was made and determines the winner and
// loser of the current round.
func (g *Game) CallLiar(ctx context.Context, player common.Address) (winningPlayer common.Address, losingPlayer common.Address, err error) {
	g.mu.Lock()
	defer g.mu.Unlock()

	var empty common.Address

	if player == empty {
		return empty, empty, errors.New("account provided is empty")
	}

	if _, exists := g.cups[player]; !exists {
		return empty, empty, fmt.Errorf("player [%s] does not exist in the game", player)
	}

	if g.status != StatusPlaying {
		return empty, empty, fmt.Errorf("game status is required to be playing: status[%s]", g.status)
	}

	if len(g.bets) == 0 {
		return empty, empty, fmt.Errorf("there are no bets to validate")
	}

	// Validate that the account who is making the bet is the account that
	// should be making this bet.
	currentPlayer := g.existingPlayers[g.playerTurn]
	if currentPlayer != player {
		return empty, empty, fmt.Errorf("player [%s] can't call liar now", currentPlayer)
	}

	// This call ends the round, not allowing any more bets to be made.
	g.status = StatusRoundOver

	// Hold the sum of all the dice values.
	dice := make([]int, 7)
	for _, player := range g.cups {
		for _, suit := range player.Dice {
			dice[suit]++
		}
	}

	// Capture the last bet that was made.
	lastBet := g.bets[len(g.bets)-1]

	// Identify the winner and the loser.
	switch {
	case dice[lastBet.Suit] < lastBet.Number:

		// The account who made the last bet lost.
		cup := g.cups[lastBet.Player]
		cup.Outs++
		g.cups[lastBet.Player] = cup

		g.playerLastOut = cup.Player
		g.playerLastWin = player

	default:

		// The account who called liar lost.
		cup := g.cups[player]
		cup.Outs++
		g.cups[player] = cup

		g.playerLastOut = player
		g.playerLastWin = lastBet.Player
	}

	// Not sure I want to return an error if I can't save this round
	// to the database. It just means we can't recover this game properly.
	// Since nothing happens to the bank, no one is losing money nor is
	// money held hostage.
	if err := g.storer.InsertRound(ctx, g.state()); err != nil {
		g.log.Error(ctx, "liar.store.insertRound", "id", g.id, "ERROR", err)
	}

	return g.playerLastWin, g.playerLastOut, nil
}

// NextRound updates the game state for players who are out and determining
// which player goes next. The function returns the number of players left
// in the game.
func (g *Game) NextRound(ctx context.Context) (int, error) {
	g.mu.Lock()
	defer g.mu.Unlock()

	if g.status != StatusRoundOver {
		return 0, errors.New("current round is not over")
	}

	// If an account has three outs, remove their account from game play.
	// Reset the last bet value and dice.
	var leftToPlay int
	for player, cup := range g.cups {
		cup.Dice = make([]int, 5)
		g.cups[player] = cup

		for i := range cup.Dice {
			cup.Dice[i] = 0
		}

		if cup.Outs == 3 {
			var empty common.Address
			g.existingPlayers[cup.OrderIdx] = empty
			continue
		}

		g.rollDice(ctx, player)

		leftToPlay++
	}

	// If there is only 1 player left we have a winner.
	// Reset the bets, the dice, and status.
	if leftToPlay == 1 {
		g.bets = []Bet{}
		g.status = StatusGameOver
		g.rollDice(ctx, g.playerLastWin, 0, 0, 0, 0, 0)
		return 1, nil
	}

	// Figure out who starts the next round.
	// The person who was last out should start the round unless they are out.
	if g.cups[g.playerLastOut].Outs != 3 {
		g.playerTurn = g.cups[g.playerLastOut].OrderIdx
	} else {
		g.playerTurn = g.cups[g.playerLastWin].OrderIdx
	}

	// Reset the game state.
	g.bets = []Bet{}
	g.status = StatusPlaying
	g.round++

	// Return the number of players for this round.
	return leftToPlay, nil
}

// Reconcile calculates the game pot and make the transfer to the winner.
func (g *Game) Reconcile(ctx context.Context) (*types.Transaction, *types.Receipt, error) {
	g.mu.Lock()
	defer g.mu.Unlock()

	if g.status != StatusGameOver {
		return nil, nil, fmt.Errorf("game status is required to be gameover: status[%s]", g.status)
	}

	// Find the losers.
	var losingPlayers []common.Address
	for _, cup := range g.cups {
		if g.playerLastWin != cup.Player {
			losingPlayers = append(losingPlayers, cup.Player)
		}
	}

	// Convert the anti and game fee from USD to Wei.
	antiGWei := g.converter.USD2GWei(big.NewFloat(g.anteUSD))
	gameFeeGWei := g.converter.USD2GWei(big.NewFloat(g.anteUSD))

	// Log the winner and losers.
	g.log.Info(ctx, "game.reconcole", "id", g.id, "winner", g.playerLastWin)
	for _, player := range losingPlayers {
		g.log.Info(ctx, "game.reconcole", "id", g.id, "loser", player)
	}

	// Perform the reconcile against the bank.
	tx, receipt, err := g.banker.Reconcile(ctx, g.playerLastWin, losingPlayers, antiGWei, gameFeeGWei)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to reconcile the game: %w", err)
	}

	g.log.Info(ctx, "game.reconcole.contract", "id", g.id, "tx", g.converter.CalculateTransactionDetails(tx), "receipt", g.converter.CalculateReceiptDetails(receipt, tx.GasPrice()))

	g.status = StatusReconciled
	g.round++

	// Update the player balances.
	g.log.Info(ctx, "game.reconcole.fees", "id", g.id, "anteUSD", g.anteUSD, "antiGWei", antiGWei, "gameFeeGWei", gameFeeGWei)
	for i, player := range g.players {
		balanceGwei, err := g.banker.AccountBalance(ctx, player)
		if err != nil {
			g.log.Info(ctx, "game.reconcole.updatebalance", "ERROR", err)
			continue
		}

		oldBalanceGWei := g.balancesGWei[i]

		g.balancesGWei[i] = Balance{
			Player: player,
			Amount: balanceGwei,
		}

		g.log.Info(ctx, "game.reconcole.updatebalance", "id", g.id, "player", player, "oldBlanceGWei", oldBalanceGWei, "balanceGWei", balanceGwei)
	}

	if err := g.storer.InsertRound(ctx, g.state()); err != nil {
		g.log.Error(ctx, "reconcile.store.insertRound", "id", g.id, "ERROR", err)
	}

	return tx, receipt, nil
}

// State returns a copy of the game state.
func (g *Game) State() State {
	g.mu.RLock()
	defer g.mu.RUnlock()

	return g.state()
}

func (g *Game) state() State {
	cups := make(map[common.Address]Cup)
	for k, v := range g.cups {
		cups[k] = v
	}

	existingPlayers := make([]common.Address, len(g.players))
	copy(existingPlayers, g.players)

	bets := make([]Bet, len(g.bets))
	copy(bets, g.bets)

	balances := make([]BalanceFmt, len(g.balancesGWei))
	for i, balance := range g.balancesGWei {
		balances[i] = BalanceFmt{
			Player: balance.Player,
			Amount: g.converter.GWei2USD(balance.Amount),
		}
	}

	var playerTurn common.Address
	if len(g.existingPlayers) > 0 {
		playerTurn = g.existingPlayers[g.playerTurn]
	}

	return State{
		GameID:          g.id,
		GameName:        g.id.String(),
		DateCreated:     g.dateCreated,
		Round:           g.round,
		Status:          g.status,
		PlayerLastOut:   g.playerLastOut,
		PlayerLastWin:   g.playerLastWin,
		PlayerTurn:      playerTurn,
		ExistingPlayers: existingPlayers,
		Cups:            cups,
		Bets:            bets,
		Balances:        balances,
	}
}

// QueryState retrieves the state for the current round of this game.
func (g *Game) QueryState(ctx context.Context) (State, error) {
	state, err := g.storer.QueryStateByID(ctx, g.id, g.round)
	if err != nil {
		return State{}, fmt.Errorf("query: %w", err)
	}

	return state, nil
}

// QueryStateByRound retrieves the state for the specified round of this game.
func (g *Game) QueryStateByRound(ctx context.Context, round int) (State, error) {
	state, err := g.storer.QueryStateByID(ctx, g.id, round)
	if err != nil {
		return State{}, fmt.Errorf("query: %w", err)
	}

	return state, nil
}
