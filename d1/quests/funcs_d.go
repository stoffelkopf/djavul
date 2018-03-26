//+build djavul

package quests

// #include <stdint.h>
// typedef uint32_t bool32_t;
// typedef uint8_t quest_id;
//
// static bool32_t __fastcall quests_is_active(quest_id quest_num) {
// 	bool32_t (__fastcall *f)(quest_id) = (void *)0x451831;
// 	return f(quest_num);
// }
//
// static void quests_init_the_butcher_area() {
//    void (*f)() = (void *)0x451BEA;
//    f();
// }
//
// static void __fastcall quests_init_the_curse_of_king_leoric_area(int quest_id, int xx, int yy) {
//    void (__fastcall *f)(int, int, int) = (void *)0x451C11;
//    f(quest_id, xx, yy);
// }
//
// static void __fastcall quests_init_warlord_of_blood_area(int xx, int yy) {
//    void (__fastcall *f)(int, int) = (void *)0x451C32;
//    f(xx, yy);
// }
//
// static void __fastcall quests_init_the_chamber_of_bone_area(int quest_id, int xx, int yy) {
//    void (__fastcall *f)(int, int, int) = (void *)0x451CC2;
//    f(quest_id, xx, yy);
// }
//
// static void __fastcall quests_init_odgens_sign_area(int xx, int yy) {
//    void (__fastcall *f)(int, int) = (void *)0x451D7C;
//    f(xx, yy);
// }
//
// static void __fastcall quests_init_halls_of_the_blind_area(int xx, int yy) {
//    void (__fastcall *f)(int, int) = (void *)0x451E08;
//    f(xx, yy);
// }
//
// static void __fastcall quests_init_valor_area(int xx, int yy) {
//    void (__fastcall *f)(int, int) = (void *)0x451E94;
//    f(xx, yy);
// }
//
// static void __fastcall quests_init_quest_area(int xx, int yy) {
//    void (__fastcall *f)(int, int) = (void *)0x451F20;
//    f(xx, yy);
// }
import "C"

// useGo specifies whether to use the Go implementation.
const useGo = true

// IsActive reports whether the given quest is active.
//
// NOTE: quest_num and quest_id are equivalent, as indicated by this function.
//
// PSX ref: 0x80067B70
// PSX def: unsigned char QuestStatus__Fi(int i)
//
// ref: 0x451831
func IsActive(questNum QuestID) bool {
	if useGo {
		return isActive(questNum)
	} else {
		return C.quests_is_active(C.quest_id(questNum)) == 1
	}
}

// InitTheButcherArea initializes the quest area of The Butcher.
//
// PSX ref: 0x8015ED8C
// PSX def: void DrawButcher__Fv()
//
// ref: 0x451BEA
func InitTheButcherArea() {
	if useGo {
		initTheButcherArea()
	} else {
		C.quests_init_the_butcher_area()
	}
}

// InitTheCurseOfKingLeoricArea initializes the quest area of The Curse of King
// Leoric.
//
// PSX ref: 0x8015EDD0
// PSX def: void DrawSkelKing__Fiii(int q, int x, int y)
//
// ref: 0x451C11
func InitTheCurseOfKingLeoricArea(questID QuestID, xx, yy int32) {
	if useGo {
		initTheCurseOfKingLeoricArea(questID, xx, yy)
	} else {
		C.quests_init_the_curse_of_king_leoric_area(C.int(questID), C.int(xx), C.int(yy))
	}
}

// InitWarlordOfBloodArea initializes the quest area of Warlord of Blood.
//
// PSX ref: 0x8015EE64
// PSX def: void DrawWarLord__Fii(int x, int y)
//
// ref: 0x451C32
func InitWarlordOfBloodArea(xx, yy int32) {
	C.quests_init_warlord_of_blood_area(C.int(xx), C.int(yy))
}

// InitTheChamberOfBoneArea initializes the quest area of The Chamber of Bone.
//
// PSX ref: 0x8015EF60
// PSX def: void DrawSChamber__Fiii(int q, int x, int y)
//
// ref: 0x451CC2
func InitTheChamberOfBoneArea(questID QuestID, xx, yy int32) {
	C.quests_init_the_chamber_of_bone_area(C.int(questID), C.int(xx), C.int(yy))
}

// InitOdgensSignArea initializes the quest area of Odgen's Sign.
//
// PSX ref: 0x8015F09C
// PSX def: void DrawLTBanner__Fii(int x, int y)
//
// ref: 0x451D7C
func InitOdgensSignArea(xxStart, yyStart int32) {
	if useGo {
		initOdgensSignArea(xxStart, yyStart)
	} else {
		C.quests_init_odgens_sign_area(C.int(xxStart), C.int(yyStart))
	}
}

// InitHallsOfTheBlindArea initializes the quest area of Halls of the Blind.
//
// PSX ref: 0x8015F178
// PSX def: void DrawBlind__Fii(int x, int y)
//
// ref: 0x451E08
func InitHallsOfTheBlindArea(xx, yy int32) {
	C.quests_init_halls_of_the_blind_area(C.int(xx), C.int(yy))
}

// InitValorArea initializes the quest area of Valor.
//
// PSX ref: 0x8015F254
// PSX def: void DrawBlood__Fii(int x, int y)
//
// ref: 0x451E94
func InitValorArea(xx, yy int32) {
	C.quests_init_valor_area(C.int(xx), C.int(yy))
}

// InitQuestArea initializes the given quest area.
//
// PSX ref: 0x8015F334
// PSX def: void DRLG_CheckQuests__Fii(int x, int y)
//
// ref: 0x451F20
func InitQuestArea(xx, yy int32) {
	if useGo {
		initQuestArea(xx, yy)
	} else {
		C.quests_init_quest_area(C.int(xx), C.int(yy))
	}
}
