//+build !djavul

// Global variable wrappers for gendung.cpp

package gendung

import (
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
	NPCNumMap = new([112][112]int32)

	// TileIDMap contains the tile IDs of the map.
	//
	// PSX ref: 0x800E40C4
	// PSX def: unsigned short dungeon[48][48]
	//
	// ref: 0x539608
	TileIDMap = new([40][40]uint8)

	// ObjectNumMap contains the object numbers (objects array indices) of the
	// map.
	//
	// ref: 0x539C48
	ObjectNumMap = new([112][112]int8)

	// TileIDMapBackup contains a backup of the tile IDs of the map.
	//
	// PSX ref: 0x800E52C4
	// PSX def: unsigned char pdungeon[40][40]
	//
	// ref: 0x53CD50
	TileIDMapBackup = new([40][40]uint8)

	// DeadMap contains the dead numbers (deads array indices) and dead direction
	// of the map, encoded as specified by the pseudo-code below.
	//
	//    deadNum   = DeadMap[col][row]&0x1F
	//    direction = DeadMap[col][row]>>5
	//
	// ref: 0x53D390
	DeadMap = new([112][112]int8)

	// TransparencyIndex specifies the current transparency category.
	//
	// PSX ref: 0x8011C148
	// PSX def: char TransVal
	//
	// ref: 0x5A5590
	TransparencyIndex = new(int8)

	// PieceIDMap contains the piece IDs of each tile on the map.
	//
	// ref: 0x5A5BD8
	PieceIDMap = new([112][112]int32)

	// LightingVisibleDistanceMap specifies the visible distance of light
	// effects.
	//
	// ref: 0x5B1FD8
	LightingVisibleDistanceMap = new([112][112]int8)

	// TileDefs specifies the tile definitions of the active dungeon type; (e.g.
	// levels/l1data/l1.til).
	//
	// ref: 0x5B70DC
	TileDefs = new(*til.Tile)

	// DPieceDefs specifies the dungeon piece definitions (a.k.a. miniture tiles)
	// of the active dungeon type; (e.g. levels/l1data/l1.min).
	//
	// ref: 0x5B70E0
	DPieceDefs = new(*min.Block)

	// TransparencyMap specifies the transparency at each coordinate of the map.
	//
	// PSX ref: 0x800E7A28
	// PSX def: map_info dung_map[112][112] // dTransVal struct member
	//
	// ref: 0x5B78EC
	TransparencyMap = new([112][112]int8)

	// DType specifies the active dungeon type of the current game.
	//
	// PSX ref: 0x8011C10D
	// PSX def: unsigned char leveltype
	//
	// ref: 0x5BB1ED
	DType = new(DungeonType)

	// DLvl specifies the active dungeon level of the current game.
	//
	// PSX ref: 0x8011C10C
	// PSX def: unsigned char currlevel
	//
	// ref: 0x5BB1EE
	DLvl = new(uint8)

	// TransparencyActive specifies the active transparency indices.
	//
	// PSX ref: 0x800E7928
	// PSX def: unsigned char TransList[256]
	//
	// ref: 0x5BB1F0
	TransparencyActive = new([256]bool)

	// LevelCEL points to the contents of the active tileset, which is one of
	// "levels/towndata/town.cel", "levels/l1data/l1.cel",
	// "levels/l2data/l2.cel", "levels/l3data/l3.cel" or "levels/l4data/l4.cel".
	//
	// ref: 0x5BDB0C
	LevelCEL = new(*uint8)

	// PlayerNumMap contains the player numbers (players array indices) of the
	// map.
	//
	// ref: 0x5BFEF8
	PlayerNumMap = new([112][112]int8)

	// ArchNumMap contains the arch frame numbers of the map from the special
	// tileset (e.g. "levels/l1data/l1s.cel"). Note, the special tileset of
	// Tristram (i.e. "levels/towndata/towns.cel") contains trees rather than
	// arches.
	//
	// ref: 0x5C3008
	ArchNumMap = new([112][112]int8)

	// LevelSpecialCEL points to the contents of the active special tileset,
	// which is one of "levels/towndata/towns.cel", "levels/l1data/l1s.cel" or
	// "levels/l2data/l2s.cel".
	//
	// ref: 0x5C690C
	LevelSpecialCEL = new(*uint8)

	// DFlagMap specifies flags used for dungeon generation.
	//
	// ref: 0x5C6910
	DFlagMap = new([112][112]DFlag)

	// ItemNumMap contains the item numbers (items array indices) of the map.
	//
	// ref: 0x5C9A10
	ItemNumMap = new([112][112]int8)

	// SetXx specifies the active miniset x-coordinate of the map.
	//
	// PSX ref: 0x8011C0E4
	// PSX def: int setpc_x
	//
	// ref: 0x5CF338
	SetXx = new(int32)

	// IsQuestLevel specifies whether the active level is a quest level.
	//
	// ref: 0x5CF31D
	IsQuestLevel = new(bool)

	// LvlViewY specifies the level viewpoint y-coordinate of the map.
	//
	// PSX ref: 0x8011C130
	// PSX def: int LvlViewY
	//
	// ref: 0x5CF320
	LvlViewY = new(int32)

	// LvlViewX specifies the level viewpoint x-coordinate of the map.
	//
	// PSX ref: 0x8011C12C
	// PSX def: int LvlViewX
	//
	// ref: 0x5CF324
	LvlViewX = new(int32)

	// ViewX specifies the player viewpoint x-coordinate of the map.
	//
	// PSX ref: 0x8011C114
	// PSX def: int ViewX
	//
	// ref: 0x5CF33C
	ViewX = new(int32)

	// ViewY specifies the player viewpoint y-coordinate of the map.
	//
	// PSX ref: 0x8011C118
	// PSX def: int ViewY
	//
	// ref: 0x5CF340
	ViewY = new(int32)

	// SetYy specifies the active miniset y-coordinate of the map.
	//
	// PSX ref: 0x8011C0E8
	// PSX def: int setpc_y
	//
	// ref: 0x5CF344
	SetYy = new(int32)

	// MissileNumMap contains the missile numbers (missiles array indices) of the
	// map.
	//
	// ref: 0x5CF350
	MissileNumMap = new([112][112]int8)
)
