// Package l1 implements dynamic random level generation of Cathedral maps.
package l1

import "C"

import (
	"reflect"
	"unsafe"

	"github.com/pkg/errors"
	"github.com/sanctuary/djavul/diablo"
	"github.com/sanctuary/djavul/engine"
	"github.com/sanctuary/djavul/gendung"
	"github.com/sanctuary/djavul/lighting"
	"github.com/sanctuary/djavul/multi"
	"github.com/sanctuary/djavul/quests"
	"github.com/sanctuary/formats/level/til"
)

// resetMaps resets the dungeon flag, player, NPC, dead, object, item, missile
// and arch maps.
//
// PSX ref: 0x8013CEAC
// PSX sig: void DRLG_Init_Globals__Fv()
//
// ref: 0x40ADD6
func resetMaps() {
	for x := range gendung.DFlagMap {
		for y := range gendung.DFlagMap[x] {
			gendung.DFlagMap[x][y] = 0
		}
	}
	for x := range gendung.PlayerNumMap {
		for y := range gendung.PlayerNumMap[x] {
			gendung.PlayerNumMap[x][y] = 0
		}
	}
	for x := range gendung.NPCNumMap {
		for y := range gendung.NPCNumMap[x] {
			gendung.NPCNumMap[x][y] = 0
		}
	}
	for x := range gendung.DeadMap {
		for y := range gendung.DeadMap[x] {
			gendung.DeadMap[x][y] = 0
		}
	}
	for x := range gendung.ObjectNumMap {
		for y := range gendung.ObjectNumMap[x] {
			gendung.ObjectNumMap[x][y] = 0
		}
	}
	for x := range gendung.ItemNumMap {
		for y := range gendung.ItemNumMap[x] {
			gendung.ItemNumMap[x][y] = 0
		}
	}
	for x := range gendung.MissileNumMap {
		for y := range gendung.MissileNumMap[x] {
			gendung.MissileNumMap[x][y] = 0
		}
	}
	for x := range gendung.ArchNumMap {
		for y := range gendung.ArchNumMap[x] {
			gendung.ArchNumMap[x][y] = 0
		}
	}
	var dist int8
	if !*lighting.Disabled {
		if *diablo.LightingFlag4 == 0 {
			dist = 15
		} else {
			dist = 3
		}
	}
	for x := range gendung.LightingVisibleDistanceMap {
		for y := range gendung.LightingVisibleDistanceMap[x] {
			gendung.LightingVisibleDistanceMap[x][y] = dist
		}
	}
}

// randomizeStoneFloor randomizes floor tiles.
//
// PSX ref: 0x8013CAC4
// PSX sig: void DRLG_L1Floor__Fv()
//
// ref: 0x40AF65
func randomizeStoneFloor() {
	for yy := 0; yy < 40; yy++ {
		for xx := 0; xx < 40; xx++ {
			if FlagMap[xx][yy] != 0 {
				continue
			}
			if TileID(gendung.TileIDMap[xx][yy]) == Floor {
				switch engine.RandCap(0, 3) {
				case 0:
					// Keep Floor tile.
				case 1:
					gendung.TileIDMap[xx][yy] = uint8(Floor2)
				case 2:
					gendung.TileIDMap[xx][yy] = uint8(Floor3)
				}
			}
		}
	}
}

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

// createDungeon creates a random Cathedral dungeon based on the given seed and
// level entry.
//
// PSX ref: 0x80140E64
// PSX sig: void CreateL5Dungeon__FUii(unsigned int rseed, int entry)
//
// ref: 0x40B229
func createDungeon(seed, entry int32) {
	engine.SetSeed(seed)
	gendung.InitTransparency()
	gendung.InitSetPiece()
	LoadSinglePlayerQuestDun()
	GenerateDungeon(entry)
	InitPieceIDMap()
	FreeSinglePlayerQuestDun()
	InitArches()
	gendung.MarkSetPiece()
}

// loadSinglePlayerQuestDun loads tile IDs from the dungeon file of the active
// single player quest level.
//
// PSX ref: 0x8013CDA0
// PSX sig: void DRLG_LoadL1SP__Fv()
//
// ref: 0x40B276
func loadSinglePlayerQuestDun() {
	*SinglePlayerQuestDunLoaded = false
	switch {
	case quests.IsActive(quests.TheButcher):
		*SinglePlayerQuestDun = (*uint8)(engine.MemLoadFile(unsafe.Pointer(C.CString(`Levels\L1Data\rnd6.DUN`)), nil))
		*SinglePlayerQuestDunLoaded = true
	case quests.IsActive(quests.TheCurseOfKingLeoric):
		if *multi.MaxPlayers == 1 {
			*SinglePlayerQuestDun = (*uint8)(engine.MemLoadFile(unsafe.Pointer(C.CString(`Levels\L1Data\SKngDO.DUN`)), nil))
			*SinglePlayerQuestDunLoaded = true
		}
	case quests.IsActive(quests.OgdensSign):
		*SinglePlayerQuestDun = (*uint8)(engine.MemLoadFile(unsafe.Pointer(C.CString(`Levels\L1Data\Banner2.DUN`)), nil))
		*SinglePlayerQuestDunLoaded = true
	}
}

// freeSinglePlayerQuestDun frees the dungeon file of the active single player
// quest level.
//
// PSX ref: 0x8013CE7C
// PSX sig: void DRLG_FreeL1SP__Fv()
//
// ref: 0x40B2F4
func freeSinglePlayerQuestDun() {
	ptr := *SinglePlayerQuestDun
	*SinglePlayerQuestDun = nil
	engine.MemFree(unsafe.Pointer(ptr))
}

// generateDungeon generates a Cathedral dungeon based on the given level entry.
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
		InitTileBitMap()
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

// initShadows initializes arch and bar shadows.
//
// PSX ref: 0x8013C190
// PSX sig: void DRLG_L1Shadows__Fv()
//
// ref: 0x40B699
func initShadows() {
	// Add arch shadows based on pre-defined lookup table.
	for yy := 1; yy < 40; yy++ {
		for xx := 1; xx < 40; xx++ {
			bottom := Base[gendung.TileIDMap[xx][yy]]
			left := Base[gendung.TileIDMap[xx-1][yy]]
			right := Base[gendung.TileIDMap[xx][yy-1]]
			top := Base[gendung.TileIDMap[xx-1][yy-1]]
			for _, shadow := range Shadows {
				// The bottom base of a shadow is always set.
				if shadow.BottomBase != bottom {
					continue
				}
				if shadow.TopBase != 0 && shadow.TopBase != top {
					continue
				}
				if shadow.RightBase != 0 && shadow.RightBase != right {
					continue
				}
				if shadow.LeftBase != 0 && shadow.LeftBase != left {
					continue
				}
				if shadow.Top != 0 && FlagMap[xx-1][yy-1] == 0 {
					gendung.TileIDMap[xx-1][yy-1] = uint8(shadow.Top)
				}
				if shadow.Right != 0 && FlagMap[xx][yy-1] == 0 {
					gendung.TileIDMap[xx][yy-1] = uint8(shadow.Right)
				}
				if shadow.Left != 0 && FlagMap[xx-1][yy] == 0 {
					gendung.TileIDMap[xx-1][yy] = uint8(shadow.Left)
				}
			}
		}
	}
	// Add shadows for bar tiles.
	for yy := 1; yy < 40; yy++ {
		for xx := 1; xx < 40; xx++ {
			if FlagMap[xx-1][yy] != 0 {
				continue
			}
			switch TileID(gendung.TileIDMap[xx][yy]) {
			case BarSwBarSe, BarEndSw, BarSw, BarSwWallSe, BarSwArchSe, BarSwDoorSe:
				switch TileID(gendung.TileIDMap[xx-1][yy]) {
				case FloorShadowArchSwRight:
					gendung.TileIDMap[xx-1][yy] = uint8(FloorShadowBarSwRight)
				case ArchSeShadowArchSwRight:
					gendung.TileIDMap[xx-1][yy] = uint8(ArchSeShadowBarSwRight)
				case WallSeShadowArchSwRight:
					gendung.TileIDMap[xx-1][yy] = uint8(WallSeShadowBarSwRight)
				}
			}
		}
	}
}

