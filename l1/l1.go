// Package l1 implements dynamic random level generation of cathedral maps.
package l1

import (
	"unsafe"

	"github.com/sanctuary/djavul/gendung"
	"github.com/sanctuary/djavul/quests"
)

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
		gendung.InitTransparency()

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
		FloorTransparency()

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
				gendung.CopyTransparency(x, y+1, x, y)
				gendung.CopyTransparency(x+1, y+1, x+1, y)
			}
			xx++
		}
		yy++
	}

	// Fix transparency, dirt and corners.
	FixTransparency()
	FixDirt()
	FixCorners()

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
	ResetMaps()

	// Initialize quest area.
	quests.InitQuestArea(*gendung.SetXx, *gendung.SetYy)
}
