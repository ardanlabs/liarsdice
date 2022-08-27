package game

import (
	"context"
	"errors"
	"fmt"
	"math/big"
	"testing"
	"time"

	"github.com/ardanlabs/liarsdice/foundation/smart/currency"
	"github.com/ethereum/go-ethereum/core/types"
)

const (
	OwnerAddress   = "0x6327A38415C53FFb36c11db55Ea74cc9cB4976Fd"
	Player1Address = "0x0070742ff6003c3e809e78d524f0fe5dcc5ba7f7"
	Player2Address = "0x8e113078adf6888b7ba84967f299f29aece24c55"
)

// =============================================================================

// mockBank represents an in-memory bank for testing.
type mockBank struct {
	owner    string
	balances map[string]*big.Int
}

// newBank constructs a mock bank for use with the game package.
func newBank() *mockBank {
	converter := currency.NewDefaultConverter()

	balances := map[string]*big.Int{
		OwnerAddress:   converter.USD2Wei(big.NewFloat(1_000_000)),
		Player1Address: converter.USD2Wei(big.NewFloat(1_000)),
		Player2Address: converter.USD2Wei(big.NewFloat(1_000)),
	}

	return &mockBank{
		owner:    OwnerAddress,
		balances: balances,
	}
}

// AccountBalance implements the game.Banker interface and will return the
// account balance for the specified account.
func (mb *mockBank) AccountBalance(ctx context.Context, account string) (GWei *big.Float, err error) {
	amountWei, exists := mb.balances[account]
	if !exists {
		return nil, fmt.Errorf("account %q does not exist", account)
	}

	return currency.Wei2GWei(amountWei), nil
}

// Reconcile implements the game.Banker interface and will reconcile a game with
// the same logic the bank smart contract is using.
func (mb *mockBank) Reconcile(ctx context.Context, winningAccount string, losingAccounts []string, anteGWei *big.Float, gameFeeGWei *big.Float) (*types.Transaction, *types.Receipt, error) {

	// The smart contract deals in wei.
	anteWei := currency.GWei2Wei(anteGWei)
	gameFeeWei := currency.GWei2Wei(gameFeeGWei)

	// Add the ante for each player to the pot. The initialization is
	// for the winner's ante.
	pot := anteWei
	for _, account := range losingAccounts {
		if mb.balances[account].Cmp(anteWei) == -1 {
			pot = big.NewInt(0).Add(pot, mb.balances[account])
			mb.balances[account] = big.NewInt(0)
		} else {
			pot = big.NewInt(0).Add(pot, anteWei)
			mb.balances[account] = big.NewInt(0).Sub(mb.balances[account], anteWei)
		}
	}

	// This should not happen but check to see if the pot is 0 because none
	// of the losers had an account balance.
	if pot.Cmp(big.NewInt(0)) == 0 {
		return nil, nil, errors.New("pot is zero")
	}

	// This should not happen but check there is enough in the pot to cover
	// the game fee.
	if pot.Cmp(gameFeeWei) == -1 {
		fmt.Printf("pot less than fee: winner[0] owner[%d]\n", pot)
		mb.balances[mb.owner] = big.NewInt(0).Add(mb.balances[mb.owner], pot)
		return nil, nil, nil
	}

	// Take the game fee from the pot and give the winner the remaining pot
	// and the owner the game fee.
	pot = big.NewInt(0).Sub(pot, gameFeeWei)
	mb.balances[winningAccount] = big.NewInt(0).Add(mb.balances[winningAccount], pot)
	mb.balances[mb.owner] = big.NewInt(0).Add(mb.balances[mb.owner], gameFeeWei)

	fmt.Printf("winner[%d] owner[%d]", pot, gameFeeWei)

	return nil, nil, nil
}

// =============================================================================

func TestSuccessGamePlay(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	converter := currency.NewDefaultConverter()
	bank := newBank()

	g, err := New(ctx, converter, bank, "player1", 0)
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
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	converter := currency.NewDefaultConverter()
	bank := newBank()

	g, err := New(ctx, converter, bank, "player1", 0)
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
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	converter := currency.NewDefaultConverter()
	bank := newBank()

	g, err := New(ctx, converter, bank, "owner", 0)
	if err != nil {
		t.Fatal("not expecting error creating game")
	}

	err = g.StartGame("owner")
	if err == nil {
		t.Fatal("expecting error trying to start a game without enough players")
	}
}

func TestWrongPlayerTryingToPlay(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	converter := currency.NewDefaultConverter()
	bank := newBank()

	g, err := New(ctx, converter, bank, "owner", 0)
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
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	converter := currency.NewDefaultConverter()
	bank := newBank()

	_, err := New(ctx, converter, bank, "owner", 100)
	if err == nil {
		t.Fatal("expecting error adding player without balance")
	}
}
