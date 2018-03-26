//+build !djavul

// Global variable wrappers for automap.cpp

package automap

// Global variables.
var (
	// Enabled specifies whether the automap is enabled.
	//
	// ref: 0x4B7E48
	Enabled = new(bool)

	// Discovered tracks the explored areas of the map.
	//
	// ref: 0x4B7E6C
	Discovered = new([40][40]bool)
)
