package embed

import (
	_ "embed"
)

//go:embed files/config.default.json
// DefaultConfig contains the contents of the default config file.
var DefaultConfig string
