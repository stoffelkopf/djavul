//+build !djavul

package gendung

// InitTransparency initializes transparency.
//
// PSX ref: 0x8015A070
// PSX def: void DRLG_InitTrans__Fv()
//
// ref: 0x41944A
func InitTransparency() {
	initTransparency()
}

// MakeRectTransparent makes the given rectangle transparent.
//
// PSX ref: 0x800578DC
// PSX def: void DRLG_MRectTrans__Fiiii(int x1, int y1, int x2, int y2)
//
// ref: 0x419477
func MakeRectTransparent(xxStart, yyStart, xxEnd, yyEnd int) {
	makeRectTransparent(xxStart, yyStart, xxEnd, yyEnd)
}

// RectTransparent makes the given rectangle transparent.
//
// PSX ref: 0x8015A0E4
// PSX def: void DRLG_RectTrans__Fiiii(int x1, int y1, int x2, int y2)
//
// ref: 0x4194D0
func RectTransparent(xStart, yStart, xEnd, yEnd int) {
	rectTransparent(xStart, yStart, xEnd, yEnd)
}

// CopyTransparency copies transparency from the source to the destination
// coordinate.
//
// PSX ref: 0x8015A158
// PSX def: void DRLG_CopyTrans__Fiiii(int sx, int sy, int dx, int dy)
//
// ref: 0x419515
func CopyTransparency(srcX, srcY, dstX, dstY int) {
	copyTransparency(srcX, srcY, dstX, dstY)
}

// InitSetPiece initializes the placement variables of the set piece (quest
// dungeon).
//
// PSX ref: 0x8015A2A4
// PSX def: void DRLG_InitSetPC__Fv()
//
// ref: 0x4195A2
func InitSetPiece() {
	initSetPiece()
}

// MarkSetPiece marks the area of the set piece (quest dungeon).
//
// PSX ref: 0x8015A2BC
// PSX def: void DRLG_SetPC__Fv()
//
// ref: 0x4195B9
func MarkSetPiece() {
	markSetPiece()
}
