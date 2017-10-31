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

// InitQuestArea initializes the given quest area.
//
// PSX ref: 0x8015F334
// PSX def: void DRLG_CheckQuests__Fii(int x, int y)
//
// ref: 0x451F20
func InitQuestArea(xx, yy int32) {
	C.quests_init_quest_area(C.int(xx), C.int(yy))
}
