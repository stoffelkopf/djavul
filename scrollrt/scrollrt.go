// Package scrollrt implements rendering.
package scrollrt

import (
	"fmt"
	"log"

	"github.com/sanctuary/djavul/internal/proto"
)

// drawMain renders the specified parts of the UI and the game world on screen.
//
// ref: 0x456124
func drawMain(height int, updateDescriptionBox, updateLifeOrb, updateManaOrb, updateBelt, updateControlButtons bool) {
	//fmt.Println("draw signal sent")
	fmt.Println("height:", height)
	fmt.Println("updateDescriptionBox:", updateDescriptionBox)
	fmt.Println("updateLifeOrb:", updateLifeOrb)
	fmt.Println("updateManaOrb:", updateManaOrb)
	fmt.Println("updateBelt:", updateBelt)
	fmt.Println("updateControlButtons:", updateControlButtons)
	if err := proto.SendUpdateScreen(); err != nil {
		log.Fatalf("%+v", err)
	}
}
