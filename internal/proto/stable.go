// Stateful connection.

package proto

import (
	"encoding/gob"

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
