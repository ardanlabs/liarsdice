package game_test

import (
	"bytes"
	"context"
	"fmt"
	"math/big"
	"os"
	"testing"
	"time"

	"github.com/ardanlabs/ethereum"
	"github.com/ardanlabs/ethereum/currency"
	scbank "github.com/ardanlabs/liarsdice/business/contract/go/bank"
	"github.com/ardanlabs/liarsdice/business/core/bank"
	"github.com/ardanlabs/liarsdice/business/core/game"
	"github.com/ardanlabs/liarsdice/foundation/logger"
	"github.com/ardanlabs/liarsdice/foundation/web"
	"github.com/ethereum/go-ethereum/common"
)

var (
	backend    *ethereum.SimulatedBackend
	ownerClt   *ethereum.Client
	player1Clt *ethereum.Client
	player2Clt *ethereum.Client
)

func TestMain(m *testing.M) {
	var err error
	backend, err = ethereum.CreateSimulatedBackend(3, true, big.NewInt(100))
	if err != nil {
		fmt.Println("create backend", err)
		os.Exit(1)
	}
	defer backend.Close()

	ownerClt, err = ethereum.NewClient(backend, backend.PrivateKeys[0])
	if err != nil {
		fmt.Println("create ownerClt client", err)
		os.Exit(1)
	}

	player1Clt, err = ethereum.NewClient(backend, backend.PrivateKeys[1])
	if err != nil {
		fmt.Println("create player1Clt client", err)
		os.Exit(1)
	}

	player2Clt, err = ethereum.NewClient(backend, backend.PrivateKeys[2])
	if err != nil {
		fmt.Println("create player2Clt client", err)
		os.Exit(1)
	}

	m.Run()
}

