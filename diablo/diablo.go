// Package diablo implements top-level functions of the Diablo 1 game.
package diablo

import "C"

import (
	"fmt"
	"unsafe"

	"github.com/sanctuary/djavul/engine"
	"github.com/sanctuary/djavul/gendung"
	"github.com/sanctuary/formats/level/min"
	"github.com/sanctuary/formats/level/til"
)

// loadLevelGraphics loads the tile graphics of the active dungeon type.
//
// PSX ref: 0x80038930
// PSX def: void LoadLvlGFX__Fv()
//
// ref: 0x40A391
func loadLevelGraphics() {
	switch *gendung.DType {
	case gendung.Tristram:
		*gendung.LevelCEL = engine.MemLoadFile(unsafe.Pointer(C.CString(`Levels\TownData\Town.CEL`)), nil)
		*gendung.TileDefs = (*til.Tile)(unsafe.Pointer(engine.MemLoadFile(unsafe.Pointer(C.CString(`Levels\TownData\Town.TIL`)), nil)))
		*gendung.DPieceDefs = (*min.Block)(unsafe.Pointer(engine.MemLoadFile(unsafe.Pointer(C.CString(`Levels\TownData\Town.MIN`)), nil)))
		*gendung.LevelSpecialCEL = engine.MemLoadFile(unsafe.Pointer(C.CString(`Levels\TownData\TownS.CEL`)), nil)
	case gendung.Cathedral:
		*gendung.LevelCEL = engine.MemLoadFile(unsafe.Pointer(C.CString(`Levels\L1Data\L1.CEL`)), nil)
		*gendung.TileDefs = (*til.Tile)(unsafe.Pointer(engine.MemLoadFile(unsafe.Pointer(C.CString(`Levels\L1Data\L1.TIL`)), nil)))
		*gendung.DPieceDefs = (*min.Block)(unsafe.Pointer(engine.MemLoadFile(unsafe.Pointer(C.CString(`Levels\L1Data\L1.MIN`)), nil)))
		*gendung.LevelSpecialCEL = engine.MemLoadFile(unsafe.Pointer(C.CString(`Levels\L1Data\L1S.CEL`)), nil)
	case gendung.Catacombs:
		*gendung.LevelCEL = engine.MemLoadFile(unsafe.Pointer(C.CString(`Levels\L2Data\L2.CEL`)), nil)
		*gendung.TileDefs = (*til.Tile)(unsafe.Pointer(engine.MemLoadFile(unsafe.Pointer(C.CString(`Levels\L2Data\L2.TIL`)), nil)))
		*gendung.DPieceDefs = (*min.Block)(unsafe.Pointer(engine.MemLoadFile(unsafe.Pointer(C.CString(`Levels\L2Data\L2.MIN`)), nil)))
		*gendung.LevelSpecialCEL = engine.MemLoadFile(unsafe.Pointer(C.CString(`Levels\L2Data\L2S.CEL`)), nil)
	case gendung.Caves:
		*gendung.LevelCEL = engine.MemLoadFile(unsafe.Pointer(C.CString(`Levels\L3Data\L3.CEL`)), nil)
		*gendung.TileDefs = (*til.Tile)(unsafe.Pointer(engine.MemLoadFile(unsafe.Pointer(C.CString(`Levels\L3Data\L3.TIL`)), nil)))
		*gendung.DPieceDefs = (*min.Block)(unsafe.Pointer(engine.MemLoadFile(unsafe.Pointer(C.CString(`Levels\L3Data\L3.MIN`)), nil)))
		*gendung.LevelSpecialCEL = engine.MemLoadFile(unsafe.Pointer(C.CString(`Levels\L1Data\L1S.CEL`)), nil)
	case gendung.Hell:
		*gendung.LevelCEL = engine.MemLoadFile(unsafe.Pointer(C.CString(`Levels\L4Data\L4.CEL`)), nil)
		*gendung.TileDefs = (*til.Tile)(unsafe.Pointer(engine.MemLoadFile(unsafe.Pointer(C.CString(`Levels\L4Data\L4.TIL`)), nil)))
		*gendung.DPieceDefs = (*min.Block)(unsafe.Pointer(engine.MemLoadFile(unsafe.Pointer(C.CString(`Levels\L4Data\L4.MIN`)), nil)))
		*gendung.LevelSpecialCEL = engine.MemLoadFile(unsafe.Pointer(C.CString(`Levels\L2Data\L2S.CEL`)), nil)
	default:
		panic(fmt.Errorf("unknown dungeon type %v", *gendung.DType))
	}
}
