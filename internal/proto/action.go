package proto

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"log"

	"github.com/pkg/errors"
)

// TODO: Make front-end agnostic. Instead of referencing buttons, pixel
// coordinates, etc; reference move actions, map coordinates, etc.

// EngineAction is a front-end action for the Diablo 1 game engine.
type EngineAction interface {
	// isEngineAction ensures that only front-end actions can be assigned to the
	// EngineAction interface.
	isEngineAction()
}

// ButtonPressedAction signals to the game engine that a mouse button or
// keyboard key has been pressed on the front-end.
type ButtonPressedAction struct {
	// Mouse button or keyboard key pressed.
	Button int32 // pixelgl.Button
	// Mouse position on screen.
	X, Y int32
}

// ButtonReleasedAction signals to the game engine that a mouse button or
// keyboard key has been released on the front-end.
type ButtonReleasedAction struct {
	// Mouse button or keyboard key released.
	Button int32 // pixelgl.Button
	// Mouse position on screen.
	X, Y int32
}

// isEngineAction ensures that only front-end actions can be assigned to the
// EngineAction interface.
func (ButtonPressedAction) isEngineAction()  {}
func (ButtonReleasedAction) isEngineAction() {}

// NewAction encodes the given action into a network packet.
func NewAction(action EngineAction) *PacketUnstable {
	pkt := &PacketUnstable{}
	switch action := action.(type) {
	case ButtonPressedAction:
		pkt.Cmd = CmdButtonPressedAction
	case ButtonReleasedAction:
		pkt.Cmd = CmdButtonReleasedAction
	default:
		panic(fmt.Errorf("support for action %T not yet implemented", action))
	}
	buf := &bytes.Buffer{}
	if err := binary.Write(buf, binary.LittleEndian, action); err != nil {
		die(err)
	}
	pkt.Data = buf.Bytes()
	return pkt
}

// Actions is a channel of actions received from the front-end.
var Actions chan EngineAction

// ### [ Helper functions ] ####################################################

// die kills the application when a fatal error has occurred.
func die(err error) {
	log.Fatalf("%+v", errors.WithStack(err))
}
