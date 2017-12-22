//+build djavul

package main

import "C"

import (
	"unsafe"

	"github.com/sanctuary/djavul/sound"
)

// --- [ sound ] ---------------------------------------------------------------

//export PlayFile
func PlayFile(file unsafe.Pointer, volumeDelta, pan int) {
	sound.PlayFile(file, volumeDelta, pan)
}
