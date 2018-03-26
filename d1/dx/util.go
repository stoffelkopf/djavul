//+build djavul

// Utility functions for interacting with dx.cpp.

package dx

import (
	"image/color"

	"github.com/sanctuary/djavul/internal/ddraw"
)

// PalFromEntries converts the given DirectDraw palette colours to a
// corresponding Go palette.
func PalFromEntries(entries []ddraw.PALETTEENTRY) color.Palette {
	pal := make(color.Palette, len(entries))
	for i, e := range entries {
		c := color.RGBA{
			R: uint8(e.Red),
			G: uint8(e.Green),
			B: uint8(e.Blue),
			A: 0xFF,
		}
		pal[i] = c
	}
	return pal
}
