package main

import (
	"github.com/faiface/pixel/pixelgl"
	"github.com/sanctuary/djavul/internal/proto"
)

// buttonPressed returns an action for the Diablo 1 game engine, based on the
// button pressed window event.
func buttonPressed(win *pixelgl.Window, e ButtonPressedEvent) proto.ButtonPressedAction {
	return proto.ButtonPressedAction{
		Button: int32(e.Button),
		X:      int32(e.Pos.X),
		Y:      int32(e.Pos.Y),
	}
}

// buttonReleased returns an action for the Diablo 1 game engine, based on the
// button released window event.
func buttonReleased(win *pixelgl.Window, e ButtonReleasedEvent) proto.ButtonReleasedAction {
	return proto.ButtonReleasedAction{
		Button: int32(e.Button),
		X:      int32(e.Pos.X),
		Y:      int32(e.Pos.Y),
	}
}
