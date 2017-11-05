//+build djavul

// Global variable wrappers for engine.cpp

package engine

import "unsafe"

// Global variables.
var (
	// InitialSeed represents the initial global seed of the game.
	//
	// ref: 0x52B974
	InitialSeed = (*int32)(unsafe.Pointer(uintptr(0x52B974)))

	// Seed represents the global seed of the game.
	//
	// PSX ref: 0x8011C7C4
	// PSX def: long sglGameSeed
	//
	// ref: 0x52B97C
	Seed = (*int32)(unsafe.Pointer(uintptr(0x52B97C)))

	// SeedCount specifies the number of invokations to engine_rand.
	//
	// PSX ref: 0x8011B85C
	// PSX def: int SeedCount
	//
	// ref: 0x52B998
	SeedCount = (*int32)(unsafe.Pointer(uintptr(0x52B998)))
)
