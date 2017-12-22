//+build djavul

package sound

// static void __fastcall sound_file_play(void *sound_file, int volume_delta, int pan) {
//    void (__fastcall *f)(void *, int, int) = (void *)0x456D60;
//    f(sound_file, volume_delta, pan);
// }
import "C"

import "unsafe"

// useGo specifies whether to use the Go implementation.
const useGo = true

// PlayFile plays the given sound file.
//
// PSX ref: 0x80077D58
// PSX def: void snd_play_snd__FP4TSFXll(struct TSFX *pSnd, long lVolume, long lPan)
func PlayFile(file unsafe.Pointer, volumeDelta, pan int) {
	if useGo {
		playFile(file, volumeDelta, pan)
	}
	C.sound_file_play(file, C.int(volumeDelta), C.int(pan))
}
