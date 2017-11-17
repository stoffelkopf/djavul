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
// static void __fastcall engine_cel_decode_frame_into_buf(uint8_t *dst_buf, uint8_t *cel_buf, int frame, int frame_width) {
//    void (__fastcall *f)(uint8_t *, uint8_t *, int, int) = (void *)0x4162B8;
//    f(dst_buf, cel_buf, frame, frame_width);
// }
//
// static void __fastcall engine_cel_decode_frame_with_header(int screen_x, int screen_y, uint8_t *cel_buf, int frame, int frame_width, int always_0, int direction) {
//    void (__fastcall *f)(int, int, uint8_t *, int, int, int, int) = (void *)0x4162DE;
//    f(screen_x, screen_y, cel_buf, frame, frame_width, always_0, direction);
// }
//
// static void __fastcall engine_cel_decode_frame_with_header_into_buf(uint8_t *dst_buf, uint8_t *cel_buf, int frame, int frame_width, int always_0, int direction) {
//    void (__fastcall *f)(uint8_t *, uint8_t *, int, int, int, int) = (void *)0x416359;
//    f(dst_buf, cel_buf, frame, frame_width, always_0, direction);
// }
//
// static void __fastcall engine_cel_decode_frame_with_light(int screen_x, int screen_y, uint8_t *cel_buf, int frame, int frame_width) {
//    void (__fastcall *f)(int, int, uint8_t *, int, int) = (void *)0x416565;
//    f(screen_x, screen_y, cel_buf, frame, frame_width);
// }
//
// static void __fastcall engine_cel_decode_frame_with_header_and_light(int screen_x, int screen_y, uint8_t *cel_buf, int frame, int frame_width, int always_0, int direction) {
//    void (__fastcall *f)(int, int, uint8_t *, int, int, int, int) = (void *)0x4165BD;
//    f(screen_x, screen_y, cel_buf, frame, frame_width, always_0, direction);
// }
//
// static void __fastcall engine_cel_decode_frame_with_header_light_and_transparency_into_buf(uint8_t *dst_buf, uint8_t *cel_buf, int frame, int frame_width, int always_0, int direction) {
//    void (__fastcall *f)(uint8_t *, uint8_t *, int, int, int, int) = (void *)0x41664B;
//    f(dst_buf, cel_buf, frame, frame_width, always_0, direction);
// }
//
// static void __fastcall engine_cel_decode_frame_with_header_and_light_not_equipable(int screen_x, int screen_y, uint8_t *cel_buf, int frame, int frame_width, int always_0, int direction, int8_t always_1) {
//    void (__fastcall *f)(int, int, uint8_t *, int, int, int, int, int8_t) = (void *)0x4166BF;
//    f(screen_x, screen_y, cel_buf, frame, frame_width, always_0, direction, always_1);
// }
//
// static void __fastcall engine_cel_decode_frame_with_header2(int screen_x, int screen_y, uint8_t *cel_buf, int frame, int frame_width, int a6, int direction) {
//    void (__fastcall *f)(int, int, uint8_t *, int, int, int, int) = (void *)0x41685A;
//    f(screen_x, screen_y, cel_buf, frame, frame_width, a6, direction);
// }
//
// static void __fastcall engine_cel_decode_frame_with_header_into_buf2(uint8_t *dst_buf, uint8_t *cel_buf, int frame, int frame_width, int a5, int direction) {
//    void (__fastcall *f)(uint8_t *, uint8_t *, int, int, int, int) = (void *)0x4168D5;
//    f(dst_buf, cel_buf, frame, frame_width, a5, direction);
// }
//
// static void __fastcall engine_cel_decode_frame_with_header_and_light2(int screen_x, int screen_y, uint8_t *cel_buf, int frame, int frame_width, int a6, int direction) {
//    void (__fastcall *f)(int, int, uint8_t *, int, int, int, int) = (void *)0x416B19;
//    f(screen_x, screen_y, cel_buf, frame, frame_width, a6, direction);
// }
//
// static void __fastcall engine_cel_decode_frame_with_header_light_and_transparency_into_buf2(uint8_t *dst_buf, uint8_t *cel_buf, int frame, int frame_width, int a5, int direction) {
//    void (__fastcall *f)(uint8_t *, uint8_t *, int, int, int, int) = (void *)0x416BA9;
//    f(dst_buf, cel_buf, frame, frame_width, a5, direction);
// }
//
// static void __fastcall engine_cel_decode_frame_with_header_and_light_not_equipable2(int screen_x, int screen_y, uint8_t *cel_buf, int frame, int frame_width, int always_0, int direction, int8_t always_1) {
//    void (__fastcall *f)(int, int, uint8_t *, int, int, int, int, int8_t) = (void *)0x416C1B;
//    f(screen_x, screen_y, cel_buf, frame, frame_width, always_0, direction, always_1);
// }
//
// static void __fastcall engine_cel_decode_frame_into_rect_of_buf(uint8_t *dst_buf, int always_0, int dst_height, int dst_width, uint8_t *cel_buf, int frame, int frame_width) {
//    void (__fastcall *f)(uint8_t *, int, int, int, uint8_t *, int, int) = (void *)0x416D3C;
//    f(dst_buf, always_0, dst_height, dst_width, cel_buf, frame, frame_width);
// }
//
// static void __fastcall engine_cel_decode_frame_with_colour(uint8_t colour, int screen_x, int screen_y, uint8_t *cel_buf, int frame, int frame_width, int a7, int direction) {
//    void (__fastcall *f)(uint8_t, int, int, uint8_t *, int, int, int, int) = (void *)0x416DC6;
//    f(colour, screen_x, screen_y, cel_buf, frame, frame_width, a7, direction);
// }
//
// static void __fastcall engine_cel_decode_frame_with_header_and_colour_highlight(uint8_t colour, int screen_x, int screen_y, uint8_t *cel_buf, int frame, int frame_width, int a7, int direction) {
//    void (__fastcall *f)(uint8_t, int, int, uint8_t *, int, int, int, int) = (void *)0x416EC0;
//    f(colour, screen_x, screen_y, cel_buf, frame, frame_width, a7, direction);
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

