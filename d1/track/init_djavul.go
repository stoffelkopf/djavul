//+build djavul

// Run in djavul.exe environment.

package track

import (
	"unsafe"

	"github.com/sanctuary/djavul/internal/types"
)

func init() {
	// Initialize pointers to global variables of djavul.exe.
	WalkTrackingActive = (*types.Bool8)(unsafe.Pointer(uintptr(0x6ABAB8)))
	Time = (*types.Time32)(unsafe.Pointer(uintptr(0x6ABAC0)))
	PrevAactive = (*types.Bool32)(unsafe.Pointer(uintptr(0x6ABAC4)))
}
