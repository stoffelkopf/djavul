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
	PrevActive = (*types.Bool32)(unsafe.Pointer(uintptr(0x6ABAC4)))

	// Initialize functions.
	//process = cProcess
	//setWalkTracking = cSetWalkTracking
	//isWalkTrackingActive = cIsWalkTrackingActive
}
