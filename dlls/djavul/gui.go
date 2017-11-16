//+build djavul

package main

import "C"

import (
	"fmt"
	"unsafe"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/sanctuary/djavul/engine"
	"golang.org/x/image/colornames"
)

// run initializes the Pixel GUI.
func run() {
	go winGUI()
	cfg := pixelgl.WindowConfig{
		Title:       "djavul",
		Bounds:      pixel.R(0, 0, 640, 480),
		Undecorated: true,
		VSync:       true,
	}
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(fmt.Errorf("unable to create new Pixel Window; %v", err))
	}
	engine.Win = win

	// Draw loop.
	win.Clear(colornames.Skyblue)
	for !win.Closed() {
		win.Update()
	}
}

// ### [ Exports ] #############################################################

//export CelDecodeFrame
func CelDecodeFrame(screenX, screenY int, celBuf unsafe.Pointer, frame, frameWidth int) {
	engine.CelDecodeFrame(screenX, screenY, celBuf, frame, frameWidth)
}

//export MemLoadFile
func MemLoadFile(path unsafe.Pointer, size *int32) unsafe.Pointer {
	return engine.MemLoadFile(path, size)
}