func Test_SuccessGamePlay(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	converter := currency.NewDefaultConverter(scbank.BankMetaData.ABI)

	bank, engine := gameSetup(t)

	player1Addr := player1Clt.Address()
	player2Addr := player2Clt.Address()

	// -------------------------------------------------------------------------
	// Define the ante for each player

	anteUSD := float64(5.0)
	anteWei := converter.USD2Wei(big.NewFloat(anteUSD))

	// -------------------------------------------------------------------------
	// Start first round

	err := engine.StartGame(ctx)
	if err != nil {
		t.Fatalf("unexpected error starting the game: %s", err)
	}

	status := engine.Info(ctx)
	if status.Status != game.StatusPlaying {
		t.Fatalf("expecting game status to be %s; got %s", game.StatusPlaying, status.Status)
	}

	// -------------------------------------------------------------------------
	// Mocked roll dice so we can validate the winner and loser

	dice := []int{6, 5, 3, 3, 3}
	engine.RollDice(ctx, player1Addr, dice...)

	dice = []int{1, 1, 4, 4, 2}
	engine.RollDice(ctx, player2Addr, dice...)

	// -------------------------------------------------------------------------
	// Game Play: Each player makes a bet and player1 calls liar.

	winnerAcct := engine.Info(ctx).PlayerTurn
	if err := engine.Bet(ctx, winnerAcct, 2, 3); err != nil {
		t.Fatalf("unexpected error making bet for player1: %s", err)
	}

	loserAcct := engine.Info(ctx).PlayerTurn
	if err := engine.Bet(ctx, loserAcct, 3, 4); err != nil {
		t.Fatalf("unexpected error making bet for player2: %s", err)
	}

	winner, loser, err := engine.CallLiar(ctx, engine.Info(ctx).PlayerTurn)
	if err != nil {
		t.Fatalf("unexpected error calling liar for player1: %s", err)
	}

	// -------------------------------------------------------------------------
	// Check winner and loser

	if winner != winnerAcct {
		t.Fatalf("expecting 'player1' to be the winner; got '%s'", winner)
	}

	if loser != loserAcct {
		t.Fatalf("expecting 'player2' to be the loser; got '%s'", loser)
	}

	status = engine.Info(ctx)

	if status.Cups[loserAcct].Outs != 1 {
		t.Fatalf("expecting 'player2' to have 1 out; got %d", status.Cups[player2Addr].Outs)
	}

	if status.Status != game.StatusRoundOver {
		t.Fatalf("expecting game status to be %s; got %s", game.StatusRoundOver, status.Status)
	}

	// -------------------------------------------------------------------------
	// Start second round

	leftToPlay, err := engine.NextRound(ctx)
	if err != nil {
		t.Fatalf("unexpected error starting new round: %s", err)
	}

	if leftToPlay != 2 {
		t.Fatalf("expecting 2 players; got %d", leftToPlay)
	}

	status = engine.Info(ctx)

	if status.Status != game.StatusPlaying {
		t.Fatalf("expecting game status to be %s; got %s", game.StatusPlaying, status.Status)
	}

	// -------------------------------------------------------------------------
	// Mocked roll dice so we can validate the winner and loser

	dice = []int{1, 2, 3, 1, 6}
	engine.RollDice(ctx, player1Addr, dice...)

	dice = []int{3, 2, 6, 5, 6}
	engine.RollDice(ctx, player2Addr, dice...)

	// -------------------------------------------------------------------------
	// Game Play : Player 2 places a bet and player 1 calls liar

	err = engine.Bet(ctx, loserAcct, 5, 1)
	if err != nil {
		t.Fatalf("unexpected error making bet for player1: %s", err)
	}

	winner, loser, err = engine.CallLiar(ctx, winnerAcct)
	if err != nil {
		t.Fatalf("unexpected error calling liar for player2: %s", err)
	}

	// -------------------------------------------------------------------------
	// Check winner and loser

	if winner != winnerAcct {
		t.Fatalf("expecting 'player1' to be the winner; got '%s'", winner)
	}

	if loser != loserAcct {
		t.Fatalf("expecting 'player2' to be the loser; got '%s'", loser)
	}

	status = engine.Info(ctx)

	if status.Cups[loserAcct].Outs != 2 {
		t.Fatalf("expecting 'player2' to have 2 out; got %d", status.Cups[player2Addr].Outs)
	}

	if status.Status != game.StatusRoundOver {
		t.Fatalf("expecting game status to be %s; got %s", game.StatusRoundOver, status.Status)
	}

	// -------------------------------------------------------------------------
	// Start third round

	leftToPlay, err = engine.NextRound(ctx)
	if err != nil {
		t.Fatalf("unexpected error starting new round: %s", err)
	}

	if leftToPlay != 2 {
		t.Fatalf("expecting 2 players; got %d", leftToPlay)
	}

	status = engine.Info(ctx)

	if status.Status != game.StatusPlaying {
		t.Fatalf("expecting game status to be %s; got %s", game.StatusPlaying, status.Status)
	}

	// -------------------------------------------------------------------------
	// Mocked roll dice so we can validate the winner and loser

	dice = []int{1, 1, 6, 1, 1}
	engine.RollDice(ctx, player1Clt.Address(), dice...)

	dice = []int{3, 3, 3, 5, 6}
	engine.RollDice(ctx, player2Clt.Address(), dice...)

	// -------------------------------------------------------------------------
	// Game Play : Player 2 makes a bet and player1 calls liar

	err = engine.Bet(ctx, loserAcct, 4, 3)
	if err != nil {
		t.Fatalf("unexpected error making bet for player2: %s", err)
	}

	winner, loser, err = engine.CallLiar(ctx, winnerAcct)
	if err != nil {
		t.Fatalf("unexpected error calling liar for player1: %s", err)
	}

	// -------------------------------------------------------------------------
	// Check winner and loser.

	if winner != winnerAcct {
		t.Fatalf("expecting 'player1' to be the winner; got '%s'", winner)
	}

	if loser != loserAcct {
		t.Fatalf("expecting 'player2' to be the loser; got '%s'", loser)
	}

	status = engine.Info(ctx)

	if status.Cups[loserAcct].Outs != 3 {
		t.Fatalf("expecting 'player2' to have 3 out; got %d", status.Cups[player2Addr].Outs)
	}

	if status.Status != game.StatusRoundOver {
		t.Fatalf("expecting game status to be %s; got %s", game.StatusRoundOver, status.Status)
	}

	// -------------------------------------------------------------------------
	// There should be only one player left, player1

	leftToPlay, err = engine.NextRound(ctx)
	if err != nil {
		t.Fatalf("unexpected error starting new round: %s", err)
	}

	if leftToPlay != 1 {
		t.Fatalf("expecting 1 player; got %d", leftToPlay)
	}

	status = engine.Info(ctx)

	if status.Status != game.StatusGameOver {
		t.Fatalf("expecting game status to be %s; got %s", game.StatusGameOver, status.Status)
	}

	if status.PlayerLastWin != winnerAcct {
		t.Fatalf("expecting 'player1' to be the LastWinAcct; got '%s'", status.PlayerLastWin)
	}

	// -------------------------------------------------------------------------
	// Reconcile the game

	if _, _, err := engine.Reconcile(ctx); err != nil {
		t.Fatalf("unexpected error reconciling the game: %s", err)
	}

	// -------------------------------------------------------------------------
	// Check balances

	engineBalance, err := bank.Balance(ctx)
	if err != nil {
		t.Fatalf("unexpected to retrieve the balance of the bank owner: %s", err)
	}

	player1Balance, err := bank.AccountBalance(ctx, winnerAcct)
	if err != nil {
		t.Fatalf("unexpected to retrieve the balance of player 1: %s", err)
	}

	player2Balance, err := bank.AccountBalance(ctx, loserAcct)
	if err != nil {
		t.Fatalf("unexpected to retrieve the balance of player 2: %s", err)
	}

	if currency.GWei2Wei(engineBalance).Cmp(anteWei) != 0 {
		t.Errorf("expecting 'engine' to have a balance of %d WEI; got %d WEI", anteWei, currency.GWei2Wei(engineBalance))
	}

	initalDepositWei := converter.USD2Wei(big.NewFloat(100))

	got := currency.GWei2Wei(player1Balance)
	exp := big.NewInt(0).Add(initalDepositWei, anteWei)
	if got.Cmp(exp) != 0 {
		t.Errorf("expecting 'player1' to have a balance of %d WEI; got %d WEI", exp, got)
	}

	got = currency.GWei2Wei(player2Balance)
	exp = big.NewInt(0).Sub(initalDepositWei, anteWei)
	if got.Cmp(exp) != 0 {
		t.Errorf("expecting 'player2' to have a balance of %d WEI; got %d WEI", exp, got)
	}

	// -------------------------------------------------------------------------
	// Validate final game state

	status = engine.Info(ctx)

	if status.Status != game.StatusReconciled {
		t.Fatalf("expecting game status to be %s; got %s", game.StatusReconciled, status.Status)
	}
}

