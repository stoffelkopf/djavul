// Stateful connection.

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

// CommandTCP specifies a TCP command to send to front-end.
type CommandTCP uint8

// TCP commands.
const (
	// CmdLoadFile specifies a file to load.
	CmdLoadFile CommandTCP = iota + 1
)

// String returns the string representation of the TCP command.
func (cmd CommandTCP) String() string {
	m := map[CommandTCP]string{
		CmdLoadFile: "CmdLoadFile",
	}
	if s, ok := m[cmd]; ok {
		return s
	}
	return fmt.Sprintf("unknown CommandTCP(%d)", uint(cmd))
}

// LoadFile specifies a file to load.
type LoadFile struct {
	Path string
}

// SendLoadFile send a load file command to the front-end.
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
