// Stateless connection.

package proto

import (
	"bytes"
	"encoding/gob"
	"fmt"

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
	// CmdDrawImage specifies an image to draw.
	CmdDrawImage CommandUDP = iota + 1
	// CmdUpdateScreen specifies that the screen should be updated.
	CmdUpdateScreen
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
	X        float64
	Y        float64
	FrameNum int
}

// drawImages is a buffer of images to draw at each frame update.
var drawImages []DrawImage

// SendDrawImage send a draw image command to the front-end.
func SendDrawImage(path string, x, y float64, frameNum int) error {
	data := DrawImage{
		Path:     path,
		X:        x,
		Y:        y,
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
