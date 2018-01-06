//+build djavul

package world

import "unsafe"

var (
	// --- [ .data section ] ----------------------------------------------------

	// TileDrawMasks specifies the draw masks used to render transparency of
	// tiles.
	//
	// ref: 0x4B327D
	TileDrawMasks = (*[96]uint32)(unsafe.Pointer(uintptr(0x4B327D)))
)