// CelDecodeFrame decodes the given CEL frame to the specified screen
// coordinate.
//
//    x = screen_x - 64
//    y = screen_y - 160
//    frameNum = frame - 1
//
// Note, the coordinates specify the bottom left corner (verified in game).
//
// Note, this function is only used to decode CEL images without frame headers.
//
// ref: 0x416274
func CelDecodeFrame(screenX, screenY int, celBuf unsafe.Pointer, frame, frameWidth int) {
	if UseGUI {
		celDecodeFrame(screenX, screenY, celBuf, frame, frameWidth)
	}
	C.engine_cel_decode_frame(C.int(screenX), C.int(screenY), (*C.uint8_t)(celBuf), C.int(frame), C.int(frameWidth))
}

// CelDecodeFrameIntoBuf decodes the given CEL frame into the specified buffer.
//
// Note, this function is only used to decode CEL images without frame headers
// (pentspn2.cel).
//
// ref: 0x4162B8
func CelDecodeFrameIntoBuf(dstBuf, celBuf unsafe.Pointer, frame, frameWidth int) {
	C.engine_cel_decode_frame_into_buf((*C.uint8_t)(dstBuf), (*C.uint8_t)(celBuf), C.int(frame), C.int(frameWidth))
}

// CelDecodeFrameWithHeader decodes the given CEL frame to the specified screen
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
func CelDecodeFrameWithHeader(screenX, screenY int, celBuf unsafe.Pointer, frame, frameWidth, always0, direction int) {
	if UseGUI {
		celDecodeFrameWithHeader(screenX, screenY, celBuf, frame, frameWidth, always0, direction)
	}
	C.engine_cel_decode_frame_with_header(C.int(screenX), C.int(screenY), (*C.uint8_t)(celBuf), C.int(frame), C.int(frameWidth), C.int(always0), C.int(direction))
}

// CelDecodeFrameWithHeaderIntoBuf decodes the given CEL frame into the
// specified buffer.
//
// Note, this function is only used to decode CEL images with frame headers
// (square.cel).
//
// ref: 0x416359
func CelDecodeFrameWithHeaderIntoBuf(dstBuf, celBuf unsafe.Pointer, frame, frameWidth, always0, direction int) {
	C.engine_cel_decode_frame_with_header_into_buf((*C.uint8_t)(dstBuf), (*C.uint8_t)(celBuf), C.int(frame), C.int(frameWidth), C.int(always0), C.int(direction))
}

// CelDecodeFrameWithLight decodes the given CEL frame to the specified screen
// coordinate, adding lighting if applicable.
//
//    x = screen_x - 64
//    y = screen_y - 160
//    frameNum = frame - 1
//
// Note, the coordinates specify the bottom left corner.
//
// Note, this function is only used to decode CEL images without frame headers
// (bigtgold.cel).
//
// ref: 0x416565
func CelDecodeFrameWithLight(screenX, screenY int, celBuf unsafe.Pointer, frame, frameWidth int) {
	C.engine_cel_decode_frame_with_light(C.int(screenX), C.int(screenY), (*C.uint8_t)(celBuf), C.int(frame), C.int(frameWidth))
}

