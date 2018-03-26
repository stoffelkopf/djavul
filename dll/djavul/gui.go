//+build djavul

package main

import "C"

import (
	"unsafe"

	"github.com/sanctuary/djavul/d1/control"
	"github.com/sanctuary/djavul/d1/engine"
	"github.com/sanctuary/djavul/d1/multi"
	"github.com/sanctuary/djavul/d1/scrollrt"
	"github.com/sanctuary/djavul/d1/world"
)

// ### [ Exports ] #############################################################

// --- [ control ] -------------------------------------------------------------

//export DrawPanel
func DrawPanel(panelX, panelY, width, height, screenX, screenY int) {
	control.DrawPanel(panelX, panelY, width, height, screenX, screenY)
}

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

// --- [ multi ] ---------------------------------------------------------------

//export ProcessNetworkPackets
func ProcessNetworkPackets() {
	multi.ProcessNetworkPackets()
}

// --- [ scrollrt ] ------------------------------------------------------------

//export DrawMain
func DrawMain(height int, updateDescriptionBox, updateLifeOrb, updateManaOrb, updateBelt, updateControlButtons bool) {
	scrollrt.DrawMain(height, updateDescriptionBox, updateLifeOrb, updateManaOrb, updateBelt, updateControlButtons)
}

// --- [ world ] ---------------------------------------------------------------

//export DrawUpperScreen
func DrawUpperScreen(dstBuf unsafe.Pointer) {
	world.DrawUpperScreen(dstBuf)
}

//export DrawLowerScreen
func DrawLowerScreen(dstBuf unsafe.Pointer) {
	world.DrawLowerScreen(dstBuf)
}
