// Utility functions for interacting with dx.cpp.

package dx

// #include <ddraw.h>
import "C"

import (
	"image/color"
	"log"

	"github.com/pkg/errors"
)

// PalFromDDP converts the given DirectDraw palette to a corresponding Go
// palette.
func PalFromDDP(ddp IDirectDrawPalette) color.Palette {
	entries := make([]C.PALETTEENTRY, 256)
	if res := ddp.GetEntries(0, entries); res != C.DD_OK {
		log.Fatalf("%+v", errors.Errorf("unable to get palette entries; %v", res))
	}
	pal := make(color.Palette, 256)
	for i, e := range entries {
		c := color.RGBA{
			R: uint8(e.peRed),
			G: uint8(e.peGreen),
			B: uint8(e.peBlue),
			A: 0xFF,
		}
		pal[i] = c
	}
	return pal
}
