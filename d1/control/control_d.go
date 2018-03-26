//+build djavul

package control

import (
	"image"
	"log"

	"github.com/sanctuary/djavul/d1/engine"
	"github.com/sanctuary/djavul/internal/proto"
)

// drawPanel renders the control panel onto screen.
//
// ref: 0x404259
func drawPanel(panelX, panelY, width, height, screenX, screenY int) {
	x, y := engine.XYFromScreenCoords(screenX, screenY)
	sr := image.Rect(panelX, panelY, panelX+width, panelY+height)
	frameNum := 0
	if err := proto.SendDrawSubimage("ctrlpan/panel8.cel", x, y, sr, frameNum); err != nil {
		log.Fatalf("%+v", err)
	}
}
