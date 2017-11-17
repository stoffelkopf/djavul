//+build djavul

package main

import "C"

import (
	"fmt"
	"unsafe"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/sanctuary/djavul/engine"
	"github.com/sanctuary/djavul/scrollrt"
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
		<-scrollrt.DrawSignal
		fmt.Println("draw signal received")
	}
}

// ### [ Exports ] #############################################################

// --- [ engine ] --------------------------------------------------------------

//export CelDecodeFrame
func CelDecodeFrame(screenX, screenY int, celBuf unsafe.Pointer, frame, frameWidth int) {
	engine.CelDecodeFrame(screenX, screenY, celBuf, frame, frameWidth)
}

//export CelDecodeFrameWithHeader
func CelDecodeFrameWithHeader(screenX, screenY int, celBuf unsafe.Pointer, frame, frameWidth, always0, direction int) {
	engine.CelDecodeFrameWithHeader(screenX, screenY, celBuf, frame, frameWidth, always0, direction)
}

//export MemLoadFile
func MemLoadFile(path unsafe.Pointer, size *int32) unsafe.Pointer {
	return engine.MemLoadFile(path, size)
}

// --- [ scrollrt ] ------------------------------------------------------------

//export DrawMainW
func DrawMainW() {
	scrollrt.DrawMainW()
}
