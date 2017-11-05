//+build ddraw

// Global variable wrappers for dx.cpp

package dx

import (
	"unsafe"

	"github.com/sanctuary/djavul/internal/ddraw"
)

// Global variables.
var (
	// DDP represents the DirectDraw palette of the system.
	//
	// ref: 0x52A51C
	DDP = (**ddraw.IDirectDrawPalette)(unsafe.Pointer(uintptr(0x52A51C)))

	// ScreenBuf contains the pixels of the screen.
	//
	// ref: 0x52A524
	ScreenBuf = (**Screen)(unsafe.Pointer(uintptr(0x52A524)))
)
