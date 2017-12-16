//+build djavul

package engine

import (
	"fmt"
	"log"
	"unsafe"

	"github.com/faiface/pixel"
	"github.com/sanctuary/djavul/dx"
	"github.com/sanctuary/djavul/internal/proto"
)

// celDecodeFrame decodes the given frame to the specified screen coordinate.
//
//    x = screenX - 64
//    y = screenY - 160
//    frameNum = frame - 1
//
// Note, the coordinates specify the bottom left corner.
//
// ref: 0x416274
func celDecodeFrame(screenX, screenY int, celBuf unsafe.Pointer, frame, frameWidth int) {
	file := getFile(celBuf)
	frameNum := frame - 1
	x, y := XYFromScreenCoords(screenX, screenY)
	if err := proto.SendDrawImage(file, x, y, frameNum); err != nil {
		log.Fatalf("%+v", err)
	}
}

// celDecodeFrameIntoBuf decodes the given CEL frame into the specified buffer.
//
// Note, this function is only used to decode CEL images without frame headers
// (pentspn2.cel).
//
// ref: 0x4162B8
func celDecodeFrameIntoBuf(dstBuf, celBuf unsafe.Pointer, frame, frameWidth int) {
	screenX, screenY := CalcScreenCoords(dstBuf)
	celDecodeFrame(screenX, screenY, celBuf, frame, frameWidth)
}

// celDecodeFrameWithHeader decodes the given CEL frame to the specified screen
// coordinate.
//
//    x = screen_x - 64
//    y = screen_y - 160
//    frameNum = frame - 1
//
// Note, the coordinates specify the bottom left corner.
//
// Note, this function is only used to decode CEL images with frame headers
// (objects, item drops, objcurs.cel, towners).
//
// ref: 0x4162DE
func celDecodeFrameWithHeader(screenX, screenY int, celBuf unsafe.Pointer, frame, frameWidth, always0, direction int) {
	relPath := getFile(celBuf)
	frameNum := frame - 1
	x, y := XYFromScreenCoords(screenX, screenY)
	if err := proto.SendDrawImage(relPath, x, y, frameNum); err != nil {
		log.Fatalf("%+v", err)
	}
}

// celDecodeFrameWithHeaderIntoBuf decodes the given CEL frame into the
// specified buffer.
//
// Note, this function is only used to decode CEL images with frame headers
// (square.cel).
//
// ref: 0x416359
func celDecodeFrameWithHeaderIntoBuf(dstBuf, celBuf unsafe.Pointer, frame, frameWidth, always0, direction int) {
	screenX, screenY := CalcScreenCoords(dstBuf)
	celDecodeFrame(screenX, screenY, celBuf, frame, frameWidth)
}

// celDecodeFrameWithLight decodes the given CEL frame to the specified screen
// coordinate, adding lighting if applicable.
//
//    x = screen_x - 64
//    y = screen_y - 160
//    frameNum = frame - 1
//
// Note, the coordinates specify the bottom left corner.
//
// Note, this function is only used to decode CEL images without frame headers
// (bigtgold.cel).
//
// ref: 0x416565
func celDecodeFrameWithLight(screenX, screenY int, celBuf unsafe.Pointer, frame, frameWidth int) {
	celDecodeFrame(screenX, screenY, celBuf, frame, frameWidth)
}

// celDecodeFrameWithHeaderAndLight decodes the given CEL frame to the specified
// screen coordinate, adding lighting if applicable.
//
//    x = screen_x - 64
//    y = screen_y - 160
//    frameNum = frame - 1
//
// Note, the coordinates specify the bottom left corner.
//
// Note, this function is only used to decode CEL images with frame headers
// (item drops, objects).
//
// ref: 0x4165BD
func celDecodeFrameWithHeaderAndLight(screenX, screenY int, celBuf unsafe.Pointer, frame, frameWidth, always0, direction int) {
	celDecodeFrame(screenX, screenY, celBuf, frame, frameWidth)
}

