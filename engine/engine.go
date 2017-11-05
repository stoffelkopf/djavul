// Package engine implements image decoding, PRNG and memory management utility
// functions.
package engine

import (
	"io/ioutil"
	"log"
	"path/filepath"
	"strings"
	"unsafe"

	"github.com/pkg/errors"
)

// setSeed sets the global seed to x.
//
// PSX ref: 0x8003DACC
// PSX def: void SetRndSeed__Fl(long s)
//
// ref: 0x417518
func setSeed(x int32) {
	*SeedCount = 0
	*Seed = x
	*InitialSeed = x
}

// rand returns a non-negative pseudo-random integer in [0, 2^31), using the
// Borland C/C++ pseudo-random number generator algorithm with a multiplier of
// 0x15A4E35 and an increment of 1.
//
// PSX ref: 0x8003DADC
// PSX def: long GetRndSeed__Fv()
//
// ref: 0x41752C
func rand() int32 {
	*SeedCount++
	*Seed = *Seed*0x15A4E35 + 1
	return abs(*Seed)
}

// randCap returns a capped non-negative pseudo-random integer in [0, max),
// using the Borland C/C++ pseudo-random number generator algorithm with a
// multiplier of 0x15A4E35 and an increment of 1.
//
// PSX ref: 0x8003DB24
// PSX def: long ENG_random__Fl(long v)
//
// ref: 0x41754B
func randCap(max int32) int32 {
	if max <= 0 {
		return 0
	}
	x := Rand()
	if max < 0xFFFF {
		x >>= 16
	}
	return x % max
}

// memLoadFile returns the contents of the given file.
//
// PSX ref: 0x80074E9C
// PSX def: unsigned char* GRL_LoadFileInMemSig__FPCcPUl(char *Name, unsigned long *Len)
//
// ref: 0x417618
func memLoadFile(path string, size *int32) *uint8 {
	// mpqDir specifies a directory containing an extracted copy of the files
	// contained within DIABDAT.MPQ. Note that the extracted files should have
	// lowercase names.
	const mpqDir = "diabdat"
	path = filepath.Join(mpqDir, strings.ToLower(path))
	buf, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatalf("unable to load file %q; %v", path, errors.WithStack(err))
	}
	if size != nil {
		*size = int32(len(buf))
	}
	return (*uint8)(unsafe.Pointer(&buf[0]))
}

// ### [ Helper functions ] ####################################################

// abs returns the absolute value of x.
func abs(x int32) int32 {
	// TODO: Figure out how to handle the most negative value.
	if x < 0 {
		return -x
	}
	return x
}
