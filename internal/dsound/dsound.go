// Package dsound provides access to DirectSound functions.
package dsound

// #include <dsound.h>
import "C"

// IDirectSoundBuffer represents a sound buffer.
type IDirectSoundBuffer C.IDirectSoundBuffer

// WAVEFORMATEX defines the format of waveform-audio data.
type WAVEFORMATEX struct {
	FormatTag       uint16
	NChannels       uint16
	NSamplesPerSec  uint32
	NAvgBytesPerSec uint32
	NBlockAlign     uint16
	BitsPerSample   uint16
	Size            uint16
}
