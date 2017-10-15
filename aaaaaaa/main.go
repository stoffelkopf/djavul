package main

import "C"
import (
	"fmt"

	"github.com/AllenDang/w32"
	"github.com/sanctuary/djavul/automap"
	"github.com/sanctuary/djavul/capture"
)

//export InitAAAAAAAAAAAAAAAA
func InitAAAAAAAAAAAAAAAA() {
	// Called from WinMain.
	fmt.Println("hello from Go :)")
}

//export Cccccccccccccccccccccccccc
func Cccccccccccccccccccccccccc() int {
	return 42
}

//export OnKeyPressAAA
func OnKeyPressAAA(key int) {
	// Called from on_key_press.
	fmt.Println("key press:", key)
	switch key {
	case 'X':
		// eXplore
		automap.Explore()
	case w32.VK_APPS, w32.VK_SNAPSHOT:
		// Capture screenshot.
		capture.Screenshot()
	}
}

func main() {}
