// Package game represents all the game play for liar's dice.
package game

import (
	"context"
	"errors"
	"fmt"
	"math/big"
	"math/rand"

	"github.com/google/uuid"
)

const (
	STATUSROUNDOVER = "roundover"
	STATUSPLAYING   = "playing"
	STATUSOPEN      = "open"
	NUMBERPLAYERS   = 2
)

// =============================================================================

// Banker interface declares the bank behaviour.
type Banker interface {
	PlayerBalance(ctx context.Context, address string) (*big.Int, error)
	Reconcile(ctx context.Context, winner string, losers []string, ante uint, gameFee uint) error
}

// =============================================================================

// Player represents a person playing the game.
type Player struct {
	Wallet string
	Outs   uint8
	Dice   []int
}

// Claim represents a claim of dice on the table.
type Claim struct {
	Wallet string
	Number int
	Suite  int
}

// =============================================================================

// Game represents a single game that is being played.
type Game struct {
	ID            string
	Status        string
	Banker        Banker
	CurrentPlayer int
	Round         int
	Players       []Player
	Claims        []Claim
}

// NewGame creates a new game.
func NewGame(banker Banker) *Game {
	return &Game{
		ID:      uuid.NewString(),
		Status:  STATUSOPEN,
		Players: []Player{},
		Banker:  banker,
	}
}

// StartGame will check if the current game can be started and update its status.
func (g *Game) StartGame() error {
	if g.Status != STATUSOPEN {
		return errors.New("game cannot be started")
	}

	if len(g.Players) < NUMBERPLAYERS {
		return errors.New("not enough players to start game")
	}

	g.Round = 1
	g.Status = STATUSPLAYING

	return nil
}

// AddPlayer adds the player to the game. The player will not be added twice.
func (g *Game) AddPlayer(wallet string) error {
	if wallet == "" {
		return errors.New("invalid wallet address")
	}

	// Check if player wallet was already added to the game.
	for _, player := range g.Players {
		if wallet == player.Wallet {
			return fmt.Errorf("player [%s] is already in the game", wallet)
		}
	}

	// Create a player based on the wallet address.
	player := Player{
		Wallet: wallet,
	}

	g.Players = append(g.Players, player)

	return nil
}

// RemovePlayer removes a player from the game.
func (g *Game) RemovePlayer(wallet string) error {
	if wallet == "" {
		return errors.New("invalid wallet address")
	}

	for i, player := range g.Players {
		if player.Wallet == wallet {
			g.Players = append(g.Players[:i], g.Players[i+1:]...)
			return nil
		}
	}

	return errors.New("player not found")
}

// CallLiar checks all the claims made so far in the round and defines a winner
//and a loser.
func (g *Game) CallLiar(wallet string) (string, string, error) {
	if wallet == "" {
		return "", "", errors.New("invalid wallet address")
	}

	// Validate if it is the player's turn..
	if g.Players[g.CurrentPlayer].Wallet != wallet {
		return "", "", fmt.Errorf("player [%s] can't make a claim now", wallet)
	}

	// Update the game status.
	g.Status = STATUSROUNDOVER

	// Hold the sum of all the dice values.
	dice := make([]int, 7)
	for _, player := range g.Players {
		for _, suite := range player.Dice {
			dice[suite]++
		}
	}

	lastClaim := g.Claims[len(g.Claims)-1]

	// If the last Claim is incorrect, the player who made it, gets an out.
	if dice[lastClaim.Suite] < lastClaim.Number {
		var loser string
		for i := range g.Players {
			if g.Players[i].Wallet == lastClaim.Wallet {
				g.Players[i].Outs++
				loser = g.Players[i].Wallet
				break
			}
		}
		return wallet, loser, nil
	}

	// Find the calling player to increment the out count.
	for i := range g.Players {
		if g.Players[i].Wallet == wallet {
			g.Players[i].Outs++
			break
		}
	}

	return lastClaim.Wallet, wallet, nil
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
		return 1, nil
	}

	// Figure out who starts the round. The person who was last out should
	// start the round.
	// var found bool
	// for i, player := range t.Game.Players {
	// 	if player.UserID == t.Game.LastOut.UserID {
	// 		t.Game.nextPlayer = i
	// 		found = true
	// 	}
	// }

	// If the person who was last out is no longer in the game, then the
	// player who won the last round starts.
	// if !found {
	// 	for i, player := range t.Game.Players {
	// 		if player.UserID == t.Game.LastWin.UserID {
	// 			t.Game.nextPlayer = i
	// 		}
	// 	}
	// }

	// Reset players dice.
	for i := range g.Players {
		g.Players[i].Dice = make([]int, 6)
	}

	// Reset the claims to start over.
	g.Claims = []Claim{}

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
	claim.Wallet = wallet

	g.Claims = append(g.Claims, claim)

	// Update the CurrentPlayer index.
	g.CurrentPlayer++
	g.CurrentPlayer = g.CurrentPlayer % len(g.Players)

	return nil
}

// RollDice will generate 5 random integer and add to the player's dice list.
func (g *Game) RollDice(wallet string) error {
	if wallet == "" {
		return errors.New("invalid wallet address")
	}

	// The game should be in the playing state.
	if g.Status != STATUSPLAYING {
		return errors.New("game is not started")
	}

	// Look for the player and roll the dice.
	var found bool
	for i := range g.Players {
		if g.Players[i].Wallet == wallet {
			found = true

			dice := make([]int, 5)
			for i := range dice {
				dice[i] = rand.Intn(6) + 1
			}

			g.Players[i].Dice = dice
			break
		}
	}

	if !found {
		return fmt.Errorf("player [%s] not found in this game", wallet)
	}

	return nil
}

// PlayerBalance returns the player's balance, by calling the banks contract
// method.
func (g *Game) PlayerBalance(ctx context.Context, wallet string) (*big.Int, error) {
	if wallet == "" {
		return nil, errors.New("invalid wallet address")
	}

	return g.Banker.PlayerBalance(ctx, wallet)
}

// Reconcile calculates the game pot and make the transfer to the winner.
func (g *Game) Reconcile(ctx context.Context, winner string, losers []string, ante uint, gameFee uint) error {
	return nil
}
