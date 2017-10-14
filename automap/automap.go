// Package automap provides access to a mini-map of the game.
package automap

import (
	"unsafe"
)

// Global variables.
var (
	// Discovered tracks the explored areas of the map.
	//
	// ref: 0x4B7E6C
	Discovered = (*[40][40]bool)(unsafe.Pointer(uintptr(0x4B7E6C)))
)
