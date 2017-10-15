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

// Screen represents the pixels of the screen.
//
// size = 0x7B000
type Screen struct {
	// offset 00000000 (122880 bytes)
	_ [160]ScreenRow
	// offset 0001E000 (368640 bytes)
	Row [480]ScreenRow
	// offset 00078000 (12288 bytes)
	_ [16]ScreenRow
}

// ScreenRow represents a single horizontal line of pixels on the screen.
//
// size = 0x300
type ScreenRow struct {
	// offset 0000 (64 bytes)
	_ [64]uint8
	// offset 0040 (640 bytes)
	Pixels [640]uint8
	// offset 02C0 (64 bytes)
	_ [64]uint8
}