// placeMiniset places the given miniset of tile IDs.
//
// PSX ref: 0x8013C5A0
// PSX sig: int DRLG_PlaceMiniSet__FPCUciiiiiii(unsigned char *miniset, int tmin, int tmax, int cx, int cy, int setview, int noquad, int ldir)
//
// ref: 0x40B881
func placeMiniset(miniset unsafe.Pointer, tmin, tmax, cx, cy int, setView bool, noquad, ldir int) int {
	// Size of the largest miniset (i.e. MinisetStairDownPoison).
	const maxSize = 1 + 1 + 6*6 + 6*6
	sh := reflect.SliceHeader{Data: uintptr(miniset), Len: maxSize, Cap: maxSize}
	mini := *(*[]uint8)(unsafe.Pointer(&sh))
	width := int(mini[0])
	height := int(mini[1])
	// Actual size of miniset.
	n := 1 + 1 + width*height + width*height
	mini = mini[:n:n]
	// Number of times to place the miniset.
	t := 1
	if tmin != tmax {
		t = tmin + int(engine.RandCap(0, int32(tmax-tmin)))
	}
	xMax := 40 - width
	yMax := 40 - height
	var x, y int
	for i := 0; i < t; i++ {
		x = int(engine.RandCap(0, int32(xMax)))
		j := 0
		y = int(engine.RandCap(0, int32(yMax)))
		for {
			found := true
			if cx != -1 && x >= cx-width && x <= cx+12 {
				x++
				found = false
			}
			if cy != -1 && y >= cy-height && y <= cy+12 {
				y++
				found = false
			}
			switch noquad {
			case 0:
				if x < cx && y < cy {
					found = false
				}
			case 1:
				if x > cx && y < cy {
					found = false
				}
			case 2:
				if x < cx && y > cy {
					found = false
				}
			case 3:
				if x > cx && y > cy {
					found = false
				}
			}
			// Locate miniset before pattern in map.
			k := 2
			for yDelta := 0; yDelta < height; yDelta++ {
				if !found {
					break
				}
				for xDelta := 0; xDelta < width; xDelta++ {
					if !found {
						break
					}
					v := mini[k]
					if v != 0 && gendung.TileIDMap[x+xDelta][y+yDelta] != v {
						found = false
					}
					if FlagMap[x+xDelta][y+yDelta] != 0 {
						found = false
					}
					k++
				}
			}
			if found {
				break
			}
			x++
			if x == xMax {
				x = 0
				y++
				if y == yMax {
					y = 0
				}
			}
			j++
			if j > 4000 {
				return -1
			}
		}
		// Add miniset to map.
		k := 2 + width*height
		for yDelta := 0; yDelta < height; yDelta++ {
			for xDelta := 0; xDelta < width; xDelta++ {
				v := mini[k]
				if v != 0 {
					gendung.TileIDMap[x+xDelta][y+yDelta] = v
				}
				k++
			}
		}
	}
	// Add transparency and entrance for Poisoned Water Supply.
	if miniset == unsafe.Pointer(MinisetStairDownPoison) {
		transIndexBak := *gendung.TransparencyIndex
		*gendung.TransparencyIndex = 0
		gendung.MakeRectTransparent(x, y+2, x+5, y+4)
		*gendung.TransparencyIndex = transIndexBak
		quests.Quests[quests.PoisonedWaterSupply].EnteranceX = int32(2*x + 21)
		quests.Quests[quests.PoisonedWaterSupply].EnteranceY = int32(2*y + 22)
	}
	if setView {
		*gendung.ViewX = int32(2*x + 19)
		*gendung.ViewY = int32(2*y + 20)
	}
	if ldir != 0 {
		*gendung.LvlViewX = int32(2*x + 19)
		*gendung.LvlViewY = int32(2*y + 20)
	}
	switch {
	case x < cx && y < cy:
		return 0
	case x > cx && y < cy:
		return 1
	case x < cx && y > cy:
		return 2
	default:
		// (x <= cx || y >= cy) && (x >= cx || y <= cy)
		return 3
	}
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

// clearFlags clears the dungeon generation flags 0x40.
//
// PSX ref: 0x8013D37C
// PSX sig: void L5ClearFlags__Fv()
//
// ref: 0x40BB18
func clearFlags() {
	for xx := 0; xx < 40; xx++ {
		for yy := 0; yy < 40; yy++ {
			FlagMap[xx][yy] &^= Flag40
		}
	}
}

// generateFirstRoom generates the first room of the dungeon.
//
// PSX ref: 0x8013D7FC
// PSX sig: void L5firstRoom__Fv()
//
// ref: 0x40BB33
func generateFirstRoom() {
	switch engine.RandCap(0, 2) {
	case 0:
		// Vertical rooms.
		yyStart := 1
		yyEnd := 39
		*VertRoom1 = engine.RandCap(0, 2) == 1
		*VertRoom2 = engine.RandCap(0, 2) == 1
		*VertRoom3 = engine.RandCap(0, 2) == 1
		if !(*VertRoom1 && *VertRoom3) {
			*VertRoom2 = true
		}
		if *VertRoom1 {
			AddRoom(15, 1, 10, 10)
		} else {
			yyStart = 18
		}
		if *VertRoom2 {
			AddRoom(15, 15, 10, 10)
		}
		if *VertRoom3 {
			AddRoom(15, 29, 10, 10)
		} else {
			yyEnd = 22
		}
		for yy := yyStart; yy < yyEnd; yy++ {
			gendung.TileIDMap[17][yy] = uint8(WallSw)
			gendung.TileIDMap[18][yy] = uint8(WallSw)
			gendung.TileIDMap[19][yy] = uint8(WallSw)
			gendung.TileIDMap[20][yy] = uint8(WallSw)
			gendung.TileIDMap[21][yy] = uint8(WallSw)
			gendung.TileIDMap[22][yy] = uint8(WallSw)
		}
		if *VertRoom1 {
			GenerateRoom(15, 1, 10, 10, false)
		}
		if *VertRoom2 {
			GenerateRoom(15, 15, 10, 10, false)
		}
		if *VertRoom3 {
			GenerateRoom(15, 29, 10, 10, false)
		}
		*HorizRoom1 = false
		*HorizRoom2 = false
		*HorizRoom3 = false
	case 1:
		// Horizontal rooms.
		xxStart := 1
		xxEnd := 39
		*HorizRoom1 = engine.RandCap(0, 2) == 1
		*HorizRoom2 = engine.RandCap(0, 2) == 1
		*HorizRoom3 = engine.RandCap(0, 2) == 1
		if !(*HorizRoom1 && *HorizRoom3) {
			*HorizRoom2 = true
		}
		if *HorizRoom1 {
			AddRoom(1, 15, 10, 10)
		} else {
			xxStart = 18
		}
		if *HorizRoom2 {
			AddRoom(15, 15, 10, 10)
		}
		if *HorizRoom3 {
			AddRoom(29, 15, 10, 10)
		} else {
			xxEnd = 22
		}
		for xx := xxStart; xx < xxEnd; xx++ {
			gendung.TileIDMap[xx][17] = uint8(WallSw)
			gendung.TileIDMap[xx][18] = uint8(WallSw)
			gendung.TileIDMap[xx][19] = uint8(WallSw)
			gendung.TileIDMap[xx][20] = uint8(WallSw)
			gendung.TileIDMap[xx][21] = uint8(WallSw)
			gendung.TileIDMap[xx][22] = uint8(WallSw)
		}
		if *HorizRoom1 {
			GenerateRoom(1, 15, 10, 10, true)
		}
		if *HorizRoom2 {
			GenerateRoom(15, 15, 10, 10, true)
		}
		if *HorizRoom3 {
			GenerateRoom(29, 15, 10, 10, true)
		}
		*VertRoom1 = false
		*VertRoom2 = false
		*VertRoom3 = false
	}
}

// addRoom adds walls for a room at the given area.
//
// PSX ref: 0x8013D3CC
// PSX sig: void L5drawRoom__Fiiii(int x, int y, int w, int h)
//
// ref: 0x40BD66
func addRoom(xxStart, yyStart, xxCount, yyCount int) {
	for yyDelta := 0; yyDelta < yyCount; yyDelta++ {
		for xxDelta := 0; xxDelta < xxCount; xxDelta++ {
			xx := xxStart + xxDelta
			yy := yyStart + yyDelta
			// NOTE: The original implementation contains an out-of-bounds access
			// invoked from GenerateRoom.
			if xx >= 0 && xx < 40 && yy >= 0 && yy < 40 {
				gendung.TileIDMap[xx][yy] = uint8(WallSw)
			}
		}
	}
}

// generateRoom generates a room of the given dimensions at the specified
// coordinates.
//
// PSX ref: 0x8013D4CC
// PSX sig: void L5roomGen__Fiiiii(int x, int y, int w, int h, int dir)
//
// ref: 0x40BD9D
func generateRoom(xxStart, yyStart, xxCount, yyCount int, dirVert bool) {
	for {
		v := engine.RandCap(0, 4)
		if dirVert && v != 0 || !dirVert && v == 0 {
			// Generate rooms in vertical direction.
			var (
				xxCountNew int
				yyCountNew int
				xxStartNew int
				yyStartNew int
				empty1     bool
			)
			for i := 0; i < 20; i++ {
				xxCountNew = int(uint32((engine.RandCap(0, 5) + 2)) & 0xFFFFFFFE)
				yyCountNew = int(uint32((engine.RandCap(0, 5) + 2)) & 0xFFFFFFFE)
				xxStartNew = xxStart + xxCount/2 - xxCountNew/2
				yyStartNew = yyStart - yyCountNew
				empty1 = IsAreaEmpty(xxStartNew-1, yyStartNew-1, xxCountNew+2, yyCountNew+1)
				if empty1 {
					break
				}
			}
			if empty1 {
				AddRoom(xxStartNew, yyStartNew, xxCountNew, yyCountNew)
			}
			yyStartNew2 := yyStart + yyCount
			empty2 := IsAreaEmpty(xxStartNew-1, yyStartNew2, xxCountNew+2, yyCountNew+1)
			if empty2 {
				AddRoom(xxStartNew, yyStartNew2, xxCountNew, yyCountNew)
			}
			if empty1 {
				GenerateRoom(xxStartNew, yyStartNew, xxCountNew, yyCountNew, false)
			}
			if !empty2 {
				return
			}
			dirVert = false
			xxCount = xxCountNew
			yyCount = yyCountNew
			xxStart = xxStartNew
			yyStart = yyStartNew2
		} else {
			// Generate rooms in horizontal direction.
			var (
				xxCountNew int
				yyCountNew int
				xxStartNew int
				yyStartNew int
				empty1     bool
			)
			for j := 0; j < 20; j++ {
				xxCountNew = int(uint32((engine.RandCap(0, 5) + 2)) & 0xFFFFFFFE)
				yyCountNew = int(uint32((engine.RandCap(0, 5) + 2)) & 0xFFFFFFFE)
				xxStartNew = xxStart - xxCountNew
				yyStartNew = yyStart + yyCount/2 - yyCountNew/2
				// NOTE: BUG in original. yy and xx swapped in 3rd and 4th argument.
				//
				// Note to self. It would be interest to check which seeds generate
				// faulty maps based on this bug.
				empty1 = IsAreaEmpty(xxStartNew-1, yyStartNew-1, yyCountNew+2, xxCountNew+1)
				if empty1 {
					break
				}
			}
			if empty1 {
				AddRoom(xxStartNew, yyStartNew, xxCountNew, yyCountNew)
			}
			xxStartNew2 := xxStart + xxCount
			empty2 := IsAreaEmpty(xxStartNew2, yyStartNew-1, xxCountNew+1, yyCountNew+2)
			if empty2 {
				AddRoom(xxStartNew2, yyStartNew, xxCountNew, yyCountNew)
			}
			if empty1 {
				GenerateRoom(xxStartNew, yyStartNew, xxCountNew, yyCountNew, true)
			}
			if !empty2 {
				return
			}
			dirVert = true
			xxCount = xxCountNew
			yyCount = yyCountNew
			xxStart = xxStartNew2
			yyStart = yyStartNew
		}
	}
}

// isAreaEmpty reports whether the given area is empty.
//
// PSX ref: 0x8013D438
// PSX sig: unsigned char L5checkRoom__Fiiii(int x, int y, int width, int height)
//
// ref: 0x40BFA4
func isAreaEmpty(xxStart, yyStart, xxCount, yyCount int) bool {
	if xxCount <= 0 || yyCount <= 0 {
		return true
	}
	for yy := yyStart; yy < yyStart+yyCount; yy++ {
		if yy < 0 || yy >= 40 {
			return false
		}
		for xx := xxStart; xx < xxStart+xxCount; xx++ {
			if xx < 0 || xx >= 40 {
				return false
			}
			if gendung.TileIDMap[xx][yy] != 0 {
				return false
			}
		}
	}
	return true
}

// getArea returns the number of walls on the map.
//
// PSX ref: 0x8013DB9C
// PSX sig: long L5GetArea__Fv()
//
// ref: 0x40C008
func getArea() int {
	n := 0
	for xx := 0; xx < 40; xx++ {
		for yy := 0; yy < 40; yy++ {
			if TileID(gendung.TileIDMap[xx][yy]) == WallSw {
				n++
			}
		}
	}
	return n
}

// initTileBitMap initializes a tile ID map of twice the size, repeating each
// tile in blocks of 4.
//
// PSX ref: 0x8013DBFC
// PSX sig: void L5makeDungeon__Fv()
//
// ref: 0x40C02A
func initTileBitMap() {
	x := 0
	for xx := 0; xx < 40; xx++ {
		y := 0
		for yy := 0; yy < 40; yy++ {
			tileID := TileID(gendung.TileIDMap[xx][yy])
			TileBitMap[x][y] = tileID
			TileBitMap[x][y+1] = tileID
			TileBitMap[x+1][y] = tileID
			TileBitMap[x+1][y+1] = tileID
			y += 2
		}
		x += 2
	}
}

// generatePattern replaces tile ID patterns based on a lookup table.
//
// PSX ref: 0x8013DC88
// PSX sig: void L5makeDmt__Fv()
//
// ref: 0x40C06E
func generatePattern() {
	// Fill the entire tile ID map with dirt.
	for xx := 0; xx < 40; xx++ {
		for yy := 0; yy < 40; yy++ {
			gendung.TileIDMap[xx][yy] = uint8(Dirt)
		}
	}

	// Fill the tile ID map based on pattern matching of 2x2 areas. Basically
	// each cell in the tile bit map either contains a SW wall or it doesn't.
	// This property allows us to create a bitfield of the various patterns
	// representable by nearby tiles. The total number of patterns is 16, so we
	// use a simple lookup table for the result.
	x := 1
	for xx := 0; xx < 40-1; xx++ {
		y := 1
		for yy := 0; yy < 40-1; yy++ {
			pattern := uint8(TileBitMap[x+1][y+1])
			pattern = 2*pattern + uint8(TileBitMap[x][y+1])
			pattern = 2*pattern + uint8(TileBitMap[x+1][y])
			pattern = 2*pattern + uint8(TileBitMap[x][y])
			gendung.TileIDMap[xx][yy] = uint8(PatternLookup[pattern])
			y += 2
		}
		x += 2
	}
}

// addWall adds wall, arch or bar tile IDs.
//
// PSX ref: 0x8013E458
// PSX sig: void L5AddWall__Fv()
//
// ref: 0x40C0E0
func addWall() {
	for yy := 0; yy < 40; yy++ {
		for xx := 0; xx < 40; xx++ {
			if FlagMap[xx][yy] == 0 {
				switch TileID(gendung.TileIDMap[xx][yy]) {
				case WallSw:
					// NOTE: rand_cap(100) < 100 always holds true.
					_ = engine.RandCap(0, 100)
					height := GetVertWallSpace(xx, yy)
					if height != -1 {
						AddVertWall(xx, yy, WallSw, height)
					}
				case WallSe:
					_ = engine.RandCap(0, 100)
					width := GetHorizWallSpace(xx, yy)
					if width != -1 {
						AddHorizWall(xx, yy, WallSe, width)
					}
				case ArchNeArchNw:
					_ = engine.RandCap(0, 100)
					width := GetHorizWallSpace(xx, yy)
					if width != -1 {
						AddHorizWall(xx, yy, WallSe, width)
					}
					_ = engine.RandCap(0, 100)
					height := GetVertWallSpace(xx, yy)
					if height != -1 {
						AddVertWall(xx, yy, WallSw, height)
					}
				case WallEndSw:
					_ = engine.RandCap(0, 100)
					width := GetHorizWallSpace(xx, yy)
					if width != -1 {
						AddHorizWall(xx, yy, WallSwWallSe, width)
					}
				case WallEndSe:
					_ = engine.RandCap(0, 100)
					height := GetVertWallSpace(xx, yy)
					if height != -1 {
						AddVertWall(xx, yy, WallSwWallSe, height)
					}
				}
			}
		}
	}
}

// getHorizWallSpace returns the number of horizontal wall tiles that fit at the
// given coordinate.
//
// PSX ref: 0x8013DD70
// PSX sig: int L5HWallOk__Fii(int i, int j)
//
// ref: 0x40C23C
func getHorizWallSpace(xx, yy int) int {
	width := 1
	for TileID(gendung.TileIDMap[xx+width][yy]) == Floor {
		if TileID(gendung.TileIDMap[xx+width][yy-1]) != Floor {
			break
		}
		if TileID(gendung.TileIDMap[xx+width][yy+1]) != Floor {
			break
		}
		if FlagMap[xx+width][yy] != 0 {
			break
		}
		width++
	}
	ok := false
	switch TileID(gendung.TileIDMap[width+xx][yy]) {
	case ArchNeArchNw, WallSwWallSe, ArchSwArchSe, WallEndSw, WallEndSe, ArchEndNe, ArchEndNw, DirtWallSw, DirtWallSe, DirtWallNeWallNw, DirtWallSwWallSe, DirtWallEndSw, DirtWallEndSe:
		ok = true
	}
	if width == 1 || !ok {
		return -1
	}
	return width
}

// getVertWallSpace returns the number of vertical wall tiles that fit at the
// given coordinate.
//
// PSX ref: 0x8013DEAC
// PSX sig: int L5VWallOk__Fii(int i, int j)
//
// ref: 0x40C2DC
func getVertWallSpace(xx, yy int) int {
	height := 1
	for TileID(gendung.TileIDMap[xx][yy+height]) == Floor {
		if TileID(gendung.TileIDMap[xx-1][yy+height]) != Floor {
			break
		}
		if TileID(gendung.TileIDMap[xx+1][yy+height]) != Floor {
			break
		}
		if FlagMap[xx][yy+height] != 0 {
			break
		}
		height++
	}
	ok := false
	switch TileID(gendung.TileIDMap[xx][yy+height]) {
	case ArchNeArchNw, WallSwWallSe, ArchSwArchSe, WallEndSw, WallEndSe, ArchEndNe, ArchEndNw, DirtWallSw, DirtWallSe, DirtWallNeWallNw, DirtWallSwWallSe, DirtWallEndSw, DirtWallEndSe:
		ok = true
	}
	if height == 1 || !ok {
		return -1
	}
	return height
}

// addHorizWall adds a horizontal wall based on the given tile ID.
//
// PSX ref: 0x8013DFF4
// PSX sig: void L5HorizWall__Fiici(int i, int j, char p, int dx)
//
// ref: 0x40C35B
func addHorizWall(xx, yy int, tileIDFirst TileID, xxCount int) {
	// Select first and other tiles.
	var tileIDOther TileID
	switch engine.RandCap(0, 4) {
	case 0, 1:
		// Walls.
		tileIDOther = WallSe
	case 2:
		// Arches.
		tileIDOther = ArchSe
		switch tileIDFirst {
		case WallSe:
			tileIDFirst = ArchSe
		case WallSwWallSe:
			tileIDFirst = WallSwArchSe
		}
	case 3:
		// Bars.
		tileIDOther = BarSe
		switch tileIDFirst {
		case WallSe:
			tileIDFirst = BarSe
		case WallSwWallSe:
			tileIDFirst = WallSwBarSe
		}
	}

	// Select random tile.
	var tileIDRand TileID
	switch engine.RandCap(0, 6) {
	case 5:
		tileIDRand = ArchSe
	default:
		// NOTE: BUG? DoorSe is always overwritten in original. Should probably be
		// WallSe.
		tileIDRand = DoorSe
	}
	if tileIDOther == ArchSe {
		tileIDRand = ArchSe
	}

	// Place first tile.
	gendung.TileIDMap[xx][yy] = uint8(tileIDFirst)

	// Place other tiles.
	for xxDelta := 1; xxDelta < xxCount; xxDelta++ {
		gendung.TileIDMap[xx+xxDelta][yy] = uint8(tileIDOther)
	}

	// Place random tile.
	xxDelta := int(engine.RandCap(0, int32(xxCount-1))) + 1
	if tileIDRand == ArchSe {
		gendung.TileIDMap[xx+xxDelta][yy] = uint8(ArchSe)
	} else {
		FlagMap[xx+xxDelta][yy] |= FlagFavourSe
		gendung.TileIDMap[xx+xxDelta][yy] = uint8(WallSe)
	}
}

// addVertWall adds a vertical wall based on the given tile ID.
//
// PSX ref: 0x8013E22C
// PSX sig: void L5VertWall__Fiici(int i, int j, char p, int dy)
//
// ref: 0x40C449
func addVertWall(xx int, yy int, tileIDFirst TileID, yyCount int) {
	// Select first and other tiles.
	var tileIDOther TileID
	switch engine.RandCap(0, 4) {
	case 0, 1:
		// Walls.
		tileIDOther = WallSw
	case 2:
		// Arches.
		tileIDOther = ArchSw
		switch tileIDFirst {
		case WallSw:
			tileIDFirst = ArchSw
		case WallSwWallSe:
			tileIDFirst = ArchSwWallSe
		}
	case 3:
		// Bars.
		tileIDOther = BarSw
		switch tileIDFirst {
		case WallSw:
			tileIDFirst = BarSw
		case WallSwWallSe:
			tileIDFirst = BarSwWallSe
		}
	}

	// Select random tile.
	var tileIDRand TileID
	switch engine.RandCap(0, 6) {
	case 5:
		tileIDRand = ArchSw
	default:
		// NOTE: BUG? DoorSw is always overwritten in original. Should probably be
		// WallSw.
		tileIDRand = DoorSw
	}
	if tileIDOther == ArchSw {
		tileIDRand = ArchSw
	}

	// Place first tile.
	gendung.TileIDMap[xx][yy] = uint8(tileIDFirst)

	// Place other tiles.
	for yyDelta := 1; yyDelta < yyCount; yyDelta++ {
		gendung.TileIDMap[xx][yy+yyDelta] = uint8(tileIDOther)
	}

	// Place random tile.
	yyDelta := int(engine.RandCap(0, int32(yyCount-1))) + 1
	if tileIDRand == ArchSw {
		gendung.TileIDMap[xx][yy+yyDelta] = uint8(ArchSw)
	} else {
		FlagMap[xx][yy+yyDelta] |= FlagFavourSw
		gendung.TileIDMap[xx][yy+yyDelta] = uint8(WallSw)
	}
}

// fixTiles fixes tile IDs of wall edges.
//
// PSX ref: 0x8013EA28
// PSX sig: void L5tileFix__Fv()
//
// ref: 0x40C551
func fixTiles() {
	for yy := 0; yy < 40; yy++ {
		for xx := 0; xx < 40; xx++ {
			switch TileID(gendung.TileIDMap[xx][yy]) {
			case WallSw:
				if TileID(gendung.TileIDMap[xx][yy+1]) == Dirt {
					gendung.TileIDMap[xx][yy+1] = uint8(DirtWallEndSe)
				}
			case WallSe:
				if TileID(gendung.TileIDMap[xx+1][yy]) == Dirt {
					gendung.TileIDMap[xx+1][yy] = uint8(DirtWallEndSw)
				}
			case WallEndSw:
				if TileID(gendung.TileIDMap[xx+1][yy]) == Dirt {
					gendung.TileIDMap[xx+1][yy] = uint8(DirtWallEndSe)
				}
			case Floor:
				switch TileID(gendung.TileIDMap[xx+1][yy]) {
				case WallSe:
					gendung.TileIDMap[xx+1][yy] = uint8(WallEndSe)
				case Dirt:
					gendung.TileIDMap[xx+1][yy] = uint8(DirtWallSw)
				}
				switch TileID(gendung.TileIDMap[xx][yy+1]) {
				case WallSw:
					gendung.TileIDMap[xx][yy+1] = uint8(WallEndSw)
				case Dirt:
					gendung.TileIDMap[xx][yy+1] = uint8(DirtWallSe)
				}
			}
		}
	}

	for yy := 0; yy < 40; yy++ {
		for xx := 0; xx < 40; xx++ {
			switch TileID(gendung.TileIDMap[xx][yy]) {
			case WallSw:
				switch TileID(gendung.TileIDMap[xx][yy+1]) {
				case WallSe:
					gendung.TileIDMap[xx][yy+1] = uint8(WallEndSe)
				case Floor:
					gendung.TileIDMap[xx][yy+1] = uint8(ArchEndNe)
				}
			case WallSe:
				switch TileID(gendung.TileIDMap[xx+1][yy]) {
				case WallSw:
					gendung.TileIDMap[xx+1][yy] = uint8(WallEndSw)
				case Floor:
					gendung.TileIDMap[xx+1][yy] = uint8(ArchEndNw)
				case DirtWallSe:
					gendung.TileIDMap[xx+1][yy] = uint8(DirtWallSwWallSe)
				case DirtWallEndSe:
					gendung.TileIDMap[xx+1][yy] = uint8(DirtWallSwWallSe)
				}
			case ArchNeArchNw:
				if TileID(gendung.TileIDMap[xx+1][yy]) == Dirt {
					gendung.TileIDMap[xx+1][yy] = uint8(DirtWallEndSe)
				}
			case WallSwWallSe:
				if TileID(gendung.TileIDMap[xx+1][yy]) == ArchEndNe {
					gendung.TileIDMap[xx+1][yy] = uint8(ArchEndNw)
				}
			case WallEndSw:
				switch TileID(gendung.TileIDMap[xx-1][yy]) {
				case Dirt:
					gendung.TileIDMap[xx-1][yy] = uint8(DirtWallEndSe)
				case DirtWallEndSw:
					gendung.TileIDMap[xx-1][yy] = uint8(DirtWallSwWallSe)
				}
				switch TileID(gendung.TileIDMap[xx][yy+1]) {
				case WallSe:
					gendung.TileIDMap[xx][yy+1] = uint8(WallEndSe)
				case Floor:
					gendung.TileIDMap[xx][yy+1] = uint8(ArchEndNe)
				case DirtWallSw:
					gendung.TileIDMap[xx][yy+1] = uint8(DirtWallSwWallSe)
				case Dirt:
					gendung.TileIDMap[xx][yy+1] = uint8(DirtWallEndSe)
				}
				if TileID(gendung.TileIDMap[xx][yy-1]) == Dirt {
					// NOTE: The following value is always overwritten.
					//gendung.TileIDMap[xx][yy-1] = uint8(WallEndSe)
					gendung.TileIDMap[xx][yy-1] = uint8(DirtWallEndSe)
				}
			case WallEndSe:
				switch TileID(gendung.TileIDMap[xx+1][yy]) {
				case WallSw:
					gendung.TileIDMap[xx+1][yy] = uint8(WallEndSw)
				case Floor:
					gendung.TileIDMap[xx+1][yy] = uint8(ArchEndNw)
				case DirtWallSe:
					gendung.TileIDMap[xx+1][yy] = uint8(DirtWallSwWallSe)
				case Dirt:
					gendung.TileIDMap[xx+1][yy] = uint8(DirtWallEndSw)
				case DirtWallEndSe:
					gendung.TileIDMap[xx+1][yy] = uint8(DirtWallSwWallSe)
				}
				if TileID(gendung.TileIDMap[xx][yy-1]) == DirtWallEndSe {
					gendung.TileIDMap[xx][yy-1] = uint8(DirtWallSwWallSe)
				}
			case Floor:
				switch TileID(gendung.TileIDMap[xx+1][yy]) {
				case DirtWallSe:
					gendung.TileIDMap[xx+1][yy] = uint8(DirtWallSwWallSe)
				case Dirt:
					gendung.TileIDMap[xx+1][yy] = uint8(DirtWallNeWallNw)
				case DirtWallEndSe:
					gendung.TileIDMap[xx+1][yy] = uint8(DirtWallSwWallSe)
				}
				if TileID(gendung.TileIDMap[xx][yy+1]) == ArchEndNe {
					gendung.TileIDMap[xx][yy+1] = uint8(ArchEndNw)
				}
			case DirtWallSw:
				if TileID(gendung.TileIDMap[xx][yy+1]) == WallSe {
					gendung.TileIDMap[xx][yy+1] = uint8(WallEndSe)
				}
				if TileID(gendung.TileIDMap[xx][yy-1]) == DirtWallEndSe {
					gendung.TileIDMap[xx][yy-1] = uint8(DirtWallSwWallSe)
				}
			case DirtWallSe:
				switch TileID(gendung.TileIDMap[xx+1][yy]) {
				case WallSw:
					gendung.TileIDMap[xx+1][yy] = uint8(WallEndSw)
				case Dirt:
					gendung.TileIDMap[xx+1][yy] = uint8(DirtWallNeWallNw)
				}
				if TileID(gendung.TileIDMap[xx-1][yy]) == DirtWallEndSw {
					gendung.TileIDMap[xx-1][yy] = uint8(DirtWallSwWallSe)
				}
			case DirtWallSwWallSe:
				if TileID(gendung.TileIDMap[xx+1][yy]) == WallSw {
					gendung.TileIDMap[xx+1][yy] = uint8(WallEndSw)
				}
				if TileID(gendung.TileIDMap[xx][yy+1]) == WallSe {
					gendung.TileIDMap[xx][yy+1] = uint8(WallEndSe)
				}
			case DirtWallEndSw:
				if TileID(gendung.TileIDMap[xx-1][yy]) == Dirt {
					gendung.TileIDMap[xx-1][yy] = uint8(DirtWallSe)
				}
			}
		}
	}

	for yy := 0; yy < 40; yy++ {
		for xx := 0; xx < 40; xx++ {
			switch TileID(gendung.TileIDMap[xx][yy]) {
			case WallSe:
				if TileID(gendung.TileIDMap[xx+1][yy]) == DirtWallSe {
					gendung.TileIDMap[xx+1][yy] = uint8(DirtWallSwWallSe)
				}
			case WallSwWallSe:
				if TileID(gendung.TileIDMap[xx][yy+1]) == WallSe {
					gendung.TileIDMap[xx][yy+1] = uint8(WallEndSe)
				}
			case DirtWallSw:
				if TileID(gendung.TileIDMap[xx][yy+1]) == Dirt {
					gendung.TileIDMap[xx][yy+1] = uint8(DirtWallNeWallNw)
				}
			}
		}
	}
}

// decorate decorates the dungeon with tapestry tile IDs.
//
// PSX ref: 0x8013F2EC
// PSX sig: void DRLG_L5Subs__Fv()
//
// ref: 0x40C8C0
func decorate() {
	for yy := 0; yy < 40; yy++ {
		for xx := 0; xx < 40; xx++ {
			if engine.RandCap(0, 4) != 0 {
				continue
			}
			plain := Plain[gendung.TileIDMap[xx][yy]]
			if plain == 0 {
				continue
			}
			if FlagMap[xx][yy] != 0 {
				continue
			}
			tileID := TileID(0xFF)
			for i := engine.RandCap(0, 16); i >= 0; {
				tileID++
				// NOTE: 206 in original, which is a bug. Should be 207.
				if tileID == 207 {
					tileID = 0
				}
				if plain == Plain[tileID] {
					i--
				}
			}
			switch tileID {
			case TapestryWallSw1:
				if Plain[gendung.TileIDMap[xx][yy-1]] == WallSw2 && FlagMap[xx][yy-1] == 0 {
					gendung.TileIDMap[xx][yy] = uint8(TapestryWallSw1)
					gendung.TileIDMap[xx][yy-1] = uint8(TapestryWallSw2)
				} else {
					gendung.TileIDMap[xx][yy] = uint8(WallSw2)
				}
			case TapestryWallSe1:
				if Plain[gendung.TileIDMap[xx+1][yy]] == WallSe2 && FlagMap[xx+1][yy] == 0 {
					gendung.TileIDMap[xx][yy] = uint8(TapestryWallSe1)
					gendung.TileIDMap[xx+1][yy] = uint8(TapestryWallSe2)
				} else {
					gendung.TileIDMap[xx][yy] = uint8(WallSe2)
				}
			default:
				gendung.TileIDMap[xx][yy] = uint8(tileID)
			}
		}
	}
}

// generateChambers generates chambers.
//
// PSX ref: 0x8013F5F8
// PSX sig: void L5FillChambers__Fv()
//
// ref: 0x40C99D
func generateChambers() {
	// Generate horizontal chambers.
	v := 1
	if *HorizRoom1 {
		GenerateChamber(0, 14, false, false, false, true)
	}
	if *HorizRoom1 && *HorizRoom2 && !*HorizRoom3 {
		GenerateChamber(14, 14, false, false, true, false)
	}
	if *HorizRoom2 && *HorizRoom3 {
		GenerateChamber(14, 14, false, false, false, true)
	}
	if *HorizRoom1 && *HorizRoom2 && *HorizRoom3 {
		GenerateChamber(14, 14, false, false, true, true)
	}
	if !*HorizRoom1 && *HorizRoom2 && !*HorizRoom3 {
		GenerateChamber(14, 14, false, false, false, false)
	}
	if *HorizRoom3 {
		GenerateChamber(28, 14, false, false, true, false)
	}
	if *HorizRoom1 && *HorizRoom2 {
		GenerateHall(12, 18, 14, 18)
	}
	if *HorizRoom2 && *HorizRoom3 {
		GenerateHall(26, 18, 28, 18)
	}
	if *HorizRoom1 && !*HorizRoom2 && *HorizRoom3 {
		GenerateHall(12, 18, 28, 18)
	}

	// Generate vertical chambers.
	if *VertRoom1 {
		GenerateChamber(14, 0, false, true, false, false)
	}
	if *VertRoom1 && *VertRoom2 && !*VertRoom3 {
		GenerateChamber(14, 14, true, false, false, false)
	}
	if *VertRoom2 && *VertRoom3 {
		GenerateChamber(14, 14, false, true, false, false)
	}
	if *VertRoom1 && *VertRoom2 && *VertRoom3 {
		GenerateChamber(14, 14, true, true, false, false)
	}
	if !*VertRoom1 && *VertRoom2 && !*VertRoom3 {
		GenerateChamber(14, 14, false, false, false, false)
	}
	if *VertRoom3 {
		GenerateChamber(14, 28, true, false, false, false)
	}
	if *VertRoom1 && *VertRoom2 {
		GenerateHall(18, 12, 18, 14)
	}
	if *VertRoom2 && *VertRoom3 {
		GenerateHall(18, 26, 18, 28)
	}
	if *VertRoom1 && !*VertRoom2 && *VertRoom3 {
		GenerateHall(18, 12, 18, 28)
	}

	// Early exit for non-quest maps.
	if !*SinglePlayerQuestDunLoaded {
		return
	}

	// Generate horizontal quest areas.
	if !*VertRoom1 && !*VertRoom2 && !*VertRoom3 {
		if !*HorizRoom1 && *HorizRoom2 && *HorizRoom3 {
			if engine.RandCap(0, 2) != 0 {
				v = 2
			}
		}
		if *HorizRoom1 && *HorizRoom2 && !*HorizRoom3 {
			if engine.RandCap(0, 2) != 0 {
				v = 0
			}
		}
		if *HorizRoom1 && !*HorizRoom2 && *HorizRoom3 {
			if engine.RandCap(0, 2) != 0 {
				v = 0
			} else {
				v = 2
			}
		}
		if *HorizRoom1 && *HorizRoom2 && *HorizRoom3 {
			v = int(engine.RandCap(0, 3))
		}
		switch v {
		case 0:
			InitQuestDun(2, 16)
		case 1:
			InitQuestDun(16, 16)
		case 2:
			InitQuestDun(30, 16)
		}
		return
	}

	// Generate vertical quest areas.
	if !*VertRoom1 && *VertRoom2 && *VertRoom3 {
		if engine.RandCap(0, 2) != 0 {
			v = 2
		}
	}
	if *VertRoom1 && *VertRoom2 && !*VertRoom3 {
		if engine.RandCap(0, 2) != 0 {
			v = 0
		}
	}
	if *VertRoom1 && !*VertRoom2 && *VertRoom3 {
		if engine.RandCap(0, 2) != 0 {
			v = 0
		} else {
			v = 2
		}
	}
	if *VertRoom1 && *VertRoom2 && *VertRoom3 {
		v = int(engine.RandCap(0, 3))
	}
	switch v {
	case 0:
		InitQuestDun(16, 2)
	case 1:
		InitQuestDun(16, 16)
	case 2:
		InitQuestDun(16, 30)
	}
}

// generateChamber generates a chamber at the given coordiates with columns on
// the specified sides.
//
// PSX ref: 0x8013E6B4
// PSX sig: void DRLG_L5GChamber__Fiiiiii(int sx, int sy, int topflag, int bottomflag, int leftflag, int rightflag)
//
// ref: 0x40CD86
func generateChamber(xxStart, yyStart int, topRight, bottomLeft, topLeft, bottomRight bool) {
	if topRight {
		gendung.TileIDMap[xxStart+2][yyStart] = uint8(ArchSe)
		gendung.TileIDMap[xxStart+3][yyStart] = uint8(ArchSe)
		gendung.TileIDMap[xxStart+4][yyStart] = uint8(ArchNeArchNw)
		gendung.TileIDMap[xxStart+7][yyStart] = uint8(ArchEndSe)
		gendung.TileIDMap[xxStart+8][yyStart] = uint8(ArchSe)
		gendung.TileIDMap[xxStart+9][yyStart] = uint8(WallSe)
	}
	if bottomLeft {
		gendung.TileIDMap[xxStart+2][yyStart+11] = uint8(WallSwArchSe)
		gendung.TileIDMap[xxStart+3][yyStart+11] = uint8(ArchSe)
		gendung.TileIDMap[xxStart+4][yyStart+11] = uint8(ArchEndSw)
		gendung.TileIDMap[xxStart+7][yyStart+11] = uint8(ArchSwArchSe)
		gendung.TileIDMap[xxStart+8][yyStart+11] = uint8(ArchSe)
		if TileID(gendung.TileIDMap[xxStart+9][yyStart+11]) != WallSwWallSe {
			gendung.TileIDMap[xxStart+9][yyStart+11] = uint8(DirtWallSwWallSe)
		}
	}
	if topLeft {
		gendung.TileIDMap[xxStart][yyStart+2] = uint8(ArchSw)
		gendung.TileIDMap[xxStart][yyStart+3] = uint8(ArchSw)
		gendung.TileIDMap[xxStart][yyStart+4] = uint8(ArchNeArchNw)
		gendung.TileIDMap[xxStart][yyStart+7] = uint8(ArchEndSw)
		gendung.TileIDMap[xxStart][yyStart+8] = uint8(ArchSw)
		gendung.TileIDMap[xxStart][yyStart+9] = uint8(WallSw)
	}
	if bottomRight {
		gendung.TileIDMap[xxStart+11][yyStart+2] = uint8(ArchSwWallSe)
		gendung.TileIDMap[xxStart+11][yyStart+3] = uint8(ArchSw)
		gendung.TileIDMap[xxStart+11][yyStart+4] = uint8(ArchEndSe)
		gendung.TileIDMap[xxStart+11][yyStart+7] = uint8(ArchSwArchSe)
		gendung.TileIDMap[xxStart+11][yyStart+8] = uint8(ArchSw)
		if TileID(gendung.TileIDMap[xxStart+11][yyStart+9]) != WallSwWallSe {
			gendung.TileIDMap[xxStart+11][yyStart+9] = uint8(DirtWallSwWallSe)
		}
	}
	for yyDelta := 1; yyDelta < 10; yyDelta++ {
		for xxDelta := 1; xxDelta < 10; xxDelta++ {
			FlagMap[xxStart+xxDelta][yyStart+yyDelta] |= Flag40
			gendung.TileIDMap[xxStart+xxDelta][yyStart+yyDelta] = uint8(Floor)
		}
	}
	gendung.TileIDMap[xxStart+4][yyStart+4] = uint8(Column)
	gendung.TileIDMap[xxStart+7][yyStart+4] = uint8(Column)
	gendung.TileIDMap[xxStart+4][yyStart+7] = uint8(Column)
	gendung.TileIDMap[xxStart+7][yyStart+7] = uint8(Column)
}

// generateHall generates a hall of columns and arches.
//
// PSX ref: 0x8013E974
// PSX sig: void DRLG_L5GHall__Fiiii(int x1, int y1, int x2, int y2)
//
// ref: 0x40CEC7
func generateHall(xxStart, yyStart, xxEnd, yyEnd int) {
	if yyStart == yyEnd {
		for xx := xxStart; xx < xxEnd; xx++ {
			gendung.TileIDMap[xx][yyStart] = uint8(ArchSe)
			gendung.TileIDMap[xx][yyStart+3] = uint8(ArchSe)
		}
	} else {
		for yy := yyStart; yy < yyEnd; yy++ {
			gendung.TileIDMap[xxStart][yy] = uint8(ArchSw)
			gendung.TileIDMap[xxStart+3][yy] = uint8(ArchSw)
		}
	}
}

// initQuestDun initializes tile IDs based on the loaded quest dungeon file.
//
// PSX ref: 0x8013F4F8
// PSX sig: void DRLG_L5SetRoom__Fii(int rx1, int ry1)
//
// ref: 0x40CF17
func initQuestDun(xxStart, yyStart int) {
	// Maximum number of elements contained in the single player quest DUN file
	// (i.e. banner2.dun).
	const maxSize = 1 + 1 + 8*8
	sh := reflect.SliceHeader{Data: uintptr(unsafe.Pointer(*SinglePlayerQuestDun)), Len: maxSize, Cap: maxSize}
	qdun := *(*[]uint16)(unsafe.Pointer(&sh))
	width := int(qdun[0])
	height := int(qdun[1])
	// Actual size of single player quest DUN.
	n := 1 + 1 + width*height
	qdun = qdun[:n:n]
	// Place DUN.
	*gendung.SetXx = int32(xxStart)
	*gendung.SetYy = int32(yyStart)
	*gendung.SetWidth = int32(width)
	*gendung.SetHeight = int32(height)
	i := 2
	for yy := yyStart; yy < yyStart+height; yy++ {
		for xx := xxStart; xx < xxStart+width; xx++ {
			tileID := qdun[i]
			if tileID != 0 {
				FlagMap[xx][yy] |= FlagDone
				gendung.TileIDMap[xx][yy] = uint8(tileID)
			} else {
				gendung.TileIDMap[xx][yy] = uint8(Floor)
			}
			i++
		}
	}
}

// floorTransparency adds transparency to concealing walls.
//
// PSX ref: 0x8014016C
// PSX sig: void DRLG_L5FloodTVal__Fv()
//
// ref: 0x40CF9C
func floorTransparency() {
	y := 16
	for yy := 0; yy < 40; yy++ {
		x := 16
		for xx := 0; xx < 40; xx++ {
			if TileID(gendung.TileIDMap[xx][yy]) == Floor && gendung.TransparencyMap[x][y] == 0 {
				FloorTransparencyRecursive(xx, yy, x, y, 0)
				*gendung.TransparencyIndex++
			}
			x += 2
		}
		y += 2
	}
}

// floorTransparencyRecursive recursively adds transparency to concealing walls.
//
// PSX ref: 0x8013FCE4
// PSX sig: void DRLG_L5FTVR__Fiiiii(int i, int j, int x, int y, int d)
//
// ref: 0x40D00B
func floorTransparencyRecursive(xx, yy, x, y, direction int) {
	for gendung.TransparencyMap[x][y] == 0 && TileID(gendung.TileIDMap[xx][yy]) == Floor {
		transIndex := *gendung.TransparencyIndex
		gendung.TransparencyMap[x][y] = transIndex
		gendung.TransparencyMap[x+1][y] = transIndex
		gendung.TransparencyMap[x][y+1] = transIndex
		gendung.TransparencyMap[x+1][y+1] = transIndex
		FloorTransparencyRecursive(xx+1, yy, x+2, y, 1)
		FloorTransparencyRecursive(xx-1, yy, x-2, y, 2)
		FloorTransparencyRecursive(xx, yy+1, x, y+2, 3)
		FloorTransparencyRecursive(xx, yy-1, x, y-2, 4)
		FloorTransparencyRecursive(xx-1, yy-1, x-2, y-2, 5)
		FloorTransparencyRecursive(xx+1, yy-1, x+2, y-2, 6)
		FloorTransparencyRecursive(xx-1, yy+1, x-2, y+2, 7)
		direction = 8
		x += 2
		y += 2
		xx++
		yy++
	}
	transIndex := *gendung.TransparencyIndex
	switch direction {
	case 1:
		gendung.TransparencyMap[x][y] = transIndex
		gendung.TransparencyMap[x][y+1] = transIndex
	case 2:
		gendung.TransparencyMap[x+1][y] = transIndex
		gendung.TransparencyMap[x+1][y+1] = transIndex
	case 3:
		gendung.TransparencyMap[x][y] = transIndex
		gendung.TransparencyMap[x+1][y] = transIndex
	case 4:
		gendung.TransparencyMap[x][y+1] = transIndex
		gendung.TransparencyMap[x+1][y+1] = transIndex
	case 5:
		gendung.TransparencyMap[x+1][y+1] = transIndex
	case 6:
		gendung.TransparencyMap[x][y+1] = transIndex
	case 7:
		gendung.TransparencyMap[x+1][y] = transIndex
	case 8:
		gendung.TransparencyMap[x][y] = transIndex
	}
}

// fixTransparency fixes transparency close to dirt tile IDs after dungeon
// generation.
//
// PSX ref: 0x80140264
// PSX sig: void DRLG_L5TransFix__Fv()
//
// ref: 0x40D1FB
func fixTransparency() {
	y := 16
	for yy := 0; yy < 40; yy++ {
		x := 16
		for xx := 0; xx < 40; xx++ {
			switch TileID(gendung.TileIDMap[xx][yy]) {
			case DirtWallSw:
				gendung.TransparencyMap[x+1][y] = gendung.TransparencyMap[x][y]
				gendung.TransparencyMap[x+1][y+1] = gendung.TransparencyMap[x][y]
			case DirtWallSe:
				gendung.TransparencyMap[x][y+1] = gendung.TransparencyMap[x][y]
				gendung.TransparencyMap[x+1][y+1] = gendung.TransparencyMap[x][y]
			case DirtWallNeWallNw:
				gendung.TransparencyMap[x+1][y] = gendung.TransparencyMap[x][y]
				gendung.TransparencyMap[x][y+1] = gendung.TransparencyMap[x][y]
				gendung.TransparencyMap[x+1][y+1] = gendung.TransparencyMap[x][y]
			case DirtWallEndSw:
				// NOTE: BUG? original checks outside of buffer for yy = 0 in
				// gendung.TileIDMap[xx][yy-1].
				if yy-1 >= 0 && TileID(gendung.TileIDMap[xx][yy-1]) == DirtWallSw {
					gendung.TransparencyMap[x+1][y] = gendung.TransparencyMap[x][y]
					gendung.TransparencyMap[x+1][y+1] = gendung.TransparencyMap[x][y]
				}
			case DirtWallEndSe:
				if xx+1 < 40 && TileID(gendung.TileIDMap[xx+1][yy]) == DirtWallSe {
					gendung.TransparencyMap[x][y+1] = gendung.TransparencyMap[x][y]
					gendung.TransparencyMap[x+1][y+1] = gendung.TransparencyMap[x][y]
				}
			}
			x += 2
		}
		y += 2
	}
}

// fixDirt fixes dirt tile IDs after dungeon generation.
//
// PSX ref: 0x801406A8
// PSX sig: void DRLG_L5DirtFix__Fv()
//
// ref: 0x40D283
func fixDirt() {
	for yy := 0; yy < 40; yy++ {
		for xx := 0; xx < 40; xx++ {
			switch TileID(gendung.TileIDMap[xx][yy]) {
			case DirtWallSw:
				if TileID(gendung.TileIDMap[xx][yy+1]) != DirtWallSw {
					gendung.TileIDMap[xx][yy] = uint8(DirtWallSwDirt)
				}
			case DirtWallSe:
				if TileID(gendung.TileIDMap[xx+1][yy]) != DirtWallSe {
					gendung.TileIDMap[xx][yy] = uint8(DirtWallSeDirt)
				}
			case DirtWallSwWallSe:
				if TileID(gendung.TileIDMap[xx+1][yy]) != DirtWallSe {
					gendung.TileIDMap[xx][yy] = uint8(DirtWallSwWallSeDirt)
				}
				if TileID(gendung.TileIDMap[xx][yy+1]) != DirtWallSw {
					gendung.TileIDMap[xx][yy] = uint8(DirtWallSwWallSeDirt)
				}
			case DirtWallEndSw:
				if TileID(gendung.TileIDMap[xx][yy+1]) != DirtWallSw {
					gendung.TileIDMap[xx][yy] = uint8(DirtWallEndSwDirt)
				}
			case DirtWallEndSe:
				if TileID(gendung.TileIDMap[xx+1][yy]) != DirtWallSe {
					gendung.TileIDMap[xx][yy] = uint8(DirtWallEndSeDirt)
				}
			}
		}
	}
}

// fixCorners fixes corner and arch tile IDs after dungeon generation.
//
// PSX ref: 0x80140824
// PSX sig: void DRLG_L5CornerFix__Fv()
//
// ref: 0x40D2EF
func fixCorners() {
	for yy := 1; yy < 39; yy++ {
		for xx := 1; xx < 39; xx++ {
			if FlagMap[xx][yy]&FlagDone == 0 &&
				TileID(gendung.TileIDMap[xx][yy]) == ArchEndNw &&
				TileID(gendung.TileIDMap[xx-1][yy]) == Floor &&
				TileID(gendung.TileIDMap[xx][yy-1]) == WallSw {
				FlagMap[xx][yy-1] &= FlagDone
				gendung.TileIDMap[xx][yy] = uint8(ArchEndNe)
			}
			if TileID(gendung.TileIDMap[xx][yy]) == DirtWallSwWallSeDirt &&
				TileID(gendung.TileIDMap[xx+1][yy]) == Floor &&
				TileID(gendung.TileIDMap[xx][yy+1]) == WallSw {
				gendung.TileIDMap[xx][yy] = uint8(ArchEndSw)
			}
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
	sh := reflect.SliceHeader{Data: data, Len: n, Cap: n}
	return *(*[]til.Tile)(unsafe.Pointer(&sh))
}
