package progress

import (
	"os"
	"time"

	"github.com/briandowns/spinner"
)

var s *spinner.Spinner

// Start the progress.
func Start() {
	s = spinner.New(spinner.CharSets[35], 250*time.Millisecond, spinner.WithWriter(os.Stderr))
	s.Suffix = " Anonymizing data..."
	s.Start()
}

// Stop the progress.
func Stop() {
	s.Stop()
}
