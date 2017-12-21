// The djavul tool is a graphical front-end to the Diablo 1 game engine.
package main

import (
	"log"
	"os"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/mewkiz/pkg/term"
	"github.com/pkg/errors"
)

var (
	// dbg represents a logger with the "djavul:" prefix, which logs debug
	// messages to standard error.
	dbg = log.New(os.Stderr, term.BlueBold("djavul:")+" ", 0)
	// warn represents a logger with the "djavul:" prefix, which logs warnings to
	// standard error.
	warn = log.New(os.Stderr, term.RedBold("djavul:")+" ", 0)
)

func main() {
	pixelgl.Run(run)
}

// run is invoked from the main thread by PixelGL.
func run() {
	if err := front(); err != nil {
		die(err)
	}
}

// front initiates a graphical front-end to the Diablo 1 game engine.
func front() error {
	// Create window.
	icons := loadIcon()
	cfg := pixelgl.WindowConfig{
		Title:  "Djavul",
		Icon:   icons,
		Bounds: pixel.R(0, 0, 640, 480),
		VSync:  true,
	}
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		return errors.WithStack(err)
	}
	// Initiate the event loop.
	loop(win)
	return nil
}

// ### [ Helper functions ] ####################################################

// die kills the application when a fatal error has occurred.
func die(err error) {
	log.Fatalf("%+v", errors.WithStack(err))
}
