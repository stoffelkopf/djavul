//+build djavul

// Global variable wrappers for diablo.cpp

package diablo

import "unsafe"

// Global variables.
var (
	// LightingFlag4 specifies flags used for light effects.
	//
	// PSX ref: 0x8011B797
	// PSX def: unsigned char light4flag
	//
	// ref: 0x525728
	LightingFlag4 = (*uint32)(unsafe.Pointer(uintptr(0x525728)))
)
