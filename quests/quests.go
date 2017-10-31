// Package quests implements quest functions.
package quests

import (
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
