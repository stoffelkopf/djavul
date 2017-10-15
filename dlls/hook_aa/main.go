package main

import "C"

import (
	"fmt"

	"github.com/AllenDang/w32"
	"github.com/sanctuary/djavul/automap"
	"github.com/sanctuary/djavul/capture"
)

//export Init_AAAAAAAAAAAAAAA
func Init_AAAAAAAAAAAAAAA() {
	// Called from WinMain.
	fmt.Println("hello from Go :)")
}

//export OnKeyPress_AA
func OnKeyPress_AA(key int) {
	// Called from diablo_on_key_press.
	fmt.Println("key press:", key)
	switch key {
	case 'X':
		// eXplore
		automap.Explore()
	}
}

//export OnKeyRelease_AAAAAAAAAAAAA
func OnKeyRelease_AAAAAAAAAAAAA(key int) {
	// Called from diablo_on_key_release.
	fmt.Println("key release:", key)
	switch key {
	case w32.VK_APPS, w32.VK_SNAPSHOT:
		// Capture screenshot.
		capture.Screenshot()
	}
}

func main() {}
