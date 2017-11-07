// Package quests implements quest functions.
package quests

import "C"

import (
	"log"
	"reflect"
	"unsafe"

	"github.com/sanctuary/djavul/engine"
	"github.com/sanctuary/djavul/gendung"
	"github.com/sanctuary/djavul/multi"
)

// isActive reports whether the given quest is active.
//
// NOTE: quest_num and quest_id are equivalent, as indicated by this function.
//
// PSX ref: 0x80067B70
// PSX def: unsigned char QuestStatus__Fi(int i)
//
// ref: 0x451831
func isActive(questNum QuestID) bool {
	if !Quests[questNum].Active {
		return false
	}
	if *gendung.DLvl != Quests[questNum].DLvl {
		return false
	}
	if *gendung.IsQuestLevel {
		return false
	}
	// Multiplayer quests:
	// * The Butcher
	// * The Curse of King Leoric
	// * Archbishop Lazarus
	// * Diablo
	if *multi.MaxPlayers != 1 && !QuestsData[questNum].Multiplayer {
		return false
	}
	return true
}

// initTheButcherArea initializes the quest area of The Butcher.
//
// PSX ref: 0x8015ED8C
// PSX def: void DrawButcher__Fv()
//
// ref: 0x451BEA
func initTheButcherArea() {
	gendung.RectTransparent(int(*gendung.SetXx*2+19), int(*gendung.SetYy*2+19), int(*gendung.SetXx*2+26), int(*gendung.SetYy*2+26))
}

// initTheCurseOfKingLeoricArea initializes the quest area of The Curse of King
// Leoric.
//
// PSX ref: 0x8015EDD0
// PSX def: void DrawSkelKing__Fiii(int q, int x, int y)
//
// ref: 0x451C11
func initTheCurseOfKingLeoricArea(questID QuestID, xx, yy int32) {
	Quests[questID].EnteranceX = 2*xx + 28
	Quests[questID].EnteranceY = 2*yy + 23
}

// initWarlordOfBloodArea initializes the quest area of Warlord of Blood.
//
// PSX ref: 0x8015EE64
// PSX def: void DrawWarLord__Fii(int x, int y)
//
// ref: 0x451C32
func initWarlordOfBloodArea(xx, yy int32) {
	// TODO: Implement initWarlordOfBloodArea.
	log.Print("note: quests.initWarlordOfBloodArea not yet implemented.")
}

// initTheChamberOfBoneArea initializes the quest area of The Chamber of Bone.
//
// PSX ref: 0x8015EF60
// PSX def: void DrawSChamber__Fiii(int q, int x, int y)
//
// ref: 0x451CC2
func initTheChamberOfBoneArea(questID QuestID, xx, yy int32) {
	// TODO: Implement initTheChamberOfBoneArea.
	log.Print("note: quests.initTheChamberOfBoneArea not yet implemented.")
}

// initOdgensSignArea initializes the quest area of Odgen's Sign.
//
// PSX ref: 0x8015F09C
// PSX def: void DrawLTBanner__Fii(int x, int y)
//
// ref: 0x451D7C
func initOdgensSignArea(xxStart, yyStart int32) {
	buf := engine.MemLoadFile(unsafe.Pointer(C.CString(`Levels\L1Data\Banner1.DUN`)), nil)
	// Maximum number of elements contained in the single player quest DUN file
	// (i.e. banner2.dun).
	const maxSize = 1 + 1 + 8*8
	sh := reflect.SliceHeader{Data: uintptr(unsafe.Pointer(buf)), Len: maxSize, Cap: maxSize}
	qdun := *(*[]uint16)(unsafe.Pointer(&sh))
	width := int32(qdun[0])
	height := int32(qdun[1])
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
			if qdun[i] != 0 {
				gendung.TileIDMapBackup[xx][yy] = uint8(qdun[i])
			}
			i++
		}
	}
}

// initHallsOfTheBlindArea initializes the quest area of Halls of the Blind.
//
// PSX ref: 0x8015F178
// PSX def: void DrawBlind__Fii(int x, int y)
//
// ref: 0x451E08
func initHallsOfTheBlindArea(xx, yy int32) {
	// TODO: Implement initHallsOfTheBlindArea.
	log.Print("note: quests.initHallsOfTheBlindArea not yet implemented.")
}

// initValorArea initializes the quest area of Valor.
//
// PSX ref: 0x8015F254
// PSX def: void DrawBlood__Fii(int x, int y)
//
// ref: 0x451E94
func initValorArea(xx, yy int32) {
	// TODO: Implement initValorArea.
	log.Print("note: quests.initValorArea not yet implemented.")
}

// initQuestArea initializes the given quest area.
//
// PSX ref: 0x8015F334
// PSX def: void DRLG_CheckQuests__Fii(int x, int y)
//
// ref: 0x451F20
func initQuestArea(xx, yy int32) {
	questID := QuestID(0)
	for _, quest := range Quests {
		if IsActive(questID) {
			switch quest.ID {
			case TheButcher:
				InitTheButcherArea()
			case OgdensSign:
				InitOdgensSignArea(xx, yy) // TODO: Add test case
			case HallsOfTheBlind:
				InitHallsOfTheBlindArea(xx, yy)
			case Valor:
				InitValorArea(xx, yy)
			case WarlordOfBlood:
				InitWarlordOfBloodArea(xx, yy)
			case TheCurseOfKingLeoric:
				InitTheCurseOfKingLeoricArea(questID, xx, yy)
			case TheChamberOfBone:
				InitTheChamberOfBoneArea(questID, xx, yy)
			}
		}
		questID++
	}
}
