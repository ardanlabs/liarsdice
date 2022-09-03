package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/ardanlabs/liarsdice/app/cli/liars/board"
	"github.com/ardanlabs/liarsdice/app/cli/liars/engine"
	"github.com/ardanlabs/liarsdice/app/cli/liars/settings"
)

const (
	keyStorePath = "zarf/ethereum/keystore/"
	passPhrase   = "123"
)

func main() {
	if err := run(); err != nil {
		fmt.Println("ERROR:", err)
		os.Exit(1)
	}
}

func run() error {

	// =========================================================================
	// Parse flags for settings.

	flags, args, err := settings.Parse()
	if err != nil {
		return fmt.Errorf("parsing arguments: %w", err)
	}

	if _, exists := flags["h"]; exists {
		settings.PrintUsage()
		return nil
	}

	// =========================================================================
	// Establish a client connection to the game engine.

	eng := engine.New(args.Engine)
	token, err := eng.Connect(keyStorePath, args.Address, passPhrase)
	if err != nil {
		return fmt.Errorf("connect to game engine: %w", err)
	}

	// =========================================================================
	// Initialize the board and display the configuration and token information.

	board, err := initalizeBoard(eng, token)
	if err != nil {
		return err
	}
	defer board.Shutdown()

	// =========================================================================
	// Establish a websocket connection to capture the game events.

	events := func(event string, address string) {
		message := fmt.Sprintf("type: %s  addr: %s", event, board.FmtAddress(address))
		board.PrintMessage(message)

		switch event {
		case "join":
			status, err := eng.QueryStatus()
			if err != nil {
				return
			}
			board.PrintStatus(status)
		}
	}
	teardown, err := eng.Events(events)
	if err != nil {
		return err
	}
	defer teardown()

	// =========================================================================
	// Start or join the game.

	status, err := eng.QueryStatus()
	if err != nil {
		status, err = eng.NewGame()
		if err != nil {
			return err
		}
	}

	var found bool
	for _, address := range status.CupsOrder {
		if strings.EqualFold(address, args.Address) {
			found = true
			break
		}
	}

	if !found {
		status, err = eng.JoinGame()
		if err != nil {
			return err
		}
	}
	board.PrintStatus(status)

	// balance, err := eng.Balance()
	// if err != nil {
	// 	return err
	// }
	// board.AddPlayer(token.Address, balance)

	// board.ActivePlayer("0x6327A38415C53FFb36c11db55Ea74cc9cB4976Fd")

	// if err := board.SetDice([]int{1, 3, 2, 5, 5}); err != nil {
	// 	return err
	// }

	<-board.StartEventLoop()
	return nil
}

// initalizeBoard draws the board with the configuation.
func initalizeBoard(engine *engine.Engine, token engine.Token) (*board.Board, error) {
	config, err := engine.Configuration()
	if err != nil {
		return nil, fmt.Errorf("get game configuration: %w", err)
	}

	board, err := board.New()
	if err != nil {
		return nil, err
	}

	board.Init()
	board.SetSettings(engine.URL(), config.Network, config.ChainID, config.ContractID, token.Address)

	return board, nil
}
