// Stateless connection.

package proto

import (
	"bytes"
	"encoding/gob"
	"image"

	"github.com/pkg/errors"
)

// === [ Streaming connection ] ================================================
//
// Any data which is time critical and may be lost without affecting the core
// logic of the game engine (such as the rendering of frames, playback of
// sounds).
//
// Unstable packet loss is insignificant: drop.
//
// Spoofing doesn't matter.

// PacketUnstable represents an unstable packet.
type PacketUnstable struct {
	Cmd  CommandUnstable
	Data []byte
}

var (
	// EncUnstable holds a Gob encoder for the unstable connection to the
	// frontend.
	EncUnstable *gob.Encoder
	// DecUnstable holds a Gob decoder for the unstable connection to the
	// frontend.
	DecUnstable *gob.Decoder
)

//go:generate stringer -type CommandUnstable

// CommandUnstable specifies a unstable command to send to frontend.
type CommandUnstable uint8

// Unstable commands.
const (
	// Engine events.

	// CmdDrawImage specifies an image to draw.
	CmdDrawImage CommandUnstable = iota + 1
	// CmdUpdateScreen specifies that the screen should be updated.
	CmdUpdateScreen
	// CmdPlaySound specifies a sound to play.
	CmdPlaySound

	// Frontend actions.

	// CmdButtonPressedAction specifies a mouse button or keyboard key pressed on
	// the frontend.
	CmdButtonPressedAction
	// CmdButtonPressedAction specifies a mouse button or keyboard key released
	// on the frontend.
	CmdButtonReleasedAction
)

// DrawImage specifies an image to draw at a given coordinate.
type DrawImage struct {
	Path     string
	Dp       image.Point
	Sr       image.Rectangle
	FrameNum int
}

// drawImages is a buffer of images to draw at each frame update.
var drawImages []DrawImage

// SendDrawImage send a draw image command to the frontend.
func SendDrawImage(path string, x, y int, frameNum int) error {
	data := DrawImage{
		Path:     path,
		Dp:       image.Pt(x, y),
		FrameNum: frameNum,
	}
	drawImages = append(drawImages, data)
	return nil
}

// SendDrawSubimage send a draw image command to the frontend, for the specified
// subimage.
func SendDrawSubimage(path string, x, y int, sr image.Rectangle, frameNum int) error {
	data := DrawImage{
		Path:     path,
		Dp:       image.Pt(x, y),
		Sr:       sr,
		FrameNum: frameNum,
	}
	drawImages = append(drawImages, data)
	return nil
}

// SendUpdateScreen send an update screen command to the frontend.
func SendUpdateScreen() error {
	buf := &bytes.Buffer{}
	enc := gob.NewEncoder(buf)
	if err := enc.Encode(&drawImages); err != nil {
		return errors.WithStack(err)
	}
	drawImages = drawImages[:0]
	pkg := PacketUnstable{
		Cmd:  CmdUpdateScreen,
		Data: buf.Bytes(),
	}
	if err := EncUnstable.Encode(&pkg); err != nil {
		return errors.WithStack(err)
	}
	return nil
}

// PlaySound specifies a sound to play with a given volume and panning.
type PlaySound struct {
	Path        string
	VolumeDelta int
	Pan         int
}

// SendPlaySound send a play sound command to the frontend.
func SendPlaySound(path string, volumeDelta, pan int) error {
	buf := &bytes.Buffer{}
	enc := gob.NewEncoder(buf)
	data := PlaySound{
		Path:        path,
		VolumeDelta: volumeDelta,
		Pan:         pan,
	}
	if err := enc.Encode(&data); err != nil {
		return errors.WithStack(err)
	}
	pkg := PacketUnstable{
		Cmd:  CmdPlaySound,
		Data: buf.Bytes(),
	}
	if err := EncUnstable.Encode(&pkg); err != nil {
		return errors.WithStack(err)
	}
	return nil
}
