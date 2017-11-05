//+build djavul

// Global variable wrappers for multi.cpp

package multi

import "unsafe"

// Global variables.
var (
	// MaxPlayers specifies the maximum number of players in a game, where 1
	// represents a single player game and 4 represents a multi player game.
	//
	// ref: 0x679660
	MaxPlayers = (*uint8)(unsafe.Pointer(uintptr(0x679660)))
)
