package diablo

// #include <stdio.h>
// #include <windows.h>
//
// static void diablo_WinMain(HINSTANCE hInstance, HINSTANCE hPrevInstance, char *szCmdLine, int nShowCmd) {
//    int (*f)(HINSTANCE, HINSTANCE, char *, int) = (void*)0x408B4A;
//    int status = f(hInstance, hPrevInstance, szCmdLine, nShowCmd);
//    printf("exit status: %d\n", status);
//    exit(status);
// }
//
// static void diablo_load_level_graphics() {
//    void (*f)() = (void *)0x40A391;
//    f();
// }
import "C"

import (
	"unsafe"

	"github.com/AllenDang/w32"
)

// WinMain is the main function of the game.
//
// ref: 0x408B4A
func WinMain(inst, prev w32.HINSTANCE, args string, showCmd int) {
	i := *(*C.HINSTANCE)(unsafe.Pointer(&inst))
	p := *(*C.HINSTANCE)(unsafe.Pointer(&prev))
	C.diablo_WinMain(i, p, C.CString(args), C.int(showCmd))
}

// LoadLevelGraphics loads the tile graphics of the active dungeon type.
//
// PSX ref: 0x80038930
// PSX def: void LoadLvlGFX__Fv()
//
// ref: 0x40A391
func LoadLevelGraphics() {
	C.diablo_load_level_graphics()
}
