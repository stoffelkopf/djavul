//+build !djavul

package engine

import "C"

import (
	"unsafe"
)

// SetSeed sets the global seed to x.
//
// PSX ref: 0x8003DACC
// PSX def: void SetRndSeed__Fl(long s)
//
// ref: 0x417518
func SetSeed(x int32) {
	setSeed(x)
}

// Rand returns a non-negative pseudo-random integer in [0, 2^31), using the
// Borland C/C++ pseudo-random number generator algorithm with a multiplier of
// 0x15A4E35 and an increment of 1.
//
// PSX ref: 0x8003DADC
// PSX def: long GetRndSeed__Fv()
//
// ref: 0x41752C
func Rand() int32 {
	return rand()
}

// RandCap returns a capped non-negative pseudo-random integer in [0, max),
// using the Borland C/C++ pseudo-random number generator algorithm with a
// multiplier of 0x15A4E35 and an increment of 1.
//
// PSX ref: 0x8003DB24
// PSX def: long ENG_random__Fl(long v)
//
// ref: 0x41754B
func RandCap(unused, max int32) int32 {
	return randCap(max)
}

// MemFree frees the given memory space.
//
// PSX ref: 0x8003DBDC
// PSX def: void mem_free_dbg__FPv(void *p)
//
// ref: 0x4175E8
func MemFree(ptr unsafe.Pointer) {
	// NOTE: Being garbage collected, MemFree is a no-op in Go.
}

// MemLoadFile returns the contents of the given file.
//
// PSX ref: 0x80074E9C
// PSX def: unsigned char* GRL_LoadFileInMemSig__FPCcPUl(char *Name, unsigned long *Len)
//
// ref: 0x417618
func MemLoadFile(path unsafe.Pointer, size *int32) unsafe.Pointer {
	addr := memLoadFile(path, size)
	files[addr] = goPath(path)
	return addr
}
