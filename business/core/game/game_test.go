package game_test

import (
	"context"
	"fmt"
	"math/big"
	"os"
	"runtime/debug"
	"testing"
	"time"

	"github.com/ardanlabs/ethereum"
	"github.com/ardanlabs/ethereum/currency"
	scbank "github.com/ardanlabs/liarsdice/business/contract/go/bank"
	"github.com/ardanlabs/liarsdice/business/core/bank"
	"github.com/ardanlabs/liarsdice/business/core/game"
	"github.com/ardanlabs/liarsdice/business/core/game/stores/gamedb"
	"github.com/ardanlabs/liarsdice/business/data/dbtest"
	"github.com/ardanlabs/liarsdice/foundation/docker"
	"github.com/ethereum/go-ethereum/common"
)

var (
	c          *docker.Container
	backend    *ethereum.SimulatedBackend
	ownerClt   *ethereum.Client
	player1Clt *ethereum.Client
	player2Clt *ethereum.Client
)

func TestMain(m *testing.M) {
	code, err := run(m)
	if err != nil {
		fmt.Println(err)
	}

	os.Exit(code)
}

func run(m *testing.M) (int, error) {
	var err error

	c, err = dbtest.StartDB()
	if err != nil {
		return 1, fmt.Errorf("starting database: %w", err)
	}
	defer dbtest.StopDB(c)

	backend, err = ethereum.CreateSimulatedBackend(3, true, big.NewInt(100))
	if err != nil {
		return 1, fmt.Errorf("create backend: %w", err)
	}
	defer backend.Close()

	ownerClt, err = ethereum.NewClient(backend, backend.PrivateKeys[0])
	if err != nil {
		return 1, fmt.Errorf("create ownerClt: %w", err)
	}

	player1Clt, err = ethereum.NewClient(backend, backend.PrivateKeys[1])
	if err != nil {
		return 1, fmt.Errorf("create player1Clt: %w", err)
	}

	player2Clt, err = ethereum.NewClient(backend, backend.PrivateKeys[2])
	if err != nil {
		return 1, fmt.Errorf("create player2Clt: %w", err)
	}

	return m.Run(), nil
}

