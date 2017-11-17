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
	"path/filepath"
	"strings"
	"unsafe"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/pkg/errors"
)

// Win is the Pixel window handler.
var Win *pixelgl.Window

// celDecodeFrame decodes the given frame to the specified screen coordinate.
//
//    x = screenX - 64
//    y = screenY - 160
//    frameNum = frame - 1
//
// Note, the coordinates specify the bottom left corner.
//
// ref: 0x416274
func celDecodeFrame(screenX, screenY int, celBuf unsafe.Pointer, frame, frameWidth int) {
	file := getFile(celBuf)
	pics := getPictures(file)
	frameNum := frame - 1
	pic := pics[frameNum]
	sprite := pixel.NewSprite(pic, pic.Bounds())
	const screenHeight = 480
	x := float64(screenX - 64)
	y := screenHeight - float64(screenY-160) - 1
	sprite.Draw(Win, pixel.IM.Moved(pic.Bounds().Center().Add(pixel.V(x, y))))
}

// celDecodeFrameWithHeader decodes the given CEL frame to the specified screen
// coordinate.
//
//    x = screen_x - 64
//    y = screen_y - 160
//    frameNum = frame - 1
//
// Note, the coordinates specify the bottom left corner.
//
// Note, this function is only used to decode CEL images with frame headers
// (objects, item drops, objcurs.cel, towners).
//
// ref: 0x4162DE
func celDecodeFrameWithHeader(screenX, screenY int, celBuf unsafe.Pointer, frame, frameWidth, always0, direction int) {
	// TODO: Handle light tables for celDecodeFrameWithHeader.
	celDecodeFrame(screenX, screenY, celBuf, frame, frameWidth)
}

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
func memLoadFile(path unsafe.Pointer, size *int32) unsafe.Pointer {
	p := absPath(goPath(path))
	fmt.Println("engine.MemLoadFile:", p)
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

// pictures maps from relative file path to decoded image frames.
var pictures = make(map[string][]pixel.Picture)

// getPictures returns the pictures associated with the given file path.
func getPictures(relPath string) []pixel.Picture {
	pics, ok := pictures[relPath]
	if !ok {
		panic(fmt.Errorf("unable to locate decoded image frames of %q", relPath))
	}
	return pics
}

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

// abs returns the absolute value of x.
func abs(x int32) int32 {
	// TODO: Figure out how to handle the most negative value.
	if x < 0 {
		return -x
	}
	return x
}
