package commands

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/ardanlabs/liarsdice/foundation/smart/contract"
)

const usage = `Usage:
	verify -t 0x46e40587966f02f5dff2cc63d3ff29a01e963a5360cf05094b54ad9dbc230dd3
	verify -b 0x8e113078adf6888b7ba84967f299f29aece24c55 -c 0xE7811C584E23419e1952fa3158DEED345901bd0e

Options:
	-t, --tx       Show transaction details for the specified transaction hash.
	-b, --balance  Show the smart contract balance for the specified account.
	-c, --contract Provides the contract id for required calls.
	-n, --network  Sets the network to use. Default: zarf/ethereum/geth.ipc
`

// PrintUsage displays the usage information.
func PrintUsage(log *log.Logger) {
	log.Print(usage)
}

// =============================================================================

// flags represent the values from the command line.
type Flags struct {
	Network    string
	TX         string
	Balance    string
	ContractID string
}

// Parse will parse the environment variables and command line flags. The command
// line flags will overwrite environment variables. Validation takes place.
func Parse() (Flags, error) {
	flag.Usage = func() { fmt.Fprintf(os.Stderr, "%s\n", usage) }

	f := Flags{
		Network: contract.NetworkLocalhost,
	}
	parseCmdline(&f)

	if err := validateFlags(f); err != nil {
		return Flags{}, err
	}

	return f, nil
}

// parseCmdline will parse all the command line flags.
// The default value is set to the values parsed by the environment variables.
func parseCmdline(f *Flags) *Flags {
	flag.StringVar(&f.Network, "n", f.Network, "transaction details for the specified tx hash")
	flag.StringVar(&f.Network, "network", f.Network, "transaction details for the specified tx hash")
	flag.StringVar(&f.TX, "t", f.TX, "transaction details for the specified tx hash")
	flag.StringVar(&f.TX, "tx", f.TX, "transaction details for the specified tx hash")
	flag.StringVar(&f.Balance, "b", f.Balance, "the balance of the specified account")
	flag.StringVar(&f.Balance, "balance", f.Balance, "the balance of the specified account")
	flag.StringVar(&f.ContractID, "c", f.ContractID, "the id of the smart contract")
	flag.StringVar(&f.ContractID, "contract", f.ContractID, "the id of the smart contract")

	flag.Parse()

	return f
}

// validateFlags performs a sanity check of the provided flag information.
func validateFlags(f Flags) error {
	switch {
	case f.Balance != "":
		if f.ContractID == "" {
			return errors.New("missing contract id")
		}
	}

	return nil
}
