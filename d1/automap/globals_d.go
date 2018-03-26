//+build djavul

// Global variable wrappers for automap.cpp

package automap

import "unsafe"

// Global variables.
var (
	// Enabled specifies whether the automap is enabled.
	//
	// ref: 0x4B7E48
	Enabled = (*bool)(unsafe.Pointer(uintptr(0x4B7E48)))

	// Discovered tracks the explored areas of the map.
	//
	// ref: 0x4B7E6C
	Discovered = (*[40][40]bool)(unsafe.Pointer(uintptr(0x4B7E6C)))
)
