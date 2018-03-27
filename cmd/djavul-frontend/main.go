// The djavul tool is a graphical front-end to the Diablo 1 game engine.
package main

import (
	"flag"
	"log"
	"os"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/mewkiz/pkg/term"
	"github.com/pkg/errors"
	"github.com/sanctuary/djavul/internal/proto"
)

var (
	// dbg represents a logger with the "djavul-frontend:" prefix, which logs
	// debug messages to standard error.
	dbg = log.New(os.Stderr, term.BlueBold("djavul-frontend:")+" ", 0)
	// warn represents a logger with the "djavul-frontend:" prefix, which logs
	// warnings to standard error.
	warn = log.New(os.Stderr, term.RedBold("djavul-frontend:")+" ", 0)
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
	// Parse command line arguments.
	var (
		// npipe specifies whether to use named pipes for IPC.
		npipe bool
	)
	flag.BoolVar(&npipe, "npipe", false, "use named pipes for IPC")
	flag.Parse()

	// Create window.
	icons := loadIcon()
	cfg := pixelgl.WindowConfig{
		Title:       "Djavul",
		Icon:        icons,
		Bounds:      pixel.R(0, 0, 640, 480),
		Undecorated: true,
		VSync:       true,
	}
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		return errors.WithStack(err)
	}

	// Create handler for stable and unstable IPC connections.
	var (
		stable   proto.IPC
		unstable proto.IPC
	)
	if npipe {
		stable = proto.NewStableNamedPipe(".")
		unstable = proto.NewUnstableNamedPipe(".")
	} else {
		stable = proto.NewStableTCP("localhost")
		unstable = proto.NewUnstableTCP("localhost")
	}

	// Initiate the event loop.
	loop(win, stable, unstable)

	return nil
}

// ### [ Helper functions ] ####################################################

// die kills the application when a fatal error has occurred.
func die(err error) {
	log.Fatalf("%+v", errors.WithStack(err))
}
