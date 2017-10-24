// Package l1 implements dynamic random level generation of cathedral maps.
package l1

import (
	"reflect"
	"unsafe"

	"github.com/pkg/errors"
	"github.com/sanctuary/djavul/engine"
	"github.com/sanctuary/djavul/gendung"
	"github.com/sanctuary/djavul/quests"
	"github.com/sanctuary/formats/level/til"
)

// initPieceIDMap initializes the dungeon piece ID map.
//
// PSX ref: 0x8013CBA8
// PSX sig: void DRLG_L1Pass3__Fv()
//
// ref: 0x40AFB3
func initPieceIDMap() {
	// Initialize the entire dungeon piece ID map with dirt.
	tiles := getTiles()
	tile := tiles[Dirt-1]
	for x := 0; x < 112-1; x += 2 {
		for y := 0; y < 112-1; y += 2 {
			gendung.PieceIDMap[x][y] = int32(tile.Top) + 1
			gendung.PieceIDMap[x+1][y] = int32(tile.Right) + 1
			gendung.PieceIDMap[x][y+1] = int32(tile.Left) + 1
			gendung.PieceIDMap[x+1][y+1] = int32(tile.Bottom) + 1
		}
	}
	// Initialize the visible tiles of the dungeon piece ID map based on the tile
	// ID map. The visible tiles are located at (16, 16) <= coordinate < (96,
	// 96).
	x := 16
	for xx := 0; xx < 40; xx++ {
		y := 16
		for yy := 0; yy < 40; yy++ {
			tileID := (*gendung.TileIDMap)[xx][yy]
			if tileID == 0 {
				panic(errors.Errorf("uninitialized tile ID at (%d, %d)", xx, yy))
			}
			tile := tiles[tileID-1]
			gendung.PieceIDMap[x][y] = int32(tile.Top) + 1
			gendung.PieceIDMap[x+1][y] = int32(tile.Right) + 1
			gendung.PieceIDMap[x][y+1] = int32(tile.Left) + 1
			gendung.PieceIDMap[x+1][y+1] = int32(tile.Bottom) + 1
			y += 2
		}
		x += 2
	}
}

// initArches initializes arches.
//
// PSX ref: 0x8013CF5C
// PSX sig: void DRLG_InitL1Vals__Fv()
//
// ref: 0x40B0A5
func initArches() {
	for x := 0; x < 112; x++ {
		for y := 0; y < 112; y++ {
			switch DPieceID(gendung.PieceIDMap[x][y]) {
			case ArchSwArchSe_left, ArchSwDoorSe_left, BloodArchSw_left, ArchSwShadowArchSeLeft_left, ArchSwWallSe3_left, EnteranceSw1_left:
				gendung.ArchNumMap[x][y] = int8(ArchIDSw)
			case ArchSwArchSe_right, BrokenArchSe_right, ArchSeShadowArchSwRight_right, ArchSeShadowBarSwRight_right, WallSw3ArchSe_right, EnteranceSe1_right:
				gendung.ArchNumMap[x][y] = int8(ArchIDSe)
			case BrokenArchSw1_left:
				gendung.ArchNumMap[x][y] = int8(ArchIDSwBroken2)
			case BrokenArchSw2_left:
				gendung.ArchNumMap[x][y] = int8(ArchIDSw2)
			}
		}
	}
}

// createDungeon creates a random cathedral dungeon based on the given seed and
// level entry.
//
// PSX ref: 0x80140E64
// PSX sig: void CreateL5Dungeon__FUii(unsigned int rseed, int entry)
//
// ref: 0x40B229
func createDungeon(seed, entry int32) {
	engine.SetSeed(seed)
	gendung.InitTransparency() // TODO: add test case
	gendung.InitSetPiece()     // TODO: add test case
	LoadQuestDun()             // TODO: add test case
	GenerateDungeon(entry)
	InitPieceIDMap()
	FreeQuestDun() // NOTE: not tested; only used for cleanup
	InitArches()
	gendung.MarkSetPiece() // TODO: add test case
}

