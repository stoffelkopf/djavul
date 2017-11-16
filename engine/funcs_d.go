//+build djavul

package engine

// #include <stdio.h>
// #include <stdint.h>
//
// static void __fastcall engine_cel_decode_frame(int screen_x, int screen_y, uint8_t *cel_buf, int frame, int frame_width) {
//    void (__fastcall *f)(int, int, uint8_t *, int, int) = (void *)0x416274;
//    f(screen_x, screen_y, cel_buf, frame, frame_width);
// }
//
// static void __fastcall engine_set_seed(int32_t x) {
//    void (__fastcall *f)(int32_t) = (void *)0x417518;
//    f(x);
// }
//
// static int32_t engine_rand() {
//    int32_t (*f)() = (void *)0x41752C;
//    return f();
// }
//
// static int32_t __fastcall engine_rand_cap(int unused, int32_t max) {
//    int32_t (__fastcall *f)(int, int32_t) = (void *)0x41754B;
//    return f(unused, max);
// }
//
// static void __fastcall engine_mem_free(void *ptr) {
//    void (__fastcall *f)(void *) = (void *)0x4175E8;
//    f(ptr);
// }
//
// static void * __fastcall engine_mem_load_file(char *file_path, int *size) {
//    void * (__fastcall *f)(char *, int *) = (void *)0x417618;
//    return f(file_path, size);
// }
import "C"

import (
	"fmt"
	"log"
	"path/filepath"
	"unsafe"

	"github.com/faiface/pixel"
	"github.com/pkg/errors"
	"github.com/sanctuary/formats/image/cel"
	"github.com/sanctuary/formats/image/cel/config"
)

const (
	// useGo specifies whether to use the Go implementation.
	useGo = true
	// UseGUI specifies whether to use the Go GUI implementation.
	UseGUI = true
)

// CelDecodeFrame decodes the given frame to the specified screen coordinate.
//
//    x = screenX - 64
//    y = screenY - 160
//    frameNum = frame - 1
//
// ref: 0x416274
func CelDecodeFrame(screenX, screenY int, celBuf unsafe.Pointer, frame, frameWidth int) {
	if UseGUI {
		celDecodeFrame(screenX, screenY, celBuf, frame, frameWidth)
	}
	C.engine_cel_decode_frame(C.int(screenX), C.int(screenY), (*C.uint8_t)(celBuf), C.int(frame), C.int(frameWidth))
}

// SetSeed sets the global seed to x.
//
// PSX ref: 0x8003DACC
// PSX def: void SetRndSeed__Fl(long s)
//
// ref: 0x417518
func SetSeed(x int32) {
	if useGo {
		setSeed(x)
	} else {
		C.engine_set_seed(C.int32_t(x))
	}
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
	if useGo {
		return rand()
	} else {
		return int32(C.engine_rand())
	}
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
	if useGo {
		return randCap(max)
	} else {
		return int32(C.engine_rand_cap(C.int(unused), C.int32_t(max)))
	}
}

// MemFree frees the given memory space.
//
// PSX ref: 0x8003DBDC
// PSX def: void mem_free_dbg__FPv(void *p)
//
// ref: 0x4175E8
func MemFree(ptr unsafe.Pointer) {
	C.engine_mem_free(ptr)
}

// MemLoadFile returns the contents of the given file.
//
// PSX ref: 0x80074E9C
// PSX def: unsigned char* GRL_LoadFileInMemSig__FPCcPUl(char *Name, unsigned long *Len)
//
// ref: 0x417618
func MemLoadFile(path unsafe.Pointer, size *int32) unsafe.Pointer {
	var addr unsafe.Pointer
	if useGo {
		addr = memLoadFile(path, size)
	} else {
		buf := C.engine_mem_load_file((*C.char)(path), (*C.int32_t)(unsafe.Pointer(size)))
		addr = unsafe.Pointer(buf)
	}
	file := goPath(path)
	files[addr] = file
	if UseGUI {
		switch filepath.Ext(file) {
		case ".cel":
			if err := loadPics(file); err != nil {
				log.Fatalf("+%v", err)
			}
		}
	}
	return addr
}

// ### [ Helper functions ] ####################################################

// loadPics loads the frames of the given CEL image.
func loadPics(relPath string) error {
	name := filepath.Base(relPath)
	conf, err := config.Get(name)
	if err != nil {
		return errors.Errorf("unable to locate image config for %q; %v", name, err)
	}
	fmt.Println("decoding CEL image:", relPath)
	palPath := "levels/towndata/town.pal"
	if len(conf.Pals) > 0 {
		// TODO: Figure out how to handle multiple palettes.
		palPath = conf.Pals[0]
	}
	pal, err := cel.ParsePal(absPath(palPath))
	if err != nil {
		return errors.Errorf("unable to parse palette %q; %v", palPath, err)
	}
	imgs, err := cel.DecodeAll(absPath(relPath), pal)
	if err != nil {
		return errors.Errorf("unable to decode CEL image %q; %v", relPath, err)
	}
	var pics []pixel.Picture
	for _, img := range imgs {
		pic := pixel.PictureDataFromImage(img)
		pics = append(pics, pic)
	}
	pictures[relPath] = pics
	return nil
}
