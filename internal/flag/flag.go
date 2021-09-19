package flag

import (
	"fmt"
	"os"

	"github.com/spf13/pflag"
)

func Parse(version, commit, date string) *string {
	helpText := `Anonymize MySQLDump is a database anonymization tool.
Version: ` + version + `
Commit: ` + commit + `
Date: ` + date + `

Usage:
  anonymize-mysqldump [flags]

Flags:
  --help -h      Outputs help text and exits.
  --config -c    The path to a custom config file.

Config:
  The anonymizer will use a default config suitable for WordPress, but you can override this by providing your own.`

	flagConfigFile := pflag.String("config", "", "Path to config file.")
	flagHelp := pflag.BoolP("help", "h", false, "")

	pflag.Parse()

	if *flagHelp {
		fmt.Println(helpText)
		os.Exit(1)
	}

	return flagConfigFile
}
