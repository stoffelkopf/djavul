package pfile

import "github.com/sanctuary/djavul/player"

// HeroInfo specifies information about the hero to be presented in the user
// interface.
//
// PSX def:
//    typedef struct _uiheroinfo {
//       struct _uiheroinfo* next;
//       char name[16];
//       unsigned short level;
//       unsigned char heroclass;
//       unsigned char herorank;
//       unsigned short strength;
//       unsigned short magic;
//       unsigned short dexterity;
//       unsigned short vitality;
//       unsigned long gold;
//       unsigned char hassaved;
//       unsigned char spawned;
//    } TUIHEROINFO, _uiheroinfo;
type HeroInfo struct {
	// offset: 0000 (4 bytes)
	Next *HeroInfo
	// offset: 0004 (16 bytes)
	Name [16]byte
	// offset: 0014 (2 bytes)
	CLvl int16
	// offset: 0016 (1 bytes)
	PlayerClass player.Class
	// offset: 0017 (1 bytes)
	Difficulty int8 // TODO: use difficulty enum
	// offset: 0018 (2 bytes)
	StrCur int16
	// offset: 001A (2 bytes)
	MagCur int16
	// offset: 001C (2 bytes)
	DexCur int16
	// offset: 001E (2 bytes)
	VitCur int16
	// offset: 0020 (4 bytes)
	GoldTotal int32
	// offset: 0024 (4 bytes)
	HasSave bool32
	// offset: 0028 (4 bytes)
	Spawned bool32
}

type bool32 int32
