// Global variable wrappers for dx.cpp

package dx

// #include <ddraw.h>
//
// HRESULT DDP_GetEntries(IDirectDrawPalette *ddp, DWORD flags, DWORD base, DWORD len, PALETTEENTRY *entries) {
//    return ddp->lpVtbl->GetEntries(ddp, flags, base, len, entries);
// }
//
// HRESULT DDP_SetEntries(IDirectDrawPalette *ddp, DWORD flags, DWORD base, DWORD len, PALETTEENTRY *entries) {
//    return ddp->lpVtbl->SetEntries(ddp, flags, base, len, entries);
// }
import "C"

import (
	"unsafe"
)

// Global variables.
var (
	// DDP represents the DirectDraw palette of the system.
	//
	// ref: 0x52A51C
	DDP = (*IDirectDrawPalette)(unsafe.Pointer(uintptr(0x52A51C)))

	// ScreenBuf contains the pixels of the screen.
	//
	// ref: 0x52A524
	ScreenBuf = (**Screen)(unsafe.Pointer(uintptr(0x52A524)))
)

// IDirectDrawPalette represents a palette of colours.
type IDirectDrawPalette struct {
	p *C.IDirectDrawPalette
}

// GetEntries retrieves the colours of the palette.
func (ddp IDirectDrawPalette) GetEntries(base C.DWORD, entries []C.PALETTEENTRY) C.HRESULT {
	return C.DDP_GetEntries(ddp.p, 0, base, C.DWORD(len(entries)), &entries[0])
}

// SetEntries sets the colours of the palette.
func (ddp IDirectDrawPalette) SetEntries(base C.DWORD, entries []C.PALETTEENTRY) C.HRESULT {
	return C.DDP_SetEntries(ddp.p, 0, base, C.DWORD(len(entries)), &entries[0])
}

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
