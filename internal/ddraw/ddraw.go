//+build djavul

// Package ddraw provides access to DirectDraw functions.
package ddraw

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

	"github.com/pkg/errors"
)

// PALETTEENTRY represents a colour in a palette.
type PALETTEENTRY struct {
	Red   uint8
	Green uint8
	Blue  uint8
	Flags uint8
}

// IDirectDrawPalette represents a palette of colours.
type IDirectDrawPalette C.IDirectDrawPalette

// GetEntries retrieves the colours of the palette.
func (ddp *IDirectDrawPalette) GetEntries(base C.DWORD, entries []PALETTEENTRY) error {
	p := (*C.IDirectDrawPalette)(ddp)
	es := *(*[]C.PALETTEENTRY)(unsafe.Pointer(&entries))
	if res := C.DDP_GetEntries(p, 0, base, C.DWORD(len(es)), &es[0]); res != C.DD_OK {
		return errors.Errorf("unable to get palette entries; %v", res)
	}
	return nil
}

// SetEntries sets the colours of the palette.
func (ddp *IDirectDrawPalette) SetEntries(base C.DWORD, entries []PALETTEENTRY) error {
	p := (*C.IDirectDrawPalette)(ddp)
	es := *(*[]C.PALETTEENTRY)(unsafe.Pointer(&entries))
	if res := C.DDP_SetEntries(p, 0, base, C.DWORD(len(es)), &es[0]); res != C.DD_OK {
		return errors.Errorf("unable to set palette entries; %v", res)
	}
	return nil
}
