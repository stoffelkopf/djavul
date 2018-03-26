//+build !djavul

// Global variable wrappers for diablo.cpp

package diablo

// Global variables.
var (
	// LightingFlag4 specifies flags used for light effects.
	//
	// PSX ref: 0x8011B797
	// PSX def: unsigned char light4flag
	//
	// ref: 0x525728
	LightingFlag4 = new(uint32)
)
