package game_test

import (
	"context"
	"errors"
	"fmt"
	"math/big"
	"testing"
	"time"

	"github.com/ardanlabs/liarsdice/business/core/game"
	"github.com/ardanlabs/liarsdice/foundation/smart/currency"
	"github.com/ethereum/go-ethereum/core/types"
)

const (
	Owner   = "owner"
	Player1 = "player1"
	Player2 = "player2"
)

// =============================================================================

// mockBank represents an in-memory bank for testing.
type mockBank struct {
	Owner    string
	Starting map[string]*big.Int
	Balances map[string]*big.Int
}

// newMockBank constructs a mock bank for use with the game package.
func newMockBank() *mockBank {
	converter := currency.NewDefaultConverter()

	ownerBalanceWei := converter.USD2Wei(big.NewFloat(0.0))
	player1BalanceWei := converter.USD2Wei(big.NewFloat(100.00))
	player2BalanceWei := converter.USD2Wei(big.NewFloat(100.00))

	starting := map[string]*big.Int{
		Owner:   ownerBalanceWei,
		Player1: player1BalanceWei,
		Player2: player2BalanceWei,
	}

	balances := map[string]*big.Int{
		Owner:   ownerBalanceWei,
		Player1: player1BalanceWei,
		Player2: player2BalanceWei,
	}

	return &mockBank{
		Owner:    Owner,
		Starting: starting,
		Balances: balances,
	}
}

// AccountBalance implements the game.Banker interface and will return the
// account balance for the specified account.
func (mb *mockBank) AccountBalance(ctx context.Context, account string) (GWei *big.Float, err error) {
	amountWei, exists := mb.Balances[account]
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
		if mb.Balances[account].Cmp(anteWei) == -1 {
			pot = big.NewInt(0).Add(pot, mb.Balances[account])
			mb.Balances[account] = big.NewInt(0)
		} else {
			pot = big.NewInt(0).Add(pot, anteWei)
			mb.Balances[account] = big.NewInt(0).Sub(mb.Balances[account], anteWei)
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
		mb.Balances[mb.Owner] = big.NewInt(0).Add(mb.Balances[mb.Owner], pot)
		return nil, nil, nil
	}

	// Take the game fee from the pot and give the winner the remaining pot
	// and the owner the game fee.
	pot = big.NewInt(0).Sub(pot, gameFeeWei)
	mb.Balances[winningAccount] = big.NewInt(0).Add(mb.Balances[winningAccount], pot)
	mb.Balances[mb.Owner] = big.NewInt(0).Add(mb.Balances[mb.Owner], gameFeeWei)

	return nil, nil, nil
}

// =============================================================================

