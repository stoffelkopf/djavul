package quests

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
	ID QuestID
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
	ID QuestID
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
	_ uint32
}
