package engine

// #include <stdint.h>
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
	"unsafe"
)

// useGo specifies whether to use the Go implementation.
const useGo = true

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
func MemLoadFile(path unsafe.Pointer, size *int32) *uint8 {
	buf := C.engine_mem_load_file((*C.char)(path), (*C.int32_t)(unsafe.Pointer(size)))
	return (*uint8)(unsafe.Pointer(buf))
}
