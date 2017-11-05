//+build !djavul

package l1

import (
	"unsafe"
)

// ResetMaps resets the dungeon flag, player, NPC, dead, object, item, missile
// and arch maps.
//
// PSX ref: 0x8013CEAC
// PSX sig: void DRLG_Init_Globals__Fv()
//
// ref: 0x40ADD6
func ResetMaps() {
	resetMaps()
}

// RandomizeStoneFloor randomizes floor tiles.
//
// PSX ref: 0x8013CAC4
// PSX sig: void DRLG_L1Floor__Fv()
//
// ref: 0x40AF65
func RandomizeStoneFloor() {
	randomizeStoneFloor()
}

// InitPieceIDMap initializes the dungeon piece ID map.
//
// PSX ref: 0x8013CBA8
// PSX sig: void DRLG_L1Pass3__Fv()
//
// ref: 0x40AFB3
func InitPieceIDMap() {
	initPieceIDMap()
}

// InitArches initializes arches.
//
// PSX ref: 0x8013CF5C
// PSX sig: void DRLG_InitL1Vals__Fv()
//
// ref: 0x40B0A5
func InitArches() {
	initArches()
}

// CreateDungeon creates a random Cathedral dungeon based on the given seed and
// level entry.
//
// PSX ref: 0x80140E64
// PSX sig: void CreateL5Dungeon__FUii(unsigned int rseed, int entry)
//
// ref: 0x40B229
func CreateDungeon(seed, entry int32) {
	createDungeon(seed, entry)
}

// LoadSinglePlayerQuestDun loads tile IDs from the dungeon file of the active
// single player quest level.
//
// PSX ref: 0x8013CDA0
// PSX sig: void DRLG_LoadL1SP__Fv()
//
// ref: 0x40B276
func LoadSinglePlayerQuestDun() {
	loadSinglePlayerQuestDun()
}

// FreeSinglePlayerQuestDun frees the dungeon file of the active single player
// quest level.
//
// PSX ref: 0x8013CE7C
// PSX sig: void DRLG_FreeL1SP__Fv()
//
// ref: 0x40B2F4
func FreeSinglePlayerQuestDun() {
	freeSinglePlayerQuestDun()
}

// GenerateDungeon generates a Cathedral dungeon based on the given level entry.
//
// PSX ref: 0x80140930
// PSX sig: void DRLG_L5__Fi(int entry)
//
// ref: 0x40B306
func GenerateDungeon(entry int32) {
	generateDungeon(entry)
}

// PlaceDoor places a door at the given coordinate.
//
// PSX ref: 0x8013BCB0
// PSX sig: void DRLG_PlaceDoor__Fii(int x, int y)
//
// ref: 0x40B56F
func PlaceDoor(xx, yy int) {
	placeDoor(xx, yy)
}

// InitShadows initializes arch and bar shadows.
//
// PSX ref: 0x8013C190
// PSX sig: void DRLG_L1Shadows__Fv()
//
// ref: 0x40B699
func InitShadows() {
	initShadows()
}

// PlaceMiniset places the given miniset of tile IDs.
//
// PSX ref: 0x8013C5A0
// PSX sig: int DRLG_PlaceMiniSet__FPCUciiiiiii(unsigned char *miniset, int tmin, int tmax, int cx, int cy, int setview, int noquad, int ldir)
//
// ref: 0x40B881
func PlaceMiniset(miniset unsafe.Pointer, tmin, tmax, cx, cy int, setView bool, noquad, ldir int) int {
	return placeMiniset(miniset, tmin, tmax, cx, cy, setView, noquad, ldir)
}

// Reset resets the tile ID and the dungeon flag maps.
//
// PSX ref: 0x8013D2F8
// PSX sig: void InitL5Dungeon__Fv()
//
// ref: 0x40BAF6
func Reset() {
	reset()
}

// ClearFlags clears the dungeon generation flags 0x40.
//
// PSX ref: 0x8013D37C
// PSX sig: void L5ClearFlags__Fv()
//
// ref: 0x40BB18
func ClearFlags() {
	clearFlags()
}

// GenerateFirstRoom generates the first room of the dungeon.
//
// PSX ref: 0x8013D7FC
// PSX sig: void L5firstRoom__Fv()
//
// ref: 0x40BB33
func GenerateFirstRoom() {
	generateFirstRoom()
}

// AddRoom adds walls for a room at the given area.
//
// PSX ref: 0x8013D3CC
// PSX sig: void L5drawRoom__Fiiii(int x, int y, int w, int h)
//
// ref: 0x40BD66
func AddRoom(xxStart, yyStart, xxCount, yyCount int) {
	addRoom(xxStart, yyStart, xxCount, yyCount)
}

// GenerateRoom generates a room of the given dimensions at the specified
// coordinates.
//
// PSX ref: 0x8013D4CC
// PSX sig: void L5roomGen__Fiiiii(int x, int y, int w, int h, int dir)
//
// ref: 0x40BD9D
func GenerateRoom(xxStart, yyStart, xxCount, yyCount int, dirHoriz bool) {
	generateRoom(xxStart, yyStart, xxCount, yyCount, dirHoriz)
}

// IsAreaEmpty reports whether the given area is empty.
//
// PSX ref: 0x8013D438
// PSX sig: unsigned char L5checkRoom__Fiiii(int x, int y, int width, int height)
//
// ref: 0x40BFA4
func IsAreaEmpty(xxStart, yyStart, xxCount, yyCount int) bool {
	return isAreaEmpty(xxStart, yyStart, xxCount, yyCount)
}

// GetArea returns the number of walls on the map.
//
// PSX ref: 0x8013DB9C
// PSX sig: long L5GetArea__Fv()
//
// ref: 0x40C008
func GetArea() int {
	return getArea()
}

