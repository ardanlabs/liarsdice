package commands

import (
	"errors"
	"flag"
	"fmt"
	"os"

	"github.com/ardanlabs/ethereum"
)

const usage = `
Usage:
	admin -d
	admin -w 0x8e113078adf6888b7ba84967f299f29aece24c55
	admin -b 0x8e113078adf6888b7ba84967f299f29aece24c55 -c 0x531130464929826c57BBBF989e44085a02eeB120
	admin -a 0x8e113078adf6888b7ba84967f299f29aece24c55 -m 1000.00 -c 0x531130464929826c57BBBF989e44085a02eeB120
	admin -r 0x8e113078adf6888b7ba84967f299f29aece24c55 -c 0x531130464929826c57BBBF989e44085a02eeB120
	admin -t 0x46e40587966f02f5dff2cc63d3ff29a01e963a5360cf05094b54ad9dbc230dd3

Options:
	-d, --deploy     Deploy the smart contract.
	-b, --balance    Show the smart contract balance.
	-w, --wallet     Show the wallet balance.
	-a, --addmoney   Deposit USD into the game contract.
	-r, --rmvmoney   Withdraw money from the game contract.
	-t, --tx         Show transaction details for the specified transaction hash.
	
	-c, --contract   Provides the contract id for required calls.
	-m, --money      Sets the amount of USD to use.
	-n, --network    Sets the network to use. Default: zarf/ethereum/geth.ipc
	-f, --filekey    Sets the private key file to use for blockchain calls. Default: 0x6327a38415c53ffb36c11db55ea74cc9cb4976fd
	-p, --passphrase Sets the pass phrase for the key file. Default: 123
	-k, --keycoin    Sets the key for the coin market cap API. Default: a8cd12fb-d056-423f-877b-659046af0aa5

	-h. --help       Show the usage information.
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
	PassPhrase       string
	TranID           string
	FileKey          string
	Address          string
	ContractID       string
	CoinMarketCapKey string
	Money            float64
}

// Parse will parse the environment variables and command line flags. The command
// line flags will overwrite environment variables. Validation takes place.
func Parse() (Flags, Args, error) {
	const (
		fileKey          = "6327a38415c53ffb36c11db55ea74cc9cb4976fd"
		passPhrase       = "123"
		coinMarketCapKey = "a8cd12fb-d056-423f-877b-659046af0aa5"
	)

	flag.Usage = func() { fmt.Fprintf(os.Stderr, "%s\n", usage) }

	args := Args{
		Network:          ethereum.NetworkLocalhost,
		FileKey:          fileKey,
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
	flag.StringVar(&args.Network, "n", args.Network, "")
	flag.StringVar(&args.Network, "network", args.Network, "")
	flag.StringVar(&args.TranID, "t", args.TranID, "")
	flag.StringVar(&args.TranID, "tx", args.TranID, "")
	flag.StringVar(&args.ContractID, "c", args.ContractID, "")
	flag.StringVar(&args.ContractID, "contract", args.ContractID, "")
	flag.StringVar(&args.PassPhrase, "p", args.PassPhrase, "")
	flag.StringVar(&args.PassPhrase, "passphrase", args.PassPhrase, "")
	flag.StringVar(&args.CoinMarketCapKey, "k", args.CoinMarketCapKey, "")
	flag.StringVar(&args.CoinMarketCapKey, "keycoin", args.CoinMarketCapKey, "")
	flag.Float64Var(&args.Money, "m", args.Money, "")
	flag.Float64Var(&args.Money, "money", args.Money, "")
	flag.StringVar(&args.Address, "f", args.FileKey, "")
	flag.StringVar(&args.Address, "filekey", args.FileKey, "")

	// For add and remove money, the bank must be using the address
	// specified for the operation.
	flag.StringVar(&args.FileKey, "a", args.FileKey, "")
	flag.StringVar(&args.FileKey, "addmoney", args.FileKey, "")
	flag.StringVar(&args.FileKey, "r", args.FileKey, "")
	flag.StringVar(&args.FileKey, "rmvmoney", args.FileKey, "")

	// For the balance and wallet, the bank must be using the contract
	// owner which is the default.
	flag.StringVar(&args.Address, "b", args.Address, "")
	flag.StringVar(&args.Address, "balance", args.Address, "")
	flag.StringVar(&args.Address, "w", args.Address, "")
	flag.StringVar(&args.Address, "wallet", args.Address, "")

	flag.Bool("h", false, "")
	flag.Bool("help", false, "")
	flag.Bool("d", false, "")
	flag.Bool("deploy", false, "")

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
		if args.Address == "" {
			return errors.New("missing address")
		}
		if args.ContractID == "" {
			return errors.New("missing contract id")
		}
		return nil
	}

	if _, exists := f["a"]; exists {
		if args.ContractID == "" {
			return errors.New("missing contract id")
		}
		if args.Money <= 0.0 {
			return errors.New("incorrect amount of USD")
		}
		return nil
	}

	if _, exists := f["r"]; exists {
		if args.ContractID == "" {
			return errors.New("missing contract id")
		}
		return nil
	}

	return nil
}
