// Package settings provides support for capturing settings.
package settings

import (
	"flag"
	"fmt"
	"os"
)

const usage = `
Usage:
	liars
	liars -a 0x8e113078adf6888b7ba84967f299f29aece24c55
	liars -e http://0.0.0.0:3000 -a 0x0070742ff6003c3e809e78d524f0fe5dcc5ba7f7

Options:
	-e, --engine     The url of the game engine. Default: http://0.0.0.0:3000
	-a, --account    The players account id. Default: 0x0070742ff6003c3e809e78d524f0fe5dcc5ba7f7
`

// PrintUsage displays the usage information.
func PrintUsage() {
	fmt.Print(usage)
}

// Flags represents the flags that were provided.
type Flags map[string]struct{}

// Args represents the values provided in the command line arguments.
type Args struct {
	Engine    string
	AccountID string
}

// Parse will parse the command line flags. The command line flags will overwrite
// the defaults.
func Parse() (Flags, Args, error) {
	const (
		engine    = "http://localhost:3000"
		accountID = "0x0070742ff6003c3e809e78d524f0fe5dcc5ba7f7"
	)

	flag.Usage = func() { fmt.Fprintf(os.Stderr, "%s\n", usage) }

	args := Args{
		Engine:    engine,
		AccountID: accountID,
	}
	flags := parseCmdline(&args)

	return flags, args, nil
}

// parseCmdline will parse all the command line flags.
// The default value is set to the values parsed by the environment variables.
func parseCmdline(args *Args) Flags {
	flag.StringVar(&args.Engine, "e", args.Engine, "")
	flag.StringVar(&args.Engine, "engine", args.Engine, "")
	flag.StringVar(&args.AccountID, "a", args.AccountID, "")
	flag.StringVar(&args.AccountID, "account", args.AccountID, "")

	flag.Bool("h", false, "show help usage")
	flag.Bool("help", false, "show help usage")

	flag.Parse()

	flags := Flags{}
	flag.Visit(func(f *flag.Flag) {
		flags[f.Name[:1]] = struct{}{}
	})

	return flags
}
