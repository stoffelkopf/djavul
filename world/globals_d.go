//+build djavul

package world

import "unsafe"

// TODO: Move global variables to their dedicated source files, based on
// address.

var (
	// TileDrawMasks specifies the draw masks used to render transparency of
	// tiles.
	//
	// ref: 0x4B327D
	TileDrawMasks = (*[96]uint32)(unsafe.Pointer(uintptr(0x4B327D)))

	// Solid_0x10_0x20_0x40_FromPieceID maps from dungeon piece ID to solidity
	// with mask 0x70.
	//
	// ref: 0x5B70E8
	Solid_0x10_0x20_0x40_FromPieceID = (*[2049]int8)(unsafe.Pointer(uintptr(0x5B70E8)))

	// SpeedCelFrameNumFromLightIndexFrameNum returns the frame number of the
	// speed CEL, an in memory decoding of level CEL frames, based on original
	// frame number and light index.
	//
	// Note, given light index 0, the original frame number is returned.
	//
	// ref: 0x5BDB10
	SpeedCelFrameNumFromLightIndexFrameNum = (*[128][16]int32)(unsafe.Pointer(uintptr(0x5BDB10)))

	// LightingMax specifies the maximum number of light entries.
	//
	// ref: 0x642A14
	LightingMax = (*int8)(unsafe.Pointer(uintptr(0x642A14)))

	// LightTableIndex specifies the current light entry.
	//
	// ref: 0x69BEF8
	LightTableIndex = (*int32)(unsafe.Pointer(uintptr(0x69BEF8)))

	// LevelCelBlock specifies the current MIN block of the level CEL file, as
	// used during rendering of the level tiles.
	//
	//    frameNum  := block&0x0FFF
	//    frameType := block&0x7000 >> 12
	//
	// ref: 0x69CF14
	LevelCelBlock = (*uint32)(unsafe.Pointer(uintptr(0x69CF14)))

	// ref: 0x69CF20
	LevelArchTypeSomething = (*int8)(unsafe.Pointer(uintptr(0x69CF20)))

	// CelTransparencyActive specifies whether transparency is active for the
	// current CEL file being decoded.
	//
	// ref: 0x69CF94
	CelTransparencyActive = (*int32)(unsafe.Pointer(uintptr(0x69CF94)))

	// LevelPieceID specifies the current dungeon piece ID of the level, as used
	// during rendering of the level tiles.
	//
	// ref: 0x69CF98
	LevelPieceID = (*int32)(unsafe.Pointer(uintptr(0x69CF98)))
)
