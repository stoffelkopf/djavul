package path

// #include <stdint.h>
// typedef int32_t bool32_t;
//
// bool32_t __fastcall player_is_valid_step(int player_num, int x, int y) {
//    bool32_t (__fastcall *f)(int player_num, int x, int y) = (void *) 0x44FDBA;
//    return f(player_num, x, y);
// }
//
// int __fastcall path_make(bool32_t (__fastcall *is_valid_step)(int entity_num, int x, int y), int entity_num, int start_x, int start_y, int target_x, int target_y, int8_t *steps) {
//    int (__fastcall *f)(bool32_t (__fastcall *is_valid_step)(int entity_num, int x, int y), int entity_num, int start_x, int start_y, int target_x, int target_y, int8_t *steps) = (void *) 0x4493D4;
//    return f(is_valid_step, entity_num, start_x, start_y, target_x, target_y, steps);
// }
import "C"
import "unsafe"

func Make(validStep func(entityNum, x, y int) bool, entityNum, startX, startY, targetX, targetY int, steps []int8) int {
	playerValidStep := (*[0]byte)(unsafe.Pointer(C.player_is_valid_step))
	n := C.path_make(playerValidStep, C.int(entityNum), C.int(startX), C.int(startY), C.int(targetX), C.int(targetY), (*C.int8_t)(unsafe.Pointer(&steps[0])))
	return int(n)
}