// InitTileBitMap initializes a tile ID map of twice the size, repeating each
// tile in blocks of 4.
//
// PSX ref: 0x8013DBFC
// PSX sig: void L5makeDungeon__Fv()
//
// ref: 0x40C02A
func InitTileBitMap() {
	initTileBitMap()

}

// GeneratePattern replaces tile ID patterns based on a lookup table.
//
// PSX ref: 0x8013DC88
// PSX sig: void L5makeDmt__Fv()
//
// ref: 0x40C06E
func GeneratePattern() {
	generatePattern()

}

// AddWall adds wall, arch or bar tile IDs.
//
// PSX ref: 0x8013E458
// PSX sig: void L5AddWall__Fv()
//
// ref: 0x40C0E0
func AddWall() {
	addWall()
}

// GetHorizWallSpace returns the number of horizontal wall tiles that fit at the
// given coordinate.
//
// PSX ref: 0x8013DD70
// PSX sig: int L5HWallOk__Fii(int i, int j)
//
// ref: 0x40C23C
func GetHorizWallSpace(xx, yy int) int {
	return getHorizWallSpace(xx, yy)
}

// GetVertWallSpace returns the number of vertical wall tiles that fit at the
// given coordinate.
//
// PSX ref: 0x8013DEAC
// PSX sig: int L5VWallOk__Fii(int i, int j)
//
// ref: 0x40C2DC
func GetVertWallSpace(xx, yy int) int {
	return getVertWallSpace(xx, yy)
}

// AddHorizWall adds a horizontal wall based on the given tile ID.
//
// PSX ref: 0x8013DFF4
// PSX sig: void L5HorizWall__Fiici(int i, int j, char p, int dx)
//
// ref: 0x40C35B
func AddHorizWall(xx, yy int, tileIDFirst TileID, xxCount int) {
	addHorizWall(xx, yy, tileIDFirst, xxCount)
}

// AddVertWall adds a vertical wall based on the given tile ID.
//
// PSX ref: 0x8013E22C
// PSX sig: void L5VertWall__Fiici(int i, int j, char p, int dy)
//
// ref: 0x40C449
func AddVertWall(xx int, yy int, tileIDFirst TileID, yyCount int) {
	addVertWall(xx, yy, tileIDFirst, yyCount)
}

// FixTiles fixes tile IDs of wall edges.
//
// PSX ref: 0x8013EA28
// PSX sig: void L5tileFix__Fv()
//
// ref: 0x40C551
func FixTiles() {
	fixTiles()
}

// Decorate decorates the dungeon with tapestry tile IDs.
//
// PSX ref: 0x8013F2EC
// PSX sig: void DRLG_L5Subs__Fv()
//
// ref: 0x40C8C0
func Decorate() {
	decorate()
}

// GenerateChambers generates chambers.
//
// PSX ref: 0x8013F5F8
// PSX sig: void L5FillChambers__Fv()
//
// ref: 0x40C99D
func GenerateChambers() {
	generateChambers()
}

// GenerateChamber generates a chamber at the given coordiates with columns on
// the specified sides.
//
// PSX ref: 0x8013E6B4
// PSX sig: void DRLG_L5GChamber__Fiiiiii(int sx, int sy, int topflag, int bottomflag, int leftflag, int rightflag)
//
// ref: 0x40CD86
func GenerateChamber(xxStart, yyStart int, topRight, bottomLeft, topLeft, bottomRight bool) {
	generateChamber(xxStart, yyStart, topRight, bottomLeft, topLeft, bottomRight)
}

// GenerateHall generates a hall of columns and arches.
//
// PSX ref: 0x8013E974
// PSX sig: void DRLG_L5GHall__Fiiii(int x1, int y1, int x2, int y2)
//
// ref: 0x40CEC7
func GenerateHall(xxStart, yyStart, xxEnd, yyEnd int) {
	generateHall(xxStart, yyStart, xxEnd, yyEnd)
}

// InitQuestDun initializes tile IDs based on the loaded quest dungeon file.
//
// PSX ref: 0x8013F4F8
// PSX sig: void DRLG_L5SetRoom__Fii(int rx1, int ry1)
//
// ref: 0x40CF17
func InitQuestDun(xxStart, yyStart int) {
	initQuestDun(xxStart, yyStart)
}

// FloorTransparency adds transparency to concealing walls.
//
// PSX ref: 0x8014016C
// PSX sig: void DRLG_L5FloodTVal__Fv()
//
// ref: 0x40CF9C
func FloorTransparency() {
	floorTransparency()
}

// FloorTransparencyRecursive recursively adds transparency to concealing walls.
//
// PSX ref: 0x8013FCE4
// PSX sig: void DRLG_L5FTVR__Fiiiii(int i, int j, int x, int y, int d)
//
// ref: 0x40D00B
func FloorTransparencyRecursive(xx, yy, x, y, direction int) {
	floorTransparencyRecursive(xx, yy, x, y, direction)
}

// FixTransparency fixes transparency close to dirt tile IDs after dungeon
// generation.
//
// PSX ref: 0x80140264
// PSX sig: void DRLG_L5TransFix__Fv()
//
// ref: 0x40D1FB
func FixTransparency() {
	fixTransparency()
}

// FixDirt fixes dirt tile IDs after dungeon generation.
//
// PSX ref: 0x801406A8
// PSX sig: void DRLG_L5DirtFix__Fv()
//
// ref: 0x40D283
func FixDirt() {
	fixDirt()
}

// FixCorners fixes corner and arch tile IDs after dungeon generation.
//
// PSX ref: 0x80140824
// PSX sig: void DRLG_L5CornerFix__Fv()
//
// ref: 0x40D2EF
func FixCorners() {
	fixCorners()
}
