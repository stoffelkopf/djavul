//+build djavul

// Global variable wrappers for drlg_l1.cpp

package l1

import "unsafe"

// Global variables.
var (
	// Shadows contains shadows for 2x2 blocks of base tile IDs in the Cathedral.
	//
	// PSX ref: 0x8013A050
	// PSX def: ShadowStruct SPATS[37]
	//
	// ref: 0x479C24
	Shadows = (*[37]Shadow)(unsafe.Pointer(uintptr(0x479C24)))

	// Base maps tile IDs to their corresponding base tile ID.
	//
	// PSX ref: 0x8013A154
	// PSX def: unsigned char BSTYPES[206]
	//
	// ref: 0x479D28
	Base = (*[207]TileID)(unsafe.Pointer(uintptr(0x479D28)))

	// Plain maps tile IDs to their corresponding undecorated tile ID.
	//
	// PSX ref: 0x8013A224
	// PSX def: unsigned char L5BTYPES[206]
	//
	// ref: 0x479DF8
	Plain = (*[207]TileID)(unsafe.Pointer(uintptr(0x479DF8)))

	// MinisetStairUp1 is a 4x4 miniset of tile IDs representing a staircase
	// going up.
	//
	// PSX ref: 0x8013A2F4
	// PSX def: unsigned char STAIRSUP[34]
	//
	// Minisets specifies how to arrange tile IDs in order to form set areas of
	// dungeons (e.g. staircases). Below follows a pseudo-code description of the
	// miniset format.
	//
	//    // A miniset defines the set area of a dungeon in terms of before and
	//    // after areas of tile IDs.
	//    type miniset struct {
	//       // Width of miniset area.
	//       width uint8
	//       // Height of miniset area.
	//       height uint8
	//       // Tile IDs before transformation.
	//       before [width][height]uint8
	//       // Tile IDs after transformation.
	//       after [width][height]uint8
	//    }
	//
	// ref: 0x479EC8
	MinisetStairUp1 = (*[34]uint8)(unsafe.Pointer(uintptr(0x479EC8)))

	// MinisetStairUp2 is a 4x4 miniset of tile IDs representing a staircase
	// going up.
	//
	// PSX ref: 0x8013A318
	// PSX def: unsigned char L5STAIRSUP[34]
	//
	// Minisets specifies how to arrange tile IDs in order to form set areas of
	// dungeons (e.g. staircases). Below follows a pseudo-code description of the
	// miniset format.
	//
	//    // A miniset defines the set area of a dungeon in terms of before and
	//    // after areas of tile IDs.
	//    type miniset struct {
	//       // Width of miniset area.
	//       width uint8
	//       // Height of miniset area.
	//       height uint8
	//       // Tile IDs before transformation.
	//       before [width][height]uint8
	//       // Tile IDs after transformation.
	//       after [width][height]uint8
	//    }
	//
	// ref: 0x479EEC
	MinisetStairUp2 = (*[34]uint8)(unsafe.Pointer(uintptr(0x479EEC)))

	// MinisetStairDown is a 4x3 miniset of tile IDs representing a staircase
	// going down.
	//
	// PSX ref: 0x8013A33C
	// PSX def: unsigned char STAIRSDOWN[26]
	//
	// Minisets specifies how to arrange tile IDs in order to form set areas of
	// dungeons (e.g. staircases). Below follows a pseudo-code description of the
	// miniset format.
	//
	//    // A miniset defines the set area of a dungeon in terms of before and
	//    // after areas of tile IDs.
	//    type miniset struct {
	//       // Width of miniset area.
	//       width uint8
	//       // Height of miniset area.
	//       height uint8
	//       // Tile IDs before transformation.
	//       before [width][height]uint8
	//       // Tile IDs after transformation.
	//       after [width][height]uint8
	//    }
	//
	// ref: 0x479F10
	MinisetStairDown = (*[26]uint8)(unsafe.Pointer(uintptr(0x479F10)))

	// MinisetCandlestick is a 2x2 miniset of tile IDs representing a
	// candlestick.
	//
	// PSX ref: 0x8013A358
	// PSX def: unsigned char LAMPS[10]
	//
	// Minisets specifies how to arrange tile IDs in order to form set areas of
	// dungeons (e.g. staircases). Below follows a pseudo-code description of the
	// miniset format.
	//
	//    // A miniset defines the set area of a dungeon in terms of before and
	//    // after areas of tile IDs.
	//    type miniset struct {
	//       // Width of miniset area.
	//       width uint8
	//       // Height of miniset area.
	//       height uint8
	//       // Tile IDs before transformation.
	//       before [width][height]uint8
	//       // Tile IDs after transformation.
	//       after [width][height]uint8
	//    }
	//
	// ref: 0x479F2C
	MinisetCandlestick = (*[10]uint8)(unsafe.Pointer(uintptr(0x479F2C)))

	// MinisetStairDownPoison is a 6x6 miniset of tile IDs representing a
	// staircase going down to the Poisoned Water Supply.
	//
	// PSX ref: 0x8013A364
	// PSX def: unsigned char PWATERIN[74]
	//
	// Minisets specifies how to arrange tile IDs in order to form set areas of
	// dungeons (e.g. staircases). Below follows a pseudo-code description of the
	// miniset format.
	//
	//    // A miniset defines the set area of a dungeon in terms of before and
	//    // after areas of tile IDs.
	//    type miniset struct {
	//       // Width of miniset area.
	//       width uint8
	//       // Height of miniset area.
	//       height uint8
	//       // Tile IDs before transformation.
	//       before [width][height]uint8
	//       // Tile IDs after transformation.
	//       after [width][height]uint8
	//    }
	//
	// ref: 0x479F38
	MinisetStairDownPoison = (*[74]uint8)(unsafe.Pointer(uintptr(0x479F38)))

	// PatternLookup is a lookup table for the 16 possible patterns of a 2x2
	// area, where each cell either contains a SW door or it doesn't.
	//
	// PSX ref: 0x80139C58
	// PSX def: unsigned char L5ConvTbl[16]
	//
	// ref: 0x484778
	PatternLookup = (*[16]TileID)(unsafe.Pointer(uintptr(0x484778)))

	// TileBitMap represents a tile ID map of twice the size, repeating each tile
	// of the original map in blocks of 2x2.
	//
	// PSX ref: 0x8013A3B0
	// PSX def: unsigned char L5dungeon[80][80]
	//
	// ref: 0x525764
	TileBitMap = (*[80][80]TileID)(unsafe.Pointer(uintptr(0x525764)))

	// FlagMap contains flags used for dungeon generation of the Cathedral.
	//
	// PSX ref: 0x8011C0D8
	// PSX def: unsigned char* mydflags
	//
	// ref: 0x527064
	FlagMap = (*[40][40]Flag)(unsafe.Pointer(uintptr(0x527064)))

	// SinglePlayerQuestDunLoaded specifies whether a single player quest DUN has
	// been loaded.
	//
	// PSX ref: 0x8011C0F4
	// PSX def: unsigned char setloadflag
	//
	// ref: 0x5276A4
	SinglePlayerQuestDunLoaded = (*bool)(unsafe.Pointer(uintptr(0x5276A4)))

	// HorizRoom1 specifies whether to generate a horizontal room at position 1
	// in the Cathedral.
	//
	// PSX ref: 0x8011C8D8
	// PSX def: unsigned char HR1
	//
	// ref: 0x5276A8
	HorizRoom1 = (*bool)(unsafe.Pointer(uintptr(0x5276A8)))

	// HorizRoom2 specifies whether to generate a horizontal room at position 2
	// in the Cathedral.
	//
	// PSX ref: 0x8011C8D9
	// PSX def: unsigned char HR2
	//
	// ref: 0x5276AC
	HorizRoom2 = (*bool)(unsafe.Pointer(uintptr(0x5276AC)))

	// HorizRoom3 specifies whether to generate a horizontal room at position 3
	// in the Cathedral.
	//
	// PSX ref: 0x8011C8DA
	// PSX def: unsigned char HR3
	//
	// ref: 0x5276B0
	HorizRoom3 = (*bool)(unsafe.Pointer(uintptr(0x5276B0)))

	// VertRoom1 specifies whether to generate a vertical room at position 1 in
	// the Cathedral.
	//
	// PSX ref: 0x8011C8DB
	// PSX def: unsigned char VR1
	//
	// ref: 0x5276B4
	VertRoom1 = (*bool)(unsafe.Pointer(uintptr(0x5276B4)))

	// VertRoom2 specifies whether to generate a vertical room at position 2 in
	// the Cathedral.
	//
	// PSX ref: 0x8011C8DC
	// PSX def: unsigned char VR2
	//
	// ref: 0x5276B8
	VertRoom2 = (*bool)(unsafe.Pointer(uintptr(0x5276B8)))

	// VertRoom3 specifies whether to generate a vertical room at position 3 in
	// the Cathedral.
	//
	// PSX ref: 0x8011C8DD
	// PSX def: unsigned char VR3
	//
	// ref: 0x5276BC
	VertRoom3 = (*bool)(unsafe.Pointer(uintptr(0x5276BC)))

	// SinglePlayerQuestDun contains the contents of the single player quest
	// DUN file.
	//
	// PSX ref: 0x8011C0DC
	// PSX def: unsigned char* pSetPiece
	//
	// ref: 0x5276C0
	SinglePlayerQuestDun = (**uint8)(unsafe.Pointer(uintptr(0x5276C0)))
)
