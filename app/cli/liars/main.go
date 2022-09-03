package main

import (
	"fmt"
	"os"

	"github.com/ardanlabs/liarsdice/app/cli/liars/board"
	"github.com/ardanlabs/liarsdice/app/cli/liars/client"
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

	client := client.New(args.Engine)
	token, err := client.Connect(keyStorePath, args.Address, passPhrase)
	if err != nil {
		return fmt.Errorf("connect to game engine: %w", err)
	}

	// =========================================================================
	// Initialize the board and display the configuration and token information.

	board, err := initalizeBoard(args.Engine, client, token)
	if err != nil {
		return err
	}
	defer board.Shutdown()

	// =========================================================================
	// Establish a websocket connection to capture the game events.

	events := func(event string, address string) {
		message := fmt.Sprintf("type: %s  addr: %s", event, board.FmtAddress(address))
		board.PrintMessage(message)
	}
	teardown, err := client.Events(events)
	if err != nil {
		return err
	}
	defer teardown()

	// =========================================================================
	// Get the current game status

	status, err := client.Status()
	if err != nil {

		// No Game exists so let's create a game.
		status, err = client.NewGame()
		if err != nil {
			return err
		}
	}

	board.SetStatus(status)

	balance, err := client.Balance()
	if err != nil {
		return err
	}
	board.AddPlayer(token.Address, balance)

	// board.ActivePlayer("0x6327A38415C53FFb36c11db55Ea74cc9cB4976Fd")

	// if err := board.SetDice([]int{1, 3, 2, 5, 5}); err != nil {
	// 	return err
	// }

	<-board.StartEventLoop()
	return nil
}

// initalizeBoard draws the board with the configuation.
func initalizeBoard(engine string, client *client.Client, token client.Token) (*board.Board, error) {
	config, err := client.Configuration()
	if err != nil {
		return nil, fmt.Errorf("get game configuration: %w", err)
	}

	board, err := board.New()
	if err != nil {
		return nil, err
	}

	board.Init()
	board.SetSettings(engine, config.Network, config.ChainID, config.ContractID, token.Address)

	return board, nil
}