// CelDecodeFrameWithHeaderAndLight decodes the given CEL frame to the specified
// screen coordinate, adding lighting if applicable.
//
//    x = screen_x - 64
//    y = screen_y - 160
//    frameNum = frame - 1
//
// Note, the coordinates specify the bottom left corner.
//
// Note, this function is only used to decode CEL images with frame headers
// (item drops, objects).
//
// ref: 0x4165BD
func CelDecodeFrameWithHeaderAndLight(screenX, screenY int, celBuf unsafe.Pointer, frame, frameWidth, always0, direction int) {
	C.engine_cel_decode_frame_with_header_and_light(C.int(screenX), C.int(screenY), (*C.uint8_t)(celBuf), C.int(frame), C.int(frameWidth), C.int(always0), C.int(direction))
}

// CelDecodeFrameWithHeaderLightAndTransparencyIntoBuf decodes the given CEL
// frame into the specified buffer with added lighting and transparency.
//
// Note, this function is only used to decode CEL images with frame headers
// (objcurs.cel, level special).
//
// ref: 0x41664B
func CelDecodeFrameWithHeaderLightAndTransparencyIntoBuf(dstBuf, celBuf unsafe.Pointer, frame, frameWidth, always0, direction int) {
	C.engine_cel_decode_frame_with_header_light_and_transparency_into_buf((*C.uint8_t)(dstBuf), (*C.uint8_t)(celBuf), C.int(frame), C.int(frameWidth), C.int(always0), C.int(direction))
}

// CelDecodeFrameWithHeaderAndLightNotEquipable decodes the given CEL frame to
// the specified screen coordinate with added lighting.
//
//    x = screen_x - 64
//    y = screen_y - 160
//    frameNum = frame - 1
//
// Note, the coordinates specify the bottom left corner.
//
// Note, this function is only used to decode CEL images with frame headers
// (objcurs.cel).
//
// ref: 0x4166BF
func CelDecodeFrameWithHeaderAndLightNotEquipable(screenX, screenY int, celBuf unsafe.Pointer, frame, frameWidth, always0, direction int, always1 int8) {
	C.engine_cel_decode_frame_with_header_and_light_not_equipable(C.int(screenX), C.int(screenY), (*C.uint8_t)(celBuf), C.int(frame), C.int(frameWidth), C.int(always0), C.int(direction), C.int8_t(always1))
}

// CelDecodeFrameWithHeader2 decodes the given CEL frame to the specified screen
// coordinate.
//
//    x = screen_x - 64
//    y = screen_y - 160
//    frameNum = frame - 1
//
// Note, the coordinates specify the bottom left corner.
//
// Note, this function is only used to decode CEL images with frame headers
// (objects, objcurs.cel, item drops, towners).
//
// Note, D1DrawCursorTown (from RE Notes)
//
// ref: 0x41685A
func CelDecodeFrameWithHeader2(screenX, screenY int, celBuf unsafe.Pointer, frame, frameWidth, a6, direction int) {
	C.engine_cel_decode_frame_with_header2(C.int(screenX), C.int(screenY), (*C.uint8_t)(celBuf), C.int(frame), C.int(frameWidth), C.int(a6), C.int(direction))
}

// CelDecodeFrameWithHeaderIntoBuf2 decodes the given CEL frame into the
// specified buffer.
//
// Note, this function is only used to decode CEL images with frame headers
// (square.cel).
//
// ref: 0x4168D5
func CelDecodeFrameWithHeaderIntoBuf2(dstBuf, celBuf unsafe.Pointer, frame, frameWidth, a5, direction int) {
	C.engine_cel_decode_frame_with_header_into_buf2((*C.uint8_t)(dstBuf), (*C.uint8_t)(celBuf), C.int(frame), C.int(frameWidth), C.int(a5), C.int(direction))
}

// CelDecodeFrameWithHeaderAndLight2 decodes the given CEL frame to the
// specified screen coordinate, adding lighting if applicable.
//
//    x = screen_x - 64
//    y = screen_y - 160
//    frameNum = frame - 1
//
// Note, the coordinates specify the bottom left corner.
//
// Note, this function is only used to decode CEL images with frame headers
// (item drops, objects).
//
// Note, D1DrawObjectBaseDungeon (from RE Notes).
//
// ref: 0x416B19
func CelDecodeFrameWithHeaderAndLight2(screenX, screenY int, celBuf unsafe.Pointer, frame, frameWidth, a6, direction int) {
	C.engine_cel_decode_frame_with_header_and_light2(C.int(screenX), C.int(screenY), (*C.uint8_t)(celBuf), C.int(frame), C.int(frameWidth), C.int(a6), C.int(direction))
}