func Test_InvalidBet(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, engine := gameSetup(t)

	player1Addr := player1Clt.Address()
	player2Addr := player2Clt.Address()

	// -------------------------------------------------------------------------
	// Start first round

	err := engine.StartGame(ctx)
	if err != nil {
		t.Fatalf("unexpected error starting the game: %s", err)
	}

	status := engine.Info(ctx)
	if status.Status != game.StatusPlaying {
		t.Fatalf("expecting game status to be %s; got %s", game.StatusPlaying, status.Status)
	}

	// -------------------------------------------------------------------------
	// Mocked roll dice so we can validate the winner and loser.

	dice := []int{6, 5, 3, 3, 3}
	engine.RollDice(ctx, player1Addr, dice...)

	dice = []int{1, 1, 4, 4, 2}
	engine.RollDice(ctx, player2Addr, dice...)

	// -------------------------------------------------------------------------
	// Game Play : player 1 makes bet and player 2 makes invalid bet

	if err := engine.Bet(ctx, engine.Info(ctx).PlayerTurn, 3, 3); err != nil {
		t.Fatalf("unexpected error making bet for player1: %s", err)
	}

	engine.NextTurn(ctx)

	if err := engine.Bet(ctx, engine.Info(ctx).PlayerTurn, 2, 6); err == nil {
		t.Fatal("expecting error making an invalid bet")
	}
}

func Test_WrongPlayerTryingToPlay(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, engine := gameSetup(t)

	player1Addr := player1Clt.Address()
	player2Addr := player2Clt.Address()

	// -------------------------------------------------------------------------
	// Start first round

	err := engine.StartGame(ctx)
	if err != nil {
		t.Fatalf("unexpected error starting the game: %s", err)
	}

	status := engine.Info(ctx)
	if status.Status != game.StatusPlaying {
		t.Fatalf("expecting game status to be %s; got %s", game.StatusPlaying, status.Status)
	}

	var wrongPlayer common.Address
	switch engine.Info(ctx).PlayerTurn {
	case player1Addr:
		wrongPlayer = player2Addr
	case player2Addr:
		wrongPlayer = player1Addr
	}

	err = engine.Bet(ctx, wrongPlayer, 1, 1)
	if err == nil {
		t.Fatal("expecting error making bet with a player not in the game")
	}
}

