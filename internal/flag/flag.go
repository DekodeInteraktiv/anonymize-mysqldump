package flag

import (
	"flag"
	"fmt"
	"os"

	"github.com/spf13/pflag"
)

const (
	helpText = `Anonymize MySQLDump is a database anonymization tool.

Usage:
  anonymize-mysqldump [flags]

Flags:
  --help -h      Outputs help text and exits.
  --config -c    The path to a custom config file.

Config:
  The anonymizer will use a default config suitable for WordPress, but you can override this by providing your own.`
)

func Parse() *string {
	pflag.CommandLine.ParseErrorsWhitelist.UnknownFlags = true
	pflag.CommandLine.AddGoFlagSet(flag.CommandLine)

	flagConfigFile := pflag.String("config", "", "Path to config file.")
	flagHelp := pflag.BoolP("help", "h", false, "")

	pflag.Parse()

	if *flagHelp {
		fmt.Println(helpText)
		os.Exit(1)
	}

	return flagConfigFile
}