// celDecodeFrameWithHeaderLightAndTransparencyIntoBuf decodes the given CEL
// frame into the specified buffer with added lighting and transparency.
//
// Note, this function is only used to decode CEL images with frame headers
// (objcurs.cel, level special).
//
// ref: 0x41664B
func celDecodeFrameWithHeaderLightAndTransparencyIntoBuf(dstBuf, celBuf unsafe.Pointer, frame, frameWidth, always0, direction int) {
	screenX, screenY := CalcScreenCoords(dstBuf)
	celDecodeFrame(screenX, screenY, celBuf, frame, frameWidth)
}

// celDecodeFrameWithHeaderAndLightNotEquipable decodes the given CEL frame to
// the specified screen coordinate with added lighting.
//
//    x = screen_x - 64
//    y = screen_y - 160
//    frameNum = frame - 1
//
// Note, the coordinates specify the bottom left corner.
//
// Note, this function is only used to decode CEL images with frame headers
// (objcurs.cel).
//
// ref: 0x4166BF
func celDecodeFrameWithHeaderAndLightNotEquipable(screenX, screenY int, celBuf unsafe.Pointer, frame, frameWidth, always0, direction int, always1 int8) {
	celDecodeFrame(screenX, screenY, celBuf, frame, frameWidth)
}

// celDecodeFrameWithHeader2 decodes the given CEL frame to the specified screen
// coordinate.
//
//    x = screen_x - 64
//    y = screen_y - 160
//    frameNum = frame - 1
//
// Note, the coordinates specify the bottom left corner.
//
// Note, this function is only used to decode CEL images with frame headers
// (objects, objcurs.cel, item drops, towners).
//
// Note, D1DrawCursorTown (from RE Notes)
//
// ref: 0x41685A
func celDecodeFrameWithHeader2(screenX, screenY int, celBuf unsafe.Pointer, frame, frameWidth, a6, direction int) {
	celDecodeFrame(screenX, screenY, celBuf, frame, frameWidth)
}

// celDecodeFrameWithHeaderIntoBuf2 decodes the given CEL frame into the
// specified buffer.
//
// Note, this function is only used to decode CEL images with frame headers
// (square.cel).
//
// ref: 0x4168D5
func celDecodeFrameWithHeaderIntoBuf2(dstBuf, celBuf unsafe.Pointer, frame, frameWidth, a5, direction int) {
	screenX, screenY := CalcScreenCoords(dstBuf)
	celDecodeFrame(screenX, screenY, celBuf, frame, frameWidth)
}

// celDecodeFrameWithHeaderAndLight2 decodes the given CEL frame to the
// specified screen coordinate, adding lighting if applicable.
//
//    x = screen_x - 64
//    y = screen_y - 160
//    frameNum = frame - 1
//
// Note, the coordinates specify the bottom left corner.
//
// Note, this function is only used to decode CEL images with frame headers
// (item drops, objects).
//
// Note, D1DrawObjectBaseDungeon (from RE Notes).
//
// ref: 0x416B19
func celDecodeFrameWithHeaderAndLight2(screenX, screenY int, celBuf unsafe.Pointer, frame, frameWidth, a6, direction int) {
	celDecodeFrame(screenX, screenY, celBuf, frame, frameWidth)
}

// celDecodeFrameWithHeaderLightAndTransparencyIntoBuf2 decodes the given CEL
// frame into the specified buffer with added lighting and transparency.
//
// Note, this function is only used to decode CEL images with frame headers
// (level special).
//
// Note, D1DrawArchTile (from RE Notes).
//
// ref: 0x416BA9
func celDecodeFrameWithHeaderLightAndTransparencyIntoBuf2(dstBuf, celBuf unsafe.Pointer, frame, frameWidth, a5, direction int) {
	screenX, screenY := CalcScreenCoords(dstBuf)
	celDecodeFrame(screenX, screenY, celBuf, frame, frameWidth)
}

// celDecodeFrameWithHeaderAndLightNotEquipable2 decodes the given CEL frame to
// the specified screen coordinate with added lighting.
//
//    x = screen_x - 64
//    y = screen_y - 160
//    frameNum = frame - 1
//
// Note, the coordinates specify the bottom left corner.
//
// Note, this function is only used to decode CEL images with frame headers (objcurs.cel).
//
// ref: 0x416C1B
func celDecodeFrameWithHeaderAndLightNotEquipable2(screenX, screenY int, celBuf unsafe.Pointer, frame, frameWidth, always0, direction int, always1 int8) {
	celDecodeFrame(screenX, screenY, celBuf, frame, frameWidth)
}

