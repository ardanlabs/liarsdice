package commands

import (
	"errors"
	"flag"
	"fmt"
	"os"

	"github.com/ardanlabs/liarsdice/foundation/smart/contract"
)

const usage = `
Usage:
	admin -d -c 0xE7811C584E23419e1952fa3158DEED345901bd0e
	admin -a 1000.00 -c 0xE7811C584E23419e1952fa3158DEED345901bd0e
	admin -b 0x8e113078adf6888b7ba84967f299f29aece24c55 -c 0xE7811C584E23419e1952fa3158DEED345901bd0e
	admin -t 0x46e40587966f02f5dff2cc63d3ff29a01e963a5360cf05094b54ad9dbc230dd3

Options:
	-d, --deploy     Deploy the smart contract.
	-b, --balance    Show the smart contract balance for the specified account.
	-w, --wallet     Show the wallet balance for the keyfile account.
	-a, --addmoney   Deposit USD into the game contract.
	-r, --rmvmoney   Withdraw money from the game contract.
	-t, --tx         Show transaction details for the specified transaction hash.
	-h. --help       Show the usage information.

	-c, --contract   Provides the contract id for required calls.
	-n, --network    Sets the network to use. Default: zarf/ethereum/geth.ipc
	-f, --keyfile    Sets the path to the key file. Default: zarf/ethereum/keystore/...6327a38415c53ffb36c11db55ea74cc9cb4976fd
	-p, --passphrase Sets the pass phrase for the key file. Default: 123
	-k, --coinkey    Sets the key for the coin market cap API. Default: a8cd12fb-d056-423f-877b-659046af0aa5
`

// PrintUsage displays the usage information.
func PrintUsage() {
	fmt.Print(usage)
}

// =============================================================================

// Flags represents the flags that were provided.
type Flags map[string]struct{}

// Args represent the values for each of the specified flags.
type Args struct {
	Network          string
	KeyFile          string
	PassPhrase       string
	Hex              string
	Address          string
	ContractID       string
	CoinMarketCapKey string
	Amount           float64
}

// Parse will parse the environment variables and command line flags. The command
// line flags will overwrite environment variables. Validation takes place.
func Parse() (Flags, Args, error) {
	const (
		keyFile          = "zarf/ethereum/keystore/UTC--2022-05-12T14-47-50.112225000Z--6327a38415c53ffb36c11db55ea74cc9cb4976fd"
		passPhrase       = "123"
		coinMarketCapKey = "a8cd12fb-d056-423f-877b-659046af0aa5"
	)

	flag.Usage = func() { fmt.Fprintf(os.Stderr, "%s\n", usage) }

	args := Args{
		Network:          contract.NetworkLocalhost,
		KeyFile:          keyFile,
		PassPhrase:       passPhrase,
		ContractID:       os.Getenv("GAME_CONTRACT_ID"),
		CoinMarketCapKey: coinMarketCapKey,
	}
	flags := parseCmdline(&args)

	if err := validate(flags, args); err != nil {
		return nil, Args{}, err
	}

	return flags, args, nil
}

// parseCmdline will parse all the command line flags.
// The default value is set to the values parsed by the environment variables.
func parseCmdline(args *Args) Flags {
	flag.StringVar(&args.Network, "n", args.Network, "transaction details for the specified tx hash")
	flag.StringVar(&args.Network, "network", args.Network, "transaction details for the specified tx hash")
	flag.StringVar(&args.Hex, "t", args.Hex, "transaction details for the specified tx hash")
	flag.StringVar(&args.Hex, "tx", args.Hex, "transaction details for the specified tx hash")
	flag.StringVar(&args.Address, "b", args.Address, "balance of the specified account")
	flag.StringVar(&args.Address, "balance", args.Address, "balance of the specified account")
	flag.StringVar(&args.ContractID, "c", args.ContractID, "id of the smart contract")
	flag.StringVar(&args.ContractID, "contract", args.ContractID, "id of the smart contract")
	flag.StringVar(&args.KeyFile, "k", args.KeyFile, "path to the key file")
	flag.StringVar(&args.KeyFile, "keyfile", args.KeyFile, "path to the key file")
	flag.StringVar(&args.PassPhrase, "p", args.PassPhrase, "pass phrase for the key file")
	flag.StringVar(&args.PassPhrase, "passphrase", args.PassPhrase, "pass phrase for the key file")
	flag.StringVar(&args.CoinMarketCapKey, "m", args.CoinMarketCapKey, "key for the coin market cap api")
	flag.StringVar(&args.CoinMarketCapKey, "market", args.CoinMarketCapKey, "key for the coin market cap api")
	flag.Float64Var(&args.Amount, "a", args.Amount, "deposit money into the game contract")
	flag.Float64Var(&args.Amount, "addmoney", args.Amount, "deposit money into the game contract")

	flag.Bool("h", false, "show help usage")
	flag.Bool("help", false, "show help usage")
	flag.Bool("d", false, "deploy the smart contract")
	flag.Bool("deploy", false, "deploy the smart contract")
	flag.Bool("w", false, "show the wallet balance")
	flag.Bool("wallet", false, "show the wallet balance")
	flag.Bool("r", false, "withdraw money from the game contract")
	flag.Bool("rmvmoney", false, "withdraw money from the game contract")

	flag.Parse()

	flags := Flags{}
	flag.Visit(func(f *flag.Flag) {
		flags[f.Name[:1]] = struct{}{}
	})

	return flags
}

// validate performs a sanity check of the provided flag information.
func validate(f Flags, args Args) error {
	if _, exists := f["b"]; exists {
		if args.ContractID == "" {
			return errors.New("missing contract id")
		}
		return nil
	}

	return nil
}
