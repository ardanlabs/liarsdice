package game

import (
	"context"
	"errors"
	"math/big"
	"testing"
)

type MockedBank struct {
	value *big.Int
	err   error
}

func (m *MockedBank) Balance(ctx context.Context, account string) (*big.Int, error) {
	return m.value, m.err
}

func (m *MockedBank) Reconcile(ctx context.Context, winningAccount string, losingAccounts []string, anteWei *big.Int, gameFeeWei *big.Int) error {
	return m.err
}

func TestSuccessGamePlay(t *testing.T) {
	bank := MockedBank{
		value: big.NewInt(100),
		err:   nil,
	}

	ctx := context.Background()
	g, err := New(ctx, &bank, "player1", 0)
	if err != nil {
		t.Fatalf("unexpected error creating game: %s", err)
	}

	// =========================================================================
	// Game Setup.

	err = g.AddAccount(ctx, "player2")
	if err != nil {
		t.Fatalf("unexpected error adding player 2: %s", err)
	}

	// Check cups number.
	if len(g.cups) != 2 {
		t.Fatalf("expecting 2 players; got %d", len(g.cups))
	}

	// Start the game.
	err = g.StartGame("player1")
	if err != nil {
		t.Fatalf("unexpected error starting the game: %s", err)
	}

	// After starting the game, the status should be updated to 'StatusPlaying'
	if g.status != StatusPlaying {
		t.Fatalf("expecting game status to be %s; got %s", StatusPlaying, g.status)
	}

	// =========================================================================
	// Mocked roll dice so we can validate the winner and loser.

	player1 := g.cups["player1"]
	player1.Dice = []int{6, 5, 3, 3, 3}
	g.cups["player1"] = player1

	player2 := g.cups["player2"]
	player2.Dice = []int{1, 1, 4, 4, 2}
	g.cups["player2"] = player2

	// =========================================================================
	// Players claims.

	err = g.Claim("player1", 2, 3)
	if err != nil {
		t.Fatalf("unexpected error making claim for player1: %s", err)
	}

	err = g.Claim("player2", 3, 4)
	if err != nil {
		t.Fatalf("unexpected error making claim for player2: %s", err)
	}

	// =========================================================================
	// Player 1 calls Player 2 a liar.

	winner, loser, err := g.CallLiar("player1")
	if err != nil {
		t.Fatalf("unexpected error calling liar for player1: %s", err)
	}

	// =========================================================================
	// Check winner and loser.

	if winner != "player1" {
		t.Fatalf("expecting 'player1' to be the winner; got '%s'", winner)
	}

	if loser != "player2" {
		t.Fatalf("expecting 'player2' to be the loser; got '%s'", loser)
	}

	// Validate outs count for player2.
	player2 = g.cups["player2"]
	if player2.Outs != 1 {
		t.Fatalf("expecting 'player2' to have 1 out; got %d", player2.Outs)
	}

	if g.status != StatusRoundOver {
		t.Fatalf("expecting game status to be %s; got %s", StatusRoundOver, g.status)
	}

	// =========================================================================
	// Start second round.
	// =========================================================================

	leftToPlay, err := g.NextRound()
	if err != nil {
		t.Fatalf("unexpected error starting new round: %s", err)
	}

	if leftToPlay != 2 {
		t.Fatalf("expecting 2 players; got %d", leftToPlay)
	}

	if g.status != StatusPlaying {
		t.Fatalf("expecting game status to be %s; got %s", StatusPlaying, g.status)
	}

	// Mocked roll dice so we can validate the winner and loser.
	player1 = g.cups["player1"]
	player1.Dice = []int{1, 2, 3, 1, 6}
	g.cups["player1"] = player1

	player2 = g.cups["player2"]
	player2.Dice = []int{3, 2, 6, 5, 6}
	g.cups["player2"] = player2

	// =========================================================================
	// Players claims.

	err = g.Claim("player1", 2, 1)
	if err != nil {
		t.Fatalf("unexpected error making claim for player1: %s", err)
	}

	// =========================================================================
	// Player 2 calls Player 1 a liar.

	winner, loser, err = g.CallLiar("player2")
	if err != nil {
		t.Fatalf("unexpected error calling liar for player2: %s", err)
	}

	// =========================================================================
	// Check winner and loser.

	if winner != "player1" {
		t.Fatalf("expecting 'player1' to be the winner; got '%s'", winner)
	}

	if loser != "player2" {
		t.Fatalf("expecting 'player2' to be the loser; got '%s'", loser)
	}

	// Validate outs count for player2.
	player2 = g.cups["player2"]
	if player2.Outs != 2 {
		t.Fatalf("expecting 'player2' to have 2 outs; got %d", player2.Outs)
	}

	// =========================================================================
	// Start third round.
	// =========================================================================

	leftToPlay, err = g.NextRound()
	if err != nil {
		t.Fatalf("unexpected error starting new round: %s", err)
	}

	if leftToPlay != 2 {
		t.Fatalf("expecting 2 players; got %d", leftToPlay)
	}

	// Mocked roll dice so we can validate the winner and loser.
	player1 = g.cups["player1"]
	player1.Dice = []int{1, 1, 6, 1, 1}
	g.cups["player1"] = player1

	player2 = g.cups["player2"]
	player2.Dice = []int{3, 3, 3, 5, 6}
	g.cups["player2"] = player2

	// =========================================================================
	// Player claim.

	err = g.Claim("player2", 4, 3)
	if err != nil {
		t.Fatalf("unexpected error making claim for player2: %s", err)
	}

	// =========================================================================
	// Player 1 calls Player 2 a liar.

	winner, loser, err = g.CallLiar("player1")
	if err != nil {
		t.Fatalf("unexpected error calling liar for player1: %s", err)
	}

	// =========================================================================
	// Check winner and loser.

	if winner != "player1" {
		t.Fatalf("expecting 'player1' to be the winner; got '%s'", winner)
	}

	if loser != "player2" {
		t.Fatalf("expecting 'player2' to be the loser; got '%s'", loser)
	}

	// Validate outs count for player2.
	player2 = g.cups["player2"]
	if player2.Outs != 3 {
		t.Fatalf("expecting 'player2' to have 3 outs; got %d", player2.Outs)
	}

	// =========================================================================
	// There should be only one player left, player1
	// =========================================================================

	leftToPlay, err = g.NextRound()
	if err != nil {
		t.Fatalf("unexpected error starting new round: %s", err)
	}

	if leftToPlay != 1 {
		t.Fatalf("expecting 1 player; got %d", leftToPlay)
	}

	if g.lastWinAcct != "player1" {
		t.Fatalf("expecting 'player1' to be the LastWinAcct; got '%s'", g.lastWinAcct)
	}
}