// celDecodeFrameIntoRectOfBuf decodes the given CEL frame into a rectangle of
// the specified buffer.
//
// Note, this function is only used to decode CEL images without frame headers
// (control panel and orbs).
//
// ref: 0x416D3C
func celDecodeFrameIntoRectOfBuf(dstBuf unsafe.Pointer, always0, dstHeight, dstWidth int, celBuf unsafe.Pointer, frame, frameWidth int) {
	// TODO: implement
	fmt.Println("celDecodeFrameIntoRectOfBuf not yet implemented")
}

// celDecodeFrameWithColour decodes the given CEL frame to the specified screen
// coordinate and with the specified colour.
//
//    x = screen_x - 64
//    y = screen_y - 160
//    frameNum = frame - 1
//
// Note, the coordinates specify the bottom left corner.
//
// Note, this function is only used to decode CEL images with frame headers
// (objcurs.cel, item drops, objects, towners).
//
// ref: 0x416DC6
func celDecodeFrameWithColour(colour uint8, screenX, screenY int, celBuf unsafe.Pointer, frame, frameWidth, a7, direction int) {
	// TODO: handle colour.
	celDecodeFrame(screenX, screenY, celBuf, frame, frameWidth)
}

// celDecodeFrameWithHeaderAndColourHighlight decodes the given CEL frame to the
// specified screen coordinate with added colour highlight.
//
//    x = screen_x - 64
//    y = screen_y - 160
//    frameNum = frame - 1
//
// Note, the coordinates specify the bottom left corner.
//
// Note, this function is only used to decode CEL images with frame headers
// (objcurs.cel, item drops, objects, towners).
//
// NOTE: D1DrawObjectHilight (from RE Notes).
//
// ref: 0x416EC0
func celDecodeFrameWithHeaderAndColourHighlight(colour uint8, screenX, screenY int, celBuf unsafe.Pointer, frame, frameWidth, a7, direction int) {
	// TODO: handle colour.
	celDecodeFrame(screenX, screenY, celBuf, frame, frameWidth)
}

// ### [ Helper functions ] ####################################################

// CalcScreenCoords returns the screen x- and y-coordinates based on the given
// destination buffer.
func CalcScreenCoords(dstBuf unsafe.Pointer) (screenX, screenY int) {
	addr := uintptr(dstBuf)
	start := uintptr(unsafe.Pointer(*dx.ScreenBuf))
	offset := int(addr - start)
	screenX = offset % (64 + 640 + 64)
	screenY = offset / (64 + 640 + 64)
	return screenX, screenY
}

// CalcXY returns the x- and y-coordinates based on the given destination
// buffer.
func CalcXY(dstBuf unsafe.Pointer) (x, y int) {
	screenX, screenY := CalcScreenCoords(dstBuf)
	return XYFromScreenCoords(screenX, screenY)
}

// XYFromScreenCoords returns the x- and y-coordinates based on the given screen
// x- and y-coordinates.
func XYFromScreenCoords(screenX, screenY int) (x, y int) {
	x = screenX - 64
	y = screenY - 160
	return x, y
}

// pictures maps from relative file path to decoded image frames.
var pictures = make(map[string][]pixel.Picture)

// getPictures returns the pictures associated with the given file path.
func getPictures(relPath string) []pixel.Picture {
	pics, ok := pictures[relPath]
	if !ok {
		panic(fmt.Errorf("unable to locate decoded image frames of %q", relPath))
	}
	return pics
}

// dirPictures maps from relative file path to decoded image frames based on
// direction.
var dirPictures = make(map[string][][]pixel.Picture)

// getPicturesForDir returns the pictures associated with the given file path
// and direction.
func getPicturesForDir(relPath string, direction int) []pixel.Picture {
	dirPics, ok := dirPictures[relPath]
	if !ok {
		panic(fmt.Errorf("unable to locate decoded image frames of %q", relPath))
	}
	if direction == 8 {
		direction = 0
	}
	if len(dirPics) <= direction {
		panic(fmt.Errorf("invalid direction for %q; expected < %d, got %d", relPath, len(dirPics), direction))
	}
	return dirPics[direction]
}
