package tests

import (
	"context"

	"github.com/ardanlabs/liarsdice/contract/sol/go/contract"
	"github.com/ardanlabs/liarsdice/foundation/smartcontract/smart"
)

const (
	PrimaryKeyPath    = "UTC--2022-05-12T14-47-50.112225000Z--6327a38415c53ffb36c11db55ea74cc9cb4976fd"
	PrimaryPassPhrase = "123"

	Player1Address    = "0x0070742ff6003c3e809e78d524f0fe5dcc5ba7f7"
	Player1KeyPath    = "UTC--2022-05-13T16-59-42.277071000Z--0070742ff6003c3e809e78d524f0fe5dcc5ba7f7"
	Player1PassPhrase = "123"

	Player2Address    = "0x8e113078adf6888b7ba84967f299f29aece24c55"
	Player2KeyPath    = "UTC--2022-05-13T16-57-20.203544000Z--8e113078adf6888b7ba84967f299f29aece24c55"
	Player2PassPhrase = "123"
)

func deployContract(ctx context.Context, primaryKeyPath string, primaryPassPhrase string) (*contract.Contract, error) {
	client, err := smart.Connect(ctx, smart.NetworkHTTPLocalhost, primaryKeyPath, primaryPassPhrase)
	if err != nil {
		return nil, err
	}

	const gasLimit = 3000000
	const valueGwei = 0
	tranOpts, err := client.NewTransactOpts(ctx, gasLimit, valueGwei)
	if err != nil {
		return nil, err
	}

	address, _, _, err := contract.DeployContract(tranOpts, client.ContractBackend())
	if err != nil {
		return nil, err
	}

	contract, err := contract.NewContract(address, client.ContractBackend())
	if err != nil {
		return nil, err
	}

	return contract, nil
}
