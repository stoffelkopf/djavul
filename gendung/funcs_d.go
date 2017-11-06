//+build djavul

package gendung

// static void gendung_init_transparency() {
// 	void (*f)() = (void *)0x41944A;
// 	f();
// }
//
// static void __fastcall gendung_make_rect_transparent(int xx_start, int yy_start, int xx_end, int yy_end) {
// 	void (__fastcall *f)(int, int, int, int) = (void *)0x419477;
// 	f(xx_start, yy_start, xx_end, yy_end);
// }
//
// static void __fastcall gendung_rect_transparent(int x_start, int y_start, int x_end, int y_end) {
// 	void (__fastcall *f)(int, int, int, int) = (void *)0x4194D0;
// 	f(x_start, y_start, x_end, y_end);
// }
//
// static void __fastcall gendung_copy_transparency(int src_x, int src_y, int dst_x, int dst_y) {
// 	void (__fastcall *f)(int, int, int, int) = (void *)0x419515;
// 	f(src_x, src_y, dst_x, dst_y);
// }
//
// static void gendung_init_set_piece() {
// 	void (*f)() = (void *)0x4195A2;
// 	f();
// }
//
// static void gendung_mark_set_piece() {
// 	void (*f)() = (void *)0x4195B9;
// 	f();
// }
import "C"

// useGo specifies whether to use the Go implementation.
const useGo = true

// InitTransparency initializes transparency.
//
// PSX ref: 0x8015A070
// PSX def: void DRLG_InitTrans__Fv()
//
// ref: 0x41944A
func InitTransparency() {
	if useGo {
		initTransparency()
	} else {
		C.gendung_init_transparency()
	}
}

// MakeRectTransparent makes the given rectangle transparent.
//
// PSX ref: 0x800578DC
// PSX def: void DRLG_MRectTrans__Fiiii(int x1, int y1, int x2, int y2)
//
// ref: 0x419477
func MakeRectTransparent(xxStart, yyStart, xxEnd, yyEnd int) {
	if useGo {
		makeRectTransparent(xxStart, yyStart, xxEnd, yyEnd)
	} else {
		C.gendung_make_rect_transparent(C.int(xxStart), C.int(yyStart), C.int(xxEnd), C.int(yyEnd))
	}
}

// RectTransparent makes the given rectangle transparent.
//
// PSX ref: 0x8015A0E4
// PSX def: void DRLG_RectTrans__Fiiii(int x1, int y1, int x2, int y2)
//
// ref: 0x4194D0
func RectTransparent(xStart, yStart, xEnd, yEnd int) {
	if useGo {
		rectTransparent(xStart, yStart, xEnd, yEnd)
	} else {
		C.gendung_rect_transparent(C.int(xStart), C.int(yStart), C.int(xEnd), C.int(yEnd))
	}
}

// CopyTransparency copies transparency from the source to the destination
// coordinate.
//
// PSX ref: 0x8015A158
// PSX def: void DRLG_CopyTrans__Fiiii(int sx, int sy, int dx, int dy)
//
// ref: 0x419515
func CopyTransparency(srcX, srcY, dstX, dstY int) {
	if useGo {
		copyTransparency(srcX, srcY, dstX, dstY)
	} else {
		C.gendung_copy_transparency(C.int(srcX), C.int(srcY), C.int(dstX), C.int(dstY))
	}
}

// InitSetPiece initializes the placement variables of the set piece (quest
// dungeon).
//
// PSX ref: 0x8015A2A4
// PSX def: void DRLG_InitSetPC__Fv()
//
// ref: 0x4195A2
func InitSetPiece() {
	if useGo {
		initSetPiece()
	} else {
		C.gendung_init_set_piece()
	}
}

// MarkSetPiece marks the area of the set piece (quest dungeon).
//
// PSX ref: 0x8015A2BC
// PSX def: void DRLG_SetPC__Fv()
//
// ref: 0x4195B9
func MarkSetPiece() {
	if useGo {
		markSetPiece()
	} else {
		C.gendung_mark_set_piece()
	}
}
