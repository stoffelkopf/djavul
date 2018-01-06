//+build djavul

// Global variable wrappers for lighting.cpp

package lighting

import "unsafe"

// Global variables.
var (
	// Disabled specifies whether light effects are disabled.
	//
	// ref: 0x646A28
	Disabled = (*bool)(unsafe.Pointer(uintptr(0x646A28)))

	// Max specifies the maximum number of light effects.
	//
	// PSX ref: 0x8011B918
	// PSX def: char lightmax
	//
	// ref: 0x642A14
	Max = (*int8)(unsafe.Pointer(uintptr(0x642A14)))
)