func Test_SuccessGamePlay(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	converter := currency.NewDefaultConverter(scbank.BankMetaData.ABI)

	test := dbtest.NewTest(t, c, "SuccessGamePlay")
	defer func() {
		if r := recover(); r != nil {
			t.Log(r)
			t.Error(string(debug.Stack()))
		}
		test.Teardown()
	}()

	// -------------------------------------------------------------------------
	// Create game and check database

	bank, engine := gameSetup(t, test)

	checkDatabase(ctx, t, "first round", engine)

	// -------------------------------------------------------------------------
	// Define the ante for each player

	anteUSD := float64(5.0)
	anteWei := converter.USD2Wei(big.NewFloat(anteUSD))

	// -------------------------------------------------------------------------
	// Start first round

	t.Log("Start first round")

	err := engine.StartGame(ctx)
	if err != nil {
		t.Fatalf("unexpected error starting the game: %s", err)
	}

	state := engine.State()
	if state.Status != game.StatusPlaying {
		t.Fatalf("expecting game status to be %s; got %s", game.StatusPlaying, state.Status)
	}

	// -------------------------------------------------------------------------
	// Mocked roll dice so we can validate the winner and loser

	t.Log("Mocked roll dice so we can validate the winner and loser")

	player1Addr := player1Clt.Address()
	player2Addr := player2Clt.Address()

	dice := []int{6, 5, 3, 3, 3}
	engine.RollDice(ctx, player1Addr, dice...)

	dice = []int{1, 1, 4, 4, 2}
	engine.RollDice(ctx, player2Addr, dice...)

	// -------------------------------------------------------------------------
	// Game Play: Each player makes a bet and player1 calls liar

	t.Log("Game Play: Each player makes a bet and player1 calls liar")

	winnerAcct := engine.State().PlayerTurn
	if err := engine.Bet(ctx, winnerAcct, 2, 3); err != nil {
		t.Fatalf("unexpected error making bet for player1: %s", err)
	}

	loserAcct := engine.State().PlayerTurn
	if err := engine.Bet(ctx, loserAcct, 3, 4); err != nil {
		t.Fatalf("unexpected error making bet for player2: %s", err)
	}

	winner, loser, err := engine.CallLiar(ctx, engine.State().PlayerTurn)
	if err != nil {
		t.Fatalf("unexpected error calling liar for player1: %s", err)
	}

	// -------------------------------------------------------------------------
	// Check winner and loser

	t.Log("Check winner and loser")

	if winner != winnerAcct {
		t.Fatalf("expecting 'player1' to be the winner; got '%s'", winner)
	}

	if loser != loserAcct {
		t.Fatalf("expecting 'player2' to be the loser; got '%s'", loser)
	}

	state = engine.State()

	if state.Cups[loserAcct].Outs != 1 {
		t.Fatalf("expecting 'player2' to have 1 out; got %d", state.Cups[player2Addr].Outs)
	}

	if state.Status != game.StatusRoundOver {
		t.Fatalf("expecting game status to be %s; got %s", game.StatusRoundOver, state.Status)
	}

	// -------------------------------------------------------------------------
	// Check database

	checkDatabase(ctx, t, "first round ended", engine)

	// -------------------------------------------------------------------------
	// Start second round

	t.Log("Start second round")

	leftToPlay, err := engine.NextRound(ctx)
	if err != nil {
		t.Fatalf("unexpected error starting new round: %s", err)
	}

	if leftToPlay != 2 {
		t.Fatalf("expecting 2 players; got %d", leftToPlay)
	}

	state = engine.State()

	if state.Status != game.StatusPlaying {
		t.Fatalf("expecting game status to be %s; got %s", game.StatusPlaying, state.Status)
	}

	// -------------------------------------------------------------------------
	// Mocked roll dice so we can validate the winner and loser

	t.Log("Mocked roll dice so we can validate the winner and loser")

	dice = []int{1, 2, 3, 1, 6}
	engine.RollDice(ctx, player1Addr, dice...)

	dice = []int{3, 2, 6, 5, 6}
	engine.RollDice(ctx, player2Addr, dice...)

	// -------------------------------------------------------------------------
	// Game Play : Player 2 places a bet and player 1 calls liar

	t.Log("Game Play : Player 2 places a bet and player 1 calls liar")

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

	t.Log("Check winner and loser")

	if winner != winnerAcct {
		t.Fatalf("expecting 'player1' to be the winner; got '%s'", winner)
	}

	if loser != loserAcct {
		t.Fatalf("expecting 'player2' to be the loser; got '%s'", loser)
	}

	state = engine.State()

	if state.Cups[loserAcct].Outs != 2 {
		t.Fatalf("expecting 'player2' to have 2 out; got %d", state.Cups[player2Addr].Outs)
	}

	if state.Status != game.StatusRoundOver {
		t.Fatalf("expecting game status to be %s; got %s", game.StatusRoundOver, state.Status)
	}

	// -------------------------------------------------------------------------
	// Check database

	checkDatabase(ctx, t, "second round ended", engine)

	// -------------------------------------------------------------------------
	// Start third round

	t.Log("Start third round")

	leftToPlay, err = engine.NextRound(ctx)
	if err != nil {
		t.Fatalf("unexpected error starting new round: %s", err)
	}

	if leftToPlay != 2 {
		t.Fatalf("expecting 2 players; got %d", leftToPlay)
	}

	state = engine.State()

	if state.Status != game.StatusPlaying {
		t.Fatalf("expecting game status to be %s; got %s", game.StatusPlaying, state.Status)
	}

	// -------------------------------------------------------------------------
	// Mocked roll dice so we can validate the winner and loser

	t.Log("Mocked roll dice so we can validate the winner and loser")

	dice = []int{1, 1, 6, 1, 1}
	engine.RollDice(ctx, player1Clt.Address(), dice...)

	dice = []int{3, 3, 3, 5, 6}
	engine.RollDice(ctx, player2Clt.Address(), dice...)

	// -------------------------------------------------------------------------
	// Game Play : Player 2 makes a bet and player1 calls liar

	t.Log("Game Play : Player 2 makes a bet and player1 calls liar")

	err = engine.Bet(ctx, loserAcct, 4, 3)
	if err != nil {
		t.Fatalf("unexpected error making bet for player2: %s", err)
	}

	winner, loser, err = engine.CallLiar(ctx, winnerAcct)
	if err != nil {
		t.Fatalf("unexpected error calling liar for player1: %s", err)
	}

	// -------------------------------------------------------------------------
	// Check winner and loser

	t.Log("Check winner and loser")

	if winner != winnerAcct {
		t.Fatalf("expecting 'player1' to be the winner; got '%s'", winner)
	}

	if loser != loserAcct {
		t.Fatalf("expecting 'player2' to be the loser; got '%s'", loser)
	}

	state = engine.State()

	if state.Cups[loserAcct].Outs != 3 {
		t.Fatalf("expecting 'player2' to have 3 out; got %d", state.Cups[player2Addr].Outs)
	}

	if state.Status != game.StatusRoundOver {
		t.Fatalf("expecting game status to be %s; got %s", game.StatusRoundOver, state.Status)
	}

	// -------------------------------------------------------------------------
	// Check database

	checkDatabase(ctx, t, "third round ended", engine)

	// -------------------------------------------------------------------------
	// There should be only one player left, player1

	t.Log("There should be only one player left, player1")

	leftToPlay, err = engine.NextRound(ctx)
	if err != nil {
		t.Fatalf("unexpected error starting new round: %s", err)
	}

	if leftToPlay != 1 {
		t.Fatalf("expecting 1 player; got %d", leftToPlay)
	}

	state = engine.State()

	if state.Status != game.StatusGameOver {
		t.Fatalf("expecting game status to be %s; got %s", game.StatusGameOver, state.Status)
	}

	if state.PlayerLastWin != winnerAcct {
		t.Fatalf("expecting 'player1' to be the LastWinAcct; got '%s'", state.PlayerLastWin)
	}

	// -------------------------------------------------------------------------
	// Reconcile the game

	t.Log("Reconcile the game")

	if _, _, err := engine.Reconcile(ctx); err != nil {
		t.Fatalf("unexpected error reconciling the game: %s", err)
	}

	// -------------------------------------------------------------------------
	// Check balances

	t.Log("Check balances")

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

	t.Log("Validate final game state")

	state = engine.State()

	if state.Status != game.StatusReconciled {
		t.Fatalf("expecting game status to be %s; got %s", game.StatusReconciled, state.Status)
	}

	// -------------------------------------------------------------------------
	// Check database

	checkDatabase(ctx, t, "vaildate game", engine)
}

