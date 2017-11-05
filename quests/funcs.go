//+build !djavul

package quests

// IsActive reports whether the given quest is active.
//
// NOTE: quest_num and quest_id are equivalent, as indicated by this function.
//
// PSX ref: 0x80067B70
// PSX def: unsigned char QuestStatus__Fi(int i)
//
// ref: 0x451831
func IsActive(questNum QuestID) bool {
	return isActive(questNum)
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
