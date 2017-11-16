//+build djavul

package scrollrt

// void scrollrt_draw_main_W(void) {
//    void (*f)(void) = (void*)0x4564F9;
//    f();
// }
import "C"

import (
	"github.com/sanctuary/djavul/engine"
)

// DrawMainW renders the UI and the game world on screen.
//
// ref: 0x4564F9
func DrawMainW() {
	if engine.UseGUI {
		drawMainW()
	}
	C.scrollrt_draw_main_W()
}