// generateDungeon generates a cathedral dungeon based on the given level entry.
//
// PSX ref: 0x80140930
// PSX sig: void DRLG_L5__Fi(int entry)
//
// ref: 0x40B306
func generateDungeon(entry int32) {
	// Later levels should contain more walls.
	var minArea int
	switch *gendung.DLvl {
	case 1:
		minArea = 533
	case 2:
		minArea = 693
	case 3, 4:
		minArea = 761
	}
	for {
		gendung.InitTransparency() // TODO: add test case

		// Regenerate rooms until the minimum number of walls are present on the
		// map.
		for {
			Reset()
			GenerateFirstRoom()
			if GetArea() >= minArea {
				break
			}
		}

		// Generate chambers and add tile ID patterns.
		InitDmap()
		GeneratePattern()
		GenerateChambers()
		FixTiles()
		AddWall()
		ClearFlags()
		FloorTransparency() // TODO: add test case

		// Place staircases.
		done := true
		// Place staircase down to the Poisoned Water Supply quest level.
		if quests.IsActive(quests.PoisonedWaterSupply) {
			setView := entry == 0
			if PlaceMiniset(unsafe.Pointer(MinisetStairDownPoison), 1, 1, 0, 0, setView, -1, 0) < 0 {
				done = false
			}
			if entry != 0 {
				*gendung.ViewY--
			}
		}

		// Place staircase up from the Odgen's Sign quest level. Note, the down
		// staircase is already included in the DUN file.
		if quests.IsActive(quests.OgdensSign) {
			setView := entry == 0
			if PlaceMiniset(unsafe.Pointer(MinisetStairUp1), 1, 1, 0, 0, setView, -1, 0) < 0 {
				done = false
			}
			if entry == 1 {
				*gendung.ViewX = *gendung.SetXx*2 + 20
				*gendung.ViewY = *gendung.SetYy*2 + 28
			} else if entry != 0 {
				*gendung.ViewY--
			}
			if done {
				break
			}
			continue
		}

		// Place staircases up from and down to level.
		if entry != 0 {
			// Place staircase up and down.
			if PlaceMiniset(unsafe.Pointer(MinisetStairUp2), 1, 1, 0, 0, false, -1, 0) < 0 || PlaceMiniset(unsafe.Pointer(MinisetStairDown), 1, 1, 0, 0, true, -1, 1) < 0 {
				done = false
			}
			*gendung.ViewY--
			if done {
				break
			}
		} else {
			// Place staircase up.
			if PlaceMiniset(unsafe.Pointer(MinisetStairUp2), 1, 1, 0, 0, true, -1, 0) < 0 {
				// Try again.
				continue
			}
			// Place staircase down.
			if PlaceMiniset(unsafe.Pointer(MinisetStairDown), 1, 1, 0, 0, false, -1, 1) < 0 {
				done = false
			}
			if done {
				break
			}
		}
	}

	// Fix transparency of staircases.
	yy := 0
	for y := 16; y < 96; y += 2 {
		xx := 0
		for x := 16; x < 96; x += 2 {
			if TileID(gendung.TileIDMap[xx][yy]) == StairB2 {
				gendung.CopyTransparency(x, y+1, x, y)     // TODO: add test case
				gendung.CopyTransparency(x+1, y+1, x+1, y) // TODO: add test case
			}
			xx++
		}
		yy++
	}

	// Fix transparency, dirt and corners.
	FixTransparency() // TODO: add test case
	FixDirt()
	FixCorners() // TODO: add test case

	// Place doors.
	for yy := 0; yy < 40; yy++ {
		for xx := 0; xx < 40; xx++ {
			if FlagMap[xx][yy]&0x7F != 0 {
				PlaceDoor(xx, yy)
			}
		}
	}

	// Decorate, add shadows, place candlesticks and randomize floor tiles.
	Decorate()
	InitShadows()
	PlaceMiniset(unsafe.Pointer(MinisetCandlestick), 5, 10, 0, 0, false, -1, 4)
	RandomizeStoneFloor()

	// Make backup of tile ID map.
	for yy := 0; yy < 40; yy++ {
		for xx := 0; xx < 40; xx++ {
			gendung.TileIDMapBackup[xx][yy] = gendung.TileIDMap[xx][yy]
		}
	}

	// Reset the dungeon flag, player, NPC, dead, object, item, missile and arch
	// maps.
	ResetMaps() // TODO: add test case

	// Initialize quest area.
	quests.InitQuestArea(*gendung.SetXx, *gendung.SetYy) // TODO: add test case
}

