//+build !djavul

package quests

// IsActive reports whether the given quest is active.
//
// NOTE: questNum and questID are equivalent, as indicated by this function.
//
// PSX ref: 0x80067B70
// PSX def: unsigned char QuestStatus__Fi(int i)
//
// ref: 0x451831
func IsActive(questNum QuestID) bool {
	return isActive(questNum)
}

// InitTheButcherArea initializes the quest area of The Butcher.
//
// PSX ref: 0x8015ED8C
// PSX def: void DrawButcher__Fv()
//
// ref: 0x451BEA
func InitTheButcherArea() {
	initTheButcherArea()
}

// InitTheCurseOfKingLeoricArea initializes the quest area of The Curse of King
// Leoric.
//
// PSX ref: 0x8015EDD0
// PSX def: void DrawSkelKing__Fiii(int q, int x, int y)
//
// ref: 0x451C11
func InitTheCurseOfKingLeoricArea(questID QuestID, xx, yy int32) {
	initTheCurseOfKingLeoricArea(questID, xx, yy)
}

// InitWarlordOfBloodArea initializes the quest area of Warlord of Blood.
//
// PSX ref: 0x8015EE64
// PSX def: void DrawWarLord__Fii(int x, int y)
//
// ref: 0x451C32
func InitWarlordOfBloodArea(xx, yy int32) {
	initWarlordOfBloodArea(xx, yy)
}

// InitTheChamberOfBoneArea initializes the quest area of The Chamber of Bone.
//
// PSX ref: 0x8015EF60
// PSX def: void DrawSChamber__Fiii(int q, int x, int y)
//
// ref: 0x451CC2
func InitTheChamberOfBoneArea(questID QuestID, xx, yy int32) {
	initTheChamberOfBoneArea(questID, xx, yy)
}

// InitOdgensSignArea initializes the quest area of Odgen's Sign.
//
// PSX ref: 0x8015F09C
// PSX def: void DrawLTBanner__Fii(int x, int y)
//
// ref: 0x451D7C
func InitOdgensSignArea(xx, yy int32) {
	initOdgensSignArea(xx, yy)
}

// InitHallsOfTheBlindArea initializes the quest area of Halls of the Blind.
//
// PSX ref: 0x8015F178
// PSX def: void DrawBlind__Fii(int x, int y)
//
// ref: 0x451E08
func InitHallsOfTheBlindArea(xx, yy int32) {
	initHallsOfTheBlindArea(xx, yy)
}

// InitValorArea initializes the quest area of Valor.
//
// PSX ref: 0x8015F254
// PSX def: void DrawBlood__Fii(int x, int y)
//
// ref: 0x451E94
func InitValorArea(xx, yy int32) {
	initValorArea(xx, yy)
}

// InitQuestArea initializes the given quest area.
//
// PSX ref: 0x8015F334
// PSX def: void DRLG_CheckQuests__Fii(int x, int y)
//
// ref: 0x451F20
func InitQuestArea(xx, yy int32) {
	initQuestArea(xx, yy)
}
