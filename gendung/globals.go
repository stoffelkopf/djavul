// Global variable wrappers for gendung.cpp

package gendung

import (
	"unsafe"

	"github.com/sanctuary/formats/level/til"
)

// Global variables.
var (
	// TileIDMap contains the tile IDs of the map.
	//
	// PSX ref: 0x800E40C4
	// PSX def: unsigned short dungeon[48][48]
	//
	// ref: 0x539608
	TileIDMap = (*[40][40]uint8)(unsafe.Pointer(uintptr(0x539608)))

	// TileIDMapBackup contains a backup of the tile IDs of the map.
	//
	// PSX ref: 0x800E52C4
	// PSX def: unsigned char pdungeon[40][40]
	//
	// ref: 0x53CD50
	TileIDMapBackup = (*[40][40]uint8)(unsafe.Pointer(uintptr(0x53CD50)))

	// PieceIDMap contains the piece IDs of each tile on the map.
	//
	// ref: 0x5A5BD8
	PieceIDMap = (*[112][112]int32)(unsafe.Pointer(uintptr(0x5A5BD8)))

	// TileDefs specifies the tile definitions of the active dungeon type; (e.g.
	// levels/l1data/l1.til).
	//
	// ref: 0x5B70DC
	TileDefs = (**til.Tile)(unsafe.Pointer(uintptr(0x5B70DC)))

	// DType specifies the active dungeon type of the current game.
	//
	//
	// ref: 0x5BB1ED
	DType = (*DungeonType)(unsafe.Pointer(uintptr(0x5BB1ED)))

	// DLvl specifies the active dungeon level of the current game.
	//
	// PSX ref: 0x8011C10C
	// PSX def: unsigned char currlevel
	//
	// ref: 0x5BB1EE
	DLvl = (*int8)(unsafe.Pointer(uintptr(0x5BB1EE)))

	// ArchNumMap contains the arch frame numbers of the map from the special
	// tileset (e.g. "levels/l1data/l1s.cel"). Note, the special tileset of
	// Tristram (i.e. "levels/towndata/towns.cel") contains trees rather than
	// arches.
	//
	// ref: 0x5C3008
	ArchNumMap = (*[112][112]int8)(unsafe.Pointer(uintptr(0x5C3008)))

	// SetXx specifies the active miniset x-coordinate of the map.
	//
	// PSX ref: 0x8011C0E4
	// PSX def: int setpc_x
	//
	// ref: 0x5CF338
	SetXx = (*int32)(unsafe.Pointer(uintptr(0x5CF338)))

	// LvlViewY specifies the level viewpoint y-coordinate of the map.
	//
	// PSX ref: 0x8011C130
	// PSX def: int LvlViewY
	//
	// ref: 0x5CF320
	LvlViewY = (*int32)(unsafe.Pointer(uintptr(0x5CF320)))

	// LvlViewX specifies the level viewpoint x-coordinate of the map.
	//
	// PSX ref: 0x8011C12C
	// PSX def: int LvlViewX
	//
	// ref: 0x5CF324
	LvlViewX = (*int32)(unsafe.Pointer(uintptr(0x5CF324)))

	// ViewX specifies the player viewpoint x-coordinate of the map.
	//
	// PSX ref: 0x8011C114
	// PSX def: int ViewX
	//
	// ref: 0x5CF33C
	ViewX = (*int32)(unsafe.Pointer(uintptr(0x5CF33C)))

	// ViewY specifies the player viewpoint y-coordinate of the map.
	//
	// PSX ref: 0x8011C118
	// PSX def: int ViewY
	//
	// ref: 0x5CF340
	ViewY = (*int32)(unsafe.Pointer(uintptr(0x5CF340)))

	// SetYy specifies the active miniset y-coordinate of the map.
	//
	// PSX ref: 0x8011C0E8
	// PSX def: int setpc_y
	//
	// ref: 0x5CF344
	SetYy = (*int32)(unsafe.Pointer(uintptr(0x5CF344)))
)
