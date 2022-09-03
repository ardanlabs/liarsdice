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
	token, err := eng.Connect(keyStorePath, args.AccountID, passPhrase)
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

	teardown, err := eng.Events(board.Events)
	if err != nil {
		return err
	}
	defer teardown()

	// =========================================================================
	// Game and player setup.

	status, err := eng.QueryStatus()
	switch {
	case err != nil:

		// No game exists yet so create one.
		status, err = eng.NewGame()
		if err != nil {
			return err
		}

	default:

		// See if the account already joined the game.
		var found bool
		for _, accountID := range status.CupsOrder {
			if strings.EqualFold(accountID, args.AccountID) {
				found = true
				break
			}
		}

		// If the address is not in the game yet, join.
		if !found {
			status, err = eng.JoinGame()
			if err != nil {
				return err
			}
		}
	}

	board.PrintStatus(status)

	// =========================================================================
	// Start the game loop.

	<-board.StartEventLoop()
	return nil
}

// initalizeBoard draws the board with the configuation.
func initalizeBoard(engine *engine.Engine, token engine.Token) (*board.Board, error) {
	config, err := engine.Configuration()
	if err != nil {
		return nil, fmt.Errorf("get game configuration: %w", err)
	}

	board, err := board.New(engine, token.Address)
	if err != nil {
		return nil, err
	}

	board.Init()
	board.PrintSettings(engine.URL(), config.Network, config.ChainID, config.ContractID, token.Address)

	return board, nil
}
