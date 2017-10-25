// Global variable wrappers for drlg_l1.cpp

package l1

import "unsafe"

// Global variables.
var (
	// Shadows contains shadows for 2x2 blocks of base tile IDs in the Cathedral.
	//
	// ref: 0x479C24
	Shadows = (*[37]Shadow)(unsafe.Pointer(uintptr(0x479C24)))

	// Base maps tile IDs to their corresponding base tile ID.
	//
	// ref: 0x479D28
	Base = (*[207]TileID)(unsafe.Pointer(uintptr(0x479D28)))

	// Plain maps tile IDs to their corresponding undecorated tile ID.
	//
	// ref: 0x479DF8
	Plain = (*[207]TileID)(unsafe.Pointer(uintptr(0x479DF8)))

	// MinisetStairUp1 is a 4x4 miniset of tile IDs representing a staircase
	// going up.
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
	// of the original map in blocks of 4.
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
)

// A Shadow contains the shadows for a 2x2 block of base tile IDs.
//
// PSX def:
//
//    typedef struct ShadowStruct {
//       unsigned char strig;
//       unsigned char s1;
//       unsigned char s2;
//       unsigned char s3;
//       unsigned char nv1;
//       unsigned char nv2;
//       unsigned char nv3;
//    } ShadowStruct;
type Shadow struct {
	// Shadow trigger base tile ID.
	BottomBase TileID
	TopBase    TileID
	RightBase  TileID
	LeftBase   TileID
	// Replacement tile IDs.
	Top   TileID
	Right TileID
	Left  TileID
}
