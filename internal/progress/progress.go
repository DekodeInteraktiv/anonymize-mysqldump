package progress

import (
	"os"
	"time"

	"github.com/briandowns/spinner"
)

var s *spinner.Spinner

// Start progress bar.
func Start() {
	s = spinner.New(spinner.CharSets[35], 100*time.Millisecond, spinner.WithWriter(os.Stderr))
	s.Suffix = " Anonymizing data..."
	s.Start()
}

func Stop() {
	s.Stop()
}
