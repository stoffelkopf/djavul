//+build !djavul

// Global variable wrappers for lighting.cpp

package lighting

// Global variables.
var (
	// Disabled specifies whether light effects are disabled.
	//
	// ref: 0x646A28
	Disabled = new(bool)

	// Max specifies the maximum light effects.
	//
	// PSX ref: 0x8011B918
	// PSX def: char lightmax
	//
	// ref: 0x642A14
	Max = new(int8)
)
