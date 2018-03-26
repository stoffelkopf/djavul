// Package multi implements multiplayer functions.
package multi

import (
	"log"
	"os"

	"github.com/mewkiz/pkg/term"
)

var (
	// dbg represents a logger with the "multi:" prefix, which logs debug
	// messages to standard error.
	dbg = log.New(os.Stderr, term.BlueBold("multi:")+" ", 0)
	// warn represents a logger with the "multi:" prefix, which logs warnings to
	// standard error.
	warn = log.New(os.Stderr, term.RedBold("multi:")+" ", 0)
)
