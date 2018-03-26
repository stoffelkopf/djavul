//+build djavul

package sound

import "C"

import (
	"log"
	"strings"
	"unsafe"

	"github.com/sanctuary/djavul/internal/proto"
)

// playFile plays the given sound file.
//
// PSX ref: 0x80077D58
// PSX def: void snd_play_snd__FP4TSFXll(struct TSFX *pSnd, long lVolume, long lPan)
func playFile(file unsafe.Pointer, volumeDelta, pan int) {
	f := (*File)(file)
	path := goPath(f.Path)
	dbg.Println("play sound:", path)
	// TODO: calcualte absolute volume:
	//    volume = sound_volume + volume_delta
	if err := proto.SendPlaySound(path, volumeDelta, pan); err != nil {
		log.Fatalf("%+v", err)
	}
}

// ### [ Helper functions ] ####################################################

// goPath returns an equivalent Go string of the given file path.
func goPath(path unsafe.Pointer) string {
	p := C.GoString((*C.char)(path))
	p = strings.Replace(p, "\\", "/", -1)
	return strings.ToLower(p)
}
