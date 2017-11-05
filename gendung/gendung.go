// Package gendung implements dungeon generation utility functions.
package gendung

// initTransparency initializes transparency.
//
// PSX ref: 0x8015A070
// PSX def: void DRLG_InitTrans__Fv()
//
// ref: 0x41944A
func initTransparency() {
	for x := range TransparencyMap {
		for y := range TransparencyMap {
			TransparencyMap[x][y] = 0
		}
	}
	for i := range TransparencyActive {
		TransparencyActive[i] = false
	}
	*TransparencyIndex = 1
}

// makeRectTransparent makes the given rectangle transparent.
//
// PSX ref: 0x800578DC
// PSX def: void DRLG_MRectTrans__Fiiii(int x1, int y1, int x2, int y2)
//
// ref: 0x419477
func makeRectTransparent(xxStart, yyStart, xxEnd, yyEnd int) {
	for y := 2*yyStart + 17; y < 2*yyEnd+17; y++ {
		for x := 2*xxStart + 17; x < 2*xxEnd+17; x++ {
			TransparencyMap[x][y] = *TransparencyIndex
		}
	}
	*TransparencyIndex++
}

// copyTransparency copies transparency from the source to the destination
// coordinate.
//
// PSX ref: 0x8015A158
// PSX def: void DRLG_CopyTrans__Fiiii(int sx, int sy, int dx, int dy)
//
// ref: 0x419515
func copyTransparency(srcX, srcY, dstX, dstY int) {
	TransparencyMap[dstX][dstY] = TransparencyMap[srcX][srcY]
}

// initSetPiece initializes the placement variables of the set piece (quest
// dungeon).
//
// PSX ref: 0x8015A2A4
// PSX def: void DRLG_InitSetPC__Fv()
//
// ref: 0x4195A2
func initSetPiece() {
	*SetXx = 0
	*SetYy = 0
	*SetWidth = 0
	*SetHeight = 0
}

// markSetPiece marks the area of the set piece (quest dungeon).
//
// PSX ref: 0x8015A2BC
// PSX def: void DRLG_SetPC__Fv()
//
// ref: 0x4195B9
func markSetPiece() {
	yStart := *SetYy*2 + 16
	yCount := *SetHeight * 2
	xStart := *SetXx*2 + 16
	xCount := *SetWidth * 2
	for y := yStart; y < yStart+yCount; y++ {
		for x := xStart; x < xStart+xCount; x++ {
			DFlagMap[x][y] |= DFlag08
		}
	}
}
