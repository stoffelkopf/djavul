// Global variables of track.cpp.

package track

import "github.com/sanctuary/djavul/internal/types"

// --- [ .bss section ] --------------------------------------------------------

// Uninitialized global variables.
var (
	// WalkTrackingActive specifies whether continued walk actions are tracked
	// (i.e. when the left mouse button is held down).
	//
	// ref: 0x6ABAB8
	WalkTrackingActive *types.Bool8

	// Time specifies the time when walk tracking was activated.
	//
	// ref: 0x6ABAC0
	Time *types.Time32

	// PrevActive specifies whether walk tracking was active in the previous game
	// step.
	//
	// ref: 0x6ABAC4
	PrevActive *types.Bool32
)
