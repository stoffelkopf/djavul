//+build !djavul

package diablo

// LoadLevelGraphics loads the tile graphics of the active dungeon type.
//
// PSX ref: 0x80038930
// PSX def: void LoadLvlGFX__Fv()
//
// ref: 0x40A391
func LoadLevelGraphics() {
	loadLevelGraphics()
}
