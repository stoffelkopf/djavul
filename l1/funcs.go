package l1

// #include <stdint.h>
// typedef int32_t bool32_t;
// typedef uint8_t l1_tile_id;
//
// static void drlg_l1_reset_maps(void) {
//    void (*f)(void) = (void*)0x40ADD6;
//    f();
// }
//
// static void __fastcall drlg_l1_load_dun(char *dun_path, int view_x, int view_y) {
//    void (__fastcall *f)(char *, int, int) = (void*)0x40AE79;
//    f(dun_path, view_x, view_y);
// }
//
// static void drlg_l1_randomize_stone_floor(void) {
//    void (*f)(void) = (void*)0x40AF65;
//    f();
// }
//
// static void drlg_l1_init_piece_id_map(void) {
//    void (*f)(void) = (void*)0x40AFB3;
//    f();
// }
//
// static void drlg_l1_init_arches(void) {
//    void (*f)(void) = (void*)0x40B0A5;
//    f();
// }
//
// static void __fastcall drlg_l1_preload_dun(char *dun_path, int view_x, int view_y) {
//    void (__fastcall *f)(char *, int, int) = (void*)0x40B160;
//    f(dun_path, view_x, view_y);
// }
//
// static void __fastcall drlg_l1_create_dungeon(uint32_t seed, int entry) {
//    void (__fastcall *f)(uint32_t, int) = (void*)0x40B229;
//    f(seed, entry);
// }
//
// static void drlg_l1_load_quest_dun(void) {
//    void (*f)(void) = (void*)0x40B276;
//    f();
// }
//
// static void drlg_l1_free_quest_dun(void) {
//    void (*f)(void) = (void*)0x40B2F4;
//    f();
// }
//
// static void drlg_l1_generate_dungeon(void) {
//    void (*f)(void) = (void*)0x40B306;
//    f();
// }
//
// static void __fastcall drlg_l1_place_door(int xx, int yy) {
// 	void (__fastcall *f)(int, int) = (void *)0x40B56F;
// 	f(xx, yy);
// }
//
// static void drlg_l1_init_shadows() {
// 	void (*f)() = (void *)0x40B699;
// 	f();
// }
//
// static int __fastcall drlg_l1_place_miniset(uint8_t *miniset, int tmin, int tmax, int cx, int cy, bool32_t set_view, int noquad, int ldir) {
// 	int (__fastcall *f)(uint8_t *, int, int, int, int, bool32_t, int, int) = (void *)0x40B881;
// 	return f(miniset, tmin, tmax, cx, cy, set_view, noquad, ldir);
// }
//
// static void drlg_l1_reset() {
// 	void (*f)() = (void *)0x40BAF6;
// 	f();
// }
//
// static void drlg_l1_clear_flags() {
// 	void (*f)() = (void *)0x40BB18;
// 	f();
// }
//
// static void drlg_l1_generate_first_room() {
// 	void (*f)() = (void *)0x40BB33;
// 	f();
// }
//
// static void __fastcall drlg_l1_add_room(int xx_start, int yy_start, int xx_count, int yy_count) {
// 	void (__fastcall *f)(int, int, int, int) = (void *)0x40BD66;
// 	f(xx_start, yy_start, xx_count, yy_count);
// }
//
// static void __fastcall drlg_l1_generate_room(int xx_start, int yy_start, int xx_count, int yy_count, int dir) {
// 	void (__fastcall *f)(int, int, int, int, int) = (void *)0x40BD9D;
// 	f(xx_start, yy_start, xx_count, yy_count, dir);
// }
//
// static bool32_t __fastcall drlg_l1_is_area_empty(int xx_start, int yy_start, int xx_count, int yy_count) {
// 	bool32_t (__fastcall *f)(int, int, int, int) = (void *)0x40BFA4;
// 	f(xx_start, yy_start, xx_count, yy_count);
// }
//
// static int drlg_l1_get_area() {
// 	int (*f)() = (void *)0x40C008;
// 	return f();
// }
//
// static void drlg_l1_init_dmap() {
// 	void (*f)() = (void *)0x40C02A;
// 	f();
// }
//
// static void drlg_l1_generate_pattern() {
// 	void (*f)() = (void *)0x40C06E;
// 	f();
// }
//
// static void drlg_l1_add_wall() {
// 	void (*f)() = (void *)0x40C0E0;
// 	f();
// }
//
// static int __fastcall drlg_l1_get_horiz_wall_space(int xx, int yy) {
// 	int (__fastcall *f)(int, int) = (void *)0x40C23C;
// 	return f(xx, yy);
// }
//
// static int __fastcall drlg_l1_get_vert_wall_space(int xx, int yy) {
// 	int (__fastcall *f)(int, int) = (void *)0x40C2DC;
// 	return f(xx, yy);
// }
//
// static void __fastcall drlg_l1_add_horiz_wall(int xx, int yy, l1_tile_id tile_id, int xx_count) {
// 	void (__fastcall *f)(int, int, l1_tile_id, int) = (void *)0x40C35B;
// 	f(xx, yy, tile_id, xx_count);
// }
//
// static void __fastcall drlg_l1_add_vert_wall(int xx, int yy, l1_tile_id tile_id, int yy_count) {
// 	void (__fastcall *f)(int, int, l1_tile_id, int) = (void *)0x40C449;
// 	f(xx, yy, tile_id, yy_count);
// }
//
// static void drlg_l1_fix_tiles() {
// 	void (*f)() = (void *)0x40C551;
// 	f();
// }
//
// static void drlg_l1_decorate() {
// 	void (*f)() = (void *)0x40C8C0;
// 	f();
// }
//
// static void drlg_l1_generate_chambers() {
// 	void (*f)() = (void *)0x40C99D;
// 	f();
// }
//
// static void __fastcall drlg_l1_generate_chamber(int xx_start, int yy_start, bool32_t top_right, bool32_t bottom_left, bool32_t top_left, bool32_t bottom_right) {
// 	void (__fastcall *f)(int, int, bool32_t, bool32_t, bool32_t, bool32_t) = (void *)0x40CD86;
// 	f(xx_start, yy_start, top_right, bottom_left, top_left, bottom_right);
// }
//
// static void __fastcall drlg_l1_generate_hall(int xx_start, int yy_start, int xx_end, int yy_end) {
// 	void (__fastcall *f)(int, int, int, int) = (void *)0x40CEC7;
// 	f(xx_start, yy_start, xx_end, yy_end);
// }
//
// static void __fastcall drlg_l1_init_quest_dun(int xx_start, int yy_start) {
// 	void (__fastcall *f)(int, int) = (void *)0x40CF17;
// 	f(xx_start, yy_start);
// }
//
// static void drlg_l1_floor_transparency() {
// 	void (*f)() = (void *)0x40CF9C;
// 	f();
// }
//
// static void __fastcall drlg_l1_floor_transparency_recursive(int xx, int yy, int x, int y, int direction) {
// 	void (__fastcall *f)(int, int, int, int, int) = (void *)0x40D00B;
// 	f(xx, yy, x, y, direction);
// }
//
// static void drlg_l1_fix_transparency() {
// 	void (*f)() = (void *)0x40D1FB;
// 	f();
// }
//
// static void drlg_l1_fix_dirt() {
// 	void (*f)() = (void *)0x40D283;
// 	f();
// }
//
// static void drlg_l1_fix_corners() {
// 	void (*f)() = (void *)0x40D2EF;
// 	f();
// }
import "C"

