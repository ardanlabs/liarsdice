package commands

import (
	"errors"
	"flag"
	"fmt"
	"os"

	"github.com/ardanlabs/liarsdice/foundation/smart/contract"
)

const usage = `Usage:
	admin -d 
	admin -t 0x46e40587966f02f5dff2cc63d3ff29a01e963a5360cf05094b54ad9dbc230dd3
	admin -b 0x8e113078adf6888b7ba84967f299f29aece24c55 -c 0xE7811C584E23419e1952fa3158DEED345901bd0e

Options:
	-d, --deploy     Deploy the smart contract.
	-t, --tx         Show transaction details for the specified transaction hash.
	-b, --balance    Show the smart contract balance for the specified account.
	-w, --wallet     Show the wallet balance for the keyfile account.
	-x, --xdraw      Withdraw money from the contract to the keyfile account.

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

// Values represent the values for each of the specified flags.
type Values struct {
	Network          string
	KeyFile          string
	PassPhrase       string
	Hex              string
	Address          string
	ContractID       string
	CoinMarketCapKey string
}

// Parse will parse the environment variables and command line flags. The command
// line flags will overwrite environment variables. Validation takes place.
func Parse() (Flags, Values, error) {
	const (
		keyFile          = "zarf/ethereum/keystore/UTC--2022-05-12T14-47-50.112225000Z--6327a38415c53ffb36c11db55ea74cc9cb4976fd"
		passPhrase       = "123"
		coinMarketCapKey = "a8cd12fb-d056-423f-877b-659046af0aa5"
	)

	flag.Usage = func() { fmt.Fprintf(os.Stderr, "%s\n", usage) }

	v := Values{
		Network:          contract.NetworkLocalhost,
		KeyFile:          keyFile,
		PassPhrase:       passPhrase,
		CoinMarketCapKey: coinMarketCapKey,
	}
	flags := parseCmdline(&v)

	if err := validate(flags, v); err != nil {
		return nil, Values{}, err
	}

	return flags, v, nil
}

// parseCmdline will parse all the command line flags.
// The default value is set to the values parsed by the environment variables.
func parseCmdline(v *Values) Flags {
	flag.StringVar(&v.Network, "n", v.Network, "transaction details for the specified tx hash")
	flag.StringVar(&v.Network, "network", v.Network, "transaction details for the specified tx hash")
	flag.StringVar(&v.Hex, "t", v.Hex, "transaction details for the specified tx hash")
	flag.StringVar(&v.Hex, "tx", v.Hex, "transaction details for the specified tx hash")
	flag.StringVar(&v.Address, "b", v.Address, "balance of the specified account")
	flag.StringVar(&v.Address, "balance", v.Address, "balance of the specified account")
	flag.StringVar(&v.ContractID, "c", v.ContractID, "id of the smart contract")
	flag.StringVar(&v.ContractID, "contract", v.ContractID, "id of the smart contract")
	flag.StringVar(&v.KeyFile, "k", v.KeyFile, "path to the key file")
	flag.StringVar(&v.KeyFile, "keyfile", v.KeyFile, "path to the key file")
	flag.StringVar(&v.PassPhrase, "p", v.PassPhrase, "pass phrase for the key file")
	flag.StringVar(&v.PassPhrase, "passphrase", v.PassPhrase, "pass phrase for the key file")
	flag.StringVar(&v.CoinMarketCapKey, "m", v.CoinMarketCapKey, "key for the coin market cap api")
	flag.StringVar(&v.CoinMarketCapKey, "market", v.CoinMarketCapKey, "key for the coin market cap api")

	flag.Bool("d", false, "deploy the smart contract")
	flag.Bool("deploy", false, "deploy the smart contract")
	flag.Bool("w", false, "show the wallet balance")
	flag.Bool("wallet", false, "show the wallet balance")
	flag.Bool("x", false, "withdraw money from the contract")
	flag.Bool("xdraw", false, "withdraw money from the contract")

	flag.Parse()

	flags := Flags{}
	flag.Visit(func(f *flag.Flag) {
		flags[f.Name[:1]] = struct{}{}
	})

	return flags
}

// validate performs a sanity check of the provided flag information.
func validate(f Flags, v Values) error {
	if _, exists := f["b"]; exists {
		if v.ContractID == "" {
			return errors.New("missing contract id")
		}
		return nil
	}

	return nil
}
