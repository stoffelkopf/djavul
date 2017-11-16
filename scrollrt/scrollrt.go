// Package scrollrt implements rendering.
package scrollrt

import (
	"fmt"
)

// drawMainW renders the UI and the game world on screen.
//
// ref: 0x4564F9
func drawMainW() {
	fmt.Println("draw signal sent")
	DrawSignal <- struct{}{}
}

// DrawSignal signals when to draw.
var DrawSignal = make(chan struct{})
