// Global variable wrappers for drlg_l1.cpp

package l1

import "unsafe"

// Global variables.
var (
	// FlagMap contains flags used for dungeon generation of the Cathedral.
	//
	// PSX ref: 0x8011C0D8
	// PSX def: unsigned char* mydflags
	//
	// ref: 0x527064
	FlagMap = (*[40][40]uint8)(unsafe.Pointer(uintptr(0x527064)))
)
