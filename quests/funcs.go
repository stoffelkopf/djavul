package quests

// #include <stdint.h>
// typedef uint32_t bool32_t;
// typedef uint8_t quest_id;
//
// static bool32_t __fastcall quests_is_active(quest_id quest_num) {
// 	bool32_t (__fastcall *f)(quest_id) = (void *)0x451831;
// 	return f(quest_num);
// }
import "C"

// quests_is_active reports whether the given quest is active.
//
// NOTE: quest_num and quest_id are equivalent, as indicated by this function.
//
// PSX ref: 0x80067B70
// PSX def: unsigned char QuestStatus__Fi(int i)
//
// ref: 0x451831
func IsActive(questNum QuestID) bool {
	return C.quests_is_active(C.quest_id(questNum)) == 1
}