func Test_InvalidBet(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	test := dbtest.NewTest(t, c, "InvalidBet")
	defer func() {
		if r := recover(); r != nil {
			t.Log(r)
			t.Error(string(debug.Stack()))
		}
		test.Teardown()
	}()

	_, engine := gameSetup(t, test)

	player1Addr := player1Clt.Address()
	player2Addr := player2Clt.Address()

	// -------------------------------------------------------------------------
	// Start first round

	err := engine.StartGame(ctx)
	if err != nil {
		t.Fatalf("unexpected error starting the game: %s", err)
	}

	state := engine.State()

	if state.Status != game.StatusPlaying {
		t.Fatalf("expecting game status to be %s; got %s", game.StatusPlaying, state.Status)
	}

	// -------------------------------------------------------------------------
	// Mocked roll dice so we can validate the winner and loser.

	dice := []int{6, 5, 3, 3, 3}
	engine.RollDice(ctx, player1Addr, dice...)

	dice = []int{1, 1, 4, 4, 2}
	engine.RollDice(ctx, player2Addr, dice...)

	// -------------------------------------------------------------------------
	// Game Play : player 1 makes bet and player 2 makes invalid bet

	if err := engine.Bet(ctx, engine.State().PlayerTurn, 3, 3); err != nil {
		t.Fatalf("unexpected error making bet for player1: %s", err)
	}

	engine.NextTurn(ctx)

	if err := engine.Bet(ctx, engine.State().PlayerTurn, 2, 6); err == nil {
		t.Fatal("expecting error making an invalid bet")
	}
}

