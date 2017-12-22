//+build djavul

package dx

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
