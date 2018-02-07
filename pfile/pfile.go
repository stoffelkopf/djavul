// Package pfile implements save file handling.
package pfile

import (
	"log"
	"os"

	"github.com/mewkiz/pkg/term"
)

var (
	// dbg represents a logger with the "pfile:" prefix, which logs debug
	// messages to standard error.
	dbg = log.New(os.Stderr, term.BlueBold("pfile:")+" ", 0)
	// warn represents a logger with the "pfile:" prefix, which logs warnings to
	// standard error.
	warn = log.New(os.Stderr, term.RedBold("pfile:")+" ", 0)
)
