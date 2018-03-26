// Package sound implements sound playback.
package sound

import (
	"log"
	"os"

	"github.com/mewkiz/pkg/term"
)

var (
	// dbg represents a logger with the "sound:" prefix, which logs debug
	// messages to standard error.
	dbg = log.New(os.Stderr, term.BlueBold("sound:")+" ", 0)
	// warn represents a logger with the "sound:" prefix, which logs warnings to
	// standard error.
	warn = log.New(os.Stderr, term.RedBold("sound:")+" ", 0)
)
