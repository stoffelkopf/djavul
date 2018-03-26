package sha1

import "unsafe"

// Global variables.
var (
	// ref: 0x69EFB0
	Contexts = (*[3]Context)(unsafe.Pointer(uintptr(0x69EFB0)))
)
