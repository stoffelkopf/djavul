// Package scrollrt implements rendering.
package scrollrt

import (
	"log"

	"github.com/sanctuary/djavul/internal/proto"
)

// drawMainW renders the UI and the game world on screen.
//
// ref: 0x4564F9
func drawMainW() {
	//fmt.Println("draw signal sent")
	if err := proto.SendUpdateScreen(); err != nil {
		log.Fatalf("%+v", err)
	}
}
