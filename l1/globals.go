// Global variable wrappers for drlg_l1.cpp

package l1

import "unsafe"

// Global variables.
var (
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

	// FlagMap contains flags used for dungeon generation of the Cathedral.
	//
	// PSX ref: 0x8011C0D8
	// PSX def: unsigned char* mydflags
	//
	// ref: 0x527064
	FlagMap = (*[40][40]uint8)(unsafe.Pointer(uintptr(0x527064)))
)