func Test_GameWithoutEnoughPlayers(t *testing.T) {
	contractID, err := deployContract()
	if err != nil {
		t.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	converter := currency.NewDefaultConverter(scbank.BankMetaData.ABI)

	var buf bytes.Buffer
	log := logger.New(&buf, logger.LevelInfo, "TEST", func(context.Context) string { return web.GetTraceID(ctx) })

	// -------------------------------------------------------------------------
	// Players need to deposit money into their accounts

	player1Bank, err := bank.New(ctx, log, backend, player1Clt.PrivateKey(), contractID)
	if err != nil {
		t.Fatalf("creating new bank for player 1: %s", err)
	}

	initalDepositGwei := converter.USD2GWei(big.NewFloat(100))
	_, _, err = player1Bank.Deposit(ctx, initalDepositGwei)
	if err != nil {
		t.Fatalf("depositing money into bank for player1: %s", err)
	}

	// -------------------------------------------------------------------------
	// Create game and add players

	bank, err := bank.New(ctx, log, backend, ownerClt.PrivateKey(), contractID)
	if err != nil {
		t.Fatalf("creating new bank for the engine: %s", err)
	}

	const anteUSD = 5.0
	game, err := game.New(ctx, log, converter, bank, player1Clt.Address(), anteUSD)
	if err != nil {
		t.Fatalf("unexpected error creating game: %s", err)
	}

	// -------------------------------------------------------------------------
	// Start the game with only 1 player

	err = game.StartGame(ctx)
	if err == nil {
		t.Fatal("expecting error trying to start a game without enough players")
	}
}

func Test_NewGameNotEnoughBalance(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	contractID, err := deployContract()
	if err != nil {
		t.Fatal(err)
	}

	converter := currency.NewDefaultConverter(scbank.BankMetaData.ABI)

	var buf bytes.Buffer
	log := logger.New(&buf, logger.LevelInfo, "TEST", func(context.Context) string { return web.GetTraceID(ctx) })

	// -------------------------------------------------------------------------
	// Players need to deposit money into their accounts

	bank, err := bank.New(ctx, log, backend, player1Clt.PrivateKey(), contractID)
	if err != nil {
		t.Fatalf("creating new bank for player 1: %s", err)
	}

	const anteUSD = 5.0
	_, err = game.New(ctx, log, converter, bank, player1Clt.Address(), anteUSD)
	if err == nil {
		t.Fatalf("expecting an error creating a game: %s", err)
	}
}

func deployContract() (common.Address, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	fmt.Println("Deploying Contract ...")
	defer fmt.Println("Deployed")

	contractID, err := smartContract(ctx)
	if err != nil {
		fmt.Println("error deploying a new contract:", err)

		var empty common.Address
		return empty, err
	}

	return contractID, nil
}

func smartContract(ctx context.Context) (common.Address, error) {
	var empty common.Address

	tranOpts, err := ownerClt.NewTransactOpts(ctx, 10_000_000, big.NewInt(0), big.NewFloat(0))
	if err != nil {
		return empty, err
	}

	address, tx, _, err := scbank.DeployBank(tranOpts, ownerClt.Backend)
	if err != nil {
		return empty, err
	}

	if _, err := ownerClt.WaitMined(ctx, tx); err != nil {
		return empty, err
	}

	return address, nil
}

func gameSetup(t *testing.T) (*bank.Bank, *game.Game) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	contractID, err := deployContract()
	if err != nil {
		t.Fatal(err)
	}

	converter := currency.NewDefaultConverter(scbank.BankMetaData.ABI)

	var buf bytes.Buffer
	log := logger.New(&buf, logger.LevelInfo, "TEST", func(context.Context) string { return web.GetTraceID(ctx) })

	// -------------------------------------------------------------------------
	// Players need to deposit money into their accounts

	player1Bank, err := bank.New(ctx, log, backend, player1Clt.PrivateKey(), contractID)
	if err != nil {
		t.Fatalf("creating new bank for player 1: %s", err)
	}

	player2Bank, err := bank.New(ctx, log, backend, player2Clt.PrivateKey(), contractID)
	if err != nil {
		t.Fatalf("creating new bank for player 2: %s", err)
	}

	initalDepositGwei := converter.USD2GWei(big.NewFloat(100))
	_, _, err = player1Bank.Deposit(ctx, initalDepositGwei)
	if err != nil {
		t.Fatalf("depositing money into bank for player1: %s", err)
	}

	_, _, err = player2Bank.Deposit(ctx, initalDepositGwei)
	if err != nil {
		t.Fatalf("depositing money into bank for player2: %s", err)
	}

	// -------------------------------------------------------------------------
	// Create game and add players

	bank, err := bank.New(ctx, log, backend, ownerClt.PrivateKey(), contractID)
	if err != nil {
		t.Fatalf("creating new bank for the engine: %s", err)
	}

	// Create a game and add player1 as first player in the game.
	const anteUSD = 5.0
	game, err := game.New(ctx, log, converter, bank, player1Clt.Address(), anteUSD)
	if err != nil {
		t.Fatalf("unexpected error creating game: %s", err)
	}

	// Add player2 as the second player in the game.
	err = game.AddAccount(ctx, player2Clt.Address())
	if err != nil {
		t.Fatalf("unexpected error adding player 2: %s", err)
	}

	status := game.Info(ctx)

	if len(status.Cups) != 2 {
		t.Fatalf("expecting 2 players; got %d", len(status.Cups))
	}

	return bank, game
}
