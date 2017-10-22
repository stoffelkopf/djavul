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
)
