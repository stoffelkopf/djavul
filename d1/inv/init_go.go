//+build !djavul

// Run natively.

package inv

import (
	"github.com/sanctuary/djavul/internal/types"
)

func init() {
	// Allocate global variables in Go.
	ScreenPos = new([73]types.Point)
	StartSlot2x2 = new([10]int32)
	IsOpen = new(types.Bool32)
	Graphics = new(*uint8)
	UpdateBelt = new(types.Bool32)
}
