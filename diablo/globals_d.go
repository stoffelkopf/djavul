//+build djavul

// Global variable wrappers for diablo.cpp

package diablo

import (
	"unsafe"

	"github.com/AllenDang/w32"
)

// Global variables.
var (
	// address: 0x525518
	//
	// Window is the window handle of the game.
	Window = (*w32.HWND)(unsafe.Pointer(uintptr(0x525518)))

	// LightingFlag4 specifies flags used for light effects.
	//
	// PSX ref: 0x8011B797
	// PSX def: unsigned char light4flag
	//
	// ref: 0x525728
	LightingFlag4 = (*uint32)(unsafe.Pointer(uintptr(0x525728)))

	// FlagRSeed specifies the seed used for dungeon generation through the -r
	// command line flag.
	//
	// ref: 0x525738
	FlagRSeed = (*int32)(unsafe.Pointer(uintptr(0x525738)))
)
