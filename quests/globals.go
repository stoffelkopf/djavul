// Global variable wrappers for quests.cpp

package quests

import "unsafe"

// Global variables.
var (
	// Quests contains the quests of the current game.
	//
	// PSX ref: 0x800DDA40
	// PSX def: QuestStruct quests[16];
	//
	// ref: 0x69BD10
	Quests = (*[16]Quest)(unsafe.Pointer(uintptr(0x69BD10)))
)

// Quest describes in-game state of any quest.
//
// PSX def:
//     typedef struct QuestStruct {
//        unsigned char _qlevel;
//        unsigned char _qtype;
//        unsigned char _qactive;
//        unsigned char _qlvltype;
//        int _qtx;
//        int _qty;
//        unsigned char _qslvl;
//        unsigned char _qidx;
//        unsigned char _qmsg;
//        unsigned char _qvar1;
//        unsigned char _qvar2;
//        unsigned char _qlog;
//        unsigned char pad_for_laz;
//     } QuestStruct;
type Quest struct {
	// offset 00000000
	DLvl uint8
	// offset 00000001
	ID uint8
	// offset 00000002
	Active bool
	// offset 00000003
	_ uint8
	// offset 00000004
	EnteranceX int32
	// offset 00000008
	EnteranceY int32
	// offset 0000000C
	QuestLevel uint8
	// offset 0000000D
	_ uint8
	// offset 0000000E
	SpeechID uint8
	// offset 0000000F
	Var1 uint8
	// offset 00000010
	_ uint8
	// offset 00000011
	_ uint8
	// offset 00000012
	_ uint8
	// offset 00000013
	_ uint8
	// offset 00000014
	SpeechSpoken int32
}
