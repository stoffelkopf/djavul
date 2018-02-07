// Package player implements player handling.
package player

import (
	"log"
	"os"

	"github.com/mewkiz/pkg/term"
)

var (
	// dbg represents a logger with the "player:" prefix, which logs debug
	// messages to standard error.
	dbg = log.New(os.Stderr, term.BlueBold("player:")+" ", 0)
	// warn represents a logger with the "player:" prefix, which logs warnings to
	// standard error.
	warn = log.New(os.Stderr, term.RedBold("player:")+" ", 0)
)
