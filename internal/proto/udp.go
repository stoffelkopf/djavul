// Stateless connection.

package proto

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"image"

	"github.com/pkg/errors"
)

// === [ Streaming connection ] ================================================
//
// Any data which is time critical and may be lost without affecting the core
// logic of the game engine (such as the rendering of frames, playback of
// sounds).
//
// UDP packet loss is insignificant: drop.
//
// Spoofing doesn't matter.

// UDP pipes.
const (
	UDPReadPipe  = `\\.\pipe\udp_r`
	UDPWritePipe = `\\.\pipe\udp_w`
)

// PacketUDP represents a UDP packet.
type PacketUDP struct {
	Cmd  CommandUDP
	Data []byte
}

var (
	// EncUDP holds a Gob encoder for the UDP connection to the front-end.
	EncUDP *gob.Encoder
	// DecUDP holds a Gob decoder for the UDP connection to the front-end.
	DecUDP *gob.Decoder
)

// CommandUDP specifies a UDP command to send to front-end.
type CommandUDP uint8

// UDP commands.
const (
	// Engine events.

	// CmdDrawImage specifies an image to draw.
	CmdDrawImage CommandUDP = iota + 1
	// CmdUpdateScreen specifies that the screen should be updated.
	CmdUpdateScreen
	// CmdPlaySound specifies a sound to play.
	CmdPlaySound

	// Front-end actions.

	// CmdButtonPressedAction specifies a mouse button or keyboard key pressed on
	// the front-end.
	CmdButtonPressedAction
	// CmdButtonPressedAction specifies a mouse button or keyboard key released
	// on the front-end.
	CmdButtonReleasedAction
)

// String returns the string representation of the UDP command.
func (cmd CommandUDP) String() string {
	m := map[CommandUDP]string{
		CmdDrawImage:    "CmdDrawImage",
		CmdUpdateScreen: "CmdUpdateScreen",
	}
	if s, ok := m[cmd]; ok {
		return s
	}
	return fmt.Sprintf("unknown CommandTCP(%d)", uint(cmd))
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

// SendDrawImage send a draw image command to the front-end.
func SendDrawImage(path string, x, y int, frameNum int) error {
	data := DrawImage{
		Path:     path,
		Dp:       image.Pt(x, y),
		FrameNum: frameNum,
	}
	drawImages = append(drawImages, data)
	return nil
}

// SendDrawSubimage send a draw image command to the front-end, for the
// specified subimage.
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

// SendUpdateScreen send an update screen command to the front-end.
func SendUpdateScreen() error {
	buf := &bytes.Buffer{}
	enc := gob.NewEncoder(buf)
	if err := enc.Encode(&drawImages); err != nil {
		return errors.WithStack(err)
	}
	drawImages = drawImages[:0]
	pkg := PacketUDP{
		Cmd:  CmdUpdateScreen,
		Data: buf.Bytes(),
	}
	if err := EncUDP.Encode(&pkg); err != nil {
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

// SendPlaySound send a play sound command to the front-end.
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
	pkg := PacketUDP{
		Cmd:  CmdPlaySound,
		Data: buf.Bytes(),
	}
	if err := EncUDP.Encode(&pkg); err != nil {
		return errors.WithStack(err)
	}
	return nil
}
