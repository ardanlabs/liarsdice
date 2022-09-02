package main

import (
	"fmt"
	"os"

	"github.com/ardanlabs/liarsdice/app/cli/liars/board"
	"github.com/ardanlabs/liarsdice/app/cli/liars/client"
)

const (
	keyFile    = "zarf/ethereum/keystore/UTC--2022-05-12T14-47-50.112225000Z--6327a38415c53ffb36c11db55ea74cc9cb4976fd"
	passPhrase = "123"
)

func main() {
	if err := run(); err != nil {
		fmt.Println("ERROR:", err)
		os.Exit(1)
	}
}

func run() error {
	client := client.New("http://0.0.0.0:3000")
	token, err := client.Connect(keyFile, passPhrase)
	if err != nil {
		return fmt.Errorf("connect to game engine: %w", err)
	}

	board, err := initalizeBoard(client, token)
	if err != nil {
		return err
	}
	defer board.Shutdown()

	// board.SetAnte(5.0)

	// board.AddPlayer("0x6327A38415C53FFb36c11db55Ea74cc9cB4976Fd", "$1,000.00")
	// board.AddPlayer("0x8e113078adf6888b7ba84967f299f29aece24c55", "$235.65")
	// board.AddPlayer("0x0070742ff6003c3e809e78d524f0fe5dcc5ba7f7", "$12,765.44")

	// board.ActivePlayer("0x6327A38415C53FFb36c11db55Ea74cc9cB4976Fd")

	// if err := board.SetDice([]int{1, 3, 2, 5, 5}); err != nil {
	// 	return err
	// }

	<-board.StartEventLoop()
	return nil
}

// initalizeBoard draws the board with the configuation.
func initalizeBoard(client *client.Client, token client.Token) (*board.Board, error) {
	config, err := client.Configuration()
	if err != nil {
		return nil, fmt.Errorf("get game configuration: %w", err)
	}

	board, err := board.New()
	if err != nil {
		return nil, err
	}

	board.Init()
	board.SetSettings(config.Network, config.ChainID, config.ContractID, token.Address)

	return board, nil
}
