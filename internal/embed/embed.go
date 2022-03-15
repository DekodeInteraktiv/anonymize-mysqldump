package embed

import (
	"embed"
)

//go:embed files/*
// Content contains a filesystem of config files.
var Content embed.FS
