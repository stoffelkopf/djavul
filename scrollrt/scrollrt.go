// Package scrollrt implements rendering.
package scrollrt

import (
	"log"

	"github.com/sanctuary/djavul/internal/proto"
)

// drawMain renders the specified parts of the UI and the game world on screen.
//
// ref: 0x456124
func drawMain(height int, updateDescriptionBox, updateLifeOrb, updateManaOrb, updateBelt, updateControlButtons bool) {
	//fmt.Println("draw signal sent")
	if err := proto.SendUpdateScreen(); err != nil {
		log.Fatalf("%+v", err)
	}
}
