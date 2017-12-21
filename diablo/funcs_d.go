//+build djavul

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
// static void diablo_on_left_mouse_click() {
//    void (*f)() = (void *)0x4093B2;
//    f();
// }
//
// static void diablo_on_left_mouse_release() {
//    void (*f)() = (void *)0x409963;
//    f();
// }
//
// static void diablo_on_right_mouse_click() {
//    void (*f)() = (void *)0x4099A8;
//    f();
// }
//
// static void __fastcall diablo_on_key_press(int keysym) {
//    void (__fastcall *f)(int) = (void*)0x409B5C;
//    f(keysym);
// }
//
// static void __fastcall diablo_on_char_press(int keysym) {
//    void (__fastcall *f)(int) = (void*)0x409F7F;
//    f(keysym);
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

// OnLeftMouseClick is an event handler invoked on left mouse button click.
//
// ref: 0x4093B2
func OnLeftMouseClick() {
	C.diablo_on_left_mouse_click()
}

// OnLeftMouseRelease is an event handler invoked on left mouse button release.
//
// ref: 0x409963
func OnLeftMouseRelease() {
	C.diablo_on_left_mouse_release()
}

// OnRightMouseClick is an event handler invoked on right mouse button click.
//
// ref: 0x4099A8
func OnRightMouseClick() {
	C.diablo_on_right_mouse_click()
}

// OnKeyPress is an event handler invoked on key press.
//
// ref: 0x409B5C
func OnKeyPress(keysym int) {
	C.diablo_on_key_press(C.int(keysym))
}

// OnCharPress is an event handler invoked on character press.
//
// ref: 0x409F7F
func OnCharPress(keysym int) {
	C.diablo_on_char_press(C.int(keysym))
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
