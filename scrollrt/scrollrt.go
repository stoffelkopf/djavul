// Package scrollrt implements rendering.
package scrollrt

import (
	"fmt"
	"log"

	"github.com/sanctuary/djavul/internal/proto"
)

// drawMainW renders the UI and the game world on screen.
//
// ref: 0x4564F9
func drawMainW() {
	fmt.Println("draw signal sent")
	if err := proto.SendUpdateScreen(); err != nil {
		log.Fatalf("%+v", err)
	}
	DrawSignal <- struct{}{}
}

// DrawSignal signals when to draw.
var DrawSignal = make(chan struct{})
