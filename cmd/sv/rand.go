package main

// seed represents the global seed of the MSVCR71 library.
var seed int32

// setSeed sets the global seed of the MSVCR71 library to s.
//
// ref: 0x46A520
func setSeed(s int32) {
	seed = s
}

// rand returns a non-negative pseudo-random integer in [0, 2^15), using the
// MSVCR71 pseudo-random number generator algorithm with a multiplier of 0x343FD
// and an increment of 2531011.
//
// ref: 0x46A530
func rand() int32 {
	seed = seed*0x343FD + 2531011
	return (seed >> 16) & 0x7FFF
}
