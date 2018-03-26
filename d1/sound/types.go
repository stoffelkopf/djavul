//+build djavul

package sound

import (
	"unsafe"

	"github.com/sanctuary/djavul/internal/dsound"
)

// File represents a WAV sound file.
//
// PSX def:
//    typedef struct TSnd {
//    } TSnd;
type File struct { // size = 0x28
	// offset 0000 (18 bytes)
	Format dsound.WAVEFORMATEX
	// offset 0014 (4 bytes)
	Len int32
	// offset 0018 (4 bytes)
	Offset int32
	// offset 001C (4 bytes)
	Path unsafe.Pointer // Sound path (NULL-terminated string)
	// offset 0020 (4 bytes)
	DSB *dsound.IDirectSoundBuffer // Direct sound buffer.
	// offset 0024 (4 bytes)
	StartTC uint32 // Start tick count.
}
