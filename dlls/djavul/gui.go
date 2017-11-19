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

	// Draw loop.
	win.Clear(colornames.Skyblue)
	for !win.Closed() {
		win.Update()
		<-scrollrt.DrawSignal
		//fmt.Println("draw signal received")
	}
}

// ### [ Exports ] #############################################################

// --- [ engine ] --------------------------------------------------------------

//export CelDecodeFrame
func CelDecodeFrame(screenX, screenY int, celBuf unsafe.Pointer, frame, frameWidth int) {
	engine.CelDecodeFrame(screenX, screenY, celBuf, frame, frameWidth)
}

//export CelDecodeFrameIntoBuf
func CelDecodeFrameIntoBuf(dstBuf, celBuf unsafe.Pointer, frame, frameWidth int) {
	engine.CelDecodeFrameIntoBuf(dstBuf, celBuf, frame, frameWidth)
}

//export CelDecodeFrameWithHeader
func CelDecodeFrameWithHeader(screenX, screenY int, celBuf unsafe.Pointer, frame, frameWidth, always0, direction int) {
	engine.CelDecodeFrameWithHeader(screenX, screenY, celBuf, frame, frameWidth, always0, direction)
}

//export CelDecodeFrameWithHeaderIntoBuf
func CelDecodeFrameWithHeaderIntoBuf(dstBuf, celBuf unsafe.Pointer, frame, frameWidth, always0, direction int) {
	engine.CelDecodeFrameWithHeaderIntoBuf(dstBuf, celBuf, frame, frameWidth, always0, direction)
}

//export CelDecodeFrameWithLight
func CelDecodeFrameWithLight(screenX, screenY int, celBuf unsafe.Pointer, frame, frameWidth int) {
	engine.CelDecodeFrameWithLight(screenX, screenY, celBuf, frame, frameWidth)
}

//export CelDecodeFrameWithHeaderAndLight
func CelDecodeFrameWithHeaderAndLight(screenX, screenY int, celBuf unsafe.Pointer, frame, frameWidth, always0, direction int) {
	engine.CelDecodeFrameWithHeaderAndLight(screenX, screenY, celBuf, frame, frameWidth, always0, direction)
}

//export CelDecodeFrameWithHeaderLightAndTransparencyIntoBuf
func CelDecodeFrameWithHeaderLightAndTransparencyIntoBuf(dstBuf, celBuf unsafe.Pointer, frame, frameWidth, always0, direction int) {
	engine.CelDecodeFrameWithHeaderLightAndTransparencyIntoBuf(dstBuf, celBuf, frame, frameWidth, always0, direction)
}

//export CelDecodeFrameWithHeaderAndLightNotEquipable
func CelDecodeFrameWithHeaderAndLightNotEquipable(screenX, screenY int, celBuf unsafe.Pointer, frame, frameWidth, always0, direction int, always1 int8) {
	engine.CelDecodeFrameWithHeaderAndLightNotEquipable(screenX, screenY, celBuf, frame, frameWidth, always0, direction, always1)
}

//export CelDecodeFrameWithHeader2
func CelDecodeFrameWithHeader2(screenX, screenY int, celBuf unsafe.Pointer, frame, frameWidth, a6, direction int) {
	engine.CelDecodeFrameWithHeader2(screenX, screenY, celBuf, frame, frameWidth, a6, direction)
}

//export CelDecodeFrameWithHeaderIntoBuf2
func CelDecodeFrameWithHeaderIntoBuf2(dstBuf, celBuf unsafe.Pointer, frame, frameWidth, a5, direction int) {
	engine.CelDecodeFrameWithHeaderIntoBuf2(dstBuf, celBuf, frame, frameWidth, a5, direction)
}

//export CelDecodeFrameWithHeaderAndLight2
func CelDecodeFrameWithHeaderAndLight2(screenX, screenY int, celBuf unsafe.Pointer, frame, frameWidth, a6, direction int) {
	engine.CelDecodeFrameWithHeaderAndLight2(screenX, screenY, celBuf, frame, frameWidth, a6, direction)
}

//export CelDecodeFrameWithHeaderLightAndTransparencyIntoBuf2
func CelDecodeFrameWithHeaderLightAndTransparencyIntoBuf2(dstBuf, celBuf unsafe.Pointer, frame, frameWidth, a5, direction int) {
	engine.CelDecodeFrameWithHeaderLightAndTransparencyIntoBuf2(dstBuf, celBuf, frame, frameWidth, a5, direction)
}

//export CelDecodeFrameWithHeaderAndLightNotEquipable2
func CelDecodeFrameWithHeaderAndLightNotEquipable2(screenX, screenY int, celBuf unsafe.Pointer, frame, frameWidth, always0, direction int, always1 int8) {
	engine.CelDecodeFrameWithHeaderAndLightNotEquipable2(screenX, screenY, celBuf, frame, frameWidth, always0, direction, always1)
}

//export CelDecodeFrameIntoRectOfBuf
func CelDecodeFrameIntoRectOfBuf(dstBuf unsafe.Pointer, always0, dstHeight, dstWidth int, celBuf unsafe.Pointer, frame, frameWidth int) {
	engine.CelDecodeFrameIntoRectOfBuf(dstBuf, always0, dstHeight, dstWidth, celBuf, frame, frameWidth)
}

//export CelDecodeFrameWithColour
func CelDecodeFrameWithColour(colour uint8, screenX, screenY int, celBuf unsafe.Pointer, frame, frameWidth, a7, direction int) {
	engine.CelDecodeFrameWithColour(colour, screenX, screenY, celBuf, frame, frameWidth, a7, direction)
}

//export CelDecodeFrameWithHeaderAndColourHighlight
func CelDecodeFrameWithHeaderAndColourHighlight(colour uint8, screenX, screenY int, celBuf unsafe.Pointer, frame, frameWidth, a7, direction int) {
	engine.CelDecodeFrameWithHeaderAndColourHighlight(colour, screenX, screenY, celBuf, frame, frameWidth, a7, direction)
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
