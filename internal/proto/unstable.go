// Stateless connection.

package proto

import (
	"encoding/gob"
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
	// Frontend actions.

	// CmdButtonPressedAction specifies a mouse button or keyboard key pressed on
	// the frontend.
	CmdButtonPressedAction CommandUnstable = iota + 1
	// CmdButtonPressedAction specifies a mouse button or keyboard key released
	// on the frontend.
	CmdButtonReleasedAction
)