// placeDoor places a door at the given coordinate.
//
// PSX ref: 0x8013BCB0
// PSX sig: void DRLG_PlaceDoor__Fii(int x, int y)
//
// ref: 0x40B56F
func placeDoor(xx, yy int) {
	if FlagMap[xx][yy]&FlagDone != 0 {
		FlagMap[xx][yy] = FlagDone
		return
	}
	flag := FlagMap[xx][yy] &^ FlagDone
	tileID := TileID(gendung.TileIDMap[xx][yy])
	if flag == FlagFavourSe || flag == FlagFavourSw {
		if xx != 1 {
			switch tileID {
			case WallSw:
				gendung.TileIDMap[xx][yy] = uint8(DoorSw)
			case WallEndSw:
				gendung.TileIDMap[xx][yy] = uint8(DoorEndSw)
			case WallSwArchSe:
				gendung.TileIDMap[xx][yy] = uint8(DoorSwArchSe)
			}
		}
		if yy != 1 {
			switch tileID {
			case WallSe:
				gendung.TileIDMap[xx][yy] = uint8(DoorSe)
			case WallEndSe:
				gendung.TileIDMap[xx][yy] = uint8(DoorEndSe)
			case ArchSwWallSe:
				gendung.TileIDMap[xx][yy] = uint8(ArchSwDoorSe)
			}
		}
	}
	if tileID == WallSwWallSe {
		switch flag {
		case FlagFavourSe:
			if yy != 1 {
				gendung.TileIDMap[xx][yy] = uint8(WallSwDoorSe)
			}
		case FlagFavourSw:
			if xx != 1 {
				gendung.TileIDMap[xx][yy] = uint8(DoorSwWallSe)
			}
		case FlagFavourSe | FlagFavourSw:
			if xx != 1 && yy != 1 {
				gendung.TileIDMap[xx][yy] = uint8(DoorSwDoorSe)
			}
		}
	}
	FlagMap[xx][yy] = FlagDone
}

// reset resets the tile ID and the dungeon flag maps.
//
// PSX ref: 0x8013D2F8
// PSX sig: void InitL5Dungeon__Fv()
//
// ref: 0x40BAF6
func reset() {
	for xx := 0; xx < 40; xx++ {
		for yy := 0; yy < 40; yy++ {
			gendung.TileIDMap[xx][yy] = 0
			FlagMap[xx][yy] = 0
		}
	}
}

// ### [ Helper functions ] ####################################################

// getTiles returns the tileset of the active dungeon type.
func getTiles() []til.Tile {
	// The tileset of town contains 342 tiles, l1 206, l2 160, l3 156, and l4
	// 137.
	var n int
	switch *gendung.DType {
	case gendung.Tristram:
		n = 342
	case gendung.Cathedral:
		n = 206
	case gendung.Catacombs:
		n = 160
	case gendung.Caves:
		n = 156
	case gendung.Hell:
		n = 137
	default:
		panic(errors.Errorf("invalid dungeon type %d", *gendung.DType))
	}
	data := (uintptr)(unsafe.Pointer(*gendung.TileDefs))
	sh := &reflect.SliceHeader{Data: data, Len: n, Cap: n}
	return *(*[]til.Tile)(unsafe.Pointer(sh))
}
