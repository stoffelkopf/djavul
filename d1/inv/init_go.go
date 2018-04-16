//+build !djavul

// Run natively.

package inv

import (
	"github.com/sanctuary/djavul/internal/parse"
	"github.com/sanctuary/djavul/internal/types"
)

func init() {
	// Allocate global variables in Go.
	ScreenPos = new([73]types.Point)
	StartSlot2x2 = new([10]int32)
	IsOpen = new(types.Bool32)
	Graphics = new(*uint8)
	UpdateBelt = new(types.Bool32)

	// Initialize read-only and read-write global variables by parsing
	// diablo.exe.
	parse.MustDataFromAddr(0x47AE60, ScreenPos)
	parse.MustDataFromAddr(0x48E9A8, StartSlot2x2)
}