func Test_WrongPlayerTryingToPlay(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	test := dbtest.NewTest(t, c, "WrongPlayerTryingToPlay")
	defer func() {
		if r := recover(); r != nil {
			t.Log(r)
			t.Error(string(debug.Stack()))
		}
		test.Teardown()
	}()

	_, engine := gameSetup(t, test)

	player1Addr := player1Clt.Address()
	player2Addr := player2Clt.Address()

	// -------------------------------------------------------------------------
	// Start first round

	err := engine.StartGame(ctx)
	if err != nil {
		t.Fatalf("unexpected error starting the game: %s", err)
	}

	state := engine.State()

	if state.Status != game.StatusPlaying {
		t.Fatalf("expecting game status to be %s; got %s", game.StatusPlaying, state.Status)
	}

	var wrongPlayer common.Address
	switch engine.State().PlayerTurn {
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
	contractID := deployContract(t)

	converter := currency.NewDefaultConverter(scbank.BankMetaData.ABI)

	test := dbtest.NewTest(t, c, "GameWithoutEnoughPlayers")
	defer func() {
		if r := recover(); r != nil {
			t.Log(r)
			t.Error(string(debug.Stack()))
		}
		test.Teardown()
	}()

	store := gamedb.NewStore(test.Log, test.DB)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// -------------------------------------------------------------------------
	// Players need to deposit money into their accounts

	player1Bank, err := bank.New(ctx, test.Log, backend, player1Clt.PrivateKey(), contractID)
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

	bank, err := bank.New(ctx, test.Log, backend, ownerClt.PrivateKey(), contractID)
	if err != nil {
		t.Fatalf("creating new bank for the engine: %s", err)
	}

	const anteUSD = 5.0
	game, err := game.New(ctx, test.Log, converter, store, bank, player1Clt.Address(), anteUSD)
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

	contractID := deployContract(t)

	converter := currency.NewDefaultConverter(scbank.BankMetaData.ABI)

	test := dbtest.NewTest(t, c, "NewGameNotEnoughBalance")
	defer func() {
		if r := recover(); r != nil {
			t.Log(r)
			t.Error(string(debug.Stack()))
		}
		test.Teardown()
	}()

	store := gamedb.NewStore(test.Log, test.DB)

	// -------------------------------------------------------------------------
	// Players need to deposit money into their accounts

	bank, err := bank.New(ctx, test.Log, backend, player1Clt.PrivateKey(), contractID)
	if err != nil {
		t.Fatalf("creating new bank for player 1: %s", err)
	}

	const anteUSD = 5.0
	_, err = game.New(ctx, test.Log, converter, store, bank, player1Clt.Address(), anteUSD)
	if err == nil {
		t.Fatalf("expecting an error creating a game: %s", err)
	}
}

// =============================================================================

func gameSetup(t *testing.T, test *dbtest.Test) (*bank.Bank, *game.Game) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	contractID := deployContract(t)

	converter := currency.NewDefaultConverter(scbank.BankMetaData.ABI)

	store := gamedb.NewStore(test.Log, test.DB)

	// -------------------------------------------------------------------------
	// Players need to deposit money into their accounts

	player1Bank, err := bank.New(ctx, test.Log, backend, player1Clt.PrivateKey(), contractID)
	if err != nil {
		t.Fatalf("creating new bank for player 1: %s", err)
	}

	player2Bank, err := bank.New(ctx, test.Log, backend, player2Clt.PrivateKey(), contractID)
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

	bank, err := bank.New(ctx, test.Log, backend, ownerClt.PrivateKey(), contractID)
	if err != nil {
		t.Fatalf("creating new bank for the engine: %s", err)
	}

	// Create a game and add player1 as first player in the game.
	const anteUSD = 5.0
	game, err := game.New(ctx, test.Log, converter, store, bank, player1Clt.Address(), anteUSD)
	if err != nil {
		t.Fatalf("unexpected error creating game: %s", err)
	}

	// Add player2 as the second player in the game.
	err = game.AddAccount(ctx, player2Clt.Address())
	if err != nil {
		t.Fatalf("unexpected error adding player 2: %s", err)
	}

	state := game.State()

	if len(state.Cups) != 2 {
		t.Fatalf("expecting 2 players; got %d", len(state.Cups))
	}

	return bank, game
}

func deployContract(t *testing.T) common.Address {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	t.Log("Deploying Contract ...")
	defer t.Log("Deployed")

	return smartContract(ctx, t)
}

func smartContract(ctx context.Context, t *testing.T) common.Address {
	tranOpts, err := ownerClt.NewTransactOpts(ctx, 10_000_000, big.NewInt(0), big.NewFloat(0))
	if err != nil {
		t.Fatal("ownerclt newtransactopts:", err)
	}

	address, tx, _, err := scbank.DeployBank(tranOpts, ownerClt.Backend)
	if err != nil {
		t.Fatal("scbank deploybank:", err)
	}

	if _, err := ownerClt.WaitMined(ctx, tx); err != nil {
		t.Fatal("ownerClt waitmined:", err)
	}

	return address
}

func checkDatabase(ctx context.Context, t *testing.T, label string, engine *game.Game) {
	t.Logf("%s: Check Database", label)

	qState, err := engine.QueryState(ctx)
	if err != nil {
		t.Fatalf("%s: unable to retrieve game state from storage: %s", label, err)
	}

	eState := engine.State()

	if qState.Status != eState.Status {
		t.Fatalf("%s: expecting game status in storage to be %s; got %s", label, eState.Status, qState.Status)
	}

	if qState.Round != eState.Round {
		t.Fatalf("%s: expecting game round in storage to be %d; got %d", label, eState.Round, qState.Round)
	}

	if qState.PlayerLastOut != eState.PlayerLastOut {
		t.Fatalf("%s: expecting game player last out in storage to be %s; got %s", label, eState.PlayerLastOut, qState.PlayerLastOut)
	}

	if qState.PlayerLastWin != eState.PlayerLastWin {
		t.Fatalf("%s: expecting game player last win in storage to be %s; got %s", label, eState.PlayerLastWin, qState.PlayerLastWin)
	}

	if qState.PlayerTurn != eState.PlayerTurn {
		t.Fatalf("%s: expecting game player turn in storage to be %s; got %s", label, eState.PlayerTurn, qState.PlayerTurn)
	}

	// if qState.ExistingPlayers != eState.PlayerTurn {
	// 	t.Fatalf("expecting game player turn in storage to be %s; got %s", label, eState.PlayerTurn, qState.PlayerTurn)
	// }
}
