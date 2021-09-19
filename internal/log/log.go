package log

import (
	"bytes"
	"log"
)

var buf *bytes.Buffer

// New returns the configured default logger.
func New() *log.Logger {
	buf = &bytes.Buffer{}
	log.SetPrefix("")
	log.SetFlags(log.Ldate | log.Ltime)
	log.SetOutput(buf)

	return log.Default()
}
