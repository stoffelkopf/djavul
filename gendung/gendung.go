// Package gendung implements dungeon generation utility functions.
package gendung

// initTransparency initializes transparency.
//
// PSX ref: 0x8015A070
// PSX def: void DRLG_InitTrans__Fv()
//
// ref: 0x41944A
func initTransparency() {
	for x := range TransparencyMap {
		for y := range TransparencyMap {
			TransparencyMap[x][y] = 0
		}
	}
	for i := range TransparencyActive {
		TransparencyActive[i] = false
	}
	*TransparencyIndex = 1
}
