//+build djavul

// Run in djavul.exe environment.

package inv

import (
	"unsafe"

	"github.com/sanctuary/djavul/internal/types"
)

func init() {
	// Initialize pointers to global variables of djavul.exe.
	ScreenPos = (*[73]types.Point)(unsafe.Pointer(uintptr(0x47AE60)))
	StartSlot2x2 = (*[10]int32)(unsafe.Pointer(uintptr(0x48E9A8)))
	IsOpen = (*types.Bool32)(unsafe.Pointer(uintptr(0x634CB8)))
	Graphics = (**uint8)(unsafe.Pointer(uintptr(0x634CBC)))
	UpdateBelt = (*types.Bool32)(unsafe.Pointer(uintptr(0x634CC0)))
}
