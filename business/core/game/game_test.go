package game_test

import (
	"context"
	"testing"

	"github.com/ardanlabs/liarsdice/business/core/game"
)

func TestGamePlay(t *testing.T) {
	g := game.New(nil)
	ctx := context.Background()

	// =========================================================================
	// Game Setup.

	// Add players.
	err := g.AddAccount(ctx, "player1")
	if err != nil {
		t.Fatalf("unexpected error adding player 1: %s", err)
	}

	err = g.AddAccount(ctx, "player2")
	if err != nil {
		t.Fatalf("unexpected error adding player 2: %s", err)
	}

	// Check cups number.
	if len(g.Info().Cups) != 2 {
		t.Fatalf("expecting 2 players; got %d", len(g.Info().Cups))
	}

	// Start the game.
	err = g.StartPlay()
	if err != nil {
		t.Fatalf("unexpected error starting the game: %s", err)
	}

	// After starting the game, the status should be updated to 'StatusPlaying'
	if g.Info().Status != game.StatusPlaying {
		t.Fatalf("expecting game status to be %s; got %s", game.StatusPlaying, g.Info().Status)
	}

	// =========================================================================
	// Mocked roll dice so we can validate the winner and loser.

	player1 := g.Info().Cups["player1"]
	player1.Dice = []int{6, 5, 3, 3, 3}
	g.Info().Cups["player1"] = player1

	player2 := g.Info().Cups["player2"]
	player2.Dice = []int{1, 1, 4, 4, 2}
	g.Info().Cups["player2"] = player2

	// =========================================================================
	// Players claims.

	err = g.Claim("player1", 2, 3)
	if err != nil {
		t.Fatalf("unexpected error making claim for player1: %s", err)
	}

	g.NextTurn()

	err = g.Claim("player2", 3, 4)
	if err != nil {
		t.Fatalf("unexpected error making claim fgor player2: %s", err)
	}

	g.NextTurn()

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
	player2 = g.Info().Cups["player2"]
	if player2.Outs != 1 {
		t.Fatalf("expecting 'player2' to have 1 out; got %d", player2.Outs)
	}

	if g.Info().Status != game.StatusRoundOver {
		t.Fatalf("expecting game status to be %s; got %s", game.StatusRoundOver, g.Info().Status)
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

	if g.Info().Status != game.StatusPlaying {
		t.Fatalf("expecting game status to be %s; got %s", game.StatusPlaying, g.Info().Status)
	}

	// Mocked roll dice so we can validate the winner and loser.
	player1 = g.Info().Cups["player1"]
	player1.Dice = []int{1, 2, 3, 1, 6}
	g.Info().Cups["player1"] = player1

	player2 = g.Info().Cups["player2"]
	player2.Dice = []int{3, 2, 6, 5, 6}
	g.Info().Cups["player2"] = player2

	// =========================================================================
	// Players claims.

	err = g.Claim("player1", 2, 1)
	if err != nil {
		t.Fatalf("unexpected error making claim for player1: %s", err)
	}

	g.NextTurn()

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
	player2 = g.Info().Cups["player2"]
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
	player1 = g.Info().Cups["player1"]
	player1.Dice = []int{1, 1, 6, 1, 1}
	g.Info().Cups["player1"] = player1

	player2 = g.Info().Cups["player2"]
	player2.Dice = []int{3, 3, 3, 5, 6}
	g.Info().Cups["player2"] = player2

	// =========================================================================
	// Player claim.

	err = g.Claim("player2", 4, 3)
	if err != nil {
		t.Fatalf("unexpected error making claim for player2: %s", err)
	}

	g.NextTurn()

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
	player2 = g.Info().Cups["player2"]
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

	if g.Info().LastWinAcct != "player1" {
		t.Fatalf("expecting 'player1' to be the LastWinAcct; got '%s'", g.Info().LastWinAcct)
	}
}
