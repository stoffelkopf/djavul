package proto

import (
	"encoding/gob"
	"fmt"

	"github.com/pkg/errors"
)

// === [ Reliable connection ] =================================================
//
// Packets which cannot be lost are sent on the reliable connection. For
// instance, loading game data in between levels. Any game statistics (such as
// life and damage).
//
// TCP packet loss is a critical failure: crash.
//
// Spoof protection needed.

var (
	// EncTCP holds a Gob encoder for the TCP connection to the front-end.
	EncTCP *gob.Encoder
	// DecTCP holds a Gob decoder for the TCP connection to the front-end.
	DecTCP *gob.Decoder
)

type CommandTCP uint8

const (
	CmdLoadFile CommandTCP = iota + 1
)

func (cmd CommandTCP) String() string {
	m := map[CommandTCP]string{
		CmdLoadFile: "CmdLoadFile",
	}
	if s, ok := m[cmd]; ok {
		return s
	}
	return fmt.Sprintf("unknown CommandTCP(%d)", uint(cmd))
}

type LoadFile struct {
	Path string
}

func SendLoadFile(path string) error {
	cmd := CmdLoadFile
	if err := EncTCP.Encode(&cmd); err != nil {
		return errors.WithStack(err)
	}
	pkg := LoadFile{
		Path: path,
	}
	if err := EncTCP.Encode(&pkg); err != nil {
		return errors.WithStack(err)
	}
	return nil
}
