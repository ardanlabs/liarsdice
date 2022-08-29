package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

const usage = `Usage:
	verify [--encrypt] (-r round)... [--armor] [-o OUTPUT] [INPUT]
	tle --decrypt [-o OUTPUT] [INPUT]

Options:
	-t, --txhash   Show transaction details for the specified transaction hash.
	-b, --balance  Show the smart contract balance for the specified account.
`

// PrintUsage displays the usage information.
func PrintUsage(log *log.Logger) {
	log.Print(usage)
}

// =============================================================================

// flags represent the values from the command line.
type Flags struct {
	TXHash  string
	Balance string
}

// Parse will parse the environment variables and command line flags. The command
// line flags will overwrite environment variables. Validation takes place.
func Parse() (Flags, error) {
	flag.Usage = func() { fmt.Fprintf(os.Stderr, "%s\n", usage) }

	var f Flags
	parseCmdline(&f)

	if err := validateFlags(f); err != nil {
		return Flags{}, err
	}

	return f, nil
}

// parseCmdline will parse all the command line flags.
// The default value is set to the values parsed by the environment variables.
func parseCmdline(f *Flags) *Flags {
	flag.StringVar(&f.TXHash, "t", f.TXHash, "transaction details for the specified tx hash")
	flag.StringVar(&f.TXHash, "txhash", f.TXHash, "transaction details for the specified tx hash")
	flag.StringVar(&f.Balance, "b", f.Balance, "the balance of the specified account")
	flag.StringVar(&f.Balance, "balance", f.Balance, "the balance of the specified account")

	flag.Parse()

	return f
}

// validateFlags performs a sanity check of the provided flag information.
func validateFlags(f Flags) error {
	return nil
}
