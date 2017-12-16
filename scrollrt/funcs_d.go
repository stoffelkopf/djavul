//+build djavul

package scrollrt

// #include <stdint.h>
// typedef int32_t bool32_t;
//
// void __fastcall scrollrt_draw_main(int height, bool32_t update_description_box, bool32_t update_life_orb, bool32_t update_mana_orb, bool32_t update_belt, bool32_t update_control_buttons) {
//    void (__fastcall *f)(int, bool32_t, bool32_t, bool32_t, bool32_t, bool32_t) = (void*)0x456124;
//    f(height, update_description_box, update_life_orb, update_mana_orb, update_belt, update_control_buttons);
// }
//
// void scrollrt_draw_main_W(void) {
//    void (*f)(void) = (void*)0x4564F9;
//    f();
// }
import "C"

import (
	"github.com/sanctuary/djavul/engine"
)

// DrawMain renders the specified parts of the UI and the game world on screen.
//
// ref: 0x456124
func DrawMain(height int, updateDescriptionBox, updateLifeOrb, updateManaOrb, updateBelt, updateControlButtons bool) {
	if engine.UseGUI {
		drawMain(height, updateDescriptionBox, updateLifeOrb, updateManaOrb, updateBelt, updateControlButtons)
	}
	C.scrollrt_draw_main(C.int(height), bool32(updateDescriptionBox), bool32(updateLifeOrb), bool32(updateManaOrb), bool32(updateBelt), bool32(updateControlButtons))
}

// DrawMainW renders the UI and the game world on screen.
//
// ref: 0x4564F9
func DrawMainW() {
	C.scrollrt_draw_main_W()
}

// ### [ Helper functions ] ####################################################

// bool32 converts the given boolean value from Go to C.
func bool32(v bool) C.bool32_t {
	if v {
		return 1
	}
	return 0
}
