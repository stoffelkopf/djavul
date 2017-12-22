//+build djavul

package main

import (
	"fmt"

	"github.com/AllenDang/w32"
	"github.com/sanctuary/djavul/automap"
	"github.com/sanctuary/djavul/capture"
)

//export Init
func Init() {
	// Called from WinMain.
	fmt.Println("init: djavul.dll")
}

//export OnKeyPress
func OnKeyPress(key int) {
	// Called from diablo_on_key_press.
	fmt.Println("key press:", key)
	switch key {
	case 'X':
		// eXplore
		automap.Explore()
	}
}

//export OnKeyRelease
func OnKeyRelease(key int) {
	// Called from diablo_on_key_release.
	fmt.Println("key release:", key)
	switch key {
	case w32.VK_APPS, w32.VK_SNAPSHOT:
		// Capture screenshot.
		capture.Screenshot()
	}
}
