// Stateful connection.

package proto

import (
	"encoding/gob"
	"image"

	"github.com/pkg/errors"
)

// === [ Reliable connection ] =================================================
//
// Packets which cannot be lost are sent on the reliable connection. For
// instance, loading game data in between levels. Any game statistics (such as
// life and damage).
//
// Stable packet loss is a critical failure: crash.
//
// Spoof protection needed.

var (
	// EncStable holds a Gob encoder for the stable connection to the frontend.
	EncStable *gob.Encoder
	// DecStable holds a Gob decoder for the stable connection to the frontend.
	DecStable *gob.Decoder
)

//go:generate stringer -type CommandStable

// CommandStable specifies a stable command to send to frontend.
type CommandStable uint8

// Stable commands.
const (
	// CmdLoadFile specifies a file to load.
	CmdLoadFile CommandStable = iota + 1

	// Engine events.

	// CmdDrawImage specifies an image to draw.
	CmdDrawImage
	// CmdUpdateScreen specifies that the screen should be updated.
	CmdUpdateScreen
	// CmdPlaySound specifies a sound to play.
	CmdPlaySound
)

// LoadFile specifies a file to load.
type LoadFile struct {
	Path string
}

// SendLoadFile send a load file command to the frontend.
func SendLoadFile(path string) error {
	cmd := CmdLoadFile
	if err := EncStable.Encode(&cmd); err != nil {
		return errors.WithStack(err)
	}
	pkg := LoadFile{
		Path: path,
	}
	if err := EncStable.Encode(&pkg); err != nil {
		return errors.WithStack(err)
	}
	return nil
}

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

// UpdateScreen specifies the images to draw at a screen update.
type UpdateScreen struct {
	Imgs []DrawImage
}

// SendUpdateScreen send an update screen command to the frontend.
func SendUpdateScreen() error {
	cmd := CmdUpdateScreen
	if err := EncStable.Encode(&cmd); err != nil {
		return errors.WithStack(err)
	}
	pkg := UpdateScreen{
		Imgs: drawImages,
	}
	drawImages = drawImages[:0]
	if err := EncStable.Encode(&pkg); err != nil {
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
	cmd := CmdPlaySound
	if err := EncStable.Encode(&cmd); err != nil {
		return errors.WithStack(err)
	}
	pkg := PlaySound{
		Path:        path,
		VolumeDelta: volumeDelta,
		Pan:         pan,
	}
	if err := EncStable.Encode(&pkg); err != nil {
		return errors.WithStack(err)
	}
	return nil
}
