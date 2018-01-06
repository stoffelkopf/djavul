//+build djavul

// Global variable wrappers for gendung.cpp

package gendung

import (
	"unsafe"

	"github.com/sanctuary/formats/level/min"
	"github.com/sanctuary/formats/level/til"
)

// Global variables.
var (
	// NPCNumMap contains the NPC numbers of the map. The NPC number represents a
	// towner number (towners array index) in Tristram and a monster number
	// (monsters array index) in the dungeon.
	//
	// ref: 0x52D208
	NPCNumMap = (*[112][112]int32)(unsafe.Pointer(uintptr(0x52D208)))

	// TileIDMap contains the tile IDs of the map.
	//
	// PSX ref: 0x800E40C4
	// PSX def: unsigned short dungeon[48][48]
	//
	// ref: 0x539608
	TileIDMap = (*[40][40]uint8)(unsafe.Pointer(uintptr(0x539608)))

	// ObjectNumMap contains the object numbers (objects array indices) of the
	// map.
	//
	// ref: 0x539C48
	ObjectNumMap = (*[112][112]int8)(unsafe.Pointer(uintptr(0x539C48)))

	// TileIDMapBackup contains a backup of the tile IDs of the map.
	//
	// PSX ref: 0x800E52C4
	// PSX def: unsigned char pdungeon[40][40]
	//
	// ref: 0x53CD50
	TileIDMapBackup = (*[40][40]uint8)(unsafe.Pointer(uintptr(0x53CD50)))

	// DeadMap contains the dead numbers (deads array indices) and dead direction
	// of the map, encoded as specified by the pseudo-code below.
	//
	//    deadNum   = DeadMap[col][row]&0x1F
	//    direction = DeadMap[col][row]>>5
	//
	// ref: 0x53D390
	DeadMap = (*[112][112]int8)(unsafe.Pointer(uintptr(0x53D390)))

	// TransparencyIndex specifies the current transparency category.
	//
	// PSX ref: 0x8011C148
	// PSX def: char TransVal
	//
	// ref: 0x5A5590
	TransparencyIndex = (*int8)(unsafe.Pointer(uintptr(0x5A5590)))

	// PieceIDMap contains the piece IDs of each tile on the map.
	//
	// ref: 0x5A5BD8
	PieceIDMap = (*[112][112]int32)(unsafe.Pointer(uintptr(0x5A5BD8)))

	// LightingVisibleDistanceMap specifies the visible distance of light
	// effects.
	//
	// ref: 0x5B1FD8
	LightingVisibleDistanceMap = (*[112][112]int8)(unsafe.Pointer(uintptr(0x5B1FD8)))

	// TileDefs specifies the tile definitions of the active dungeon type; (e.g.
	// levels/l1data/l1.til).
	//
	// ref: 0x5B70DC
	TileDefs = (**til.Tile)(unsafe.Pointer(uintptr(0x5B70DC)))

	// DPieceDefs specifies the dungeon piece definitions (a.k.a. miniture tiles)
	// of the active dungeon type; (e.g. levels/l1data/l1.min).
	//
	// ref: 0x5B70E0
	DPieceDefs = (**min.Block)(unsafe.Pointer(uintptr(0x5B70E0)))

	// Solid_0x10_0x20_0x40_FromPieceID maps from dungeon piece ID to solidity
	// with mask 0x70.
	//
	// ref: 0x5B70E8
	Solid_0x10_0x20_0x40_FromPieceID = (*[2049]int8)(unsafe.Pointer(uintptr(0x5B70E8)))

	// TransparencyMap specifies the transparency at each coordinate of the map.
	//
	// PSX ref: 0x800E7A28
	// PSX def: map_info dung_map[112][112] // dTransVal struct member
	//
	// ref: 0x5B78EC
	TransparencyMap = (*[112][112]int8)(unsafe.Pointer(uintptr(0x5B78EC)))

	// DType specifies the active dungeon type of the current game.
	//
	// PSX ref: 0x8011C10D
	// PSX def: unsigned char leveltype
	//
	// ref: 0x5BB1ED
	DType = (*DungeonType)(unsafe.Pointer(uintptr(0x5BB1ED)))

	// DLvl specifies the active dungeon level of the current game.
	//
	// PSX ref: 0x8011C10C
	// PSX def: unsigned char currlevel
	//
	// ref: 0x5BB1EE
	DLvl = (*uint8)(unsafe.Pointer(uintptr(0x5BB1EE)))

	// TransparencyActive specifies the active transparency indices.
	//
	// PSX ref: 0x800E7928
	// PSX def: unsigned char TransList[256]
	//
	// ref: 0x5BB1F0
	TransparencyActive = (*[256]bool)(unsafe.Pointer(uintptr(0x5BB1F0)))

	// LevelCEL points to the contents of the active tileset, which is one of
	// "levels/towndata/town.cel", "levels/l1data/l1.cel",
	// "levels/l2data/l2.cel", "levels/l3data/l3.cel" or "levels/l4data/l4.cel".
	//
	// ref: 0x5BDB0C
	LevelCEL = (**uint8)(unsafe.Pointer(uintptr(0x5BDB0C)))

	// PlayerNumMap contains the player numbers (players array indices) of the
	// map.
	//
	// ref: 0x5BFEF8
	PlayerNumMap = (*[112][112]int8)(unsafe.Pointer(uintptr(0x5BFEF8)))

	// ArchNumMap contains the arch frame numbers of the map from the special
	// tileset (e.g. "levels/l1data/l1s.cel"). Note, the special tileset of
	// Tristram (i.e. "levels/towndata/towns.cel") contains trees rather than
	// arches.
	//
	// ref: 0x5C3008
	ArchNumMap = (*[112][112]int8)(unsafe.Pointer(uintptr(0x5C3008)))

	// LevelSpecialCEL points to the contents of the active special tileset,
	// which is one of "levels/towndata/towns.cel", "levels/l1data/l1s.cel" or
	// "levels/l2data/l2s.cel".
	//
	// ref: 0x5C690C
	LevelSpecialCEL = (**uint8)(unsafe.Pointer(uintptr(0x5C690C)))

	// DFlagMap specifies flags used for dungeon generation.
	//
	// ref: 0x5C6910
	DFlagMap = (*[112][112]DFlag)(unsafe.Pointer(uintptr(0x5C6910)))

	// ItemNumMap contains the item numbers (items array indices) of the map.
	//
	// ref: 0x5C9A10
	ItemNumMap = (*[112][112]int8)(unsafe.Pointer(uintptr(0x5C9A10)))

	// SetHeight specifies the height of the active miniset of the map.
	//
	// PSX ref: 0x8011C0F0
	// PSX def: int setpc_h
	//
	// ref: 0x5CF330
	SetHeight = (*int32)(unsafe.Pointer(uintptr(0x5CF330)))

	// SetWidth specifies the width of the active miniset of the map.
	//
	// PSX ref: 0x8011C0EC
	// PSX def: int setpc_w
	//
	// ref: 0x5CF334
	SetWidth = (*int32)(unsafe.Pointer(uintptr(0x5CF334)))

	// SetXx specifies the active miniset x-coordinate of the map.
	//
	// PSX ref: 0x8011C0E4
	// PSX def: int setpc_x
	//
	// ref: 0x5CF338
	SetXx = (*int32)(unsafe.Pointer(uintptr(0x5CF338)))

	// IsQuestLevel specifies whether the active level is a quest level.
	//
	// ref: 0x5CF31D
	IsQuestLevel = (*bool)(unsafe.Pointer(uintptr(0x5CF31D)))

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

	// MissileNumMap contains the missile numbers (missiles array indices) of the
	// map.
	//
	// ref: 0x5CF350
	MissileNumMap = (*[112][112]int8)(unsafe.Pointer(uintptr(0x5CF350)))
)
