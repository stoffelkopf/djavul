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

// PSX def:
//    typedef struct QuestData {
//       unsigned char _qdlvl;
//       char _qdmultlvl;
//       unsigned char _qlvlt;
//       unsigned char _qdtype;
//       unsigned char _qdrnd;
//       unsigned char _qslvl;
//       unsigned char _qflags;
//       int _qdmsg;
//       int _qlstr;
//    } QuestData;
type QuestData struct {
	// offset: 00000000
	_ uint8
	// offset: 00000001
	_ uint8
	// offset: 00000002
	_ uint8
	// offset: 00000003
	_ uint8
	// offset: 00000004
	_ uint8
	// offset: 00000005
	_ uint8
	// offset: 00000006
	_ uint8
	// offset: 00000007
	_ uint8
	// offset: 00000008
	Multiplayer bool
	// offset: 00000009
	_ uint8
	// offset: 0000000A
	_ uint8
	// offset: 0000000B
	_ uint8
	// offset: 0000000C
	_ uint8
	// offset: 0000000D
	_ uint8
	// offset: 0000000E
	_ uint8
	// offset: 0000000F
	_ uint8
	// offset: 00000010
	Name unsafe.Pointer
}

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
