//+build !djavul

// Run natively.

package inv

import (
	"github.com/sanctuary/djavul/internal/types"
)

func init() {
	// Allocate global variables in Go.
	ScreenPos = new([73]types.Point)
}
