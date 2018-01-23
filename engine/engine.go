// Package engine implements image decoding, PRNG and memory management utility
// functions.
package engine

// #include <stdint.h>
// #include <stdlib.h>
// #include <string.h>
//
// uint8_t * copy(uint8_t *src, int n) {
//    uint8_t *dst = malloc(n);
//    memcpy(dst, src, n);
//    return dst;
// }
import "C"

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
	"unsafe"

	"github.com/mewkiz/pkg/term"
	"github.com/pkg/errors"
)

var (
	// dbg represents a logger with the "engine:" prefix, which logs debug
	// messages to standard error.
	dbg = log.New(os.Stderr, term.BlueBold("engine:")+" ", 0)
	// warn represents a logger with the "engine:" prefix, which logs warnings to
	// standard error.
	warn = log.New(os.Stderr, term.RedBold("engine:")+" ", 0)
)

// setSeed sets the global seed to s.
//
// PSX ref: 0x8003DACC
// PSX def: void SetRndSeed__Fl(long s)
//
// ref: 0x417518
func setSeed(s int32) {
	*SeedCount = 0
	*Seed = s
	*InitialSeed = s
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
func memLoadFile(path unsafe.Pointer, size *int32) unsafe.Pointer {
	p := absPath(goPath(path))
	dbg.Println("engine.MemLoadFile:", p)
	buf, err := ioutil.ReadFile(p)
	if err != nil {
		log.Fatalf("unable to load file %q; %v", p, errors.WithStack(err))
	}
	if size != nil {
		*size = int32(len(buf))
	}
	return unsafe.Pointer(C.copy((*C.uint8_t)(unsafe.Pointer(&buf[0])), C.int(len(buf))))
}

// ### [ Helper functions ] ####################################################

// files maps from file contents pointer to file path.
var files = make(map[unsafe.Pointer]string)

// getFile returns the file path of the given file contents pointer.
func getFile(addr unsafe.Pointer) string {
	file, ok := files[addr]
	if !ok {
		panic(fmt.Errorf("unable to locate file path for address 0x%08X", uintptr(addr)))
	}
	return file
}

// absPath returns the absolute path to the given file, relative to the MPQ
// directory.
func absPath(relPath string) string {
	// mpqDir specifies a directory containing an extracted copy of the files
	// contained within DIABDAT.MPQ. Note that the extracted files should have
	// lowercase names.
	const mpqDir = "diabdat"
	return filepath.Join(mpqDir, relPath)
}

// goPath returns an equivalent Go string of the given file path.
func goPath(path unsafe.Pointer) string {
	p := C.GoString((*C.char)(path))
	p = strings.Replace(p, "\\", "/", -1)
	return strings.ToLower(p)
}

// abs returns the absolute value of v.
func abs(v int32) int32 {
	// TODO: Figure out how to handle the most negative value.
	if v < 0 {
		return -v
	}
	return v
}
