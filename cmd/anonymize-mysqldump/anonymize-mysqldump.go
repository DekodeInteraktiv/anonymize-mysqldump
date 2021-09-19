package main

import (
	"github.com/DekodeInteraktiv/anonymize-mysqldump/internal/anonymize"
)

var (
	version string
	commit  string
	date    string
)

func main() {
	anonymize.Start(version, commit, date)
}