// CelDecodeFrameWithHeaderLightAndTransparencyIntoBuf2 decodes the given CEL
// frame into the specified buffer with added lighting and transparency.
//
// Note, this function is only used to decode CEL images with frame headers
// (level special).
//
// Note, D1DrawArchTile (from RE Notes).
//
// ref: 0x416BA9
func CelDecodeFrameWithHeaderLightAndTransparencyIntoBuf2(dstBuf, celBuf unsafe.Pointer, frame, frameWidth, a5, direction int) {
	C.engine_cel_decode_frame_with_header_light_and_transparency_into_buf2((*C.uint8_t)(dstBuf), (*C.uint8_t)(celBuf), C.int(frame), C.int(frameWidth), C.int(a5), C.int(direction))
}

// CelDecodeFrameWithHeaderAndLightNotEquipable2 decodes the given CEL frame to
// the specified screen coordinate with added lighting.
//
//    x = screen_x - 64
//    y = screen_y - 160
//    frameNum = frame - 1
//
// Note, the coordinates specify the bottom left corner.
//
// Note, this function is only used to decode CEL images with frame headers (objcurs.cel).
//
// ref: 0x416C1B
func CelDecodeFrameWithHeaderAndLightNotEquipable2(screenX, screenY int, celBuf unsafe.Pointer, frame, frameWidth, always0, direction int, always1 int8) {
	C.engine_cel_decode_frame_with_header_and_light_not_equipable2(C.int(screenX), C.int(screenY), (*C.uint8_t)(celBuf), C.int(frame), C.int(frameWidth), C.int(always0), C.int(direction), C.int8_t(always1))
}

// CelDecodeFrameIntoRectOfBuf decodes the given CEL frame into a rectangle of
// the specified buffer.
//
// Note, this function is only used to decode CEL images without frame headers
// (control panel and orbs).
//
// ref: 0x416D3C
func CelDecodeFrameIntoRectOfBuf(dstBuf unsafe.Pointer, always0, dstHeight, dstWidth int, celBuf unsafe.Pointer, frame, frameWidth int) {
	C.engine_cel_decode_frame_into_rect_of_buf((*C.uint8_t)(dstBuf), C.int(always0), C.int(dstHeight), C.int(dstWidth), (*C.uint8_t)(celBuf), C.int(frame), C.int(frameWidth))
}

// CelDecodeFrameWithColour decodes the given CEL frame to the specified screen
// coordinate and with the specified colour.
//
//    x = screen_x - 64
//    y = screen_y - 160
//    frameNum = frame - 1
//
// Note, the coordinates specify the bottom left corner.
//
// Note, this function is only used to decode CEL images with frame headers
// (objcurs.cel, item drops, objects, towners).
//
// ref: 0x416DC6
func CelDecodeFrameWithColour(colour uint8, screenX, screenY int, celBuf unsafe.Pointer, frame, frameWidth, a7, direction int) {
	C.engine_cel_decode_frame_with_colour(C.uint8_t(colour), C.int(screenX), C.int(screenY), (*C.uint8_t)(celBuf), C.int(frame), C.int(frameWidth), C.int(a7), C.int(direction))
}

// CelDecodeFrameWithHeaderAndColourHighlight decodes the given CEL frame to the
// specified screen coordinate with added colour highlight.
//
//    x = screen_x - 64
//    y = screen_y - 160
//    frameNum = frame - 1
//
// Note, the coordinates specify the bottom left corner.
//
// Note, this function is only used to decode CEL images with frame headers
// (objcurs.cel, item drops, objects, towners).
//
// NOTE: D1DrawObjectHilight (from RE Notes).
//
// ref: 0x416EC0
func CelDecodeFrameWithHeaderAndColourHighlight(colour uint8, screenX, screenY int, celBuf unsafe.Pointer, frame, frameWidth, a7, direction int) {
	C.engine_cel_decode_frame_with_header_and_colour_highlight(C.uint8_t(colour), C.int(screenX), C.int(screenY), (*C.uint8_t)(celBuf), C.int(frame), C.int(frameWidth), C.int(a7), C.int(direction))
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
