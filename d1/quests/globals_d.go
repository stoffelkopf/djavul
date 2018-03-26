//+build djavul

// Global variable wrappers for quests.cpp

package quests

import "unsafe"

// Global variables.
var (
	// QuestsData contains the data related to each quest ID.
	//
	// References:
	//    * https://github.com/sanctuary/notes/blob/master/enums.h#quest_id
	//
	// ref: 0x4A1AE0
	QuestsData = (*[16]QuestData)(unsafe.Pointer(uintptr(0x4A1AE0)))

	// Quests contains the quests of the current game.
	//
	// PSX ref: 0x800DDA40
	// PSX def: QuestStruct quests[16];
	//
	// ref: 0x69BD10
	Quests = (*[16]Quest)(unsafe.Pointer(uintptr(0x69BD10)))
)
