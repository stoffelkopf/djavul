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

type CommandUDP uint8

const (
	CmdDrawImage CommandUDP = iota + 1
	CmdUpdateScreen
)

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

type DrawImage struct {
	Path     string
	X        float64
	Y        float64
	FrameNum int
}

func SendDrawImage(path string, x, y float64, frameNum int) error {
	buf := &bytes.Buffer{}
	data := DrawImage{
		Path:     path,
		X:        x,
		Y:        y,
		FrameNum: frameNum,
	}
	enc := gob.NewEncoder(buf)
	if err := enc.Encode(&data); err != nil {
		return errors.WithStack(err)
	}
	pkg := PacketUDP{
		Cmd:  CmdDrawImage,
		Data: buf.Bytes(),
	}
	if err := EncUDP.Encode(&pkg); err != nil {
		return errors.WithStack(err)
	}
	return nil
}

func SendUpdateScreen() error {
	pkg := PacketUDP{
		Cmd: CmdUpdateScreen,
	}
	if err := EncUDP.Encode(&pkg); err != nil {
		return errors.WithStack(err)
	}
	return nil
}
