//+build !djavul

// Run natively.

package track

import (
	"github.com/sanctuary/djavul/internal/types"
)

func init() {
	// Allocate global variables in Go.
	WalkTrackingActive = new(types.Bool8)
	Time = new(types.Time32)
	PrevActive = new(types.Bool32)
}