func Test_SuccessGamePlay(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	converter := currency.NewDefaultConverter()
	bank := newMockBank()

	anteUSD := float64(5.0)
	anteWei := converter.USD2Wei(big.NewFloat(anteUSD))

	// =========================================================================
	// Create game and add players
	// =========================================================================

	// Create a game and add player1 and the owner and first player in the game.
	g, err := game.New(ctx, converter, bank, Player1, anteUSD)
	if err != nil {
		t.Fatalf("unexpected error creating game: %s", err)
	}

	// Add player2 as the second player in the game.
	err = g.AddAccount(ctx, Player2)
	if err != nil {
		t.Fatalf("unexpected error adding player 2: %s", err)
	}

	status := g.Info()

	if len(status.Cups) != 2 {
		t.Fatalf("expecting 2 players; got %d", len(status.Cups))
	}

	// =========================================================================
	// Start first round
	// =========================================================================

	// Only the owner can start the game.
	err = g.StartGame(Player1)
	if err != nil {
		t.Fatalf("unexpected error starting the game: %s", err)
	}

	status = g.Info()

	if status.Status != game.StatusPlaying {
		t.Fatalf("expecting game status to be %s; got %s", game.StatusPlaying, status.Status)
	}

	// =========================================================================
	// Mocked roll dice so we can validate the winner and loser.

	dice := []int{6, 5, 3, 3, 3}
	g.RollDice(Player1, dice...)

	dice = []int{1, 1, 4, 4, 2}
	g.RollDice(Player2, dice...)

	// =========================================================================
	// Game Play.

	if err := g.Bet(Player1, 2, 3); err != nil {
		t.Fatalf("unexpected error making bet for player1: %s", err)
	}

	if err := g.Bet(Player2, 3, 4); err != nil {
		t.Fatalf("unexpected error making bet for player2: %s", err)
	}

	// Player1 calls Player2 a liar.
	winner, loser, err := g.CallLiar(Player1)
	if err != nil {
		t.Fatalf("unexpected error calling liar for player1: %s", err)
	}

	// =========================================================================
	// Check winner and loser.

	if winner != Player1 {
		t.Fatalf("expecting 'player1' to be the winner; got '%s'", winner)
	}

	if loser != Player2 {
		t.Fatalf("expecting 'player2' to be the loser; got '%s'", loser)
	}

	status = g.Info()

	if status.Cups[Player2].Outs != 1 {
		t.Fatalf("expecting 'player2' to have 1 out; got %d", status.Cups[Player2].Outs)
	}

	if status.Status != game.StatusRoundOver {
		t.Fatalf("expecting game status to be %s; got %s", game.StatusRoundOver, status.Status)
	}

	// =========================================================================
	// Start second round
	// =========================================================================

	leftToPlay, err := g.NextRound()
	if err != nil {
		t.Fatalf("unexpected error starting new round: %s", err)
	}

	if leftToPlay != 2 {
		t.Fatalf("expecting 2 players; got %d", leftToPlay)
	}

	status = g.Info()

	if status.Status != game.StatusPlaying {
		t.Fatalf("expecting game status to be %s; got %s", game.StatusPlaying, status.Status)
	}

	// =========================================================================
	// Mocked roll dice so we can validate the winner and loser.

	dice = []int{1, 2, 3, 1, 6}
	g.RollDice(Player1, dice...)

	dice = []int{3, 2, 6, 5, 6}
	g.RollDice(Player2, dice...)

	// =========================================================================
	// Game Play.

	err = g.Bet(Player1, 2, 1)
	if err != nil {
		t.Fatalf("unexpected error making bet for player1: %s", err)
	}

	winner, loser, err = g.CallLiar(Player2)
	if err != nil {
		t.Fatalf("unexpected error calling liar for player2: %s", err)
	}

	// =========================================================================
	// Check winner and loser.

	if winner != Player1 {
		t.Fatalf("expecting 'player1' to be the winner; got '%s'", winner)
	}

	if loser != Player2 {
		t.Fatalf("expecting 'player2' to be the loser; got '%s'", loser)
	}

	status = g.Info()

	if status.Cups[Player2].Outs != 2 {
		t.Fatalf("expecting 'player2' to have 2 out; got %d", status.Cups[Player2].Outs)
	}

	if status.Status != game.StatusRoundOver {
		t.Fatalf("expecting game status to be %s; got %s", game.StatusRoundOver, status.Status)
	}

	// =========================================================================
	// Start third round
	// =========================================================================

	leftToPlay, err = g.NextRound()
	if err != nil {
		t.Fatalf("unexpected error starting new round: %s", err)
	}

	if leftToPlay != 2 {
		t.Fatalf("expecting 2 players; got %d", leftToPlay)
	}

	status = g.Info()

	if status.Status != game.StatusPlaying {
		t.Fatalf("expecting game status to be %s; got %s", game.StatusPlaying, status.Status)
	}

	// =========================================================================
	// Mocked roll dice so we can validate the winner and loser.

	dice = []int{1, 1, 6, 1, 1}
	g.RollDice(Player1, dice...)

	dice = []int{3, 3, 3, 5, 6}
	g.RollDice(Player2, dice...)

	// =========================================================================
	// Game Play.

	err = g.Bet(Player2, 4, 3)
	if err != nil {
		t.Fatalf("unexpected error making bet for player2: %s", err)
	}

	winner, loser, err = g.CallLiar(Player1)
	if err != nil {
		t.Fatalf("unexpected error calling liar for player1: %s", err)
	}

	// =========================================================================
	// Check winner and loser.

	if winner != Player1 {
		t.Fatalf("expecting 'player1' to be the winner; got '%s'", winner)
	}

	if loser != Player2 {
		t.Fatalf("expecting 'player2' to be the loser; got '%s'", loser)
	}

	status = g.Info()

	if status.Cups[Player2].Outs != 3 {
		t.Fatalf("expecting 'player2' to have 3 out; got %d", status.Cups[Player2].Outs)
	}

	if status.Status != game.StatusRoundOver {
		t.Fatalf("expecting game status to be %s; got %s", game.StatusRoundOver, status.Status)
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

	status = g.Info()

	if status.Status != game.StatusGameOver {
		t.Fatalf("expecting game status to be %s; got %s", game.StatusGameOver, status.Status)
	}

	if status.LastWinAcct != Player1 {
		t.Fatalf("expecting 'player1' to be the LastWinAcct; got '%s'", status.LastWinAcct)
	}

	// =========================================================================
	// Reconcile the game
	// =========================================================================

	if _, _, err := g.Reconcile(ctx, winner); err != nil {
		t.Fatalf("unexpected error reconciling the game: %s", err)
	}

	// =========================================================================
	// Check balances.

	ownerBalance := bank.Balances[Owner]
	if ownerBalance.Cmp(anteWei) != 0 {
		t.Errorf("expecting 'owner' to have a balance of %d WEI; got %d WEI", anteWei, ownerBalance)
	}

	expBalance := big.NewInt(0).Add(bank.Starting[Player1], anteWei)
	player1Balance := bank.Balances[Player1]
	if player1Balance.Cmp(expBalance) != 0 {
		t.Errorf("expecting 'player1' to have a balance of %d WEI; got %d WEI", expBalance, player1Balance)
	}

	expBalance = big.NewInt(0).Sub(bank.Starting[Player2], anteWei)
	player2Balance := bank.Balances[Player2]
	if player2Balance.Cmp(expBalance) != 0 {
		t.Errorf("expecting 'player2' to have a balance of %d WEI; got %d WEI", expBalance, player2Balance)
	}

	// =========================================================================
	// Validate final game state
	// =========================================================================

	status = g.Info()

	if status.Status != game.StatusReconciled {
		t.Fatalf("expecting game status to be %s; got %s", game.StatusReconciled, status.Status)
	}
}

func Test_InvalidClaim(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	converter := currency.NewDefaultConverter()
	bank := newMockBank()

	anteUSD := float64(5.0)

	// =========================================================================
	// Create game and add players
	// =========================================================================

	g, err := game.New(ctx, converter, bank, Player1, anteUSD)
	if err != nil {
		t.Fatalf("unexpected error adding owner: %s", err)
	}

	err = g.AddAccount(ctx, Player2)
	if err != nil {
		t.Fatalf("unexpected error adding player 2: %s", err)
	}

	status := g.Info()

	if len(status.Cups) != 2 {
		t.Fatalf("expecting 2 players; got %d", len(status.Cups))
	}

	// =========================================================================
	// Start first round
	// =========================================================================

	// Only the owner can start the game.
	err = g.StartGame(Player1)
	if err != nil {
		t.Fatalf("unexpected error starting the game: %s", err)
	}

	status = g.Info()

	if status.Status != game.StatusPlaying {
		t.Fatalf("expecting game status to be %s; got %s", game.StatusPlaying, status.Status)
	}

	// =========================================================================
	// Mocked roll dice so we can validate the winner and loser.

	dice := []int{6, 5, 3, 3, 3}
	g.RollDice(Player1, dice...)

	dice = []int{1, 1, 4, 4, 2}
	g.RollDice(Player2, dice...)

	// =========================================================================
	// Game Play.

	if err := g.Bet(Player1, 3, 3); err != nil {
		t.Fatalf("unexpected error making bet for player1: %s", err)
	}

	g.NextTurn(Player1)

	if err := g.Bet(Player2, 2, 6); err == nil {
		t.Fatal("expecting error making an invalid bet")
	}
}

func Test_GameWithoutEnoughPlayers(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	converter := currency.NewDefaultConverter()
	bank := newMockBank()

	anteUSD := float64(5.0)

	// =========================================================================
	// Create game and add players
	// =========================================================================

	g, err := game.New(ctx, converter, bank, Player1, anteUSD)
	if err != nil {
		t.Fatalf("unexpected error adding owner: %s", err)
	}

	err = g.StartGame(Player1)
	if err == nil {
		t.Fatal("expecting error trying to start a game without enough players")
	}
}

func Test_WrongPlayerTryingToPlay(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	converter := currency.NewDefaultConverter()
	bank := newMockBank()

	anteUSD := float64(5.0)

	// =========================================================================
	// Create game and add players
	// =========================================================================

	g, err := game.New(ctx, converter, bank, Player1, anteUSD)
	if err != nil {
		t.Fatalf("unexpected error adding owner: %s", err)
	}

	err = g.AddAccount(ctx, Player2)
	if err != nil {
		t.Fatalf("unexpected error adding player 2: %s", err)
	}

	status := g.Info()

	if len(status.Cups) != 2 {
		t.Fatalf("expecting 2 players; got %d", len(status.Cups))
	}

	// =========================================================================
	// Start first round
	// =========================================================================

	// Only the owner can start the game.
	err = g.StartGame(Player1)
	if err != nil {
		t.Fatalf("unexpected error starting the game: %s", err)
	}

	status = g.Info()

	if status.Status != game.StatusPlaying {
		t.Fatalf("expecting game status to be %s; got %s", game.StatusPlaying, status.Status)
	}

	err = g.Bet(Player2, 1, 1)
	if err == nil {
		t.Fatal("expecting error making bet with a player not in the game")
	}
}

func Test_NewGameNotEnoughBalance(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	converter := currency.NewDefaultConverter()
	bank := newMockBank()

	anteUSD := float64(2000.0)

	// =========================================================================
	// Create game where account doesn't have enough money
	// =========================================================================

	if _, err := game.New(ctx, converter, bank, Owner, anteUSD); err == nil {
		t.Fatal("expecting error adding player without enough balance")
	}
}