func TestInvalidClaim(t *testing.T) {
	bank := MockedBank{
		value: big.NewInt(100),
		err:   nil,
	}

	ctx := context.Background()

	g, err := New(ctx, &bank, "player1", 0)
	if err != nil {
		t.Fatalf("unexpected error adding owner: %s", err)
	}

	// =========================================================================
	// Game Setup.

	err = g.AddAccount(ctx, "player2")
	if err != nil {
		t.Fatalf("unexpected error adding player 2: %s", err)
	}

	err = g.StartGame("player1")
	if err != nil {
		t.Fatalf("unexpected error starting game: %s", err)
	}

	// =========================================================================
	// Mock roll dice so we can validate the winner and loser.

	player1 := g.cups["player1"]
	player1.Dice = []int{6, 5, 3, 3, 3}
	g.cups["player1"] = player1

	player2 := g.cups["player2"]
	player2.Dice = []int{1, 1, 4, 4, 2}
	g.cups["player2"] = player2

	// =========================================================================
	// Players make claims.

	err = g.Claim("player1", 3, 3)
	if err != nil {
		t.Fatalf("unexpected error making claim for player1: %s", err)
	}

	g.NextTurn("player1")

	err = g.Claim("player2", 2, 6)
	if err == nil {
		t.Fatal("expecting error making an invalid claim")
	}
}

func TestGameWithoutEnoughPlayers(t *testing.T) {
	bank := MockedBank{
		value: big.NewInt(100),
		err:   nil,
	}

	ctx := context.Background()

	g, err := New(ctx, &bank, "owner", 0)
	if err != nil {
		t.Fatal("not expecting error creating game")
	}

	err = g.StartGame("owner")
	if err == nil {
		t.Fatal("expecting error trying to start a game without enough players")
	}
}

func TestWrongPlayerTryingToPlay(t *testing.T) {
	bank := MockedBank{
		value: big.NewInt(100),
		err:   nil,
	}

	ctx := context.Background()
	g, err := New(ctx, &bank, "owner", 0)
	if err != nil {
		t.Fatal("not expecting error creating game")
	}

	err = g.AddAccount(ctx, "player1")
	if err != nil {
		t.Fatalf("unexpected error: %s", err)
	}

	err = g.AddAccount(ctx, "player2")
	if err != nil {
		t.Fatalf("unexpected error: %s", err)
	}

	err = g.StartGame("owner")
	if err != nil {
		t.Fatalf("unexpected error: %s", err)
	}

	err = g.Claim("player2", 1, 1)
	if err == nil {
		t.Fatal("expecting error making claim with the wrong player")
	}
}

func TestAddAccountWithoutBalance(t *testing.T) {
	bank := MockedBank{
		value: big.NewInt(100),
		err:   nil,
	}

	ctx := context.Background()
	g, err := New(ctx, &bank, "owner", 100)
	if err != nil {
		t.Fatal("not expecting error creating game")
	}

	bank.err = errors.New("the player don't have enough balance")

	err = g.AddAccount(ctx, "player1")
	if err == nil {
		t.Fatal("expecting error adding player without balance")
	}
}