import (
	"reflect"
	"unsafe"

	"github.com/pkg/errors"
	"github.com/sanctuary/djavul/engine"
	"github.com/sanctuary/djavul/gendung"
	"github.com/sanctuary/formats/level/til"
)

// ResetMaps resets the dungeon flag, player, NPC, dead, object, item, missile
// and arch maps.
//
// PSX ref: 0x8013CEAC
// PSX sig: void DRLG_Init_Globals__Fv()
//
// ref: 0x40ADD6
func ResetMaps() {
	C.drlg_l1_reset_maps()
}

// LoadDun loads tile IDs, monsters and objects from the given dungeon file.
//
// PSX ref: 0x8013CF64
// PSX sig: void LoadL1Dungeon__FPcii(char *sFileName, int vx, int vy)
//
// ref: 0x40AE79
func LoadDun(dunPath unsafe.Pointer, viewX, viewY int32) {
	C.drlg_l1_load_dun((*C.char)(dunPath), C.int(viewX), C.int(viewY))
}

// RandomizeStoneFloor randomizes floor tiles.
//
// PSX ref: 0x8013CAC4
// PSX sig: void DRLG_L1Floor__Fv()
//
// ref: 0x40AF65
func RandomizeStoneFloor() {
	C.drlg_l1_randomize_stone_floor()
}

// InitPieceIDMap initializes the dungeon piece ID map.
//
// PSX ref: 0x8013CBA8
// PSX sig: void DRLG_L1Pass3__Fv()
//
// ref: 0x40AFB3
func InitPieceIDMap() {
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

// InitArches initializes arches.
//
// PSX ref: 0x8013CF5C
// PSX sig: void DRLG_InitL1Vals__Fv()
//
// ref: 0x40B0A5
func InitArches() {
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

// PreloadDun loads tile IDs from the given dungeon file.
//
// PSX ref: 0x8013D138
// PSX sig: void LoadPreL1Dungeon__FPcii(char *sFileName, int vx, int vy)
//
// ref: 0x40B160
func PreloadDun(dunPath *int8, viewX, viewY int32) {
	C.drlg_l1_preload_dun((*C.char)(unsafe.Pointer(dunPath)), C.int(viewX), C.int(viewY))
}

// CreateDungeon creates a random cathedral dungeon based on the given seed and
// level entry.
//
// PSX ref: 0x80140E64
// PSX sig: void CreateL5Dungeon__FUii(unsigned int rseed, int entry)
//
// ref: 0x40B229
func CreateDungeon(seed, entry int32) {
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

// LoadQuestLun loads tile IDs from the dungeon file of the active quest level.
//
// PSX ref: 0x8013CDA0
// PSX sig: void DRLG_LoadL1SP__Fv()
//
// ref: 0x40B276
func LoadQuestDun() {
	C.drlg_l1_load_quest_dun()
}

// FreeQuestDun frees the dungeon file of the active quest level.
//
// PSX ref: 0x8013CE7C
// PSX sig: void DRLG_FreeL1SP__Fv()
//
// ref: 0x40B2F4
func FreeQuestDun() {
	C.drlg_l1_free_quest_dun()
}

// GenerateDungeon generates a cathedral dungeon based on the given level entry.
//
// PSX ref: 0x80140930
// PSX sig: void DRLG_L5__Fi(int entry)
//
// ref: 0x40B306
func GenerateDungeon(entry int32) {
	C.drlg_l1_generate_dungeon()
}

// PlaceDoor places a door at the given coordinate.
//
// PSX ref: 0x8013BCB0
// PSX sig: void DRLG_PlaceDoor__Fii(int x, int y)
//
// ref: 0x40B56F
func PlaceDoor(xx, yy int) {
	C.drlg_l1_place_door(C.int(xx), C.int(yy))
}

// InitShadows initializes arch and bar shadows.
//
// PSX ref: 0x8013C190
// PSX sig: void DRLG_L1Shadows__Fv()
//
// ref: 0x40B699
func InitShadows() {
	C.drlg_l1_init_shadows()
}

// PlaceMiniset places the given miniset of tile IDs.
//
// PSX ref: 0x8013C5A0
// PSX sig: int DRLG_PlaceMiniSet__FPCUciiiiiii(unsigned char *miniset, int tmin, int tmax, int cx, int cy, int setview, int noquad, int ldir)
//
// ref: 0x40B881
func PlaceMiniset(miniset unsafe.Pointer, tmin, tmax, cx, cy int, setView bool, noquad, ldir int) int {
	return int(C.drlg_l1_place_miniset((*C.uint8_t)(miniset), C.int(tmin), C.int(tmax), C.int(cx), C.int(cy), bool32(setView), C.int(noquad), C.int(ldir)))
}

// Reset resets the tile ID and the dungeon flag maps.
//
// PSX ref: 0x8013D2F8
// PSX sig: void InitL5Dungeon__Fv()
//
// ref: 0x40BAF6
func Reset() {
	C.drlg_l1_reset()
}

// ClearFlags clears the dungeon generation flags 0x40.
//
// PSX ref: 0x8013D37C
// PSX sig: void L5ClearFlags__Fv()
//
// ref: 0x40BB18
func ClearFlags() {
	C.drlg_l1_clear_flags()
}

// GenerateFirstRoom generates the first room of the dungeon.
//
// PSX ref: 0x8013D7FC
// PSX sig: void L5firstRoom__Fv()
//
// ref: 0x40BB33
func GenerateFirstRoom() {
	C.drlg_l1_generate_first_room()
}

// AddRoom adds walls for a room at the given area.
//
// PSX ref: 0x8013D3CC
// PSX sig: void L5drawRoom__Fiiii(int x, int y, int w, int h)
//
// ref: 0x40BD66
func AddRoom(xxStart, yyStart, xxCount, yyCount int) {
	C.drlg_l1_add_room(C.int(xxStart), C.int(yyStart), C.int(xxCount), C.int(yyCount))
}

// GenerateRoom generates a room of the given dimensions at the spcified
// coordinates.
//
// PSX ref: 0x8013D4CC
// PSX sig: void L5roomGen__Fiiiii(int x, int y, int w, int h, int dir)
//
// ref: 0x40BD9D
func GenerateRoom(xxStart, yyStart, xxCount, yyCount, dir int) {
	C.drlg_l1_generate_room(C.int(xxStart), C.int(yyStart), C.int(xxCount), C.int(yyCount), C.int(dir))
}

// IsAreaEmpty reports whether the given area is empty.
//
// PSX ref: 0x8013D438
// PSX sig: unsigned char L5checkRoom__Fiiii(int x, int y, int width, int height)
//
// ref: 0x40BFA4
func IsAreaEmpty(xxStart, yyStart, xxCount, yyCount int) bool {
	return C.drlg_l1_is_area_empty(C.int(xxStart), C.int(yyStart), C.int(xxCount), C.int(yyCount)) == 1
}

// GetArea returns the number of walls on the map.
//
// PSX ref: 0x8013DB9C
// PSX sig: long L5GetArea__Fv()
//
// ref: 0x40C008
func GetArea() int {
	return int(C.drlg_l1_get_area())
}

// InitDmap initializes a dungeon tile ID map of twice the size of the dungeon,
// repeating each tile in blocks of 4.
//
// NOTE: The dmap (double map) seems to be unused.
//
// PSX ref: 0x8013DBFC
// PSX sig: void L5makeDungeon__Fv()
//
// ref: 0x40C02A
func InitDmap() {
	C.drlg_l1_init_dmap()
}

// GeneratePattern replaces tile ID patterns based on a lookup table.
//
// PSX ref: 0x8013DC88
// PSX sig: void L5makeDmt__Fv()
//
// ref: 0x40C06E
func GeneratePattern() {
	C.drlg_l1_generate_pattern()
}

// AddWall adds wall, arch or bar tile IDs.
//
// PSX ref: 0x8013E458
// PSX sig: void L5AddWall__Fv()
//
// ref: 0x40C0E0
func AddWall() {
	C.drlg_l1_add_wall()
}

// GetHorizWallSpace returns the number of horizontal wall tiles that fit at the
// given coordinate.
//
// PSX ref: 0x8013DD70
// PSX sig: int L5HWallOk__Fii(int i, int j)
//
// ref: 0x40C23C
func GetHorizWallSpace(xx, yy int) int {
	return int(C.drlg_l1_get_horiz_wall_space(C.int(xx), C.int(yy)))
}

// GetVertWallSpace returns the number of vertical wall tiles that fit at the
// given coordinate.
//
// PSX ref: 0x8013DEAC
// PSX sig: int L5VWallOk__Fii(int i, int j)
//
// ref: 0x40C2DC
func GetVertWallSpace(xx, yy int) int {
	return int(C.drlg_l1_get_vert_wall_space(C.int(xx), C.int(yy)))
}

// AddHorizWall adds a horizontal wall based on the given tile ID.
//
// PSX ref: 0x8013DFF4
// PSX sig: void L5HorizWall__Fiici(int i, int j, char p, int dx)
//
// ref: 0x40C35B
func AddHorizWall(xx, yy int, tileID TileID, xxCount int) {
	C.drlg_l1_add_horiz_wall(C.int(xx), C.int(yy), C.l1_tile_id(tileID), C.int(xxCount))
}

// AddVertWall adds a vertical wall based on the given tile ID.
//
// PSX ref: 0x8013E22C
// PSX sig: void L5VertWall__Fiici(int i, int j, char p, int dy)
//
// ref: 0x40C449
func AddVertWall(xx int, yy int, tileID TileID, yyCount int) {
	C.drlg_l1_add_vert_wall(C.int(xx), C.int(yy), C.l1_tile_id(tileID), C.int(yyCount))
}

// FixTiles fixes tile IDs of wall edges.
//
// PSX ref: 0x8013EA28
// PSX sig: void L5tileFix__Fv()
//
// ref: 0x40C551
func FixTiles() {
	C.drlg_l1_fix_tiles()
}

// Decorate decorates the dungeon with tapestry tile IDs.
//
// PSX ref: 0x8013F2EC
// PSX sig: void DRLG_L5Subs__Fv()
//
// ref: 0x40C8C0
func Decorate() {
	C.drlg_l1_decorate()
}

// GenerateChambers generates chambers.
//
// PSX ref: 0x8013F5F8
// PSX sig: void L5FillChambers__Fv()
//
// ref: 0x40C99D
func GenerateChambers() {
	C.drlg_l1_generate_chambers()
}

// GenerateChamber generates a chamber at the given coordiates with columns on
// the specified sides.
//
// PSX ref: 0x8013E6B4
// PSX sig: void DRLG_L5GChamber__Fiiiiii(int sx, int sy, int topflag, int bottomflag, int leftflag, int rightflag)
//
// ref: 0x40CD86
func GenerateChamber(xxStart, yyStart int, topRight, bottomLeft, topLeft, bottomRight bool) {
	C.drlg_l1_generate_chamber(C.int(xxStart), C.int(yyStart), bool32(topRight), bool32(bottomLeft), bool32(topLeft), bool32(bottomRight))
}

// GenerateHall generates a hall of columns and arches.
//
// PSX ref: 0x8013E974
// PSX sig: void DRLG_L5GHall__Fiiii(int x1, int y1, int x2, int y2)
//
// ref: 0x40CEC7
func GenerateHall(xxStart, yyStart, xxEnd, yyEnd int) {
	C.drlg_l1_generate_hall(C.int(xxStart), C.int(yyStart), C.int(xxEnd), C.int(yyEnd))
}

// InitQuestDun initializes tile IDs based on the loaded quest dungeon file.
//
// PSX ref: 0x8013F4F8
// PSX sig: void DRLG_L5SetRoom__Fii(int rx1, int ry1)
//
// ref: 0x40CF17
func InitQuestDun(xx_start, yy_start int) {
	C.drlg_l1_init_quest_dun(C.int(xx_start), C.int(yy_start))
}

// FloorTransparency adds transparency to concealing walls.
//
// PSX ref: 0x8014016C
// PSX sig: void DRLG_L5FloodTVal__Fv()
//
// ref: 0x40CF9C
func FloorTransparency() {
	C.drlg_l1_floor_transparency()
}

// FloorTransparencyRecursive recursively adds transparency to concealing walls.
//
// PSX ref: 0x8013FCE4
// PSX sig: void DRLG_L5FTVR__Fiiiii(int i, int j, int x, int y, int d)
//
// ref: 0x40D00B
func FloorTransparencyRecursive(xx, yy, x, y, direction int) {
	C.drlg_l1_floor_transparency_recursive(C.int(xx), C.int(yy), C.int(x), C.int(y), C.int(direction))
}

// FixTransparency fixes transparency close to dirt tile IDs after dungeon
// generation.
//
// PSX ref: 0x80140264
// PSX sig: void DRLG_L5TransFix__Fv()
//
// ref: 0x40D1FB
func FixTransparency() {
	C.drlg_l1_fix_transparency()
}

// FixDirt fixes dirt tile IDs after dungeon generation.
//
// PSX ref: 0x801406A8
// PSX sig: void DRLG_L5DirtFix__Fv()
//
// ref: 0x40D283
func FixDirt() {
	C.drlg_l1_fix_dirt()
}

// FixCorners fixes corner and arch tile IDs after dungeon generation.
//
// PSX ref: 0x80140824
// PSX sig: void DRLG_L5CornerFix__Fv()
//
// ref: 0x40D2EF
func FixCorners() {
	C.drlg_l1_fix_corners()
}

// ### [ Helper functions ] ####################################################

// bool32 converts the given boolean value from Go to C.
func bool32(v bool) C.bool32_t {
	if v {
		return 1
	}
	return 0
}

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
